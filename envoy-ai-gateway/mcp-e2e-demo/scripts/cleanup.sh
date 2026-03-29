#!/bin/bash
set -e

echo "Cleaning up..."

if [ -f .pf.pid ]; then
	kill $(cat .pf.pid) 2>/dev/null || true
	rm .pf.pid
fi

kind delete cluster --name mcp-demo 2>/dev/null || true

echo "✓ Cleanup complete"
