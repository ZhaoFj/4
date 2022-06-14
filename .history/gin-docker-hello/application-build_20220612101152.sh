#!/bin/bash
echo 'starting'
CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o gin_docker_hello main.go
echo 'end...'