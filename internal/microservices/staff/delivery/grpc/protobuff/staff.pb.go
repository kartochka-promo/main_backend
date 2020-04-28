//protoc --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: staff.proto

package staff

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{0}
}

type SafeStaff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffID  int64                `protobuf:"varint,1,opt,name=StaffID,proto3" json:"StaffID,omitempty"`
	Name     string               `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Email    string               `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	EditedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=EditedAt,proto3" json:"EditedAt,omitempty"`
	Photo    string               `protobuf:"bytes,5,opt,name=Photo,proto3" json:"Photo,omitempty"`
	IsOwner  bool                 `protobuf:"varint,6,opt,name=IsOwner,proto3" json:"IsOwner,omitempty"`
	CafeId   int64                `protobuf:"varint,7,opt,name=CafeId,proto3" json:"CafeId,omitempty"`
	Position string               `protobuf:"bytes,8,opt,name=Position,proto3" json:"Position,omitempty"`
}

func (x *SafeStaff) Reset() {
	*x = SafeStaff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SafeStaff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SafeStaff) ProtoMessage() {}

func (x *SafeStaff) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SafeStaff.ProtoReflect.Descriptor instead.
func (*SafeStaff) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{1}
}

func (x *SafeStaff) GetStaffID() int64 {
	if x != nil {
		return x.StaffID
	}
	return 0
}

func (x *SafeStaff) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SafeStaff) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SafeStaff) GetEditedAt() *timestamp.Timestamp {
	if x != nil {
		return x.EditedAt
	}
	return nil
}

func (x *SafeStaff) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *SafeStaff) GetIsOwner() bool {
	if x != nil {
		return x.IsOwner
	}
	return false
}

func (x *SafeStaff) GetCafeId() int64 {
	if x != nil {
		return x.CafeId
	}
	return 0
}

func (x *SafeStaff) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

var File_staff_proto protoreflect.FileDescriptor

var file_staff_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xeb,
	0x01, 0x0a, 0x09, 0x53, 0x61, 0x66, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x18, 0x0a, 0x07,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x36, 0x0a, 0x08, 0x45, 0x64, 0x69, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x45, 0x64, 0x69, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x74,
	0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x0a, 0x07, 0x49, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x49, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x61, 0x66, 0x65,
	0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x43, 0x61, 0x66, 0x65, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x44, 0x0a, 0x10,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x47, 0x52, 0x50, 0x43, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x12, 0x30, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x0c, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x10, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x53, 0x61, 0x66, 0x65, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x73, 0x74, 0x61, 0x66, 0x66, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_staff_proto_rawDescOnce sync.Once
	file_staff_proto_rawDescData = file_staff_proto_rawDesc
)

func file_staff_proto_rawDescGZIP() []byte {
	file_staff_proto_rawDescOnce.Do(func() {
		file_staff_proto_rawDescData = protoimpl.X.CompressGZIP(file_staff_proto_rawDescData)
	})
	return file_staff_proto_rawDescData
}

var file_staff_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_staff_proto_goTypes = []interface{}{
	(*Empty)(nil),               // 0: staff.Empty
	(*SafeStaff)(nil),           // 1: staff.SafeStaff
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_staff_proto_depIdxs = []int32{
	2, // 0: staff.SafeStaff.EditedAt:type_name -> google.protobuf.Timestamp
	0, // 1: staff.StaffGRPCHandler.GetFromSession:input_type -> staff.Empty
	1, // 2: staff.StaffGRPCHandler.GetFromSession:output_type -> staff.SafeStaff
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_staff_proto_init() }
func file_staff_proto_init() {
	if File_staff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_staff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_staff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SafeStaff); i {
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
			RawDescriptor: file_staff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_staff_proto_goTypes,
		DependencyIndexes: file_staff_proto_depIdxs,
		MessageInfos:      file_staff_proto_msgTypes,
	}.Build()
	File_staff_proto = out.File
	file_staff_proto_rawDesc = nil
	file_staff_proto_goTypes = nil
	file_staff_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StaffGRPCHandlerClient is the client API for StaffGRPCHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StaffGRPCHandlerClient interface {
	GetFromSession(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SafeStaff, error)
}

type staffGRPCHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewStaffGRPCHandlerClient(cc grpc.ClientConnInterface) StaffGRPCHandlerClient {
	return &staffGRPCHandlerClient{cc}
}

func (c *staffGRPCHandlerClient) GetFromSession(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SafeStaff, error) {
	out := new(SafeStaff)
	err := c.cc.Invoke(ctx, "/staff.StaffGRPCHandler/GetFromSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffGRPCHandlerServer is the server API for StaffGRPCHandler service.
type StaffGRPCHandlerServer interface {
	GetFromSession(context.Context, *Empty) (*SafeStaff, error)
}

// UnimplementedStaffGRPCHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedStaffGRPCHandlerServer struct {
}

func (*UnimplementedStaffGRPCHandlerServer) GetFromSession(context.Context, *Empty) (*SafeStaff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFromSession not implemented")
}

func RegisterStaffGRPCHandlerServer(s *grpc.Server, srv StaffGRPCHandlerServer) {
	s.RegisterService(&_StaffGRPCHandler_serviceDesc, srv)
}

func _StaffGRPCHandler_GetFromSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffGRPCHandlerServer).GetFromSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staff.StaffGRPCHandler/GetFromSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffGRPCHandlerServer).GetFromSession(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _StaffGRPCHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "staff.StaffGRPCHandler",
	HandlerType: (*StaffGRPCHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFromSession",
			Handler:    _StaffGRPCHandler_GetFromSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staff.proto",
}
