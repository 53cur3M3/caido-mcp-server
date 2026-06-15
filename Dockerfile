# syntax=docker/dockerfile:1

FROM golang:1.24 AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG VERSION=dev
ARG TOOL=mcp
ARG TARGETOS=linux
ARG TARGETARCH=amd64

RUN set -eux; \
    mkdir -p /out; \
    case "${TOOL}" in \
      mcp|server) \
        ext=""; \
        if [ "${TARGETOS}" = "windows" ]; then ext=".exe"; fi; \
        CGO_ENABLED=0 GOOS="${TARGETOS}" GOARCH="${TARGETARCH}" \
          go build -ldflags="-s -w -X main.version=${VERSION}" \
          -o "/out/caido-mcp-server-${TARGETOS}-${TARGETARCH}${ext}" ./cmd/mcp; \
        ;; \
      cli) \
        ext=""; \
        if [ "${TARGETOS}" = "windows" ]; then ext=".exe"; fi; \
        CGO_ENABLED=0 GOOS="${TARGETOS}" GOARCH="${TARGETARCH}" \
          go build -ldflags="-s -w" \
          -o "/out/caido-cli-${TARGETOS}-${TARGETARCH}${ext}" ./cmd/cli; \
        ;; \
      all) \
        bash ./scripts/build.sh "${VERSION}"; \
        cp dist/* /out/; \
        ;; \
      *) \
        echo "Unknown TOOL: ${TOOL} (use 'mcp', 'cli', or 'all')" >&2; \
        exit 1; \
        ;; \
    esac

FROM scratch AS artifacts
COPY --from=build /out/ /
