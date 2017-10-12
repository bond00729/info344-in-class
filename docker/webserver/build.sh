#! /usr/bin/env bash
set -e
GOOS=linux go build
docker build -t bond00729/testserver .
docker push bond00729/testserver