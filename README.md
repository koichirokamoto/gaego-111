# Google App Engine Go 1.11 sample

## Prerequisite

- Intall Go
- Install Google Cloud SDK

## Configuration

Set your project configuration.

`gcloud config configurations create gaego111`

Check created configuration is active.

`gcloud config configurations list --format="value(name)" --filter="is_active: true"`

Set configurations.

`gcloud init`

## Deploy

`gcloud app deploy --version=first-app app/servers/default/app.yaml app/servers/apis/app.yaml`

Or, run `./scripts/deploy.sh`

## Dispatch

`./scripts/dispatch.yaml`

## Check

`gcloud app browse`
