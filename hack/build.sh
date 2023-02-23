#!/usr/bin/env sh
set -euo pipefail
IFS=$'\n\t'

# build basic golang server
CGO_ENABLED=0 go build -o ./server ./cmd/server/.

# build the Web Assebly binary we'll serve
GOOS=js GOARCH=wasm go build -o ./jwt.wasm ./cmd/wasm/.
