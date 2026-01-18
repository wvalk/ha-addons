#!/usr/bin/env bash
source /usr/lib/bashio/bashio.sh

TARGET_URL=$(bashio::config 'target_url')
bashio::log.info "Starting reverse proxy with target URL: ${TARGET_URL}"

export TARGET_URL
exec /app/reverse-proxy
