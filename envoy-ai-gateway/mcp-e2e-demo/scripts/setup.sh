#!/bin/bash
set -e

echo "=== MCP OAuth E2E Demo Setup ==="

# 1. Create KIND cluster
echo "1. Creating KIND cluster..."
if ! kind get clusters | grep -q mcp-demo; then
	kind create cluster --name mcp-demo
fi

# 2. Install Envoy Gateway
echo "2. Installing Envoy Gateway..."
ENVOY_GATEWAY_VERSION=${ENVOY_GATEWAY_VERSION:-v1.6.3}
helm upgrade --install eg oci://docker.io/envoyproxy/gateway-helm \
	--version "$ENVOY_GATEWAY_VERSION" \
	--namespace envoy-gateway-system \
	--create-namespace \
	-f k8s/envoy-gateway-values.yaml \
	--wait

# 2a. Install AI Gateway CRDs
echo "2a. Installing AI Gateway CRDs..."
helm upgrade --install aieg-crd oci://docker.io/envoyproxy/ai-gateway-crds-helm \
	--version v0.0.0-latest \
	--namespace envoy-ai-gateway-system \
	--create-namespace

# 2b. Install AI Gateway
echo "2b. Installing AI Gateway..."
helm upgrade --install aieg oci://docker.io/envoyproxy/ai-gateway-helm \
	--version v0.0.0-latest \
	--namespace envoy-ai-gateway-system \
	--create-namespace
kubectl wait --timeout=5m -n envoy-ai-gateway-system deployment/ai-gateway-controller --for=condition=Available

# 3. Build images
echo "3. Building Docker images..."
docker build -t mcp:local mcp-server/
docker build -t oauth:local oauth-server/
kind load docker-image mcp:local --name mcp-demo
kind load docker-image oauth:local --name mcp-demo

# 4. Deploy base services
echo "4. Deploying base services..."
kubectl apply -f k8s/all.yaml
kubectl wait --for=condition=available --timeout=300s deployment/mcp deployment/oauth

# 6. Create JWKS ConfigMap for MCP OAuth
echo "6. Creating JWKS ConfigMap..."
kubectl port-forward service/oauth 8081:8080 >/tmp/pf-oauth.log 2>&1 &
OAUTH_PF_PID=$!
sleep 2
JWKS=$(curl -fsS http://localhost:8081/oidc/.well-known/jwks.json || true)
kill $OAUTH_PF_PID >/dev/null 2>&1 || true
if [ -z "$JWKS" ]; then
	echo "Error: Failed to fetch JWKS from oauth service" >&2
	cat /tmp/pf-oauth.log >&2 || true
	exit 1
fi
kubectl create configmap jwks-configmap \
	--from-literal=jwks="$JWKS" \
	--dry-run=client -o yaml | kubectl apply -f -

# 6. Deploy MCPRoute with OAuth
echo "6. Deploying MCPRoute with OAuth..."
kubectl apply -f k8s/mcp-route.yaml

# 7. Get OAuth credentials
echo "7. Getting OAuth credentials..."
sleep 3
kubectl logs deployment/oauth | grep -E "CLIENT_ID|CLIENT_SECRET|ISSUER" >.env
cat .env

# 8. Wait for gateway
echo "8. Waiting for gateway..."
sleep 10 # Wait for gateway service to be created

# 9. Port forward
echo "9. Setting up port forward..."
GATEWAY_SVC=$(kubectl get svc -n envoy-gateway-system --selector=gateway.envoyproxy.io/owning-gateway-name=aigw-run -o jsonpath='{.items[0].metadata.name}')
echo "Gateway service: $GATEWAY_SVC"
kubectl port-forward -n envoy-gateway-system service/$GATEWAY_SVC 8080:8080 >/dev/null 2>&1 &
PF_PID=$!
echo $PF_PID >.pf.pid
sleep 2

echo ""
echo "✓ Setup complete!"
echo ""
echo "Gateway: http://localhost:8080/mcp"
echo "OAuth credentials saved to .env"
echo ""
echo "Next steps:"
echo "  ./scripts/test-working.sh    # Test MCP tools"
echo "  ./scripts/cleanup.sh          # Cleanup everything"
echo ""
echo "Note: Port forward is running in background (PID in .pf.pid)"
