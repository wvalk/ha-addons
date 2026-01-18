# Step 1: Define runtime base image (must be first line in HA add-ons)
ARG BUILD_FROM
FROM $BUILD_FROM AS runtime

# Step 2: Runtime dependencies (already has bash and bashio in HA base)
WORKDIR /app

# Step 3: Build Go binary in a separate builder stage
FROM golang:1.22-alpine AS builder

WORKDIR /app
ENV GO111MODULE=on

COPY src/ ./src/
WORKDIR /app/src

RUN go build -o /app/reverse-proxy main.go

# Step 4: Copy Go binary and run.sh into runtime container
COPY --from=builder /app/reverse-proxy /app/reverse-proxy
COPY run.sh /app/run.sh

RUN chmod +x /app/reverse-proxy /app/run.sh

# Step 5: Start the application
CMD ["/app/run.sh"]
