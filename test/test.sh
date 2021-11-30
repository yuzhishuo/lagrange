#!/usr/bin/bash

for ((i=1; i<1001 ;i++)) ; do
    rand_key=$(head -c 5 /dev/urandom | od -A n -t x | tr -d ' ')
    rand_val=$(head -c 5 /dev/urandom | od -A n -t x | tr -d ' ')
    curl -H "Content-Type: application/json" -X POST -d '{  "Key": "'"${rand_key}"'", "Value": "'"${rand_val}"'"  }' "localhost:8071/kv"
done