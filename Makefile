

build:
	go build -o echo -v github.com/covarity/echo/cmd/echo 

test:
	go test ./...

benchmark:
	go test -bench=. -benchmem -count=10 ./... > bench.txt
	benchstat bench.txt

proto.install:
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	cp -r ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google third_party/
	mkdir -p third_party/protoc-gen-swagger/options
	cp ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options/annotations.proto third_party/protoc-gen-swagger/options
	cp ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options/openapiv2.proto third_party/protoc-gen-swagger/options

proto.gen:
	mkdir -p pkg/api/v1
	protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 task.proto
	mkdir -p api/swagger/v1
	protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 task.proto
	protoc --proto_path=api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:pkg/api/v1 task.proto
	protoc --proto_path=api/proto/v1 --proto_path=third_party --swagger_out=logtostderr=true:api/swagger/v1 task.proto

