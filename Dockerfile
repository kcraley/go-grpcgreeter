# Arguments controlling container image the Go will build in.
ARG GO_VERSION=1.18
ARG GO_DISTRO=alpine
ARG GO_WORKDIR=/go/src/github.com/kcraley/go-grpcgreeter

# Compile the binary in a build stage.
FROM golang:${GO_VERSION}-${GO_DISTRO} as build
ARG GO_WORKDIR
WORKDIR ${GO_WORKDIR}}

COPY go.mod ./
COPY go.sum ./
RUN set -ex && \
    go mod download

COPY ./ ./
RUN set -ex && \
    CGO_ENABLED=0 go build -o ${GO_WORKDIR}/bin/greeter

# Build secondary production container
FROM gcr.io/distroless/static-debian11
ARG GO_WORKDIR
WORKDIR ${GO_WORKDIR}

COPY --from=build ${GO_WORKDIR}/bin/greeter /greeter
ENTRYPOINT [ "/greeter" ]
