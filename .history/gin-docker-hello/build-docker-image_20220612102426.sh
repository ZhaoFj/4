#!/bin/bash

docker build -t hanta/gin_docker_hello:v1.0.0 -f Dockerfile
docker tag hanta/gin_docker_hello:v1.0.0 hanta/gin_docker_hello:v1.0.0