IMAGE_TAG ?= alpha
NAMESPACE ?= open-microtools

# Run a real basic sample based on an example token from jwt.io
.PHONY: sample
sample:
	go run . eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c

# sample with payload encoded in base64 instead of base64url
.PHONY: sample-nonurl
sample-nonurl:
	go run . eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lPz8/IiwiaWF0IjoxNTE2MjM5MDIyfQ.K6s7vE/2ZRUY6JQ7CbeGMn77U02AhqDd+wnK/wQ1Q9c

.PHONY: sample-all
sample-all: sample sample-nonurl

# run unit tests exclude wasm folder
.PHONY: test
test:
	go test -v $(shell go list ./... | grep -v /cmd/wasm) -coverprofile=coverage.out

# create test coverage func report
.PHONY: test-coverage-html
test-coverage-html:test
	go tool cover -html=coverage.out

# create test coverage html report
.PHONY: test-coverage-func
test-coverage-func:test
	go tool cover -func coverage.out

# run go fmt
.PHONY: fmt
fmt:
	go fmt ./...

# run go vet exclude wasm folder
.PHONY: vet
vet: 
	go vet $(shell go list ./... | grep -v /cmd/wasm)

# Build the binary and put it in a bin dir
.PHONY: build
build: fmt vet test
	go build -o bin/ .

# compile go code to wasm binary
.PHONY: wasm
wasm: fmt vet test
	GOOS=js GOARCH=wasm go build -o ./assets/jwt.wasm ./cmd/wasm/.

# run go server
run: wasm
	#
	# please browse to http://localhost:8080 to view the served page
	#
	go run cmd/server/main.go

.PHONY: docker-build
docker-build:
	docker build -f ./docker/Dockerfile . --tag jwt-server:$(IMAGE_TAG)

# Kustomize targets
.PHONY: deploy-nodeport
deploy-nodeport: 
	kustomize build deployments/overlays/nodeport | kubectl apply -f -

.PHONY: deploy-ingress
deploy-ingress:
	kustomize build deployments/overlays/ingress | kubectl apply -f -

.PHONY: deploy-base
deploy-base:
	kustomize build deployments/base | kubectl apply -n $(NAMESPACE) -f -

# Local testing functions
.PHONY: port-forward
port-forward:
	./hack/port-forward-to-poc.sh
