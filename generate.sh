#!/bin/bash

OPENAPI_GENERATOR_VERSION=v6.2.0

docker run --user ${UID:?}:$(id -g) --rm -v "$(pwd):/input" -v "$(pwd)/skipr:/output" openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION:?} \
generate -i /input/openapi.yml -g go-gin-server -o /output --global-property debugOperations=true -c /input/config.yml

if [ -n "$SUDO_UID" ]
then
    chown ${SUDO_UID:?}:${SUDO_GID:?} -R skipr
fi
