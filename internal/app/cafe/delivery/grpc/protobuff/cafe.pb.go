// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: cafe.proto

package cafe

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cafe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_cafe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_cafe_proto_rawDescGZIP(), []int{0}
}

func (x *ID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Cafe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CafeID      int64                `protobuf:"varint,1,opt,name=CafeID,proto3" json:"CafeID,omitempty"`
	CafeName    string               `protobuf:"bytes,2,opt,name=CafeName,proto3" json:"CafeName,omitempty"`
	Address     string               `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	Description string               `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	StaffID     int64                `protobuf:"varint,5,opt,name=StaffID,proto3" json:"StaffID,omitempty"`
	OpenTime    *timestamp.Timestamp `protobuf:"bytes,6,opt,name=OpenTime,proto3" json:"OpenTime,omitempty"`
	CloseTime   *timestamp.Timestamp `protobuf:"bytes,7,opt,name=CloseTime,proto3" json:"CloseTime,omitempty"`
	Photo       string               `protobuf:"bytes,8,opt,name=Photo,proto3" json:"Photo,omitempty"`
}

func (x *Cafe) Reset() {
	*x = Cafe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cafe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cafe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cafe) ProtoMessage() {}

func (x *Cafe) ProtoReflect() protoreflect.Message {
	mi := &file_cafe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cafe.ProtoReflect.Descriptor instead.
func (*Cafe) Descriptor() ([]byte, []int) {
	return file_cafe_proto_rawDescGZIP(), []int{1}
}

func (x *Cafe) GetCafeID() int64 {
	if x != nil {
		return x.CafeID
	}
	return 0
}

func (x *Cafe) GetCafeName() string {
	if x != nil {
		return x.CafeName
	}
	return ""
}

func (x *Cafe) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Cafe) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Cafe) GetStaffID() int64 {
	if x != nil {
		return x.StaffID
	}
	return 0
}

func (x *Cafe) GetOpenTime() *timestamp.Timestamp {
	if x != nil {
		return x.OpenTime
	}
	return nil
}

func (x *Cafe) GetCloseTime() *timestamp.Timestamp {
	if x != nil {
		return x.CloseTime
	}
	return nil
}

func (x *Cafe) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

var File_cafe_proto protoreflect.FileDescriptor

var file_cafe_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x66, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x98, 0x02,
	0x0a, 0x04, 0x43, 0x61, 0x66, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x61, 0x66, 0x65, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x43, 0x61, 0x66, 0x65, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x43, 0x61, 0x66, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x43, 0x61, 0x66, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x66, 0x66, 0x49,
	0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x74, 0x61, 0x66, 0x66, 0x49, 0x44,
	0x12, 0x36, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x4f, 0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x32, 0x3a, 0x0a, 0x0f, 0x43, 0x61, 0x66, 0x65,
	0x47, 0x52, 0x50, 0x43, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0c, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2e, 0x49, 0x44, 0x1a, 0x0e, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x43, 0x61, 0x66, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x63, 0x61, 0x66, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cafe_proto_rawDescOnce sync.Once
	file_cafe_proto_rawDescData = file_cafe_proto_rawDesc
)

func file_cafe_proto_rawDescGZIP() []byte {
	file_cafe_proto_rawDescOnce.Do(func() {
		file_cafe_proto_rawDescData = protoimpl.X.CompressGZIP(file_cafe_proto_rawDescData)
	})
	return file_cafe_proto_rawDescData
}

var file_cafe_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cafe_proto_goTypes = []interface{}{
	(*ID)(nil),                  // 0: customer.ID
	(*Cafe)(nil),                // 1: customer.Cafe
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_cafe_proto_depIdxs = []int32{
	2, // 0: customer.Cafe.OpenTime:type_name -> google.protobuf.Timestamp
	2, // 1: customer.Cafe.CloseTime:type_name -> google.protobuf.Timestamp
	0, // 2: customer.CafeGRPCHandler.GetByID:input_type -> customer.ID
	1, // 3: customer.CafeGRPCHandler.GetByID:output_type -> customer.Cafe
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cafe_proto_init() }
func file_cafe_proto_init() {
	if File_cafe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cafe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_cafe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cafe); i {
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
			RawDescriptor: file_cafe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cafe_proto_goTypes,
		DependencyIndexes: file_cafe_proto_depIdxs,
		MessageInfos:      file_cafe_proto_msgTypes,
	}.Build()
	File_cafe_proto = out.File
	file_cafe_proto_rawDesc = nil
	file_cafe_proto_goTypes = nil
	file_cafe_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CafeGRPCHandlerClient is the client API for CafeGRPCHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CafeGRPCHandlerClient interface {
	GetByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Cafe, error)
}

type cafeGRPCHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewCafeGRPCHandlerClient(cc grpc.ClientConnInterface) CafeGRPCHandlerClient {
	return &cafeGRPCHandlerClient{cc}
}

func (c *cafeGRPCHandlerClient) GetByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Cafe, error) {
	out := new(Cafe)
	err := c.cc.Invoke(ctx, "/customer.CafeGRPCHandler/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CafeGRPCHandlerServer is the server API for CafeGRPCHandler service.
type CafeGRPCHandlerServer interface {
	GetByID(context.Context, *ID) (*Cafe, error)
}

// UnimplementedCafeGRPCHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedCafeGRPCHandlerServer struct {
}

func (*UnimplementedCafeGRPCHandlerServer) GetByID(context.Context, *ID) (*Cafe, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}

func RegisterCafeGRPCHandlerServer(s *grpc.Server, srv CafeGRPCHandlerServer) {
	s.RegisterService(&_CafeGRPCHandler_serviceDesc, srv)
}

func _CafeGRPCHandler_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CafeGRPCHandlerServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.CafeGRPCHandler/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CafeGRPCHandlerServer).GetByID(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _CafeGRPCHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customer.CafeGRPCHandler",
	HandlerType: (*CafeGRPCHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByID",
			Handler:    _CafeGRPCHandler_GetByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cafe.proto",
}
