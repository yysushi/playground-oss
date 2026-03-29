# MCP E2E Demo

Minimal end-to-end demonstration of MCP (Model Context Protocol) server running behind Envoy AI Gateway on KIND.

## Architecture

```
Client
  ↓ HTTP/JSON-RPC
Envoy AI Gateway (KIND)
  ↓ HTTPRoute
MCP Server (2 tools: echo, sum)
```

**What Works:**
- ✅ MCP Server with 2 tools (echo, sum)
- ✅ Envoy Gateway routing via HTTPRoute
- ✅ Full MCP JSON-RPC protocol support
- ✅ KIND cluster with automated setup

**What's Documented (but requires more config):**
- ⚠️ MCPRoute (needs MCP proxy sidecar - see STATUS.md)
- ⚠️ OAuth integration (mockoidc ready, integration pending)

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

## Quick Start

```bash
# 1. Setup (creates cluster, builds images, deploys everything - takes ~3 min)
./scripts/setup.sh

# 2. Test MCP calls
./scripts/test-working.sh

# 3. Cleanup
./scripts/cleanup.sh
```

## What It Demonstrates

1. **MCP Server**: Simple Go implementation (70 lines!)
2. **Gateway Routing**: Envoy AI Gateway proxying MCP protocol  
3. **Tool Execution**: echo and sum tools via JSON-RPC
4. **KIND Setup**: Production-like local Kubernetes environment

## Manual Testing

```bash
# List tools
curl -s -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"tools/list","id":1}' | jq

# Call echo tool
curl -s -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"echo","arguments":{"text":"Hello!"}},"id":2}' | jq

# Call sum tool  
curl -s -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"sum","arguments":{"a":10,"b":32}},"id":3}' | jq
```

## Files

```
mcp-e2e-demo/
├── mcp-server/       # Simple MCP server (Go)
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
kubectl get gateway,mcproute,backend
```

## Key Differences from Complex Demos

- ✅ Single YAML file (not multiple)
- ✅ No TLS (simpler)
- ✅ Free OAuth (mockoidc, not external provider)
- ✅ Minimal tools (just echo + sum)
- ✅ One-command setup
- ✅ Clear test script

## Next Steps

- Add more MCP tools
- Add authorization policies (scope-based access)
- Add observability (traces, metrics)
- Use real OAuth provider (Auth0, Keycloak)
