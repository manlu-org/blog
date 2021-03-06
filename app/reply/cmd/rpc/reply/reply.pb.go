// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: reply.proto

package reply

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

type GetReplyByArticleIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleId int64 `protobuf:"varint,1,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
}

func (x *GetReplyByArticleIdReq) Reset() {
	*x = GetReplyByArticleIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reply_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReplyByArticleIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReplyByArticleIdReq) ProtoMessage() {}

func (x *GetReplyByArticleIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_reply_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReplyByArticleIdReq.ProtoReflect.Descriptor instead.
func (*GetReplyByArticleIdReq) Descriptor() ([]byte, []int) {
	return file_reply_proto_rawDescGZIP(), []int{0}
}

func (x *GetReplyByArticleIdReq) GetArticleId() int64 {
	if x != nil {
		return x.ArticleId
	}
	return 0
}

type ReplyModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content   string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *ReplyModel) Reset() {
	*x = ReplyModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reply_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyModel) ProtoMessage() {}

func (x *ReplyModel) ProtoReflect() protoreflect.Message {
	mi := &file_reply_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyModel.ProtoReflect.Descriptor instead.
func (*ReplyModel) Descriptor() ([]byte, []int) {
	return file_reply_proto_rawDescGZIP(), []int{1}
}

func (x *ReplyModel) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReplyModel) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ReplyModel) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetReplyByArticleIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply []*ReplyModel `protobuf:"bytes,1,rep,name=reply,proto3" json:"reply,omitempty"`
}

func (x *GetReplyByArticleIdResp) Reset() {
	*x = GetReplyByArticleIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reply_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReplyByArticleIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReplyByArticleIdResp) ProtoMessage() {}

func (x *GetReplyByArticleIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_reply_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReplyByArticleIdResp.ProtoReflect.Descriptor instead.
func (*GetReplyByArticleIdResp) Descriptor() ([]byte, []int) {
	return file_reply_proto_rawDescGZIP(), []int{2}
}

func (x *GetReplyByArticleIdResp) GetReply() []*ReplyModel {
	if x != nil {
		return x.Reply
	}
	return nil
}

var File_reply_proto protoreflect.FileDescriptor

var file_reply_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x37, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x42, 0x79, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x55, 0x0a,
	0x0a, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x42, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x42, 0x79, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x27, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x5d, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x54, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x79, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x79, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x79, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reply_proto_rawDescOnce sync.Once
	file_reply_proto_rawDescData = file_reply_proto_rawDesc
)

func file_reply_proto_rawDescGZIP() []byte {
	file_reply_proto_rawDescOnce.Do(func() {
		file_reply_proto_rawDescData = protoimpl.X.CompressGZIP(file_reply_proto_rawDescData)
	})
	return file_reply_proto_rawDescData
}

var file_reply_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_reply_proto_goTypes = []interface{}{
	(*GetReplyByArticleIdReq)(nil),  // 0: reply.GetReplyByArticleIdReq
	(*ReplyModel)(nil),              // 1: reply.ReplyModel
	(*GetReplyByArticleIdResp)(nil), // 2: reply.GetReplyByArticleIdResp
}
var file_reply_proto_depIdxs = []int32{
	1, // 0: reply.GetReplyByArticleIdResp.reply:type_name -> reply.ReplyModel
	0, // 1: reply.reply.GetReplyByArticleId:input_type -> reply.GetReplyByArticleIdReq
	2, // 2: reply.reply.GetReplyByArticleId:output_type -> reply.GetReplyByArticleIdResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_reply_proto_init() }
func file_reply_proto_init() {
	if File_reply_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reply_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReplyByArticleIdReq); i {
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
		file_reply_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyModel); i {
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
		file_reply_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReplyByArticleIdResp); i {
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
			RawDescriptor: file_reply_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reply_proto_goTypes,
		DependencyIndexes: file_reply_proto_depIdxs,
		MessageInfos:      file_reply_proto_msgTypes,
	}.Build()
	File_reply_proto = out.File
	file_reply_proto_rawDesc = nil
	file_reply_proto_goTypes = nil
	file_reply_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReplyClient is the client API for Reply service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReplyClient interface {
	GetReplyByArticleId(ctx context.Context, in *GetReplyByArticleIdReq, opts ...grpc.CallOption) (*GetReplyByArticleIdResp, error)
}

type replyClient struct {
	cc grpc.ClientConnInterface
}

func NewReplyClient(cc grpc.ClientConnInterface) ReplyClient {
	return &replyClient{cc}
}

func (c *replyClient) GetReplyByArticleId(ctx context.Context, in *GetReplyByArticleIdReq, opts ...grpc.CallOption) (*GetReplyByArticleIdResp, error) {
	out := new(GetReplyByArticleIdResp)
	err := c.cc.Invoke(ctx, "/reply.reply/GetReplyByArticleId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplyServer is the server API for Reply service.
type ReplyServer interface {
	GetReplyByArticleId(context.Context, *GetReplyByArticleIdReq) (*GetReplyByArticleIdResp, error)
}

// UnimplementedReplyServer can be embedded to have forward compatible implementations.
type UnimplementedReplyServer struct {
}

func (*UnimplementedReplyServer) GetReplyByArticleId(context.Context, *GetReplyByArticleIdReq) (*GetReplyByArticleIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReplyByArticleId not implemented")
}

func RegisterReplyServer(s *grpc.Server, srv ReplyServer) {
	s.RegisterService(&_Reply_serviceDesc, srv)
}

func _Reply_GetReplyByArticleId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReplyByArticleIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplyServer).GetReplyByArticleId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reply.reply/GetReplyByArticleId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplyServer).GetReplyByArticleId(ctx, req.(*GetReplyByArticleIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Reply_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reply.reply",
	HandlerType: (*ReplyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReplyByArticleId",
			Handler:    _Reply_GetReplyByArticleId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reply.proto",
}
