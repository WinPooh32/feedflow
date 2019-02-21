#!/usr/bin/env bash

#https://github.com/OpenAPITools/openapi-generator
openapi-generator generate -i api-oas3.yaml -g typescript-axios -o ../client-ts-axios