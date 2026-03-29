# Envoy Gateway OIDC Demo

Simple POC showing OIDC authentication with Envoy Gateway.

## What This Does

- Runs a mock OIDC server (mockoidc)
- Runs a backend service (nginx)
- Configures Envoy Gateway to protect the backend with OIDC auth

## Quick Start

```bash
./setup.sh      # Deploy everything
./cleanup.sh    # Remove everything
```

## Architecture

```
Client → Envoy Gateway (OIDC) → Backend (nginx)
                ↓
         Mock OIDC Server
```

Super simple. No complexity.

## Files

- `mockoidc-simple.yaml` - OIDC server deployment (43 lines)
- `echo-k8s.yaml` - Backend deployment (33 lines)  
- `oidc.yaml` - Gateway + OIDC config (51 lines)
- `Dockerfile` - OIDC server build (14 lines)
- `main.go` - OIDC server code

**Total: 141 lines of YAML**

## How It Works

1. `setup.sh` builds Docker image
2. Loads image into KIND cluster
3. Deploys OIDC server + backend
4. Configures Envoy Gateway with OIDC

Done in ~10 seconds.

## Testing

After `./setup.sh`:

```bash
# Terminal 1: Port forward Envoy Gateway
kubectl -n envoy-gateway-system port-forward service/envoy-default-eg-e41e7b31 8443:443

# Terminal 2: Port forward OIDC server  
kubectl -n default port-forward service/mockoidc 8888:8888

# Terminal 3: Test
curl -vk https://localhost:8443/myapp 2>&1 | grep -i location
# Should redirect to OIDC login
```

Or open in browser: `https://localhost:8443/myapp`

## Requirements

- kubectl
- helm
- docker
- kind
- envsubst
- openssl

## Why So Simple?

This is a POC. It proves OIDC works with Envoy Gateway.

Not production. Not hardened. Just simple and works.

See `SIMPLIFICATION.md` for details on why we kept it minimal.

## Learning Resources

Want to understand how things work?

- **SERVICE_EXPLAINED.md** - Deep dive into Kubernetes Services (how the mockoidc service routes traffic)
- **SERVICE_QUICK_REF.md** - Quick reference card for Services
- **SIMPLIFICATION.md** - Why we kept configs minimal

These explain the "how" and "why" behind the YAML.
