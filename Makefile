.PHONY: generate

generate:
	mkdir -p pkg/user_v1
	protoc --proto_path vendor.protogen --proto_path api/user_v1 --go_out=pkg/user_v1 --go_opt=paths=import \
          --go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=import \
          --grpc-gateway_out=pkg/user_v1 \
          --grpc-gateway_opt=logtostderr=true \
          --grpc-gateway_opt=paths=import \
          --validate_out lang=go:pkg/user_v1 \
          api/user_v1/user.proto
	mv pkg/user_v1/github.com/iamtonydev/user-service-api/pkg/user_v1/* pkg/user_v1/
	rm -rf pkg/user_v1/github.com/

PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google/protobuf ]; then \
			git clone https://github.com/protocolbuffers/protobuf vendor.protogen/protobuf &&\
			mkdir -p  vendor.protogen/google/protobuf &&\
			mv vendor.protogen/protobuf/src/google/protobuf/*.proto vendor.protogen/google/protobuf &&\
			rm -rf vendor.protogen/protobuf ;\
		fi