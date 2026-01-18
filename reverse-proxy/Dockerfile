ARG BUILD_FROM
FROM $BUILD_FROM

# Install Go and tools
RUN apk add --no-cache go git build-base bash curl

# Build Go application
WORKDIR /app/src
COPY src/ ./
RUN go build -o /app/reverse-proxy main.go

# Copy run script
COPY run.sh /app/run.sh
RUN chmod +x /app/run.sh /app/reverse-proxy

CMD [ "/app/run.sh" ]