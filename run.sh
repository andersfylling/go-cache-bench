#!/bin/sh

docker build -t go-cache-bench .
docker run -t go-cache-bench