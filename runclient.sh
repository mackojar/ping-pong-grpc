#!/bin/bash

# run regular server
docker run -d --name pingpong -p 8090:8090 grpcpingpong server --msg "custom return message" --port 8090
