

build:
	go build -o echo -v github.com/covarity/echo/cmd/echo 

test:
	go test ./...

benchmark:
	go test -bench=. -benchmem -count=10 ./... > bench.txt
	benchstat bench.txt

proto.install:
	go get -u github.com/golang/protobuf
	go get -u github.com/gogo/protobuf/protoc-gen-gogo
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/gogo/protobuf/protoc-gen-gogoslick
	go get -u github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
	go get -u github.com/rakyll/statik
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get github.com/gogo/protobuf/gogoproto
	cp -r ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google third_party/
	mkdir -p third_party/protoc-gen-swagger/options
	cp ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options/annotations.proto third_party/protoc-gen-swagger/options
	cp ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options/openapiv2.proto third_party/protoc-gen-swagger/options
	mkdir -p third_party/gogo 
	cp -r ${GOPATH}/src/github.com/gogo/protobuf/gogoproto third_party/gogo

proto.gen:
	mkdir -p pkg/api/v1
	protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 task.proto
	mkdir -p api/swagger/v1
	protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 task.proto
	protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 health.proto
	protoc --proto_path=api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:pkg/api/v1 task.proto
	protoc --proto_path=api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:pkg/api/v1 health.proto
	protoc --proto_path=api/proto/v1 --proto_path=third_party --swagger_out=logtostderr=true:api/swagger/v1 task.proto
	protoc --proto_path=api/proto/v1 --proto_path=third_party --swagger_out=logtostderr=true:api/swagger/v1 health.proto
adapter.gen:
	protoc --proto_path=api/adapter/model/v1 \
				--proto_path=third_party \
				--proto_path=api/ \
				--plugin=grpc \
				--go_out=paths=source_relative:api/adapter/model/v1 extensions.proto request.proto template.proto

adapter.config.gen:
	for adapter in tcp; do \
		protoc \
			--proto_path=adapters/$$adapter/config \
			-Ithird_party \
			--gogoslick_out=plugins=grpc,paths=source_relative:adapters/$$adapter/config \
			config.proto; \
	done

template.gen:
	protoc --proto_path=api/adapter/model/v1 \
		-I./third_party \
		--proto_path=api/ \
		--proto_path=templates/synthetic \
		--go_out=plugins=grpc:templates/synthetic synthetic_handler_service.proto


run.agent:
	go build -o echo -v github.com/covarity/echo/cmd/echo
	./echo tcp --d

test.agent:
	curl -XPOST localhost:3001/v1/task -d '{"task": { "protocol": "HTTP"}}'
debug:
	cd cmd/echo; dlv debug