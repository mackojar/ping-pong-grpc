#!/bin/bash

# first run docker-compose -f dc-ping.yaml up

docker exec -ti grpctlsclient ping-pong-grpc client --port 8090
