# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

ENV GO111MODULE=on

COPY src/ ./src/
WORKDIR /app/src

RUN go build -o /app/reverse-proxy main.go

# Runtime stage
FROM alpine:3.20

WORKDIR /app

# Add bash to support run.sh
RUN apk add --no-cache bash

COPY --from=builder /app/reverse-proxy /app/reverse-proxy
COPY run.sh /app/run.sh

RUN chmod +x /app/run.sh /app/reverse-proxy

EXPOSE 8080

CMD ["/app/run.sh"]
