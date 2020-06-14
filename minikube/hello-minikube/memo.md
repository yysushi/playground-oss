# Minikube

<https://kubernetes.io/ja/docs/tutorials/hello-minikube/>

a sample app on Kubernetes using Minikube

- deploy an app
- run the app
- view the app logs

## Steps

### Create a Minikube cluster

`minikube status`

`minikube start`

`minikube stop`

`minikube dashboard`

### Create a Deployment

`kubectl get deployments`

`kubectl get pods`

`kubectl get events`

`kubectl config view`

[pod](https://kubernetes.io/ja/docs/concepts/workloads/pods/pod/)

[deployment](https://kubernetes.io/ja/docs/concepts/workloads/controllers/deployment/)

[kubectl](https://kubernetes.io/docs/reference/kubectl/overview://kubernetes.io/docs/reference/kubectl/overview/)

### Create a Service

`kubectl expose deployment hello-node --type=LoadBalancer --port=8080`

`kubectl get services`

`minikube service hello-node`

[service](https://kubernetes.io/ja/docs/concepts/services-networking/service/)

### Enable addons

`minikube addons list`

### Clean up

`kubectl delete service hello-node`

`kubectl delete deployment hello-node`

`minikube stop`

`minikube delete`

## Log

```shell-session
koketani:minikube (minikube %)$ minikube status
host: Stopped
kubelet:
apiserver:
kubeconfig:

koketani:minikube (minikube %)$ minikube start
🎉  minikube 1.11.0 is available! Download it: https://github.com/kubernetes/minikube/releases/tag/v1.11.0
💡  To disable this notice, run: 'minikube config set WantUpdateNotification false'

🙄  minikube v1.6.2 on Darwin 10.14.6
✨  Selecting 'hyperkit' driver from existing profile (alternates: [virtualbox])
💡  Tip: Use 'minikube start -p <name>' to create a new cluster, or 'minikube delete' to delete this one.
🔄  Starting existing hyperkit VM for "minikube" ...
⌛  Waiting for the host to be provisioned ...
🐳  Preparing Kubernetes v1.17.0 on Docker '19.03.5' ...
🚀  Launching Kubernetes ...
🏄  Done! kubectl is now configured to use "minikube"

koketani:minikube (minikube %)$ minikube status
host: Running
kubelet: Running
apiserver: Running
kubeconfig: Configured

koketani:minikube (minikube %)$ minikube dashboard
🤔  Verifying dashboard health ...
🚀  Launching proxy ...
🤔  Verifying proxy health ...
🎉  Opening http://127.0.0.1:50708/api/v1/namespaces/kubernetes-dashboard/services/http:kubernetes-dashboard:/proxy/ in your default browser...
```

```shell-session

koketani:minikube (minikube %)$ kubectl get deployments
No resources found in default namespace.

koketani:minikube (minikube %)$ kubectl create deployment hello-node --image=k8s.gcr.io/echoserver:1.4
deployment.apps/hello-node created

koketani:minikube (minikube %)$ kubectl get deployments
NAME         READY   UP-TO-DATE   AVAILABLE   AGE
hello-node   0/1     1            0           3s

koketani:minikube (minikube %)$ kubectl get deployments
NAME         READY   UP-TO-DATE   AVAILABLE   AGE
hello-node   1/1     1            1           12s

koketani:minikube (minikube %)$ kubectl get pods
NAME                          READY   STATUS    RESTARTS   AGE
hello-node-7dc7987866-fm5m6   1/1     Running   0          25s

koketani:minikube (minikube %)$ kubectl get events
LAST SEEN   TYPE     REASON                    OBJECT                             MESSAGE
<unknown>   Normal   Scheduled                 pod/hello-node-7dc7987866-fm5m6    Successfully assigned default/hello-node-7dc7987866-fm5m6 to minikube
51s         Normal   Pulling                   pod/hello-node-7dc7987866-fm5m6    Pulling image "k8s.gcr.io/echoserver:1.4"
44s         Normal   Pulled                    pod/hello-node-7dc7987866-fm5m6    Successfully pulled image "k8s.gcr.io/echoserver:1.4"
44s         Normal   Created                   pod/hello-node-7dc7987866-fm5m6    Created container echoserver
43s         Normal   Started                   pod/hello-node-7dc7987866-fm5m6    Started container echoserver
51s         Normal   SuccessfulCreate          replicaset/hello-node-7dc7987866   Created pod: hello-node-7dc7987866-fm5m6
51s         Normal   ScalingReplicaSet         deployment/hello-node              Scaled up replica set hello-node-7dc7987866 to 1
27m         Normal   Starting                  node/minikube                      Starting kubelet.
27m         Normal   NodeHasSufficientMemory   node/minikube                      Node minikube status is now: NodeHasSufficientMemory
27m         Normal   NodeHasNoDiskPressure     node/minikube                      Node minikube status is now: NodeHasNoDiskPressure
27m         Normal   NodeHasSufficientPID      node/minikube                      Node minikube status is now: NodeHasSufficientPID
27m         Normal   NodeAllocatableEnforced   node/minikube                      Updated Node Allocatable limit across pods
27m         Normal   Starting                  node/minikube                      Starting kube-proxy.
27m         Normal   RegisteredNode            node/minikube                      Node minikube event: Registered Node minikube in Controller
15m         Normal   Starting                  node/minikube                      Starting kubelet.
15m         Normal   NodeHasSufficientMemory   node/minikube                      Node minikube status is now: NodeHasSufficientMemory
15m         Normal   NodeHasNoDiskPressure     node/minikube                      Node minikube status is now: NodeHasNoDiskPressure
15m         Normal   NodeHasSufficientPID      node/minikube                      Node minikube status is now: NodeHasSufficientPID
15m         Normal   NodeAllocatableEnforced   node/minikube                      Updated Node Allocatable limit across pods
15m         Normal   Starting                  node/minikube                      Starting kube-proxy.
15m         Normal   RegisteredNode            node/minikube                      Node minikube event: Registered Node minikube in Controller

koketani:minikube (minikube %)$

koketani:minikube (minikube %)$ kubectl config view
apiVersion: v1
clusters:
- cluster:
    certificate-authority: /Users/koketani/.minikube/ca.crt
    server: https://192.168.64.3:8443
  name: minikube
contexts:
- context:
    cluster: minikube
    user: minikube
  name: minikube
current-context: minikube
kind: Config
preferences: {}
users:
- name: minikube
  user:
    client-certificate: /Users/koketani/.minikube/client.crt
    client-key: /Users/koketani/.minikube/client.key
```

```shell-session
koketani:minikube (minikube %)$ kubectl get services
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   6d5h

koketani:minikube (minikube %)$ kubectl expose deployment hello-node --type=LoadBalancer --port=8080
service/hello-node exposed

koketani:minikube (minikube %)$ kubectl get services
NAME         TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
hello-node   LoadBalancer   10.96.237.36   <pending>     8080:31104/TCP   13s
kubernetes   ClusterIP      10.96.0.1      <none>        443/TCP          6d5h

koketani:minikube (minikube %)$ minikube service hello-node
|-----------|------------|-------------|---------------------------|
| NAMESPACE |    NAME    | TARGET PORT |            URL            |
|-----------|------------|-------------|---------------------------|
| default   | hello-node |             | http://192.168.64.3:31104 |
|-----------|------------|-------------|---------------------------|
🎉  Opening service default/hello-node in default browser...

koketani:minikube (minikube %)$ curl -i http://192.168.64.3:31104
HTTP/1.1 200 OK
Server: nginx/1.10.0
Date: Sun, 14 Jun 2020 04:21:34 GMT
Content-Type: text/plain
Transfer-Encoding: chunked
Connection: keep-alive

CLIENT VALUES:
client_address=172.17.0.1
command=GET
real path=/
query=nil
request_version=1.1
request_uri=http://192.168.64.3:8080/

SERVER VALUES:
server_version=nginx: 1.10.0 - lua: 10001

HEADERS RECEIVED:
accept=*/*
host=192.168.64.3:31104
user-agent=curl/7.54.0
BODY:
-no body in request-

koketani:minikube (minikube %)$ minikube service list
|----------------------|---------------------------|---------------------------|-----|
|      NAMESPACE       |           NAME            |        TARGET PORT        | URL |
|----------------------|---------------------------|---------------------------|-----|
| default              | hello-node                | http://192.168.64.3:31104 |
| default              | kubernetes                | No node port              |
| kube-system          | kube-dns                  | No node port              |
| kubernetes-dashboard | dashboard-metrics-scraper | No node port              |
| kubernetes-dashboard | kubernetes-dashboard      | No node port              |
|----------------------|---------------------------|---------------------------|-----|
```

```shell-sesison

koketani:hello-minikube (minikube %)$ minikube addons list
- addon-manager: enabled
- dashboard: enabled
- default-storageclass: enabled
- efk: disabled
- freshpod: disabled
- gvisor: disabled
- helm-tiller: disabled
- ingress: disabled
- ingress-dns: disabled
- logviewer: disabled
- metrics-server: disabled
- nvidia-driver-installer: disabled
- nvidia-gpu-device-plugin: disabled
- registry: disabled
- registry-creds: disabled
- storage-provisioner: enabled
- storage-provisioner-gluster: disabled
koketani:hello-minikube (minikube %)$
koketani:hello-minikube (minikube %)$ minikube addons enable heapster

💣  enable failed: property name "heapster" not found

😿  minikube is exiting due to an error. If the above message is not useful, open an issue:
👉  https://github.com/kubernetes/minikube/issues/new/choose
```

```shell-session

koketani:minikube (minikube %)$ kubectl delete service hello-node
service "hello-node" deleted
koketani:minikube (minikube %)$ kubectl get services
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   6d6h
koketani:minikube (minikube %)$ kubectl get deployments
NAME         READY   UP-TO-DATE   AVAILABLE   AGE
hello-node   1/1     1            1           67m
koketani:minikube (minikube %)$ kubectl delete deployment hello-node
deployment.apps "hello-node" deleted
koketani:minikube (minikube %)$ kubectl get deployments
No resources found in default namespace.
koketani:minikube (minikube %)$ minikube stop
✋  Stopping "minikube" in hyperkit ...
🛑  "minikube" stopped.
koketani:minikube (minikube %)$ minikube status
host: Stopped
kubelet:
apiserver:
kubeconfig:
koketani:minikube (minikube %)$ minikube delete
🔥  Deleting "minikube" in hyperkit ...
💔  The "minikube" cluster has been deleted.
🔥  Successfully deleted profile "minikube"
koketani:minikube (minikube %)$ minikube status
host:
kubelet:
apiserver:
kubeconfig:
```
