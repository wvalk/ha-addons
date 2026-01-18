#!/usr/bin/with-contenv bashio

echo "Hello world!"

echo "printing message from config: "; cat /data/config.json | jq '.message' 