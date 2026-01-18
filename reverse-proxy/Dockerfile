# First line: Home Assistant add-on base injected by Supervisor
FROM ghcr.io/home-assistant/amd64-addon-base:13.2.0

# Install Go compiler for building the binary
RUN apk add --no-cache go build-base git

# Set working directory
WORKDIR /app

# Copy Go source code
COPY src/ ./src/
WORKDIR /app/src

# Build Go binary
RUN go build -o /app/reverse-proxy main.go

# Copy run.sh
COPY run.sh /app/run.sh
RUN chmod +x /app/run.sh /app/reverse-proxy

# Start the add-on
CMD ["/app/run.sh"]
