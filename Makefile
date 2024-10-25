PROTO_REPO_PATH=$(shell go list -m -f '{{ .Dir }}' github.com/faruqii/faruqi-protos)
PROTO_OUT_DIR=.

.PHONY: proto

proto-products:
	protoc -I="$(PROTO_REPO_PATH)" \
		--go_out=$(PROTO_OUT_DIR) \
		--go-grpc_out=$(PROTO_OUT_DIR) \
		"$(PROTO_REPO_PATH)/proto/products/*.proto"
