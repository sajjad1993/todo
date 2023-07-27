gen-grpc:
	@protoc \
		--go_opt=paths=source_relative --go_out=internal/pkg/rpc/proto/user \
		--go-grpc_opt=paths=source_relative --go-grpc_out=internal/pkg/rpc/proto/user \
		api/protobuf/user.proto
	@protoc \
    		--go_opt=paths=source_relative --go_out=internal/pkg/rpc/proto/auth \
    		--go-grpc_opt=paths=source_relative --go-grpc_out=internal/pkg/rpc/proto/auth \
    		api/protobuf/auth.proto
	@protoc \
    		--go_opt=paths=source_relative --go_out=internal/pkg/rpc/proto/todo_list \
    		--go-grpc_opt=paths=source_relative --go-grpc_out=internal/pkg/rpc/proto/todo_list \
    		api/protobuf/todo_list.proto



install_tools:
	@apt update
	@apt install -y protoc-gen-go
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

clean-grpc:
	@rm -rf pkg/rpc
