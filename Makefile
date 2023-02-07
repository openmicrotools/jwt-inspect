
# Run a real basic sample based on an example token from jwt.io
sample:
	go run . eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c

# Build the binary and put it in a bin dir
build:
	go build -o bin/ .

# compile go code to wasm binary
wasm:
	GOOS=js GOARCH=wasm go build -o ./assets/jwt.wasm ./cmd/wasm/.

# run go server
run:
	go run cmd/server/main.go
