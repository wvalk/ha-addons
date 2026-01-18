# First line: Home Assistant add-on base injected by Supervisor
ARG BUILD_FROM
FROM $BUILD_FROM

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
