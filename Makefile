

build:
	go build -o echo -v github.com/covarity/echo/cmd/echo 

test:
	go test ./...

benchmark:
	go test -bench=. -benchmem -count=10 ./... > bench.txt
	benchstat bench.txt

proto.install:
	$(eval grpcGatewayVersion := $(shell cat go.mod| grep -E -o 'grpc-ecosystem/grpc-gateway .*\b'| grep -E -o 'v[0-9]+.[0-9]+.[0-9]+'))
	cp -r ${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@$(grpcGatewayVersion)/third_party/googleapis/google third_party/
	mkdir -p third_party/protoc-gen-swagger/options
	cp ${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@$(grpcGatewayVersion)/protoc-gen-swagger/options/annotations.proto third_party/protoc-gen-swagger/options
	cp ${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@$(grpcGatewayVersion)/protoc-gen-swagger/options/openapiv2.proto third_party/protoc-gen-swagger/options
	mkdir -p third_party/gogo 
	$(eval protobufVersion := $(shell cat go.mod| grep -E -o 'gogo/protobuf .*\b'| grep -E -o 'v[0-9]+.[0-9]+.[0-9]+'))
	cp -r ${GOPATH}/pkg/mod/github.com/gogo/protobuf@$(protobufVersion)/gogoproto third_party/gogo

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
	
policy.gen:
	protoc --proto_path=api/policy/v1alpha1/ \
				--proto_path=third_party \
				--proto_path=api/ \
				--plugin=grpc \
				--go_out=paths=source_relative:api/policy/v1alpha1 cfg.proto type.proto
	protoc --proto_path=api/policy/v1alpha1/ \
				--proto_path=third_party \
				--proto_path=api/ \
				--plugin=grpc \
				--go_out=paths=source_relative:api/policy/v1alpha1 value_type.proto
adapter.gen:
	protoc --proto_path=api/adapter/model/v1 \
				--proto_path=third_party \
				--proto_path=api/ \
				--plugin=grpc \
				--go_out=paths=source_relative:api/adapter/model/v1 extensions.proto request.proto template.proto

adapter.config.gen:
	for adapter in tcp http; do \
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