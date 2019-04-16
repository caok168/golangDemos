// Code generated by protoc-gen-go. DO NOT EDIT.
// source: faceTask.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ReplyType int32

const (
	ReplyType_UNKNOWN    ReplyType = 0
	ReplyType_TASKRESULT ReplyType = 1
	ReplyType_TASKSTATUS ReplyType = 2
)

var ReplyType_name = map[int32]string{
	0: "UNKNOWN",
	1: "TASKRESULT",
	2: "TASKSTATUS",
}
var ReplyType_value = map[string]int32{
	"UNKNOWN":    0,
	"TASKRESULT": 1,
	"TASKSTATUS": 2,
}

func (x ReplyType) String() string {
	return proto.EnumName(ReplyType_name, int32(x))
}
func (ReplyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_faceTask_698c8d77d1d8704b, []int{0}
}

type FaceExtractRequest struct {
	TaskId               uint32   `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	DeviceDesc           uint32   `protobuf:"varint,2,opt,name=deviceDesc,proto3" json:"deviceDesc,omitempty"`
	SourceImagePath      string   `protobuf:"bytes,3,opt,name=sourceImagePath,proto3" json:"sourceImagePath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaceExtractRequest) Reset()         { *m = FaceExtractRequest{} }
func (m *FaceExtractRequest) String() string { return proto.CompactTextString(m) }
func (*FaceExtractRequest) ProtoMessage()    {}
func (*FaceExtractRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_faceTask_698c8d77d1d8704b, []int{0}
}
func (m *FaceExtractRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaceExtractRequest.Unmarshal(m, b)
}
func (m *FaceExtractRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaceExtractRequest.Marshal(b, m, deterministic)
}
func (dst *FaceExtractRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaceExtractRequest.Merge(dst, src)
}
func (m *FaceExtractRequest) XXX_Size() int {
	return xxx_messageInfo_FaceExtractRequest.Size(m)
}
func (m *FaceExtractRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FaceExtractRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FaceExtractRequest proto.InternalMessageInfo

func (m *FaceExtractRequest) GetTaskId() uint32 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *FaceExtractRequest) GetDeviceDesc() uint32 {
	if m != nil {
		return m.DeviceDesc
	}
	return 0
}

func (m *FaceExtractRequest) GetSourceImagePath() string {
	if m != nil {
		return m.SourceImagePath
	}
	return ""
}

type FaceExtractReply struct {
	TaskId uint32 `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	// Types that are valid to be assigned to StatusOrResult:
	//	*FaceExtractReply_Result
	//	*FaceExtractReply_Status
	StatusOrResult       isFaceExtractReply_StatusOrResult `protobuf_oneof:"statusOrResult"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *FaceExtractReply) Reset()         { *m = FaceExtractReply{} }
func (m *FaceExtractReply) String() string { return proto.CompactTextString(m) }
func (*FaceExtractReply) ProtoMessage()    {}
func (*FaceExtractReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_faceTask_698c8d77d1d8704b, []int{1}
}
func (m *FaceExtractReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaceExtractReply.Unmarshal(m, b)
}
func (m *FaceExtractReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaceExtractReply.Marshal(b, m, deterministic)
}
func (dst *FaceExtractReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaceExtractReply.Merge(dst, src)
}
func (m *FaceExtractReply) XXX_Size() int {
	return xxx_messageInfo_FaceExtractReply.Size(m)
}
func (m *FaceExtractReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FaceExtractReply.DiscardUnknown(m)
}

var xxx_messageInfo_FaceExtractReply proto.InternalMessageInfo

func (m *FaceExtractReply) GetTaskId() uint32 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

type isFaceExtractReply_StatusOrResult interface {
	isFaceExtractReply_StatusOrResult()
}

type FaceExtractReply_Result struct {
	Result *FaceExtractResult `protobuf:"bytes,2,opt,name=result,proto3,oneof"`
}

type FaceExtractReply_Status struct {
	Status *Status `protobuf:"bytes,3,opt,name=status,proto3,oneof"`
}

func (*FaceExtractReply_Result) isFaceExtractReply_StatusOrResult() {}

func (*FaceExtractReply_Status) isFaceExtractReply_StatusOrResult() {}

func (m *FaceExtractReply) GetStatusOrResult() isFaceExtractReply_StatusOrResult {
	if m != nil {
		return m.StatusOrResult
	}
	return nil
}

func (m *FaceExtractReply) GetResult() *FaceExtractResult {
	if x, ok := m.GetStatusOrResult().(*FaceExtractReply_Result); ok {
		return x.Result
	}
	return nil
}

func (m *FaceExtractReply) GetStatus() *Status {
	if x, ok := m.GetStatusOrResult().(*FaceExtractReply_Status); ok {
		return x.Status
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FaceExtractReply) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FaceExtractReply_OneofMarshaler, _FaceExtractReply_OneofUnmarshaler, _FaceExtractReply_OneofSizer, []interface{}{
		(*FaceExtractReply_Result)(nil),
		(*FaceExtractReply_Status)(nil),
	}
}

func _FaceExtractReply_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FaceExtractReply)
	// statusOrResult
	switch x := m.StatusOrResult.(type) {
	case *FaceExtractReply_Result:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Result); err != nil {
			return err
		}
	case *FaceExtractReply_Status:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Status); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("FaceExtractReply.StatusOrResult has unexpected type %T", x)
	}
	return nil
}

