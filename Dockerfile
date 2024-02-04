FROM --platform=${BUILDPLATFORM} golang:1.21-bookworm@sha256:135ba6c2ae2a9a022f53f3eee24ef213bc62e876811cd9582b91045168206873 AS builder

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

FROM gcr.io/distroless/static:nonroot@sha256:112a87f19e83c83711cc81ce8ed0b4d79acd65789682a6a272df57c4a0858534

COPY --from=builder /src/version /bin/version

ARG CREATED
ARG AUTHORS="Koen van Zuijlen <8818390+kvanzuijlen@users.noreply.github.com>"
ARG URL="https://github.com/kvanzuijlen/version"
ARG DOCUMENTATION="https://github.com/kvanzuijlen/version"
ARG SOURCE="https://github.com/kvanzuijlen/version"
ARG VERSION
ARG REVISION
ARG VENDOR="kvanzuijlen"
ARG LICENSES="GNU GPLv3"
ARG TITLE="version"
ARG DESCRIPTION="A simple CLI tool to deal with version numbers."
ARG BASE_IMAGE="gcr.io/distroless/static:nonroot"
ARG BASE_IMAGE_DIGEST="sha256:112a87f19e83c83711cc81ce8ed0b4d79acd65789682a6a272df57c4a0858534"

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
