# Using Overlays to make Consistent Changes

You can customize the namespace, labels, image tag, and a number of other settings by simply configuring the kustomization.yaml in this folder. You will need to update the configuration per overlay you wish to amend.

To change the image tag name, please modify it in the ../base/deployment.yaml file first, and then update the kustomization.yaml files accordingly. Only the image tag can be quickly changed out with the current configuration of the kustomization.yaml files.

To build and deploy an overlay:
    1. From project root
        - ```make deploy-base``` applies the base directory contents. This is suitable for local testing with port-forwarding.
        - ```make deploy-nodeport``` applies the nodeport directory contents. This includes a service patch to add a nodeport port configuration to the service. 
        - ```make deploy-ingress``` applies the ingress directory contents. This includes an ingress resource for the project, and allows it to be accessed with ingress using a selected hostname.
    2. From this (overlays) folder:
        - For the selected overlay run ```kustomize build <base/nodeport/ingress> | kubectl apply -f -```. 