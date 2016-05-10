FROM golang:1.6-alpine

COPY . $GOPATH/src/github.com/denderello/ping-pong-grpc

WORKDIR $GOPATH/src/github.com/denderello/ping-pong-grpc

RUN apk add --update \
    make \
    musl-dev \
    git \
    gcc \
  && \
    make \
  && \
    mv build/ping-pong-grpc $GOPATH/bin
  

ENTRYPOINT ["ping-pong-grpc"]

EXPOSE 8080
