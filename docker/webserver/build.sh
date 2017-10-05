#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t pavk/testserver .
docker push pavk/testserver
go clean