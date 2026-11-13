FROM --platform=$BUILDPLATFORM golang:1.25.5-alpine@sha256:ac09a5f469f307e5da71e766b0bd59c9c49ea460a528cc3e6686513d64a6f1fb AS builder

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

FROM gcr.io/distroless/static:nonroot@sha256:aafe993665bf6200db38a02c11a9351e782a84b6a45e7f33b6695638c88e3931

COPY --from=builder /src/version /bin/version

ENTRYPOINT ["/bin/version"]
