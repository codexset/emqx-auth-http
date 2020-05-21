#!/bin/sh
# Login docker
echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin
# Build Golang Application
go build -o dist/emqx-auth-http
# Build docker image
docker build . -t kainonly/emqx-auth-http:latest
# Push docker image
docker push kainonly/emqx-auth-http:latest