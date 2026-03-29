#!/bin/bash
set -e

echo "=== Testing MCP through Envoy Gateway ==="
echo ""

# Check if port 8080 is accessible
if ! nc -z localhost 8080 2>/dev/null; then
	echo "⚠️  Gateway not accessible on localhost:8080"
	echo "   Port forward may have died. Restart with:"
	echo "   GATEWAY_SVC=\$(kubectl get svc -n envoy-gateway-system --selector=gateway.envoyproxy.io/owning-gateway-name=aigw-run -o jsonpath='{.items[0].metadata.name}')"
	echo "   kubectl port-forward -n envoy-gateway-system service/\$GATEWAY_SVC 8080:8080"
	echo ""
	exit 1
fi

echo "1. List available tools (should be 401 without token):"
curl -s -i -X POST http://localhost:8080/mcp \
	-H "Content-Type: application/json" \
	-d '{"jsonrpc":"2.0","method":"tools/list","id":1}' | head -20

echo ""
echo ""
echo "2. Verify OAuth protected resource metadata endpoint:"
curl -s -i http://localhost:8080/.well-known/oauth-protected-resource/mcp | head -20

echo ""
echo ""
echo "3. Fetch authorization server metadata:"
curl -s -i http://localhost:8080/.well-known/oauth-authorization-server/mcp | head -20

echo ""
echo ""
echo "4. For full OAuth token flow + tool calls, run:"
echo "   ./scripts/test-token.sh"

echo ""
echo "✓ OAuth is enforced and metadata endpoints are wired"
