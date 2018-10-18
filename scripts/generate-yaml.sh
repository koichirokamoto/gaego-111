#!/bin/bash

API_KEY=$1

cat > app/servers/apis/app.yaml << EOF
runtime: go111
api_version: go1

service: apis

handlers:
  - url: /.*
    script: auto
    secure: always
    redirect_http_response_code: 301

env_variables:
  WETHER_API_KEY: "$API_KEY"
EOF
