#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t pavk/zipsvr .
docker push pavk/zipsvr
go clean 