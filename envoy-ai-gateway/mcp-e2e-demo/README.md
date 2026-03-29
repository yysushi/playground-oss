# MCP E2E Demo

Minimal end-to-end demonstration of MCP (Model Context Protocol) with OAuth-protected MCPRoute behind Envoy AI Gateway on KIND.

## Architecture

```
Client
  ↓ HTTP/JSON-RPC (streamable)
Envoy AI Gateway (KIND)
  ↓ MCPRoute + OAuth
MCP Server (go-sdk streamable HTTP)
```

**What Works:**
- ✅ MCPRoute with OAuth enforcement
- ✅ OAuth metadata endpoints wired
- ✅ Streamable HTTP MCP server (go-sdk)
- ✅ Tool listing via MCP session
- ✅ KIND cluster with automated setup

## Prerequisites

```bash
# macOS
brew install kind kubectl helm jq docker

# Verify
kind version
kubectl version --client
helm version
jq --version
```

Envoy AI Gateway requires Envoy Gateway v1.5.0 or higher. The setup script defaults to v1.6.3 and applies the official AI Gateway Envoy Gateway values in `mcp-e2e-demo/k8s/envoy-gateway-values.yaml`.

## Quick Start

```bash
# 1. Setup (creates cluster, builds images, deploys everything - takes ~3 min)
./scripts/setup.sh

# 2. Test OAuth metadata + 401 enforcement
./scripts/test-working.sh

# 3. Test with a real mockoidc token (session + tools)
./scripts/test-token.sh

# 4. Cleanup
./scripts/cleanup.sh
```

## What It Demonstrates

1. **MCP Server**: Streamable HTTP server using `go-sdk/mcp`
2. **Gateway Routing**: Envoy AI Gateway MCPRoute
3. **OAuth Enforcement**: JWT validation with local JWKS
4. **Metadata Endpoints**: Protected resource + auth server metadata
5. **Tool Execution**: echo and sum via MCP session
6. **KIND Setup**: Production-like local Kubernetes environment

## Manual Testing

```bash
# Fetch OAuth metadata
curl -s http://localhost:8080/.well-known/oauth-protected-resource/mcp | jq
curl -s http://localhost:8080/.well-known/oauth-authorization-server/mcp | jq

For a real token flow, use:

```bash
./scripts/test-token.sh
```
```

## Files

```
mcp-e2e-demo/
├── mcp-server/       # Streamable MCP server (go-sdk)
├── oauth-server/     # mockoidc OAuth provider
├── k8s/all.yaml      # All K8s resources (Gateway, MCPRoute, etc.)
└── scripts/          # Setup, test, cleanup scripts
```

## Troubleshooting

```bash
# Check deployments
kubectl get pods

# View logs
kubectl logs deployment/mcp
kubectl logs deployment/oauth
kubectl logs -n envoy-gateway-system deployment/envoy-gateway

# Check gateway
kubectl get gatewayclass,gateway,mcproute,backend
```

## Key Differences from Complex Demos

- ✅ One YAML for core resources + one for MCPRoute
- ✅ No TLS (simpler)
- ✅ Free OAuth (mockoidc, not external provider)
- ✅ Minimal tools (echo + sum)
- ✅ One-command setup
- ✅ Clear test script

## Next Steps

- Add more MCP tools
- Add authorization policies (scope-based access)
- Add observability (traces, metrics)
- Use real OAuth provider (Auth0, Keycloak)
