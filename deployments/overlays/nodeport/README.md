# Nodeport Kustomization
This kustomize set will produce a manifest to apply to a Kuberenetes cluster without ingress traffic configured. It can be used to test any stage of deployment to that type of cluster, and will not produce a working endpoint. 

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

You can also use port-forward to access the pod or service if you do not want to set up KinD or other local k8s testing service with a specialized configuration.