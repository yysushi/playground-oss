# Kubernetes Service Explained

## Your Service Definition

```yaml
apiVersion: v1
kind: Service
metadata:
  name: mockoidc
  namespace: default
spec:
  selector:
    app: mockoidc
  ports:
  - name: http
    port: 80
    targetPort: 8080
  - name: http-alt
    port: 8888
    targetPort: 8888
```

---

## What is a Service?

A **Service** is like a **load balancer** + **DNS name** for your pods.

### The Problem It Solves

Pods have dynamic IPs that change when they restart:
```
Pod mockoidc-abc123 → IP: 10.244.0.5  (dies)
Pod mockoidc-xyz789 → IP: 10.244.0.9  (new pod, new IP!)
```

How do other apps find your pod if the IP keeps changing?

**Answer:** Use a Service! It gives you a stable DNS name and IP.

---

## How Your Service Works

### 1. **Selector** - Find the Pods

```yaml
selector:
  app: mockoidc
```

This says: "Find all pods with the label `app: mockoidc`"

The Service looks for pods like this:
```yaml
# In your Deployment
template:
  metadata:
    labels:
      app: mockoidc  # ← Service finds this!
```

**Result:** Service automatically tracks all matching pods, even as they come and go.

---

### 2. **Ports** - The Traffic Routing

You have **two ports** defined:

#### Port 1: http
```yaml
- name: http
  port: 80
  targetPort: 8080
```

**What this means:**
```
External Request → Service:80 → Pod:8080
```

- `port: 80` - "When someone calls me on port 80..."
- `targetPort: 8080` - "...forward to pod's port 8080"

#### Port 2: http-alt
```yaml
- name: http-alt
  port: 8888
  targetPort: 8080
```

**What this means:**
```
External Request → Service:8888 → Pod:8080
```

- `port: 8888` - "When someone calls me on port 8888..."
- `targetPort: 8080` - "...forward to pod's port 8080"

---

## Visual Example

```
┌─────────────────────────────────────┐
│  Other Pods / Services              │
└──────────┬──────────────┬───────────┘
           │              │
           │              │
    Call port 80    Call port 8888
           │              │
           ▼              ▼
┌──────────────────────────────────────┐
│  Service: mockoidc                   │
│  DNS: mockoidc.default.svc.cluster.local
│  ClusterIP: 10.96.10.81              │
│                                      │
│  Port 80 → targetPort 8080           │
│  Port 8888 → targetPort 8080         │
└──────────┬───────────────────────────┘
           │
           │ selector: app=mockoidc
           │
           ▼
┌──────────────────────────────────────┐
│  Pod: mockoidc-6958f548d4-m4nzk     │
│  Labels: app=mockoidc                │
│  IP: 10.244.0.83 (changes!)          │
│                                      │
│  Container listening on port 8080    │
└──────────────────────────────────────┘
```

---

## Why Two Ports to Same Target?

```yaml
ports:
  - port: 80 → targetPort: 8080      # Standard HTTP
  - port: 8888 → targetPort: 8080    # Alternative access
```

Both ports go to the **same container port (8080)**.

### Use Cases:

1. **Different access methods**
   - Port 80: Used by Envoy Gateway (standard HTTP)
   - Port 8888: Direct access for port-forwarding

2. **Migration / Compatibility**
   - Old clients use port 8888
   - New clients use port 80
   - Both work during transition

3. **In your setup:**
   ```bash
   # Access via port 80 (through Envoy)
   curl http://mockoidc.default.svc.cluster.local:80/oidc
   
   # Access via port 8888 (direct, for testing)
   kubectl port-forward svc/mockoidc 8888:8888
   curl http://localhost:8888/oidc
   ```

---

## Service DNS Name

The Service automatically gets a DNS name:

```
mockoidc.default.svc.cluster.local
```

Format: `{service-name}.{namespace}.svc.cluster.local`

**Any pod in your cluster can use this name:**

```bash
# From any pod:
curl http://mockoidc.default.svc.cluster.local:80/oidc
curl http://mockoidc.default.svc.cluster.local:8888/oidc

# Short form (same namespace):
curl http://mockoidc:80/oidc
curl http://mockoidc:8888/oidc
```

---

## How the Selector Works

### Step-by-step:

