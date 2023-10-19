####################################################################################################
# Base Stage: Install Dependencies
####################################################################################################
FROM alpine:3.18 AS base
RUN apk add --no-cache ca-certificates
ENV PATH /usr/local/go/bin:$PATH
ENV GOLANG_VERSION 1.21.3
RUN set -eux; \
    apk add --no-cache --virtual .fetch-deps gnupg; \
    arch="$(apk --print-arch)"; \
    url=; \
    case "$arch" in \
    'x86_64') \
    url='https://dl.google.com/go/go1.21.3.linux-amd64.tar.gz'; \
    sha256='1241381b2843fae5a9707eec1f8fb2ef94d827990582c7c7c32f5bdfbfd420c8'; \
    ;; \
    'armhf') \
    url='https://dl.google.com/go/go1.21.3.linux-armv6l.tar.gz'; \
    sha256='a1ddcaaf0821a12a800884c14cb4268ce1c1f5a0301e9060646f1e15e611c6c7'; \
    ;; \
    'armv7') \
    url='https://dl.google.com/go/go1.21.3.linux-armv6l.tar.gz'; \
    sha256='a1ddcaaf0821a12a800884c14cb4268ce1c1f5a0301e9060646f1e15e611c6c7'; \
    ;; \
    'aarch64') \
    url='https://dl.google.com/go/go1.21.3.linux-arm64.tar.gz'; \
    sha256='fc90fa48ae97ba6368eecb914343590bbb61b388089510d0c56c2dde52987ef3'; \
    ;; \
    'x86') \
    url='https://dl.google.com/go/go1.21.3.linux-386.tar.gz'; \
    sha256='fb209fd070db500a84291c5a95251cceeb1723e8f6142de9baca5af70a927c0e'; \
    ;; \
    'ppc64le') \
    url='https://dl.google.com/go/go1.21.3.linux-ppc64le.tar.gz'; \
    sha256='3b0e10a3704f164a6e85e0377728ec5fd21524fabe4c925610e34076586d5826'; \
    ;; \
    'riscv64') \
    url='https://dl.google.com/go/go1.21.3.linux-riscv64.tar.gz'; \
    sha256='67d14d3e513e505d1ec3ea34b55641c6c29556603c7899af94045c170c1c0f94'; \
    ;; \
    's390x') \
    url='https://dl.google.com/go/go1.21.3.linux-s390x.tar.gz'; \
    sha256='4c78e2e6f4c684a3d5a9bdc97202729053f44eb7be188206f0627ef3e18716b6'; \
    ;; \
    *) echo >&2 "error: unsupported architecture '$arch' (likely packaging update needed)"; exit 1 ;; \
    esac; \
    build=; \
    if [ -z "$url" ]; then \
    build=1; \
    url='https://dl.google.com/go/go1.21.3.src.tar.gz'; \
    sha256='186f2b6f8c8b704e696821b09ab2041a5c1ee13dcbc3156a13adcf75931ee488'; \
    echo >&2; \
    echo >&2 "warning: current architecture ($arch) does not have a compatible Go binary release; will be building from source"; \
    echo >&2; \
    fi; \
    \
    wget -O go.tgz.asc "$url.asc"; \
    wget -O go.tgz "$url"; \
    echo "$sha256 *go.tgz" | sha256sum -c -; \
    \
    GNUPGHOME="$(mktemp -d)"; export GNUPGHOME; \
    gpg --batch --keyserver keyserver.ubuntu.com --recv-keys 'EB4C 1BFD 4F04 2F6D DDCC  EC91 7721 F63B D38B 4796'; \
    gpg --batch --keyserver keyserver.ubuntu.com --recv-keys '2F52 8D36 D67B 69ED F998  D857 78BD 6547 3CB3 BD13'; \
    gpg --batch --verify go.tgz.asc go.tgz; \
    gpgconf --kill all; \
    rm -rf "$GNUPGHOME" go.tgz.asc; \
    \
    tar -C /usr/local -xzf go.tgz; \
    rm go.tgz; \
    \
    if [ -n "$build" ]; then \
    apk add --no-cache --virtual .build-deps \
    bash \
    gcc \
    go \
    musl-dev \
    ; \
    \
    export GOCACHE='/tmp/gocache'; \
    \
    ( \
    cd /usr/local/go/src; \
    export GOROOT_BOOTSTRAP="$(go env GOROOT)" GOHOSTOS="$GOOS" GOHOSTARCH="$GOARCH"; \
    if [ "${GOARCH:-}" = '386' ]; then \
    export CGO_CFLAGS='-fno-stack-protector'; \
    fi; \
    ./make.bash; \
    ); \
    \
    apk del --no-network .build-deps; \
    \
    rm -rf \
    /usr/local/go/pkg/*/cmd \
    /usr/local/go/pkg/bootstrap \
    /usr/local/go/pkg/obj \
    /usr/local/go/pkg/tool/*/api \
    /usr/local/go/pkg/tool/*/go_bootstrap \
    /usr/local/go/src/cmd/dist/dist \
    "$GOCACHE" \
    ; \
    fi; \
    \
    apk del --no-network .fetch-deps; \
    \
    go version
ENV GOTOOLCHAIN=local
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 1777 "$GOPATH"
WORKDIR $GOPATH

####################################################################################################
# Build Stage: Compile and Strip Binary
####################################################################################################
FROM base AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /app
COPY . .
RUN go mod download && \
    go build -o web-server main.go && \
    chmod +x web-server && \
    apk add --no-cache --upgrade "upx=4.0.2-r0" && \
    upx -q -o /app/web-server-compressed /app/web-server && \
    rm -rf /var/cache/apk/* && \
    rm -rf /tmp/* && \
    rm -rf /var/log/*

####################################################################################################
# Final Stage: Minimal Image
####################################################################################################
FROM scratch
LABEL project="os-container-project"
LABEL description="A minimal container image for a web server"
WORKDIR /app
COPY --from=builder /app/web-server-compressed /app/web-server
COPY --from=builder /app/data /app/data
EXPOSE 3000
CMD ["/app/web-server"]