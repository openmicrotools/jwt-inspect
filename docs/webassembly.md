# WebAssembly

## Quick start

1. Compile go code to wasm `make wasm`
1. Run go web server `make run`
1. Go to browser and type <http://localhost:8080>

## Development notes

### Regarding IDE error in cmd/wasm/main.go

VS code error: `error while importing syscall/js: build constraints exclude all Go files in /usr/local/go/src/syscall/js`

The reason for this is that `syscall/js` package should be complied on `wasm` architecture and `js` os, like the command when we run to compile wasm binary: `GOOS=js GOARCH=wasm go build -o ./assets/jwt.wasm ./cmd/wasm/`. However IDE(VS code) isn't aware of the setting. It is not a concern error just a VS code error.

The solution to fix this error is to add `wasm/.vscode/settings`

```json
{
  "go.toolsEnvVars": {
    "GOOS": "js",
    "GOARCH": "wasm"
  }
}
```

and configure VS code with [multi-root-workspaces](https://code.visualstudio.com/docs/editor/multi-root-workspaces)
