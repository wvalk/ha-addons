#!/usr/bin/with-contenv bashio

echo "Hello world!"

echo "printing message from config: "; cat config.json | jq '.message' 