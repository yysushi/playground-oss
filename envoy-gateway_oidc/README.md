# Envoy Gateway OIDC E2E Demo

Minimal end-to-end demonstration of OIDC authentication with Envoy Gateway on KIND.

## Architecture

```
Client (Browser/curl)
  ↓ HTTPS request
Envoy Gateway (with OIDC SecurityPolicy)
  ↓ OAuth2 redirect
Mock OIDC Server (mockoidc)
  ↓ token validation
Backend (nginx echo service)
```

**What Works:**
- ✅ Mock OIDC server (mockoidc) - free OAuth provider
- ✅ Envoy Gateway with OIDC SecurityPolicy
- ✅ TLS-enabled gateway (self-signed cert)
- ✅ OAuth2 authorization code flow with PKCE
- ✅ Protected backend (nginx)

## Prerequisites

```bash
# macOS
brew install kind kubectl helm docker openssl

# Linux (Ubuntu/Debian)
# Install kind, kubectl, helm, docker, openssl via apt

# Also needed
brew install gettext  # for envsubst

# Verify
kind version
kubectl version --client
helm version
docker --version
openssl version
```

## Quick Start

```bash
# 1. Setup (creates cluster, builds images, deploys - takes ~2 min)
./scripts/setup.sh

# 2. In separate terminals, start port forwards:
# Terminal 2:
kubectl -n envoy-gateway-system port-forward \
  service/$(kubectl get svc -n envoy-gateway-system \
    --selector=gateway.envoyproxy.io/owning-gateway-name=eg \
    -o jsonpath='{.items[0].metadata.name}') 8443:443

# Terminal 3:
kubectl port-forward service/mockoidc 8888:8888

# 3. Test OIDC flow
./scripts/test.sh

# 4. Or test in browser
open https://localhost:8443/myapp

# 5. Cleanup
./scripts/cleanup.sh
```

## What It Demonstrates

1. **OIDC Authentication**: Full OAuth2 authorization code flow with PKCE
2. **SecurityPolicy**: Envoy Gateway's OIDC security configuration
3. **Free OAuth Provider**: mockoidc (no external service needed)
4. **TLS Termination**: Self-signed certificates for HTTPS
5. **Protected Routes**: Backend only accessible after authentication

## Manual Testing

### Test unauthenticated request (should redirect):
```bash
curl -Lk https://localhost:8443/myapp -v 2>&1 | grep -i location
# Should show redirect to /oauth2/authorize
```

### Test OIDC discovery:
```bash
curl -s http://localhost:8888/oidc/.well-known/openid-configuration | jq
```

### Browser flow:
1. Open `https://localhost:8443/myapp`
2. You'll be redirected to OIDC login page
3. Enter any username/password (mockoidc accepts all credentials)
4. After login, you'll see the nginx welcome page

## Files

```
envoy-gateway_oidc/
├── oidc-server/           # Mock OIDC server
│   ├── main.go           # mockoidc wrapper with domain support
│   ├── Dockerfile
│   └── go.mod/go.sum
├── k8s/                   # Kubernetes manifests
│   ├── manifests.yaml    # OIDC server + backend deployments
│   ├── oidc.yaml         # Gateway + OIDC SecurityPolicy
│   └── echo-k8s.yaml     # Backend service (nginx)
├── scripts/
│   ├── setup.sh          # Automated deployment
│   ├── test.sh           # OIDC flow tests
│   └── cleanup.sh        # Teardown
└── README.md             # This file
```

## Troubleshooting

```bash
# Check deployments
kubectl get pods

# View OIDC server logs
kubectl logs deployment/mockoidc

# Check gateway status
kubectl get gateway,httproute,securitypolicy

# Verify OIDC credentials
kubectl logs deployment/mockoidc | grep -E "CLIENT_ID|ISSUER"
```

## Key Technical Details

### OIDC Flow
1. Client requests `/myapp` without auth
2. Gateway redirects to OIDC `/oauth2/authorize`
3. User logs in at OIDC server
4. OIDC redirects back with authorization code
5. Gateway exchanges code for access token
6. Gateway validates token and proxies to backend

### SecurityPolicy Configuration
```yaml
apiVersion: gateway.envoyproxy.io/v1alpha1
kind: SecurityPolicy
spec:
  targetRefs:
  - name: eg
  oidc:
    provider:
      issuer: "http://mockoidc.default.svc.cluster.local:8888/oidc"
    clientID: "${CLIENT_ID}"
    clientSecret:
      name: oidc-client-secret
    redirectURL: "https://localhost:8443/myapp/oauth2/callback"
```

## Learning Resources

- **SERVICE_EXPLAINED.md** - Deep dive into Kubernetes Services
- **SERVICE_QUICK_REF.md** - Quick reference for Services
- **SIMPLIFICATION.md** - Why configs are minimal

## Next Steps

- Add real OIDC provider (Keycloak, Auth0, Okta)
- Add JWT validation and claims-based authorization
- Add observability (metrics, traces)
- Use real TLS certificates (Let's Encrypt)

## Comparison to Other Demos

**vs mcp-e2e-demo:**
- This: OIDC authentication focus
- MCP: MCP protocol + gateway routing focus
- Both: Use KIND, free services, automated setup

**Shared patterns:**
- ✅ One-command setup
- ✅ Clear directory structure
- ✅ Working test scripts
- ✅ Free/self-hosted components
