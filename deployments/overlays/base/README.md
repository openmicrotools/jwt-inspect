# Nodeport Kustomization
This kustomize set will produce a manifest to apply to a Kuberenetes cluster without ingress traffic configured. It can be used to test any stage of deployment to that type of cluster, and will not produce a working endpoint. To access this, you must use it with a port-forward to either the service or pod. Then you can route to localhost:PORT to access the pod.

Use this for port-forward only testing, it should not be used for production.
