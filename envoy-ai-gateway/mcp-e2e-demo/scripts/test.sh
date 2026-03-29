#!/bin/bash
set -e

source .env

echo "=== Testing MCP with OAuth ==="

# 1. Get token
echo "1. Getting OAuth token..."
TOKEN=$(curl -s -X POST http://localhost:8080/oidc/token \
	-u "$CLIENT_ID:$CLIENT_SECRET" \
	-d "grant_type=client_credentials&scope=openid" | jq -r .access_token)

echo "Token: ${TOKEN:0:20}..."

# 2. Initialize MCP
echo "2. Initializing MCP..."
curl -X POST http://localhost:8080/mcp \
	-H "Authorization: Bearer $TOKEN" \
	-H "Content-Type: application/json" \
	-d '{"jsonrpc":"2.0","method":"initialize","id":1}' | jq

# 3. List tools
echo "3. Listing tools..."
curl -X POST http://localhost:8080/mcp \
	-H "Authorization: Bearer $TOKEN" \
	-H "Content-Type: application/json" \
	-d '{"jsonrpc":"2.0","method":"tools/list","id":2}' | jq

# 4. Call echo tool
echo "4. Calling echo tool..."
curl -X POST http://localhost:8080/mcp \
	-H "Authorization: Bearer $TOKEN" \
	-H "Content-Type: application/json" \
	-d '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"echo","arguments":{"text":"Hello MCP!"}},"id":3}' | jq

# 5. Call sum tool
echo "5. Calling sum tool..."
curl -X POST http://localhost:8080/mcp \
	-H "Authorization: Bearer $TOKEN" \
	-H "Content-Type: application/json" \
	-d '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"sum","arguments":{"a":5,"b":3}},"id":4}' | jq

echo ""
echo "✓ All tests passed!"
