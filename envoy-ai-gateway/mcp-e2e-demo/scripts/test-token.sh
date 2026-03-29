#!/bin/bash
set -e

echo "=== Testing MCP with Real mockoidc Token ==="
echo ""

if ! nc -z localhost 8080 2>/dev/null; then
	echo "⚠️  Gateway not accessible on localhost:8080"
	echo "   Port forward may have died. Restart with:"
	echo "   GATEWAY_SVC=\$(kubectl get svc -n envoy-gateway-system --selector=gateway.envoyproxy.io/owning-gateway-name=aigw-run -o jsonpath='{.items[0].metadata.name}')"
	echo "   kubectl port-forward -n envoy-gateway-system service/\$GATEWAY_SVC 8080:8080"
	echo ""
	exit 1
fi

echo "1. Port-forwarding mockoidc..."
kubectl port-forward service/oauth 8081:8080 >/tmp/pf-oauth.log 2>&1 &
OAUTH_PF_PID=$!
trap 'kill $OAUTH_PF_PID >/dev/null 2>&1 || true' EXIT
sleep 2

CLIENT_ID=$(kubectl logs deployment/oauth | awk -F': ' '/CLIENT_ID:/{print $2; exit}')
CLIENT_SECRET=$(kubectl logs deployment/oauth | awk -F': ' '/CLIENT_SECRET:/{print $2; exit}')

echo "2. Requesting token via auth code + PKCE..."
CODE_VERIFIER_CHALLENGE=$(
	python - <<'PY'
import base64, hashlib, secrets
ver=secrets.token_urlsafe(32)
chal=base64.urlsafe_b64encode(hashlib.sha256(ver.encode()).digest()).rstrip(b'=').decode()
print(ver)
print(chal)
PY
)
CODE_VERIFIER=$(echo "$CODE_VERIFIER_CHALLENGE" | head -1)
CODE_CHALLENGE=$(echo "$CODE_VERIFIER_CHALLENGE" | tail -1)
STATE=$(
	python - <<'PY'
import secrets
print(secrets.token_urlsafe(12))
PY
)
AUTH_URL="http://localhost:8081/oidc/authorize?response_type=code&client_id=$CLIENT_ID&redirect_uri=http://localhost:8081/callback&scope=openid&state=$STATE&code_challenge=$CODE_CHALLENGE&code_challenge_method=S256"
LOCATION=$(curl -s -D - -o /dev/null "$AUTH_URL" | awk '/^Location:/{print $2}' | tr -d '\r')
CODE=$(LOCATION="$LOCATION" python -c 'import os,urllib.parse;loc=os.environ["LOCATION"];qs=urllib.parse.parse_qs(urllib.parse.urlparse(loc).query);print(qs.get("code",[""])[0])')
if [ -z "$CODE" ]; then
	echo "Error: failed to obtain authorization code" >&2
	exit 1
fi
TOKEN_RESPONSE=$(curl -s -X POST http://localhost:8081/oidc/token \
	-d "grant_type=authorization_code" \
	-d "code=$CODE" \
	-d "redirect_uri=http://localhost:8081/callback" \
	-d "code_verifier=$CODE_VERIFIER" \
	-d "client_id=$CLIENT_ID" \
	-d "client_secret=$CLIENT_SECRET")
ACCESS_TOKEN=$(TOKEN_RESPONSE="$TOKEN_RESPONSE" python -c 'import os,json;print(json.loads(os.environ["TOKEN_RESPONSE"])["access_token"])')

if [ -z "$ACCESS_TOKEN" ]; then
	echo "Error: failed to obtain access token" >&2
	exit 1
fi

echo "3. Initialize MCP session..."
SESSION=$(curl -s -D - http://localhost:8080/mcp -X POST \
	-H "Authorization: Bearer $ACCESS_TOKEN" \
	-H "Content-Type: application/json" \
	-H "mcp-protocol-version: 2025-06-18" \
	-d '{"jsonrpc":"2.0","method":"initialize","id":1}' | awk '/^mcp-session-id:/{print $2}' | tr -d '\r')
if [ -z "$SESSION" ]; then
	echo "Error: failed to obtain MCP session ID" >&2
	exit 1
fi

echo "4. List tools with session:"
curl -s -X POST http://localhost:8080/mcp \
	-H "Authorization: Bearer $ACCESS_TOKEN" \
	-H "Content-Type: application/json" \
	-H "mcp-protocol-version: 2025-06-18" \
	-H "Mcp-Session-Id: $SESSION" \
	-d '{"jsonrpc":"2.0","method":"tools/list","id":2}'

echo ""
echo "5. Call echo tool:"
curl -s -X POST http://localhost:8080/mcp \
	-H "Authorization: Bearer $ACCESS_TOKEN" \
	-H "Content-Type: application/json" \
	-H "mcp-protocol-version: 2025-06-18" \
	-H "Mcp-Session-Id: $SESSION" \
	-d '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"mcp-backend__echo","arguments":{"text":"hello"}},"id":3}'

echo ""
echo "6. Call sum tool:"
curl -s -X POST http://localhost:8080/mcp \
	-H "Authorization: Bearer $ACCESS_TOKEN" \
	-H "Content-Type: application/json" \
	-H "mcp-protocol-version: 2025-06-18" \
	-H "Mcp-Session-Id: $SESSION" \
	-d '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"mcp-backend__sum","arguments":{"a":42,"b":58}},"id":4}'

echo ""
echo "✓ Real token flow works (session + tool calls)"