1. **Service created with selector:**
   ```yaml
   selector:
     app: mockoidc
   ```

2. **Kubernetes looks for matching pods:**
   ```bash
   $ kubectl get pods -l app=mockoidc
   NAME                        READY   STATUS    RESTARTS   AGE
   mockoidc-6958f548d4-m4nzk   1/1     Running   0          5m
   ```

3. **Service creates Endpoints:**
   ```bash
   $ kubectl get endpoints mockoidc
   NAME       ENDPOINTS                          AGE
   mockoidc   10.244.0.83:8080                   5m
   ```

4. **Traffic flows:**
   ```
   Request to mockoidc:80
     → Service forwards to 10.244.0.83:8080
     → Pod receives request
   ```

---

## Testing Your Service

### 1. Check the Service
```bash
$ kubectl get svc mockoidc
NAME       TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)           AGE
mockoidc   ClusterIP   10.96.10.81    <none>        80/TCP,8888/TCP   5m
```

### 2. Check Endpoints (Pod IPs)
```bash
$ kubectl get endpoints mockoidc
NAME       ENDPOINTS         AGE
mockoidc   10.244.0.83:8080  5m
```

If endpoints are empty, the selector doesn't match any pods!

### 3. Test from Another Pod
```bash
$ kubectl run test --rm -it --image=curlimages/curl -- sh
/ $ curl http://mockoidc:80/oidc/.well-known/openid-configuration
{
  "issuer": "http://mockoidc.default.svc.cluster.local:8888/oidc",
  ...
}
```

### 4. Port Forward to Your Machine
```bash
# Forward service port 8888 to localhost:8888
$ kubectl port-forward svc/mockoidc 8888:8888

# In another terminal:
$ curl http://localhost:8888/oidc/.well-known/openid-configuration
```

---

## Common Mistakes

### ❌ Wrong: Selector doesn't match

```yaml
# Service
selector:
  app: mockoidc

# Deployment
labels:
  app: mock-oidc  # Typo! Won't match
```

**Fix:** Labels must match exactly.

### ❌ Wrong: Wrong targetPort

```yaml
# Service
targetPort: 80

# Container
ports:
  - containerPort: 8080  # Container listening on 8080, not 80!
```

**Fix:** targetPort must match containerPort.

### ❌ Wrong: Forgot port name

```yaml
ports:
  - port: 80        # No name
  - port: 8888      # No name
```

**Best practice:** Always name your ports for clarity.

---

## Service Types

Your service doesn't specify a type, so it defaults to `ClusterIP`:

```yaml
spec:
  type: ClusterIP  # Default (only accessible within cluster)
```

### Other types:

```yaml
# NodePort - Expose on each node's IP
type: NodePort
ports:
  - port: 80
    targetPort: 8080
    nodePort: 30080  # Access via <NodeIP>:30080

# LoadBalancer - Cloud provider LB (AWS ELB, GCP LB, etc.)
type: LoadBalancer
# Gets external IP automatically

# ExternalName - DNS alias (no selector)
type: ExternalName
externalName: my-external-service.example.com
```

For your POC, `ClusterIP` is perfect - keeps it simple and internal.

---

## Key Takeaways

1. **Service = Stable DNS + IP** for your pods
2. **Selector** finds pods by labels
3. **port** = External-facing port (what clients call)
4. **targetPort** = Container port (where traffic goes)
5. **Multiple ports** = Multiple ways to access same container
6. **DNS name** = `{service}.{namespace}.svc.cluster.local`
7. **ClusterIP** = Internal only (default, used here)

---

## Your Service in Action

```yaml
apiVersion: v1
kind: Service
metadata:
  name: mockoidc              # DNS: mockoidc.default.svc.cluster.local
  namespace: default
spec:
  selector:
    app: mockoidc             # Find pods with this label
  ports:
  - name: http
    port: 80                  # Service port
    targetPort: 8080          # Container port
  - name: http-alt
    port: 8888                # Alternative service port
    targetPort: 8080          # Same container port
```

**Result:**
- Stable name: `mockoidc.default.svc.cluster.local`
- Two ways to access: port 80 or 8888
- Both go to pod's port 8080
- Works even if pod IP changes

Simple, elegant, powerful. That's Kubernetes Services! 🎯
