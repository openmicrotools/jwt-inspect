# Nodeport Kustomization
This kustomize set will produce a manifest to apply to a Kuberenetes cluster without ingress traffic configured. It can be used to test any stage of deployment to that type of cluster, and will not produce a working endpoint. To access this, you must use it with a port-forward to either the service or pod. Then you can route to localhost:PORT to access the pod.

Use this for local testing (KinD), or anywhere there is not ingress configured. 

For your KinD configuration, please save this to a file and add this as via flag --config=```<file-name>.yaml``` when creating the cluster (```kind create cluster --config <file-name>.yaml```)

```yaml
apiVersion: kind.x-k8s.io/v1alpha4
kind: Cluster
nodes:
- role: control-plane
  extraPortMappings:
  - containerPort: 30443
    hostPort: 30443
    listenAddress: "0.0.0.0"
    protocol: tcp
    - role: worker
```

Access it via localhost:30443 (the node port it is served on)