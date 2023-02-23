FROM cgr.dev/chainguard/go:latest AS builder

WORKDIR /src

RUN mkdir jwt-inspect

WORKDIR /src/jwt-inspect

COPY . .

RUN hack/build.sh

# FROM gcr.io/distroless/static-debian11:nonroot
FROM cgr.dev/chainguard/static:latest

WORKDIR /

COPY assets/ assets/
COPY --from=builder /src/jwt-inspect/server .
COPY --from=builder /src/jwt-inspect/jwt.wasm assets/jwt.wasm

EXPOSE 8080

# RUN ls -lah /server

ENTRYPOINT [ "/server" ]
