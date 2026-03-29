# Service Quick Reference Card

## Your Service at a Glance

```yaml
apiVersion: v1
kind: Service
metadata:
  name: mockoidc                    # ← Service name
  namespace: default                # ← Namespace
spec:
  selector:
    app: mockoidc                   # ← Find pods with this label
  ports:
  - name: http
    port: 80         # ← External port (what you call)
    targetPort: 8080 # ← Container port (where it goes)
  - name: http-alt
    port: 8888       # ← Alternative external port
    targetPort: 8080 # ← Same container port
```

---

## Key Concepts (5 Second Version)

| Concept | What It Does |
|---------|--------------|
| **Service** | Stable DNS name + load balancer for pods |
| **selector** | Finds pods by matching labels |
| **port** | Port clients use to connect |
| **targetPort** | Port on the container |
| **ClusterIP** | Service type (internal-only, default) |

---

## The Magic Formula

```
Service Port  →  targetPort  →  Container Port
    80        →     8080     →      8080
   8888       →     8080     →      8080
```

**Both external ports (80, 8888) go to the SAME container port (8080)**

---

## DNS Name

```
mockoidc.default.svc.cluster.local
   ↑       ↑       ↑      ↑
 name  namespace  svc  cluster
```

Short forms (from same namespace):
- `mockoidc`
- `mockoidc.default`
- `mockoidc.default.svc`

---

## Real-World Flow

```
┌─────────┐
│ Client  │
└────┬────┘
     │ curl http://mockoidc:80/oidc
     ▼
┌──────────────────┐
│ Service          │
│ mockoidc         │
│ ClusterIP:       │
│ 10.96.197.10     │
└────┬─────────────┘
     │ selector: app=mockoidc
     ▼
┌──────────────────┐
│ Pod              │
│ mockoidc-xxx     │
│ IP: 10.244.0.89  │
│ Port: 8080       │
└──────────────────┘
```

---

## Common Commands

```bash
# Get service info
kubectl get svc mockoidc

# Get endpoints (pod IPs)
kubectl get endpoints mockoidc

# Describe full details
kubectl describe svc mockoidc

# Test from another pod
kubectl run test --rm -it --image=curlimages/curl -- \
  curl http://mockoidc:80/oidc

# Port forward to localhost
kubectl port-forward svc/mockoidc 8888:8888
# Then: curl http://localhost:8888/oidc
```

---

## Troubleshooting

### No endpoints?
```bash
$ kubectl get endpoints mockoidc
NAME       ENDPOINTS   AGE
mockoidc   <none>      5m
```

**Problem:** Selector doesn't match any pods.

**Fix:** Check labels match:
```bash
# Service selector
kubectl get svc mockoidc -o jsonpath='{.spec.selector}'

# Pod labels  
kubectl get pods -l app=mockoidc --show-labels
```

### Connection refused?

**Problem:** Wrong targetPort.

**Fix:** Check container port:
```bash
kubectl get pod <pod-name> -o jsonpath='{.spec.containers[0].ports[0].containerPort}'
```

Must match service's `targetPort`.

---

## Why Two Ports?

```yaml
ports:
  - port: 80       # Standard HTTP (used by Envoy Gateway)
  - port: 8888     # Alt port (used for direct access/debugging)
```

**Same container, different entry points.**

Think of it like a building with two doors:
- Front door (port 80) - Main entrance
- Side door (port 8888) - Maintenance access

Both get you inside the same building (container:8080).

---

## Quick Test

Verify your service works:

```bash
# Inside cluster
kubectl run test --rm -it --image=busybox -- sh
/ # wget -qO- http://mockoidc:80/oidc/.well-known/openid-configuration

# From your machine (requires port-forward)
kubectl port-forward svc/mockoidc 8888:8888 &
curl http://localhost:8888/oidc/.well-known/openid-configuration
```

---

## The Bottom Line

**Service = Phone Book + Call Router**

- DNS name: `mockoidc.default.svc.cluster.local` (never changes)
- Routes calls to pods (even when pod IPs change)
- Multiple ports supported (80, 8888 → 8080)
- Automatic load balancing (if multiple pods exist)

That's it! 📞
