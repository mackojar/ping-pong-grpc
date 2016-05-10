PROJECT=ping-pong-grpc
ORGANIZATION=denderello

BUILD_PATH := $(shell pwd)/build
PROJECT_PATH := "github.com/$(ORGANIZATION)/$(PROJECT)"

BIN := $(PROJECT)

VERSION := $(shell cat VERSION)
COMMIT := $(shell git rev-parse --short HEAD)

.PHONY: all clean protoc

SOURCE=$(shell find . -name '*.go')

all: $(BIN)

clean:
	rm -rf $(BUILD_PATH)

build:
	mkdir -p $(BUILD_PATH)

$(BIN): $(SOURCE) VERSION build
	go build \
		-o build/$(BIN) \
		-a -ldflags \
		"-X $(PROJECT_PATH)/cmd.projectVersion=$(VERSION) -X $(PROJECT_PATH)/cmd.projectCommit=$(COMMIT)"

protoc:
	mkdir -p pingpong
	protoc -I pingpong pingpong/pingpong.proto --go_out=plugins=grpc:pingpong
