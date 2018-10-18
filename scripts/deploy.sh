#! /bin/bash

VERSION=$1

gcloud app deploy \
--version=$VERSION \
app/servers/default/app.yaml \
app/servers/apis/app.yaml
