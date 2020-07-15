// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.12.3
// source: route.proto

package route

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// RouteRequest defines the request for a GetRoute call
type RouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// StartingCity is the starting point of the route
	StartCity string `protobuf:"bytes,1,opt,name=start_city,json=startCity,proto3" json:"start_city,omitempty"`
	// List of cities contains all the cities the traveler must visit
	CitiesList []string `protobuf:"bytes,2,rep,name=cities_list,json=citiesList,proto3" json:"cities_list,omitempty"`
}

func (x *RouteRequest) Reset() {
	*x = RouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteRequest) ProtoMessage() {}

func (x *RouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteRequest.ProtoReflect.Descriptor instead.
func (*RouteRequest) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{0}
}

func (x *RouteRequest) GetStartCity() string {
	if x != nil {
		return x.StartCity
	}
	return ""
}

func (x *RouteRequest) GetCitiesList() []string {
	if x != nil {
		return x.CitiesList
	}
	return nil
}

// RouteResponse is the response from a GetRoute call, it contains
// path from starting point through all the cities in the list.
type RouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routes []*RouteResponse_RoutePath `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	Error  string                     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RouteResponse) Reset() {
	*x = RouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteResponse) ProtoMessage() {}

func (x *RouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteResponse.ProtoReflect.Descriptor instead.
func (*RouteResponse) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{1}
}

func (x *RouteResponse) GetRoutes() []*RouteResponse_RoutePath {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *RouteResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type RouteResponse_RoutePath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
}

func (x *RouteResponse_RoutePath) Reset() {
	*x = RouteResponse_RoutePath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteResponse_RoutePath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteResponse_RoutePath) ProtoMessage() {}

func (x *RouteResponse_RoutePath) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteResponse_RoutePath.ProtoReflect.Descriptor instead.
func (*RouteResponse_RoutePath) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{1, 0}
}

func (x *RouteResponse_RoutePath) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

var File_route_proto protoreflect.FileDescriptor

var file_route_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a,
	0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x78, 0x0a,
	0x0d, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x1f, 0x0a, 0x09, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x32, 0x32, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x12, 0x29, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_route_proto_rawDescOnce sync.Once
	file_route_proto_rawDescData = file_route_proto_rawDesc
)

func file_route_proto_rawDescGZIP() []byte {
	file_route_proto_rawDescOnce.Do(func() {
		file_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_route_proto_rawDescData)
	})
	return file_route_proto_rawDescData
}

var file_route_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_route_proto_goTypes = []interface{}{
	(*RouteRequest)(nil),            // 0: RouteRequest
	(*RouteResponse)(nil),           // 1: RouteResponse
	(*RouteResponse_RoutePath)(nil), // 2: RouteResponse.RoutePath
}
var file_route_proto_depIdxs = []int32{
	2, // 0: RouteResponse.routes:type_name -> RouteResponse.RoutePath
	0, // 1: Route.GetRoute:input_type -> RouteRequest
	1, // 2: Route.GetRoute:output_type -> RouteResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_route_proto_init() }
func file_route_proto_init() {
	if File_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteRequest); i {
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
		file_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteResponse); i {
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
		file_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteResponse_RoutePath); i {
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
			RawDescriptor: file_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_route_proto_goTypes,
		DependencyIndexes: file_route_proto_depIdxs,
		MessageInfos:      file_route_proto_msgTypes,
	}.Build()
	File_route_proto = out.File
	file_route_proto_rawDesc = nil
	file_route_proto_goTypes = nil
	file_route_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RouteClient is the client API for Route service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouteClient interface {
	// GetRoute returns the route for the provided slice of cities with starting city
	GetRoute(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error)
}

type routeClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteClient(cc grpc.ClientConnInterface) RouteClient {
	return &routeClient{cc}
}

func (c *routeClient) GetRoute(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error) {
	out := new(RouteResponse)
	err := c.cc.Invoke(ctx, "/Route/GetRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteServer is the server API for Route service.
type RouteServer interface {
	// GetRoute returns the route for the provided slice of cities with starting city
	GetRoute(context.Context, *RouteRequest) (*RouteResponse, error)
}

// UnimplementedRouteServer can be embedded to have forward compatible implementations.
type UnimplementedRouteServer struct {
}

func (*UnimplementedRouteServer) GetRoute(context.Context, *RouteRequest) (*RouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoute not implemented")
}

func RegisterRouteServer(s *grpc.Server, srv RouteServer) {
	s.RegisterService(&_Route_serviceDesc, srv)
}

func _Route_GetRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServer).GetRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Route/GetRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServer).GetRoute(ctx, req.(*RouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Route_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Route",
	HandlerType: (*RouteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRoute",
			Handler:    _Route_GetRoute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "route.proto",
}