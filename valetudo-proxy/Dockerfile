ARG BUILD_FROM
FROM $BUILD_FROM

# Add env
ENV LANG C.UTF-8

# Install Go and tools
RUN apk add --no-cache go

#copy the source code to the container
COPY src/ src/
# Build the code
RUN go build -o /usr/local/bin/valetudo-proxy src/main.go
RUN chmod a+x /usr/local/bin/valetudo-proxy

# Copy data
COPY rootfs /