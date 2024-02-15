PROTOC ?= $(shell which protoc)
ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: go-mod
go-mod:
	@go mod download
	@go mod tidy

.PHONY: proto
proto: | go-mod
	@$(PROTOC) --proto_path=$(ROOT_DIR)/proto \
                   --go_opt=paths=source_relative \
                   --go-grpc_opt=paths=source_relative \
                   --go_out=$(ROOT_DIR)/pkg/iam/runtime \
                   --go-grpc_out=$(ROOT_DIR)/pkg/iam/runtime \
                   authentication/authentication.proto authorization/authorization.proto
