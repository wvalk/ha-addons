# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app
ENV GO111MODULE=on

COPY src/ ./src/
WORKDIR /app/src

RUN go build -o /app/reverse-proxy main.go

# Runtime stage (Home Assistant add-on base)
ARG BUILD_FROM=ghcr.io/home-assistant/amd64-addon-base:13.2.0
FROM $BUILD_FROM

WORKDIR /app

# Copy application
COPY --from=builder /app/reverse-proxy /app/reverse-proxy
COPY run.sh /app/run.sh

RUN chmod +x /app/run.sh /app/reverse-proxy

CMD ["/app/run.sh"]
