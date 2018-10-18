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

## Generete app yaml for apis service

run `./scripts/generate-yaml.sh`

If you try `/apis/temp`, visit [here](https://openweathermap.org/) and get API keys.
Then, run `./scripts/generate-yaml.sh <API keys>`

## Deploy

`gcloud app deploy --version=first-app app/servers/default/app.yaml app/servers/apis/app.yaml`

Or, run `./scripts/deploy.sh first-app`

## Dispatch

`./scripts/dispatch.yaml`

## Check

`gcloud app browse`
