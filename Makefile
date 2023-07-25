gen-grpc:
	@protoc \
		--go_opt=paths=source_relative --go_out=internal/common/rpc/user \
		--go-grpc_opt=paths=source_relative --go-grpc_out=internal/common/rpc/user \
		api/protobuf/user.proto
	@protoc \
    		--go_opt=paths=source_relative --go_out=internal/common/rpc/auth \
    		--go-grpc_opt=paths=source_relative --go-grpc_out=internal/common/rpc/auth \
    		api/protobuf/auth.proto

install_tools:
	@apt update
	@apt install -y protoc-gen-go
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

clean-grpc:
	@rm -rf pkg/rpc
