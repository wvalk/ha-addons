#!/usr/bin/with-contenv bashio

echo "Hello world!"

python3 -m http.server 8000

echo "printing message from config: "; cat /data/config.json | jq '.message' 