// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: synthetic_handler_service.proto

package synthetic

import (
	context "context"
	v1 "github.com/covarity/echo/api/adapter/model/v1"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Request message for HandleSynthetic method.
type HandleSyntheticRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Adapter specific handler configuration.
	//
	AdapterConfig *any.Any `protobuf:"bytes,2,opt,name=adapter_config,json=adapterConfig,proto3" json:"adapter_config,omitempty"`
}

func (x *HandleSyntheticRequest) Reset() {
	*x = HandleSyntheticRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synthetic_handler_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleSyntheticRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleSyntheticRequest) ProtoMessage() {}

func (x *HandleSyntheticRequest) ProtoReflect() protoreflect.Message {
	mi := &file_synthetic_handler_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleSyntheticRequest.ProtoReflect.Descriptor instead.
func (*HandleSyntheticRequest) Descriptor() ([]byte, []int) {
	return file_synthetic_handler_service_proto_rawDescGZIP(), []int{0}
}

func (x *HandleSyntheticRequest) GetAdapterConfig() *any.Any {
	if x != nil {
		return x.AdapterConfig
	}
	return nil
}

var File_synthetic_handler_service_proto protoreflect.FileDescriptor

var file_synthetic_handler_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x5f, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x73, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x1a, 0x21, 0x61, 0x64,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x67, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x16, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x53,
	0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3b, 0x0a, 0x0e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0d, 0x61,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x6f, 0x0a, 0x16,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0f, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x12, 0x21, 0x2e, 0x73, 0x79, 0x6e, 0x74,
	0x68, 0x65, 0x74, 0x69, 0x63, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x53, 0x79, 0x6e, 0x74,
	0x68, 0x65, 0x74, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x21, 0xf8,
	0xd2, 0xe4, 0x93, 0x02, 0x00, 0x82, 0xdd, 0xe4, 0x93, 0x02, 0x09, 0x73, 0x79, 0x6e, 0x74, 0x68,
	0x65, 0x74, 0x69, 0x63, 0xc8, 0xe1, 0x1e, 0x00, 0xa8, 0xe2, 0x1e, 0x00, 0xf0, 0xe1, 0x1e, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synthetic_handler_service_proto_rawDescOnce sync.Once
	file_synthetic_handler_service_proto_rawDescData = file_synthetic_handler_service_proto_rawDesc
)

func file_synthetic_handler_service_proto_rawDescGZIP() []byte {
	file_synthetic_handler_service_proto_rawDescOnce.Do(func() {
		file_synthetic_handler_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_synthetic_handler_service_proto_rawDescData)
	})
	return file_synthetic_handler_service_proto_rawDescData
}

var file_synthetic_handler_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_synthetic_handler_service_proto_goTypes = []interface{}{
	(*HandleSyntheticRequest)(nil), // 0: synthetic.HandleSyntheticRequest
	(*any.Any)(nil),                // 1: google.protobuf.Any
	(*v1.RequestResult)(nil),       // 2: adapter.model.v1.RequestResult
}
var file_synthetic_handler_service_proto_depIdxs = []int32{
	1, // 0: synthetic.HandleSyntheticRequest.adapter_config:type_name -> google.protobuf.Any
	0, // 1: synthetic.HandleSyntheticService.HandleSynthetic:input_type -> synthetic.HandleSyntheticRequest
	2, // 2: synthetic.HandleSyntheticService.HandleSynthetic:output_type -> adapter.model.v1.RequestResult
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_synthetic_handler_service_proto_init() }
func file_synthetic_handler_service_proto_init() {
	if File_synthetic_handler_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_synthetic_handler_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleSyntheticRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_synthetic_handler_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_synthetic_handler_service_proto_goTypes,
		DependencyIndexes: file_synthetic_handler_service_proto_depIdxs,
		MessageInfos:      file_synthetic_handler_service_proto_msgTypes,
	}.Build()
	File_synthetic_handler_service_proto = out.File
	file_synthetic_handler_service_proto_rawDesc = nil
	file_synthetic_handler_service_proto_goTypes = nil
	file_synthetic_handler_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HandleSyntheticServiceClient is the client API for HandleSyntheticService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HandleSyntheticServiceClient interface {
	// HandleTraceSpan is called by Mixer at request-time to deliver 'tracespan' instances to the backend.
	HandleSynthetic(ctx context.Context, in *HandleSyntheticRequest, opts ...grpc.CallOption) (*v1.RequestResult, error)
}

type handleSyntheticServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHandleSyntheticServiceClient(cc grpc.ClientConnInterface) HandleSyntheticServiceClient {
	return &handleSyntheticServiceClient{cc}
}

func (c *handleSyntheticServiceClient) HandleSynthetic(ctx context.Context, in *HandleSyntheticRequest, opts ...grpc.CallOption) (*v1.RequestResult, error) {
	out := new(v1.RequestResult)
	err := c.cc.Invoke(ctx, "/synthetic.HandleSyntheticService/HandleSynthetic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HandleSyntheticServiceServer is the server API for HandleSyntheticService service.
type HandleSyntheticServiceServer interface {
	// HandleTraceSpan is called by Mixer at request-time to deliver 'tracespan' instances to the backend.
	HandleSynthetic(context.Context, *HandleSyntheticRequest) (*v1.RequestResult, error)
}

// UnimplementedHandleSyntheticServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHandleSyntheticServiceServer struct {
}

func (*UnimplementedHandleSyntheticServiceServer) HandleSynthetic(context.Context, *HandleSyntheticRequest) (*v1.RequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleSynthetic not implemented")
}

func RegisterHandleSyntheticServiceServer(s *grpc.Server, srv HandleSyntheticServiceServer) {
	s.RegisterService(&_HandleSyntheticService_serviceDesc, srv)
}

func _HandleSyntheticService_HandleSynthetic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleSyntheticRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandleSyntheticServiceServer).HandleSynthetic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/synthetic.HandleSyntheticService/HandleSynthetic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandleSyntheticServiceServer).HandleSynthetic(ctx, req.(*HandleSyntheticRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HandleSyntheticService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "synthetic.HandleSyntheticService",
	HandlerType: (*HandleSyntheticServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleSynthetic",
			Handler:    _HandleSyntheticService_HandleSynthetic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "synthetic_handler_service.proto",
}
