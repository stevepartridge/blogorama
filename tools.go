// +build tools

package tools

import (
	_ "github.com/go-bindata/go-bindata"
	_ "github.com/gogo/protobuf/protoc-gen-gogo"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
)
