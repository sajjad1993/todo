gen-grpc:
	@mkdir -p pkg/rpc
	@protoc \
		--go_opt=paths=source_relative --go_out=internal/common/rpc \
		--go-grpc_opt=paths=source_relative --go-grpc_out=internal/common/rpc \
		api/protobuf/user.proto

install_tools:
	@apt update
	@apt install -y protoc-gen-go
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

clean-grpc:
	@rm -rf pkg/rpc
