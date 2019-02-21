#!/usr/bin/env bash

export GO_POST_PROCESS_FILE="/usr/bin/gofmt -w"

openapi-generator generate -i api-oas3.yaml -g go-gin-server -o ../server-go-gin