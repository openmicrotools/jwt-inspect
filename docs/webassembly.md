# WebAssembly

## Quick start

1. Compile go code to wasm

```bash
cd wasm
GOOS=js GOARCH=wasm go build -o ./assets/jwt.wasm .
```

1. Run go web server

```bash
cd server
go run main.go
```

1. Go to browser and type <http://localhost:8080>
