FROM --platform=$BUILDPLATFORM golang:1.25 AS build
WORKDIR /workspace
COPY go.mod go.sum .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY . .
ARG TARGETOS TARGETARCH
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o webhook -ldflags '-w -extldflags "-static"' .

FROM alpine:3.23.3 AS certs
RUN apk add -U --no-cache ca-certificates

FROM scratch
COPY --from=certs /etc/ssl/certs /etc/ssl/certs
COPY --from=build /workspace/webhook /usr/local/bin/webhook
USER 1000:1000
ENTRYPOINT ["webhook"]
