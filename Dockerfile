FROM golang:1.6-wheezy

COPY . /opt/ping-pong-grpc/src

RUN mv /opt/ping-pong-grpc/src/build/ping-pong-grpc /opt/ping-pong-grpc/ping-pong-grpc

ENTRYPOINT ["/opt/ping-pong-grpc/ping-pong-grpc"]

EXPOSE 8080
