# Envoy AI Gateway Demos

Two aligned demonstrations of Envoy Gateway features using KIND clusters.

## Quick Comparison

| Aspect | MCP E2E Demo | OIDC E2E Demo |
|--------|--------------|---------------|
| **Location** | `envoy-ai-gateway/mcp-e2e-demo/` | `envoy-gateway_oidc/` |
| **Focus** | MCP protocol gateway | OIDC authentication |
| **Protocol** | MCP JSON-RPC | OAuth2/OIDC |
| **Gateway** | Envoy AI Gateway | Envoy Gateway |
| **Backend** | Custom MCP server | nginx |
| **Auth** | None (HTTPRoute) | OIDC SecurityPolicy |
| **Testing** | Automated tests (curl + jq) | Browser + automated tests |
| **Setup Time** | ~3 min | ~2 min |

## Directory Structure (Aligned)

Both demos follow the same clean structure:

```
demo-name/
├── server/              # Custom server implementation
│   ├── main.go
│   ├── Dockerfile
│   └── go.mod/go.sum
├── k8s/                 # Kubernetes manifests
│   └── *.yaml
├── scripts/             # Automation
│   ├── setup.sh
│   ├── test.sh (or test-working.sh)
│   └── cleanup.sh
├── README.md            # User guide
├── SUMMARY.md           # What we built
└── STATUS.md            # Technical details
```

## Common Patterns

Both demos share these design principles:

### ✅ Automated Setup
- One command: `./scripts/setup.sh`
- Creates KIND cluster
- Installs Envoy Gateway
- Builds and deploys custom images
- Configures everything automatically

### ✅ Working Test Scripts
- Validate the end-to-end flow
- Show concrete examples
- Help troubleshoot issues

### ✅ Free Components
- No external paid services
- No cloud provider dependencies
- Self-hosted OIDC (mockoidc)
- Runs entirely on local machine

### ✅ Clear Documentation
- README: User guide with quick start
- SUMMARY: Architecture and what's working
- STATUS: Technical details and limitations

### ✅ Clean Teardown
- One command: `./scripts/cleanup.sh`
- Removes KIND cluster
- Cleans up all resources

## When to Use Which Demo

### Use MCP E2E Demo to learn:
- ✅ Model Context Protocol (MCP)
- ✅ How to build a minimal MCP server
- ✅ JSON-RPC protocol handling
- ✅ HTTPRoute configuration
- ✅ AI Gateway CRDs and controllers
- ✅ Backend feature configuration

**Best for:** Understanding MCP protocol and AI Gateway routing

### Use OIDC E2E Demo to learn:
- ✅ OAuth2 authorization code flow
- ✅ OIDC authentication patterns
- ✅ Envoy Gateway SecurityPolicy
- ✅ TLS termination
- ✅ Session management basics
- ✅ Protecting backends with auth

**Best for:** Understanding authentication and authorization

## Running Both Demos

The demos use different cluster names and can coexist:

```bash
# MCP Demo
cd envoy-ai-gateway/mcp-e2e-demo
./scripts/setup.sh      # Creates 'mcp-demo' cluster
./scripts/test-working.sh
./scripts/cleanup.sh

# OIDC Demo  
cd ../../envoy-gateway_oidc
./scripts/setup.sh      # Creates 'envoy-gateway-oidc' cluster
# Port forward in separate terminals
./scripts/test.sh
./scripts/cleanup.sh
```

## Architecture Comparison

### MCP E2E Demo
```
Client
  ↓ HTTP POST /mcp (JSON-RPC)
Envoy AI Gateway
  ↓ HTTPRoute
MCP Server
  ↓ Tools: echo, sum
Response
```

### OIDC E2E Demo
```
Client (Browser)
  ↓ HTTPS GET /myapp
Envoy Gateway
  ↓ No auth? Redirect to OIDC
Mock OIDC Server
  ↓ Login & token
Envoy Gateway (validates token)
  ↓ Authorized request
Backend (nginx)
```

## Technical Highlights

### MCP Demo Technical Wins
- ✅ 70-line MCP server implementation
- ✅ Direct JSON-RPC protocol demonstration
- ✅ AI Gateway controller integration
- ✅ Backend CRD feature enablement
- ✅ Clear documentation of MCPRoute challenges

### OIDC Demo Technical Wins
- ✅ Complete OAuth2 flow with PKCE
- ✅ SecurityPolicy configuration
- ✅ Dynamic credential extraction
- ✅ TLS termination setup
- ✅ Session cookie handling

## Files Overview

### MCP Demo Files (~300 lines total)
```
mcp-server/main.go:        70 lines (MCP protocol)
oauth-server/main.go:      ~50 lines (unused but ready)
k8s/all.yaml:              ~100 lines
scripts/setup.sh:          ~50 lines
scripts/test-working.sh:   ~30 lines
```

### OIDC Demo Files (~656 lines total)
```
oidc-server/main.go:       ~180 lines (OIDC + DNS rewriting)
k8s/*:                     ~200 lines
scripts/setup.sh:          ~216 lines  
scripts/test.sh:           ~60 lines
```

## Learning Path

### Beginner Path
1. Start with **MCP Demo** (simpler protocol, fewer components)
2. Understand HTTPRoute and Gateway basics
3. Move to **OIDC Demo** (more complex auth flow)
4. Understand SecurityPolicy and authentication

### Advanced Path
1. Run both demos simultaneously
2. Compare HTTPRoute vs SecurityPolicy approaches
3. Experiment with combining MCP + OIDC
4. Extend with custom tools and auth rules

## Production Considerations

### MCP Demo → Production
- Add OAuth via SecurityPolicy
- Implement MCPRoute (when sidecar config is clear)
- Add observability (traces, metrics)
- Use real tool implementations
- Add rate limiting and quotas

### OIDC Demo → Production
- Replace mockoidc with real provider (Keycloak, Auth0)
- Use proper TLS certificates
- Add session persistence (Redis)
- Implement authorization (claims, roles)
- Add observability

## Combining Both Demos

You can combine learnings from both:

```yaml
# Protected MCP endpoint with OIDC
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: mcp-protected
spec:
  parentRefs:
  - name: gateway
  rules:
  - matches:
    - path:
        value: /mcp
    backendRefs:
    - name: mcp-server
---
apiVersion: gateway.envoyproxy.io/v1alpha1
kind: SecurityPolicy
metadata:
  name: mcp-auth
spec:
  targetRefs:
  - kind: HTTPRoute
    name: mcp-protected
  oidc:
    provider:
      issuer: "https://your-oidc.example.com"
    # ... OIDC config
```

## Community and Support

Both demos are designed to be:
- ✅ Self-documenting (clear code + docs)
- ✅ Easy to modify (minimal dependencies)
- ✅ Easy to share (no secrets or external deps)
- ✅ Educational (comments and explanations)

Perfect for:
- Learning Envoy Gateway features
- Demoing to team members
- Starting point for POCs
- Understanding gateway patterns

## Next Steps

After mastering both demos:

1. **Combine features**: MCP server with OIDC auth
2. **Add observability**: Prometheus + Grafana
3. **Production setup**: Real OIDC, real TLS, HA
4. **Advanced routing**: Rate limiting, retries, circuit breaking
5. **Multi-tenancy**: Multiple backends with different auth

## Credits

Both demos use:
- **Envoy Gateway**: https://gateway.envoyproxy.io
- **Envoy AI Gateway**: https://aigateway.envoyproxy.io  
- **mockoidc**: https://github.com/oauth2-proxy/mockoidc
- **KIND**: https://kind.sigs.k8s.io

Minimal, production-ready patterns for real-world use.
