PROJECT=ping-pong-grpc
ORGANIZATION=denderello

PROJECT_PATH := "github.com/$(ORGANIZATION)/$(PROJECT)"
BIN := $(PROJECT)

VERSION := $(shell cat VERSION)
COMMIT := $(shell git rev-parse --short HEAD)

SOURCE=$(shell find . -name '*.go')
BUILD_FLAGS=-a -ldflags \
	"-X $(PROJECT_PATH)/cmd.projectVersion=$(VERSION) -X $(PROJECT_PATH)/cmd.projectCommit=$(COMMIT)"

.PHONY: install protoc

$(BIN): $(SOURCE) VERSION
	go build \
		$(BUILD_FLAGS)

install:
	go install \
		$(BUILD_FLAGS)

protoc:
	mkdir -p pingpong
	protoc -I pingpong pingpong/pingpong.proto --go_out=plugins=grpc:pingpong
