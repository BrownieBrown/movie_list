PROTO_DIR = protobuf
PROTO_FILE = pkg/$(PROTO_DIR)/player.proto
GO_OUT_DIR = .

.PHONY: all proto

all: proto

proto:
	protoc --go_out=$(GO_OUT_DIR) --go_opt=paths=source_relative \
	       --go-grpc_out=$(GO_OUT_DIR) --go-grpc_opt=paths=source_relative \
	       $(PROTO_FILE)
