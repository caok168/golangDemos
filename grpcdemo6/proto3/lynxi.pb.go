// author: samuel ko

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.6.0
// source: lynxi.proto

package lynxi

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

type TaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId       uint32 `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"`            //任务编号
	PipelineName string `protobuf:"bytes,2,opt,name=pipelineName,proto3" json:"pipelineName,omitempty"` //pipeline名称
	Uuid         string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`                 //用户uuid
	Rawinput     []byte `protobuf:"bytes,4,opt,name=rawinput,proto3" json:"rawinput,omitempty"`         //原始数据
}

func (x *TaskRequest) Reset() {
	*x = TaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lynxi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRequest) ProtoMessage() {}

func (x *TaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lynxi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRequest.ProtoReflect.Descriptor instead.
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return file_lynxi_proto_rawDescGZIP(), []int{0}
}

func (x *TaskRequest) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *TaskRequest) GetPipelineName() string {
	if x != nil {
		return x.PipelineName
	}
	return ""
}

func (x *TaskRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *TaskRequest) GetRawinput() []byte {
	if x != nil {
		return x.Rawinput
	}
	return nil
}

type ImageFeatureExtractReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid       string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` //用户uuid
	Rawbboxraw *BboxRaw `protobuf:"bytes,2,opt,name=Rawbboxraw,proto3" json:"Rawbboxraw,omitempty"`
	Feature    []string `protobuf:"bytes,3,rep,name=feature,proto3" json:"feature,omitempty"` //String形式的人脸特征
}

func (x *ImageFeatureExtractReply) Reset() {
	*x = ImageFeatureExtractReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lynxi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageFeatureExtractReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageFeatureExtractReply) ProtoMessage() {}

func (x *ImageFeatureExtractReply) ProtoReflect() protoreflect.Message {
	mi := &file_lynxi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageFeatureExtractReply.ProtoReflect.Descriptor instead.
func (*ImageFeatureExtractReply) Descriptor() ([]byte, []int) {
	return file_lynxi_proto_rawDescGZIP(), []int{1}
}

func (x *ImageFeatureExtractReply) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ImageFeatureExtractReply) GetRawbboxraw() *BboxRaw {
	if x != nil {
		return x.Rawbboxraw
	}
	return nil
}

func (x *ImageFeatureExtractReply) GetFeature() []string {
	if x != nil {
		return x.Feature
	}
	return nil
}

type BboxRaw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"` //矩形框左上角x值
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"` //矩形框左上角y值
	H int32 `protobuf:"varint,3,opt,name=h,proto3" json:"h,omitempty"` //矩形框高度
	W int32 `protobuf:"varint,4,opt,name=w,proto3" json:"w,omitempty"` //矩形框宽度
}

func (x *BboxRaw) Reset() {
	*x = BboxRaw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lynxi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BboxRaw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BboxRaw) ProtoMessage() {}

func (x *BboxRaw) ProtoReflect() protoreflect.Message {
	mi := &file_lynxi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BboxRaw.ProtoReflect.Descriptor instead.
func (*BboxRaw) Descriptor() ([]byte, []int) {
	return file_lynxi_proto_rawDescGZIP(), []int{2}
}

func (x *BboxRaw) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *BboxRaw) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *BboxRaw) GetH() int32 {
	if x != nil {
		return x.H
	}
	return 0
}

func (x *BboxRaw) GetW() int32 {
	if x != nil {
		return x.W
	}
	return 0
}

var File_lynxi_proto protoreflect.FileDescriptor

