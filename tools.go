// +build tools

package tools

import (
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/gogo/protobuf/protoc-gen-gogo"
	_ "github.com/gogo/protobuf/protoc-gen-gogoslick"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	_ "github.com/mwitkow/go-proto-validators/protoc-gen-govalidators"
	_ "github.com/rakyll/statik"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
