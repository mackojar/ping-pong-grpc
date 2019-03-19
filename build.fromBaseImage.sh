#!/bin/bash

# build image
docker build -t grpcpingpong -f Dockerfile.fromBaseImage .
