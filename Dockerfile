FROM golang:1.25.7-alpine AS go-builder
RUN apk --no-cache add tzdata ca-certificates
WORKDIR /work
RUN mkdir /data
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app

# runner
FROM alpine:3.22
WORKDIR /work
ENV TZ=Asia/Shanghai \
    PORT=8080 \
    LOG_LEVEL=info
HEALTHCHECK --interval=5s --timeout=3s --retries=5 --start-period=5s \
    CMD wget --spider http://127.0.0.1:$PORT/health || exit 1

COPY --from=go-builder /data /data
COPY --from=go-builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=go-builder /work/app /app
ENTRYPOINT ["/app"]
