FROM golang:1.6-alpine

COPY . $GOPATH/src/github.com/denderello/ping-pong-grpc

WORKDIR $GOPATH/src/github.com/denderello/ping-pong-grpc

RUN apk add --update \
    make \
    musl-dev \
    git \
    gcc \
  && make install

ENTRYPOINT ["ping-pong-grpc"]

