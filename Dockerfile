FROM --platform=$BUILDPLATFORM golang:1.27.1-alpine@sha256:cf6fca6641884b8433441b2b0652976f975e1d0fdd26d177eaaf8596087f3125 AS builder

WORKDIR /src/

ARG TARGETOS
ARG TARGETARCH
ARG VERSION

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -a -installsuffix cgo \
      -ldflags="-X github.com/version-cli/version/cmd.VERSION=${VERSION}" -o version

FROM gcr.io/distroless/static:nonroot@sha256:2b7c93f6d6648c11f0e80a48558c8f77885eb0445213b8e69a6a0d7c89fc6ae4

COPY --from=builder /src/version /bin/version

ENTRYPOINT ["/bin/version"]
