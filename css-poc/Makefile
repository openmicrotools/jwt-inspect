
.PHONY: docker-build-css-poc
docker-build-css-poc: 
	docker build -t css.poc:v1 .

.PHONY: push-css-poc
push-css-poc:
	kustomize build | kubectl apply -f -  

ifndef port-forward-to-poc
  
endif

.PHONY: port-forward-to-poc
port-forward-to-poc:
	./hack/scripts/port-forward-to-poc.sh

.PHONY: load-docker-image
load-docker-image:
	kind load docker-image css.poc:v1
