#! /bin/bash

podname=$(kubectl get pods -A | grep jwt-inspect | awk '{print $2}')
kubectl port-forward "${podname}" -n open-microtools 8080:8080
