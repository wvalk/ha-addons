#!/usr/bin/env bash
set -e

TARGET_URL="$(bashio::config 'target_url')"

if [ -z "$TARGET_URL" ]; then
    echo "Error: target_url is not set"
    exit 1
fi

export TARGET_URL

echo "Starting reverse proxy with target: $TARGET_URL"
/app/reverse-proxy
