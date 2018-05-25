# build the gateway binary
FROM golang:1.10.0-alpine as builder
LABEL intermediate=true
WORKDIR /go/src/github.com/infobloxopen/atlas-stubs/authn-stub
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/authn-stub ./cmd/

# build the gateway binary
FROM alpine:latest as runner
WORKDIR /bin
COPY --from=0 /go/src/github.com/infobloxopen/atlas-stubs/authn-stub/bin/authn-stub .
ENTRYPOINT ["authn-stub"]
