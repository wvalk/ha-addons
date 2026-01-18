# Home Assistant Supervisor injects this automatically
ARG BUILD_FROM
FROM $BUILD_FROM

# Install dependencies: Go, compiler tools, Bash, Curl
RUN apk add --no-cache go git build-base bash curl

# Install Bashio manually (Supervisor's BUILD_FROM does NOT include it)
RUN mkdir -p /usr/lib/bashio \
    && curl -sSL https://raw.githubusercontent.com/hassio-addons/bashio/main/lib/bashio.sh -o /usr/lib/bashio/bashio.sh

ENV BASHIO_LIB=/usr/lib/bashio

# Build Go application
WORKDIR /app/src
COPY src/ ./
RUN go build -o /app/reverse-proxy main.go

# Install run.sh
COPY run.sh /app/run.sh
RUN chmod +x /app/run.sh /app/reverse-proxy

WORKDIR /app
CMD ["/app/run.sh"]
