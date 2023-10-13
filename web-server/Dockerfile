#
# Test stage
#
FROM golang:1.21.3-alpine3.18 AS test
WORKDIR /os-project/build
COPY . .
SHELL ["/bin/ash", "-eo", "pipefail", "-c"]
RUN apk --no-cache --upgrade add git=2.40.1-r0 curl=8.4.0-r0 && \
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2 && \
    curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.18.0
RUN go mod download && \
    golangci-lint run && \
    gosec ./... && \
    go test -v ./...

#
# Builder stage
#
FROM test AS builder
ARG USER_ID=1000
ARG GROUP_ID=1000
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN go mod download && \
    go build -o web-server main.go && \
    chmod +x web-server && \
    rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

#
# Final stage
#
FROM busybox:1.36.1-glibc AS final
LABEL project="os-container-project"
LABEL description="This is a Go web server with Fiber framework"
WORKDIR /os-project
COPY --from=builder /os-project/build/web-server .
COPY --from=builder /os-project/build/data ./data
COPY --from=builder /os-project/build/public ./public
USER ${USER_ID}:${GROUP_ID}
ENTRYPOINT [ "./web-server" ]
EXPOSE 3000
