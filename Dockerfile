FROM --platform=${BUILDPLATFORM} golang:1.22-bookworm@sha256:b03f3ba515751657c75475b20941fef47341fccb3341c3c0b64283ff15d3fb46 AS builder

WORKDIR /src/

COPY . .

ARG VERSION

RUN case ${TARGETPLATFORM} in \
         "linux/amd64")  GOARCH=amd64  ;; \
         # arm64 and arm64v8 are equivilant in go and do not require a goarm
         # https://github.com/golang/go/wiki/GoArm
         "linux/arm64" | "linux/arm/v8")  GOARCH=arm64  ;; \
         "linux/ppc64le")  GOARCH=ppc64le  ;; \
         "linux/arm/v6") GOARCH=arm GOARM=6  ;; \
         "linux/arm/v7") GOARCH=arm GOARM=7 ;; \
    esac && \
    GOARCH=${GOARCH} VERSION=${VERSION} make build

FROM gcr.io/distroless/static:nonroot@sha256:e9ac71e2b8e279a8372741b7a0293afda17650d926900233ec3a7b2b7c22a246

COPY --from=builder /src/version /bin/version

ARG CREATED
ARG AUTHORS="Koen van Zuijlen <8818390+kvanzuijlen@users.noreply.github.com>"
ARG URL="https://github.com/version-cli/version"
ARG DOCUMENTATION="https://github.com/version-cli/version"
ARG SOURCE="https://github.com/version-cli/version"
ARG VERSION
ARG REVISION
ARG VENDOR="version-cli"
ARG LICENSES="GNU GPLv3"
ARG TITLE="version"
ARG DESCRIPTION="A simple CLI tool to deal with version numbers."
ARG BASE_IMAGE="gcr.io/distroless/static:nonroot"
ARG BASE_IMAGE_DIGEST="sha256:e9ac71e2b8e279a8372741b7a0293afda17650d926900233ec3a7b2b7c22a246"

LABEL authors="kvanzuijlen" \
    org.opencontainers.image.created=${CREATED} \
    org.opencontainers.image.authors=${AUTHORS} \
    org.opencontainers.image.url=${URL} \
    org.opencontainers.image.documentation=${DOCUMENTATION} \
    org.opencontainers.image.source=${SOURCE} \
    org.opencontainers.image.version=${VERSION} \
    org.opencontainers.image.revision=${REVISION} \
    org.opencontainers.image.vendor=${VENDOR} \
    org.opencontainers.image.licenses=${LICENSES} \
    org.opencontainers.image.ref.name=v${VERSION} \
    org.opencontainers.image.title=${TITLE} \
    org.opencontainers.image.description=${DESCRIPTION} \
    org.opencontainers.image.base.digest=${BASE_IMAGE_DIGEST} \
    org.opencontainers.image.base.name=${BASE_IMAGE}

ENTRYPOINT ["/bin/version"]
