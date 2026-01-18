ARG BUILD_FROM
FROM $BUILD_FROM

# Install Go and tools
RUN apk add --no-cache go git build-base bash curl

# Install Bashio manually (because BUILD_FROM does NOT include it)
RUN mkdir -p /usr/lib/bashio \
    && curl -sSL https://raw.githubusercontent.com/hassio-addons/bashio/main/lib/bashio.sh -o /usr/lib/bashio/bashio.sh

ENV BASHIO_LIB=/usr/lib/bashio

# Build Go application
WORKDIR /app/src
COPY src/ ./
RUN go build -o /app/reverse-proxy main.go

# Copy run script
COPY run.sh /app/run.sh
RUN chmod +x /app/run.sh /app/reverse-proxy

# Copy s6 services
COPY rootfs/ /

# No CMD needed — s6-overlay handles startup