var file_lynxi_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x79, 0x6e, 0x78, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c,
	0x79, 0x6e, 0x78, 0x69, 0x22, 0x79, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x61, 0x77, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22,
	0x78, 0x0a, 0x18, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x2e, 0x0a, 0x0a, 0x52, 0x61, 0x77, 0x62, 0x62, 0x6f, 0x78, 0x72, 0x61, 0x77, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x79, 0x6e, 0x78, 0x69, 0x2e, 0x42, 0x62, 0x6f, 0x78,
	0x52, 0x61, 0x77, 0x52, 0x0a, 0x52, 0x61, 0x77, 0x62, 0x62, 0x6f, 0x78, 0x72, 0x61, 0x77, 0x12,
	0x18, 0x0a, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x41, 0x0a, 0x07, 0x42, 0x62, 0x6f,
	0x78, 0x52, 0x61, 0x77, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79,
	0x12, 0x0c, 0x0a, 0x01, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x68, 0x12, 0x0c,
	0x0a, 0x01, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x77, 0x32, 0x49, 0x0a, 0x04,
	0x46, 0x61, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x46, 0x61, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x2e, 0x6c, 0x79, 0x6e, 0x78, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6c, 0x79, 0x6e, 0x78, 0x69, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lynxi_proto_rawDescOnce sync.Once
	file_lynxi_proto_rawDescData = file_lynxi_proto_rawDesc
)

func file_lynxi_proto_rawDescGZIP() []byte {
	file_lynxi_proto_rawDescOnce.Do(func() {
		file_lynxi_proto_rawDescData = protoimpl.X.CompressGZIP(file_lynxi_proto_rawDescData)
	})
	return file_lynxi_proto_rawDescData
}

var file_lynxi_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_lynxi_proto_goTypes = []interface{}{
	(*TaskRequest)(nil),              // 0: lynxi.TaskRequest
	(*ImageFeatureExtractReply)(nil), // 1: lynxi.ImageFeatureExtractReply
	(*BboxRaw)(nil),                  // 2: lynxi.BboxRaw
}
var file_lynxi_proto_depIdxs = []int32{
	2, // 0: lynxi.ImageFeatureExtractReply.Rawbboxraw:type_name -> lynxi.BboxRaw
	0, // 1: lynxi.Face.FaceData:input_type -> lynxi.TaskRequest
	1, // 2: lynxi.Face.FaceData:output_type -> lynxi.ImageFeatureExtractReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_lynxi_proto_init() }
func file_lynxi_proto_init() {
	if File_lynxi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lynxi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRequest); i {
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
		file_lynxi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageFeatureExtractReply); i {
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
		file_lynxi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BboxRaw); i {
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
			RawDescriptor: file_lynxi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lynxi_proto_goTypes,
		DependencyIndexes: file_lynxi_proto_depIdxs,
		MessageInfos:      file_lynxi_proto_msgTypes,
	}.Build()
	File_lynxi_proto = out.File
	file_lynxi_proto_rawDesc = nil
	file_lynxi_proto_goTypes = nil
	file_lynxi_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FaceClient is the client API for Face service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FaceClient interface {
	// 最基本的形式.
	FaceData(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*ImageFeatureExtractReply, error)
}

type faceClient struct {
	cc grpc.ClientConnInterface
}

func NewFaceClient(cc grpc.ClientConnInterface) FaceClient {
	return &faceClient{cc}
}

func (c *faceClient) FaceData(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*ImageFeatureExtractReply, error) {
	out := new(ImageFeatureExtractReply)
	err := c.cc.Invoke(ctx, "/lynxi.Face/FaceData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FaceServer is the server API for Face service.
type FaceServer interface {
	// 最基本的形式.
	FaceData(context.Context, *TaskRequest) (*ImageFeatureExtractReply, error)
}

// UnimplementedFaceServer can be embedded to have forward compatible implementations.
type UnimplementedFaceServer struct {
}

func (*UnimplementedFaceServer) FaceData(context.Context, *TaskRequest) (*ImageFeatureExtractReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FaceData not implemented")
}

func RegisterFaceServer(s *grpc.Server, srv FaceServer) {
	s.RegisterService(&_Face_serviceDesc, srv)
}

func _Face_FaceData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaceServer).FaceData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lynxi.Face/FaceData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaceServer).FaceData(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Face_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lynxi.Face",
	HandlerType: (*FaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FaceData",
			Handler:    _Face_FaceData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lynxi.proto",
}