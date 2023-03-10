# Ingress Kustomization
This kustomize set will produce a manifest to apply to a Kuberenetes cluster with ingress traffic configured. It can be used to test any stage of deployment to that type of cluster, and will produce a working endpoint (you must replace ```<HOSTNAME>``` with your own hostname). 

Access it via <HOSTNAME>, as configured in the ingress.yaml file.