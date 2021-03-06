// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: tag.proto

package tag

import (
	context "context"
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

type TagModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TagModel) Reset() {
	*x = TagModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagModel) ProtoMessage() {}

func (x *TagModel) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagModel.ProtoReflect.Descriptor instead.
func (*TagModel) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{0}
}

func (x *TagModel) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TagModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetTagListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTagListReq) Reset() {
	*x = GetTagListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagListReq) ProtoMessage() {}

func (x *GetTagListReq) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagListReq.ProtoReflect.Descriptor instead.
func (*GetTagListReq) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{1}
}

type GetTagListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []*TagModel `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *GetTagListResp) Reset() {
	*x = GetTagListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagListResp) ProtoMessage() {}

func (x *GetTagListResp) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagListResp.ProtoReflect.Descriptor instead.
func (*GetTagListResp) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{2}
}

func (x *GetTagListResp) GetTags() []*TagModel {
	if x != nil {
		return x.Tags
	}
	return nil
}

type GetTagInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetTagInfoReq) Reset() {
	*x = GetTagInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagInfoReq) ProtoMessage() {}

func (x *GetTagInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagInfoReq.ProtoReflect.Descriptor instead.
func (*GetTagInfoReq) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{3}
}

func (x *GetTagInfoReq) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetTagInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag []*TagModel `protobuf:"bytes,1,rep,name=tag,proto3" json:"tag,omitempty"`
}

func (x *GetTagInfoResp) Reset() {
	*x = GetTagInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagInfoResp) ProtoMessage() {}

func (x *GetTagInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagInfoResp.ProtoReflect.Descriptor instead.
func (*GetTagInfoResp) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{4}
}

func (x *GetTagInfoResp) GetTag() []*TagModel {
	if x != nil {
		return x.Tag
	}
	return nil
}

var File_tag_proto protoreflect.FileDescriptor

var file_tag_proto_rawDesc = []byte{
	0x0a, 0x09, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x74, 0x61, 0x67,
	0x22, 0x2e, 0x0a, 0x08, 0x54, 0x61, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x22, 0x33, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x21, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x31, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x03, 0x74,
	0x61, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54,
	0x61, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x03, 0x74, 0x61, 0x67, 0x32, 0x73, 0x0a, 0x03,
	0x74, 0x61, 0x67, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x12, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x74,
	0x61, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tag_proto_rawDescOnce sync.Once
	file_tag_proto_rawDescData = file_tag_proto_rawDesc
)

func file_tag_proto_rawDescGZIP() []byte {
	file_tag_proto_rawDescOnce.Do(func() {
		file_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_tag_proto_rawDescData)
	})
	return file_tag_proto_rawDescData
}

var file_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tag_proto_goTypes = []interface{}{
	(*TagModel)(nil),       // 0: tag.TagModel
	(*GetTagListReq)(nil),  // 1: tag.GetTagListReq
	(*GetTagListResp)(nil), // 2: tag.GetTagListResp
	(*GetTagInfoReq)(nil),  // 3: tag.GetTagInfoReq
	(*GetTagInfoResp)(nil), // 4: tag.GetTagInfoResp
}
var file_tag_proto_depIdxs = []int32{
	0, // 0: tag.GetTagListResp.tags:type_name -> tag.TagModel
	0, // 1: tag.GetTagInfoResp.tag:type_name -> tag.TagModel
	1, // 2: tag.tag.GetTagList:input_type -> tag.GetTagListReq
	3, // 3: tag.tag.GetTagInfo:input_type -> tag.GetTagInfoReq
	2, // 4: tag.tag.GetTagList:output_type -> tag.GetTagListResp
	4, // 5: tag.tag.GetTagInfo:output_type -> tag.GetTagInfoResp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tag_proto_init() }
func file_tag_proto_init() {
	if File_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagModel); i {
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
		file_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagListReq); i {
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
		file_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagListResp); i {
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
		file_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagInfoReq); i {
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
		file_tag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagInfoResp); i {
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
			RawDescriptor: file_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tag_proto_goTypes,
		DependencyIndexes: file_tag_proto_depIdxs,
		MessageInfos:      file_tag_proto_msgTypes,
	}.Build()
	File_tag_proto = out.File
	file_tag_proto_rawDesc = nil
	file_tag_proto_goTypes = nil
	file_tag_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TagClient is the client API for Tag service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TagClient interface {
	GetTagList(ctx context.Context, in *GetTagListReq, opts ...grpc.CallOption) (*GetTagListResp, error)
	GetTagInfo(ctx context.Context, in *GetTagInfoReq, opts ...grpc.CallOption) (*GetTagInfoResp, error)
}

type tagClient struct {
	cc grpc.ClientConnInterface
}

func NewTagClient(cc grpc.ClientConnInterface) TagClient {
	return &tagClient{cc}
}

func (c *tagClient) GetTagList(ctx context.Context, in *GetTagListReq, opts ...grpc.CallOption) (*GetTagListResp, error) {
	out := new(GetTagListResp)
	err := c.cc.Invoke(ctx, "/tag.tag/GetTagList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagClient) GetTagInfo(ctx context.Context, in *GetTagInfoReq, opts ...grpc.CallOption) (*GetTagInfoResp, error) {
	out := new(GetTagInfoResp)
	err := c.cc.Invoke(ctx, "/tag.tag/GetTagInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServer is the server API for Tag service.
type TagServer interface {
	GetTagList(context.Context, *GetTagListReq) (*GetTagListResp, error)
	GetTagInfo(context.Context, *GetTagInfoReq) (*GetTagInfoResp, error)
}

// UnimplementedTagServer can be embedded to have forward compatible implementations.
type UnimplementedTagServer struct {
}

func (*UnimplementedTagServer) GetTagList(context.Context, *GetTagListReq) (*GetTagListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagList not implemented")
}
func (*UnimplementedTagServer) GetTagInfo(context.Context, *GetTagInfoReq) (*GetTagInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagInfo not implemented")
}

func RegisterTagServer(s *grpc.Server, srv TagServer) {
	s.RegisterService(&_Tag_serviceDesc, srv)
}

func _Tag_GetTagList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServer).GetTagList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tag.tag/GetTagList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServer).GetTagList(ctx, req.(*GetTagListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tag_GetTagInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServer).GetTagInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tag.tag/GetTagInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServer).GetTagInfo(ctx, req.(*GetTagInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tag_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tag.tag",
	HandlerType: (*TagServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTagList",
			Handler:    _Tag_GetTagList_Handler,
		},
		{
			MethodName: "GetTagInfo",
			Handler:    _Tag_GetTagInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tag.proto",
}