func _FaceExtractReply_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FaceExtractReply)
	switch tag {
	case 2: // statusOrResult.result
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FaceExtractResult)
		err := b.DecodeMessage(msg)
		m.StatusOrResult = &FaceExtractReply_Result{msg}
		return true, err
	case 3: // statusOrResult.status
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Status)
		err := b.DecodeMessage(msg)
		m.StatusOrResult = &FaceExtractReply_Status{msg}
		return true, err
	default:
		return false, nil
	}
}

func _FaceExtractReply_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FaceExtractReply)
	// statusOrResult
	switch x := m.StatusOrResult.(type) {
	case *FaceExtractReply_Result:
		s := proto.Size(x.Result)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FaceExtractReply_Status:
		s := proto.Size(x.Status)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type FaceExtractResult struct {
	Type                 ReplyType         `protobuf:"varint,1,opt,name=type,proto3,enum=proto.ReplyType" json:"type,omitempty"`
	Uuid                 string            `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Image                []byte            `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	ImageSize            uint32            `protobuf:"varint,4,opt,name=imageSize,proto3" json:"imageSize,omitempty"`
	ImageType            string            `protobuf:"bytes,5,opt,name=imageType,proto3" json:"imageType,omitempty"`
	Eigen                []byte            `protobuf:"bytes,6,opt,name=eigen,proto3" json:"eigen,omitempty"`
	EigenSize            uint32            `protobuf:"varint,7,opt,name=eigenSize,proto3" json:"eigenSize,omitempty"`
	Attrs                map[string]string `protobuf:"bytes,8,rep,name=attrs,proto3" json:"attrs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FaceExtractResult) Reset()         { *m = FaceExtractResult{} }
func (m *FaceExtractResult) String() string { return proto.CompactTextString(m) }
func (*FaceExtractResult) ProtoMessage()    {}
func (*FaceExtractResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_faceTask_698c8d77d1d8704b, []int{2}
}
func (m *FaceExtractResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaceExtractResult.Unmarshal(m, b)
}
func (m *FaceExtractResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaceExtractResult.Marshal(b, m, deterministic)
}
func (dst *FaceExtractResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaceExtractResult.Merge(dst, src)
}
func (m *FaceExtractResult) XXX_Size() int {
	return xxx_messageInfo_FaceExtractResult.Size(m)
}
func (m *FaceExtractResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FaceExtractResult.DiscardUnknown(m)
}

var xxx_messageInfo_FaceExtractResult proto.InternalMessageInfo

func (m *FaceExtractResult) GetType() ReplyType {
	if m != nil {
		return m.Type
	}
	return ReplyType_UNKNOWN
}

func (m *FaceExtractResult) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *FaceExtractResult) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *FaceExtractResult) GetImageSize() uint32 {
	if m != nil {
		return m.ImageSize
	}
	return 0
}

func (m *FaceExtractResult) GetImageType() string {
	if m != nil {
		return m.ImageType
	}
	return ""
}

func (m *FaceExtractResult) GetEigen() []byte {
	if m != nil {
		return m.Eigen
	}
	return nil
}

func (m *FaceExtractResult) GetEigenSize() uint32 {
	if m != nil {
		return m.EigenSize
	}
	return 0
}

func (m *FaceExtractResult) GetAttrs() map[string]string {
	if m != nil {
		return m.Attrs
	}
	return nil
}

type Status struct {
	State                string   `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Code                 uint32   `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_faceTask_698c8d77d1d8704b, []int{3}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (dst *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(dst, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Status) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*FaceExtractRequest)(nil), "proto.FaceExtractRequest")
	proto.RegisterType((*FaceExtractReply)(nil), "proto.FaceExtractReply")
	proto.RegisterType((*FaceExtractResult)(nil), "proto.FaceExtractResult")
	proto.RegisterMapType((map[string]string)(nil), "proto.FaceExtractResult.AttrsEntry")
	proto.RegisterType((*Status)(nil), "proto.Status")
	proto.RegisterEnum("proto.ReplyType", ReplyType_name, ReplyType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EngineClient is the client API for Engine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EngineClient interface {
	FaceExtract(ctx context.Context, in *FaceExtractRequest, opts ...grpc.CallOption) (*FaceExtractReply, error)
}

type engineClient struct {
	cc *grpc.ClientConn
}

func NewEngineClient(cc *grpc.ClientConn) EngineClient {
	return &engineClient{cc}
}

func (c *engineClient) FaceExtract(ctx context.Context, in *FaceExtractRequest, opts ...grpc.CallOption) (*FaceExtractReply, error) {
	out := new(FaceExtractReply)
	err := c.cc.Invoke(ctx, "/proto.Engine/FaceExtract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EngineServer is the server API for Engine service.
type EngineServer interface {
	FaceExtract(context.Context, *FaceExtractRequest) (*FaceExtractReply, error)
}

func RegisterEngineServer(s *grpc.Server, srv EngineServer) {
	s.RegisterService(&_Engine_serviceDesc, srv)
}

func _Engine_FaceExtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FaceExtractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).FaceExtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Engine/FaceExtract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).FaceExtract(ctx, req.(*FaceExtractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Engine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Engine",
	HandlerType: (*EngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FaceExtract",
			Handler:    _Engine_FaceExtract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "faceTask.proto",
}

func init() { proto.RegisterFile("faceTask.proto", fileDescriptor_faceTask_698c8d77d1d8704b) }

var fileDescriptor_faceTask_698c8d77d1d8704b = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0xf3, 0xe1, 0xe2, 0x31, 0x0d, 0x66, 0x84, 0xc0, 0x54, 0x08, 0x45, 0x06, 0x89, 0x88,
	0x43, 0x0e, 0xe6, 0x12, 0xb8, 0x05, 0x08, 0xa2, 0x2a, 0xa4, 0x68, 0xed, 0x88, 0xf3, 0xe2, 0x0c,
	0xc1, 0x4a, 0x48, 0x8c, 0x77, 0x1d, 0x61, 0xfe, 0x09, 0xff, 0x88, 0x9f, 0x85, 0x76, 0xd6, 0x24,
	0x2d, 0x6d, 0x4f, 0x9e, 0x79, 0x6f, 0xe6, 0xed, 0xbc, 0x27, 0x43, 0xff, 0xab, 0xcc, 0x28, 0x95,
	0x6a, 0x35, 0x2a, 0xca, 0xad, 0xde, 0x62, 0x8f, 0x3f, 0xd1, 0x0e, 0xf0, 0x9d, 0xcc, 0x68, 0xfa,
	0x53, 0x97, 0x32, 0xd3, 0x82, 0x7e, 0x54, 0xa4, 0x34, 0xde, 0x07, 0x57, 0x4b, 0xb5, 0x3a, 0x5d,
	0x84, 0xce, 0xc0, 0x19, 0x1e, 0x8b, 0xa6, 0xc3, 0xc7, 0x00, 0x0b, 0xda, 0xe5, 0x19, 0xbd, 0x25,
	0x95, 0x85, 0x6d, 0xe6, 0x2e, 0x20, 0x38, 0x84, 0x3b, 0x6a, 0x5b, 0x95, 0x19, 0x9d, 0x7e, 0x97,
	0x4b, 0xfa, 0x24, 0xf5, 0xb7, 0xb0, 0x33, 0x70, 0x86, 0x9e, 0xf8, 0x1f, 0x8e, 0x7e, 0x3b, 0x10,
	0x5c, 0x7a, 0xb8, 0x58, 0xd7, 0x37, 0x3e, 0x1b, 0x83, 0x5b, 0x92, 0xaa, 0xd6, 0x9a, 0x9f, 0xf4,
	0xe3, 0xd0, 0x7a, 0x18, 0x5d, 0x12, 0x30, 0xfc, 0xfb, 0x96, 0x68, 0x26, 0xf1, 0x19, 0xb8, 0x4a,
	0x4b, 0x5d, 0x29, 0xbe, 0xc0, 0x8f, 0x8f, 0x9b, 0x9d, 0x84, 0x41, 0x33, 0x68, 0xe9, 0xd7, 0x01,
	0xf4, 0x6d, 0x75, 0x5e, 0x5a, 0x91, 0xe8, 0x4f, 0x1b, 0xee, 0x5e, 0x91, 0xc6, 0xa7, 0xd0, 0xd5,
	0x75, 0x41, 0x7c, 0x5a, 0x3f, 0x0e, 0x1a, 0x39, 0x3e, 0x3c, 0xad, 0x0b, 0x12, 0xcc, 0x22, 0x42,
	0xb7, 0xaa, 0xf2, 0x05, 0x1f, 0xea, 0x09, 0xae, 0xf1, 0x1e, 0xf4, 0x72, 0x63, 0x9c, 0x2f, 0xb9,
	0x2d, 0x6c, 0x83, 0x8f, 0xc0, 0xe3, 0x22, 0xc9, 0x7f, 0x51, 0xd8, 0x65, 0xbf, 0x07, 0x60, 0xcf,
	0x1a, 0xe9, 0xb0, 0xc7, 0x62, 0x07, 0xc0, 0x28, 0x52, 0xbe, 0xa4, 0x4d, 0xe8, 0x5a, 0x45, 0x6e,
	0xcc, 0x0e, 0x17, 0xac, 0x78, 0x64, 0x15, 0xf7, 0x00, 0xbe, 0x84, 0x9e, 0xd4, 0xba, 0x54, 0xe1,
	0xad, 0x41, 0x67, 0xe8, 0xc7, 0x4f, 0x6e, 0xca, 0x70, 0x34, 0x31, 0x53, 0xd3, 0x8d, 0x2e, 0x6b,
	0x61, 0x37, 0x4e, 0xc6, 0x00, 0x07, 0x10, 0x03, 0xe8, 0xac, 0xa8, 0xe6, 0x1c, 0x3c, 0x61, 0x4a,
	0x73, 0xce, 0x4e, 0xae, 0x2b, 0x6a, 0x5c, 0xdb, 0xe6, 0x55, 0x7b, 0xec, 0x44, 0x31, 0xb8, 0x36,
	0x70, 0x33, 0x63, 0x62, 0xa6, 0x66, 0xcf, 0x36, 0x26, 0xae, 0x6c, 0xbb, 0xa0, 0xe6, 0x57, 0xe2,
	0xfa, 0xf9, 0x18, 0xbc, 0x7d, 0xaa, 0xe8, 0xc3, 0xd1, 0x7c, 0x76, 0x36, 0x3b, 0xff, 0x3c, 0x0b,
	0x5a, 0xd8, 0x07, 0x48, 0x27, 0xc9, 0x99, 0x98, 0x26, 0xf3, 0x0f, 0x69, 0xe0, 0xfc, 0xeb, 0x93,
	0x74, 0x92, 0xce, 0x93, 0xa0, 0x1d, 0x7f, 0x04, 0x77, 0xba, 0x59, 0xe6, 0x1b, 0xc2, 0x37, 0xe0,
	0x5f, 0x30, 0x86, 0x0f, 0xaf, 0x33, 0xcb, 0xbf, 0xfa, 0xc9, 0x83, 0xeb, 0xa8, 0x62, 0x5d, 0x47,
	0xad, 0x2f, 0x2e, 0x33, 0x2f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x15, 0xeb, 0x3e, 0x3b,
	0x03, 0x00, 0x00,
}