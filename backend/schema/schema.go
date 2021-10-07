//go:generate protoc --go-grpc_out=require_unimplemented_servers=false:./grpcapi --go_out=./grpcapi grpcapi.proto

package schema