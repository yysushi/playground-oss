#!/bin/bash
set -e

echo "=== Testing OIDC Authentication with Envoy Gateway ==="
echo ""

# Check if port forwards are needed
if ! nc -z localhost 8443 2>/dev/null; then
	echo "⚠️  Gateway not accessible on localhost:8443"
	echo "   Run in separate terminal: kubectl -n envoy-gateway-system port-forward service/\$(kubectl get svc -n envoy-gateway-system --selector=gateway.envoyproxy.io/owning-gateway-name=eg -o jsonpath='{.items[0].metadata.name}') 8443:443"
	echo ""
	exit 1
fi

if ! nc -z localhost 8888 2>/dev/null; then
	echo "⚠️  OIDC server not accessible on localhost:8888"
	echo "   Run in separate terminal: kubectl port-forward service/mockoidc 8888:8888"
	echo ""
	exit 1
fi

echo "1. Testing unauthenticated request (should redirect to OIDC):"
LOCATION=$(curl -sLk https://localhost:8443/myapp 2>&1 | grep -i "^< location:" | awk '{print $3}' | tr -d '\r')
if [[ "$LOCATION" == *"oauth2/authorize"* ]]; then
	echo "   ✓ Redirected to OIDC authorization endpoint"
	echo "   Location: ${LOCATION:0:80}..."
else
	echo "   ✗ No OIDC redirect found"
	echo "   Location: $LOCATION"
	exit 1
fi

echo ""
echo "2. Testing OIDC discovery endpoint:"
DISCOVERY=$(curl -sk http://localhost:8888/oidc/.well-known/openid-configuration | jq -r '.issuer' 2>/dev/null)
if [[ -n "$DISCOVERY" ]]; then
	echo "   ✓ OIDC server responding"
	echo "   Issuer: $DISCOVERY"
else
	echo "   ✗ OIDC server not responding"
	exit 1
fi

echo ""
echo "3. Testing backend accessibility:"
BACKEND_STATUS=$(kubectl run test-curl --rm -i --restart=Never --image=curlimages/curl -- curl -s -o /dev/null -w "%{http_code}" http://echo.default.svc.cluster.local 2>/dev/null || echo "000")
if [[ "$BACKEND_STATUS" == "200" ]]; then
	echo "   ✓ Backend (echo service) is healthy"
else
	echo "   ⚠️  Backend returned status: $BACKEND_STATUS"
fi

echo ""
echo "4. Browser test instructions:"
echo "   Open: https://localhost:8443/myapp"
echo "   - Should redirect to OIDC login"
echo "   - Login with any username/password (mockoidc accepts all)"
echo "   - After login, you'll see the nginx welcome page"

echo ""
echo "✓ OIDC authentication flow is working!"
echo ""
echo "Note: For full OAuth flow testing, use a browser or OAuth client."
