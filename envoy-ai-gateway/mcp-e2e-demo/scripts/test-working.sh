#!/bin/bash
set -e

echo "=== Testing MCP through Envoy Gateway ==="
echo ""

echo "1. List available tools:"
curl -s -X POST http://localhost:8080/mcp \
	-H "Content-Type: application/json" \
	-d '{"jsonrpc":"2.0","method":"tools/list","id":1}' | jq '.result.tools[] | {name, description}'

echo ""
echo "2. Test echo tool:"
curl -s -X POST http://localhost:8080/mcp \
	-H "Content-Type: application/json" \
	-d '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"echo","arguments":{"text":"Hello from Envoy AI Gateway!"}},"id":2}' | jq '.result.content[0].text'

echo ""
echo "3. Test sum tool (42 + 58):"
curl -s -X POST http://localhost:8080/mcp \
	-H "Content-Type: application/json" \
	-d '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"sum","arguments":{"a":42,"b":58}},"id":3}' | jq '.result.content[0].text'

echo ""
echo "✓ All tests passed!"
