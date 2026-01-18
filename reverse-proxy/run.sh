#!/usr/bin/with-contenv bashio
TARGET_URL=$(bashio::config 'target_url')
bashio::log.info "Reverse proxy starting with target URL: ${TARGET_URL}"

export TARGET_URL

exec /app/reverse-proxy
