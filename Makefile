
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

# a better testing framework needs to be divised however adding this make rule for easy testing of the package which has unit testing
test:
	go test ./pkg/jwt/...

# Build the binary and put it in a bin dir
build:
	go build -o bin/ .

# compile go code to wasm binary
wasm:
	GOOS=js GOARCH=wasm go build -o ./assets/jwt.wasm ./cmd/wasm/.

# run go server
run: wasm
	#
	# please browse to http://localhost:8080 to view the served page
	#
	go run cmd/server/main.go
