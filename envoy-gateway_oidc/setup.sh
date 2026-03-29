#!/bin/bash

set -e

echo "========================================"
echo "Envoy Gateway OIDC Setup Script"
echo "========================================"
echo ""

# Check prerequisites
command -v kubectl >/dev/null 2>&1 || {
	echo "Error: kubectl is required but not installed. Aborting." >&2
	exit 1
}
command -v helm >/dev/null 2>&1 || {
	echo "Error: helm is required but not installed. Aborting." >&2
	exit 1
}
command -v docker >/dev/null 2>&1 || {
	echo "Error: docker is required but not installed. Aborting." >&2
	exit 1
}
command -v kind >/dev/null 2>&1 || {
	echo "Error: kind is required but not installed. Aborting." >&2
	exit 1
}
command -v envsubst >/dev/null 2>&1 || {
	echo "Error: envsubst is required but not installed. Install with 'brew install gettext'. Aborting." >&2
	exit 1
}
command -v openssl >/dev/null 2>&1 || {
	echo "Error: openssl is required but not installed. Aborting." >&2
	exit 1
}

echo "✓ Prerequisites check passed"
echo ""

# Step 1: Install Envoy Gateway
echo "Step 1/6: Installing Envoy Gateway..."
if helm list -n envoy-gateway-system | grep -q "^eg"; then
	echo "  → Envoy Gateway already installed, skipping"
else
	ENVOY_GATEWAY_VERSION=${ENVOY_GATEWAY_VERSION:-v1.6.3}
	helm install eg oci://docker.io/envoyproxy/gateway-helm \
		--version "$ENVOY_GATEWAY_VERSION" \
		--namespace envoy-gateway-system \
		--create-namespace
	echo "  → Waiting for Envoy Gateway to be ready..."
	kubectl wait --timeout=5m -n envoy-gateway-system \
		deployment/envoy-gateway --for=condition=Available
fi
echo "✓ Envoy Gateway installed"
echo ""

# Step 2: Build and Deploy Mock OIDC Server and Echo Backend
echo "Step 2/6: Building and deploying Mock OIDC Server and Echo Backend..."
echo "  → Building mockoidc Docker image..."
docker build -t mockoidc:latest . >/dev/null 2>&1
echo "  → Loading image into KIND cluster..."
kind load docker-image mockoidc:latest >/dev/null 2>&1
echo "  → Deploying to Kubernetes..."
kubectl apply -f manifests.yaml
echo "  → Waiting for deployments to be ready..."
kubectl rollout status deployment/mockoidc -n default --timeout=5m
kubectl rollout status deployment/echo -n default --timeout=5m
sleep 5 # Wait for logs to be available
echo "✓ Mock OIDC Server and Echo Backend deployed"
echo ""

# Step 3: Get Mock OIDC Credentials
echo "Step 3/6: Retrieving Mock OIDC credentials..."
POD_NAME=$(kubectl get pods -n default -l app=mockoidc --field-selector=status.phase=Running -o jsonpath='{.items[0].metadata.name}')
LOGS=$(kubectl logs -n default "$POD_NAME" 2>/dev/null || echo "")

if [ -z "$LOGS" ]; then
	echo "Error: Could not retrieve logs from mock OIDC server"
	exit 1
fi

# Extract credentials
export ISSUER=$(echo "$LOGS" | grep "^export ISSUER=" | cut -d'"' -f2)
export CLIENT_ID=$(echo "$LOGS" | grep "^export CLIENT_ID=" | cut -d'"' -f2)
export CLIENT_SECRET_BASE64_ENCODED=$(echo "$LOGS" | grep "^export CLIENT_SECRET_BASE64_ENCODED=" | cut -d'"' -f2)

if [ -z "$ISSUER" ] || [ -z "$CLIENT_ID" ] || [ -z "$CLIENT_SECRET_BASE64_ENCODED" ]; then
	echo "Error: Could not parse credentials"
	echo "Logs:"
	echo "$LOGS"
	exit 1
fi

