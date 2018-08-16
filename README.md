# Ping Pong gRPC

[![Docker Repository on Quay](https://quay.io/repository/denderello/ping-pong-grpc/status "Docker Repository on Quay")](https://quay.io/repository/denderello/ping-pong-grpc)
[![Go Report Card](https://goreportcard.com/badge/denderello/ping-pong-grpc "Go Report Card")](https://goreportcard.com/report/denderello/ping-pong-grpc)


A little ping pong program that talks via gRPC.

## Buld

Build the docker container with all necessary binaries
```
docker build -t grpcpingpong .
```

## Run server 

using docker command
```
docker run -d --name test -p 8090:8090 grpcpingpong server --msg "custom return message" --port 8090
```

using predefined docker-compose yaml
```
docker-compose -f docker-compose-sample.yaml up -d
```

## Run client
```
./ping-pong-grpc client --port 7053
```

## Server mode

You can run `ping-pong-grpc` in server mode. In this mode it will wait for
incoming `SendPing()` calls and respond with a simple `Pong` message.

You can run it in server mode like this:
```
build/ping-pong-grpc server
```

## Client mode

You can also run `ping-pong-grpc` in client mode. In this mode it will open a
connection to a instance running in server mode and send `SendPing()` calls with
a `Ping` message.

You can run a single ping rpc in client mode like this:
```
ping-pong-grpc client
```

### Cycle mode

Besides running a single ping rpc call you can also start the client in cycle
mode which will run consecutive rpc calls with a sleep duration.

You can run the client in cycle mode like this:
```
ping-pong-grpc client --cycle-mode --cycle-sleep-duration=1s
```

## Log levels

All modes support multiple log levels via the `--log-level` flag.
