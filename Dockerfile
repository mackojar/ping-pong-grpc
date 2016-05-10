FROM golang:1.6-wheezy

COPY . $GOPATH/src/github.com/denderello/ping-pong-grpc

WORKDIR $GOPATH/src/github.com/denderello/ping-pong-grpc

RUN make && mv build/ping-pong-grpc $GOPATH/bin

ENTRYPOINT ["ping-pong-grpc"]

EXPOSE 8080
