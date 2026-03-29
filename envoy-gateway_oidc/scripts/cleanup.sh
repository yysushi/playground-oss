#!/bin/bash

set -e

echo "========================================"
echo "Envoy Gateway OIDC Cleanup Script"
echo "========================================"
echo ""

read -p "This will delete all resources. Continue? (y/N): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
	echo "Cleanup cancelled"
	exit 0
fi

echo "Cleaning up resources..."
echo ""

# Delete HTTPRoutes
echo "1. Deleting HTTPRoutes..."
kubectl delete httproute myapp mockoidc-route -n default --ignore-not-found=true
echo "✓ HTTPRoutes deleted"
echo ""

# Delete SecurityPolicy
echo "2. Deleting SecurityPolicy..."
kubectl delete securitypolicy oidc-example --ignore-not-found=true
echo "✓ SecurityPolicy deleted"
echo ""

# Delete Secret
echo "3. Deleting Secret..."
kubectl delete secret mockoidc-secret -n default --ignore-not-found=true
echo "✓ Secret deleted"
echo ""

# Delete Mock OIDC
echo "4. Deleting Mock OIDC Server..."
kubectl delete -f mockoidc-simple.yaml --ignore-not-found=true
echo "✓ Mock OIDC Server deleted"
echo ""

# Delete Gateway
echo "5. Deleting Gateway..."
kubectl delete gateway eg -n default --ignore-not-found=true
echo "✓ Gateway deleted"
echo ""

# Delete GatewayClass
echo "6. Deleting GatewayClass..."
kubectl delete gatewayclass eg --ignore-not-found=true
echo "✓ GatewayClass deleted"
echo ""

# Remove /etc/hosts entry
echo "7. Removing /etc/hosts entry..."
if grep -q "mockoidc.default.svc.cluster.local" /etc/hosts; then
	echo "  → Removing entry (requires sudo)..."
	sudo sed -i.bak '/mockoidc.default.svc.cluster.local/d' /etc/hosts
	echo "✓ /etc/hosts entry removed"
else
	echo "  → No entry found in /etc/hosts, skipping"
fi
echo ""

# Optional: Uninstall Envoy Gateway
read -p "Do you want to uninstall Envoy Gateway? (y/N): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
	echo "8. Uninstalling Envoy Gateway..."
	helm uninstall eg -n envoy-gateway-system --ignore-not-found=true
	kubectl delete namespace envoy-gateway-system --ignore-not-found=true
	echo "✓ Envoy Gateway uninstalled"
	echo ""
fi

echo "========================================"
echo "Cleanup Complete!"
echo "========================================"
echo ""
echo "Note: Make sure to stop any running port-forwards (Ctrl+C)"
echo ""
