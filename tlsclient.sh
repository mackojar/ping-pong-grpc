#!/bin/bash

# first run docker-compose -f dc-tlsping.yaml up

docker exec -ti grpctlsclient ping-pong-grpc client --port 8090 --host ${GRPCS_HOST} --cert /etc/certs/ca.crt
