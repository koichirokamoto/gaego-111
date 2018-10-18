#! /bin/bash

APPSPOT=$1

cat > app/servers/dispatch.yaml << EOF
dispatch:
  - url: "$APPSPOT/public/*"
    service: default
  - url: "$APPSPOT/"
    service: default
  - url: "*/apis/*"
    service: apis
EOF

gcloud app deploy app/servers/dispatch.yaml
