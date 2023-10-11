#
# Build stage
#
FROM golang:1.21.3-alpine3.18 AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /os-project/build
COPY . .
RUN go mod download && \
    go build -o web-server main.go && \
    chmod +x web-server && \
    rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

#
# Final stage
#
FROM busybox:1.36.1-glibc AS final
WORKDIR /os-project
COPY --from=builder /os-project/build/web-server .
ENTRYPOINT [ "./web-server" ]
EXPOSE 3000