echo "Credentials extracted:"
echo "  ISSUER: $ISSUER"
echo "  CLIENT_ID: $CLIENT_ID"
echo "  CLIENT_SECRET_BASE64: ${CLIENT_SECRET_BASE64_ENCODED:0:20}..."
echo "✓ Credentials retrieved"
echo ""

# Step 4: Create TLS secret and apply OIDC Configuration
echo "Step 4/6: Creating TLS cert and applying OIDC SecurityPolicy..."
TMP_CERT_DIR=$(mktemp -d)
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
	-subj "/CN=localhost" \
	-addext "subjectAltName=DNS:localhost" \
	-keyout "$TMP_CERT_DIR/tls.key" \
	-out "$TMP_CERT_DIR/tls.crt" >/dev/null 2>&1
kubectl -n default create secret tls eg-tls \
	--key "$TMP_CERT_DIR/tls.key" \
	--cert "$TMP_CERT_DIR/tls.crt" \
	--dry-run=client -o yaml | kubectl apply -f -
rm -rf "$TMP_CERT_DIR"
envsubst <oidc.yaml | kubectl apply -f -
sleep 3 # Wait for policy to be processed
echo "✓ OIDC SecurityPolicy applied"
echo ""

# Step 5: Create HTTPRoutes
echo "Step 5/6: Creating HTTPRoutes..."
kubectl apply -f - <<EOF
---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: myapp
  namespace: default
spec:
  parentRefs:
  - name: eg
  hostnames:
  - localhost
  rules:
  - matches:
    - path:
        type: PathPrefix
        value: /myapp
    filters:
    - type: URLRewrite
      urlRewrite:
        path:
          type: ReplacePrefixMatch
          replacePrefixMatch: /
    backendRefs:
    - name: echo
      port: 80
---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: mockoidc-route
  namespace: default
spec:
  parentRefs:
  - name: eg
  hostnames:
  - localhost
  rules:
  - matches:
    - path:
        type: PathPrefix
        value: /oidc
    backendRefs:
    - name: mockoidc
      port: 80
EOF
echo "✓ HTTPRoutes created"
echo ""

# Step 6: Add /etc/hosts entry
echo "Step 6/6: Adding /etc/hosts entry..."
if grep -q "mockoidc.default.svc.cluster.local" /etc/hosts; then
	echo "  → Entry already exists in /etc/hosts, skipping"
else
	echo "  → Adding entry (requires sudo)..."
	echo "127.0.0.1 mockoidc.default.svc.cluster.local" | sudo tee -a /etc/hosts >/dev/null
fi
echo "✓ /etc/hosts entry added"
echo ""

# Verify SecurityPolicy
echo "Verifying SecurityPolicy status..."
POLICY_STATUS=$(kubectl get securitypolicy oidc-example -o jsonpath='{.status.ancestors[0].conditions[?(@.type=="Accepted")].status}')

if [ "$POLICY_STATUS" = "True" ]; then
	echo "✓ SecurityPolicy accepted!"
else
	echo "⚠ SecurityPolicy status: $POLICY_STATUS"
	echo "Check with: kubectl get securitypolicy oidc-example -o yaml"
fi
echo ""

echo "========================================"
echo "Setup Complete!"
echo "========================================"
echo ""
echo "Next steps:"
echo ""
echo "1. Start port-forwards in two separate terminals:"
echo ""
echo "   Terminal 1:"
echo "   $ kubectl -n envoy-gateway-system port-forward service/envoy-default-eg-e41e7b31 8443:443"
echo ""
echo "   Terminal 2:"
echo "   $ kubectl -n default port-forward service/mockoidc 8888:8888"
echo ""
echo "2. Test with curl:"
echo "   $ curl -vk https://localhost:8443/myapp 2>&1 | grep -i location"
echo ""
echo "3. Test with Firefox:"
echo "   Navigate to: https://localhost:8443/myapp"
echo "   (You'll need to accept the self-signed certificate warning)"
echo ""
echo "4. To cleanup:"
echo "   $ ./cleanup.sh"
echo ""
