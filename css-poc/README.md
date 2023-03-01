## Demo-ing the configmap css page

This POC is to demo the css customization via configmap.

Requirements:
KinD
Docker

How to run it:
    1. Build the container with ```make docker-build-css-poc```
    1. Using kind, stand up a cluster locally with ```kind create cluster```, and load the image manually, or using make ```load-docker-image```
    1. Apply the files using ```kustomize build | apply -f -``` or ```make push-css-poc```
    1. Port-Forward to the pod (hosted on localhost 8080), manually or use ```make port-forward-to-pod```
    1. Check localhost:8080 to see current color state
    1. Make a change to ```assets/customize.css```, suggestion: under .navbar change ```background-color: ``` to any color not in index.css for ```background-color```. You can use hex or reference a common color name(example: "red", "blue", "yellow")
    1. Push updated changes with ```kustomize build | apply -f -``` or ```make push-css-poc``` to update the configmap. Allow the pod to come back up.
    1. Port-Forward to the pod (hosted on localhost 8080), manually or use ```make port-forward-to-pod```
    1. Check localhost:8080 to see current color state. If there is no change, do a forced refresh (shift+f5). State should match new change in the config map.
