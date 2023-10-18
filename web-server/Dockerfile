####################################################################################################
# Build Stage: Compile and Strip Binary
####################################################################################################
FROM golang:1.21.3-alpine3.18 AS builder
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
COPY --from=builder /app/public /app/public
EXPOSE 3000
ENTRYPOINT ["/app/web-server"]
