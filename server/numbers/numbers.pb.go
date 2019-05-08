// Code generated by protoc-gen-go. DO NOT EDIT.
// source: numbers.proto

package numbers

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Number struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Number) Reset()         { *m = Number{} }
func (m *Number) String() string { return proto.CompactTextString(m) }
func (*Number) ProtoMessage()    {}
func (*Number) Descriptor() ([]byte, []int) {
	return fileDescriptor_874cc518860a4f8b, []int{0}
}

func (m *Number) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Number.Unmarshal(m, b)
}
func (m *Number) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Number.Marshal(b, m, deterministic)
}
func (m *Number) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Number.Merge(m, src)
}
func (m *Number) XXX_Size() int {
	return xxx_messageInfo_Number.Size(m)
}
func (m *Number) XXX_DiscardUnknown() {
	xxx_messageInfo_Number.DiscardUnknown(m)
}

var xxx_messageInfo_Number proto.InternalMessageInfo

func (m *Number) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type BeginRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BeginRequest) Reset()         { *m = BeginRequest{} }
func (m *BeginRequest) String() string { return proto.CompactTextString(m) }
func (*BeginRequest) ProtoMessage()    {}
func (*BeginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_874cc518860a4f8b, []int{1}
}

func (m *BeginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeginRequest.Unmarshal(m, b)
}
func (m *BeginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeginRequest.Marshal(b, m, deterministic)
}
func (m *BeginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeginRequest.Merge(m, src)
}
func (m *BeginRequest) XXX_Size() int {
	return xxx_messageInfo_BeginRequest.Size(m)
}
func (m *BeginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BeginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BeginRequest proto.InternalMessageInfo

type ResumeRequest struct {
	Seed                 int32    `protobuf:"varint,1,opt,name=seed,proto3" json:"seed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResumeRequest) Reset()         { *m = ResumeRequest{} }
func (m *ResumeRequest) String() string { return proto.CompactTextString(m) }
func (*ResumeRequest) ProtoMessage()    {}
func (*ResumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_874cc518860a4f8b, []int{2}
}

func (m *ResumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeRequest.Unmarshal(m, b)
}
func (m *ResumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeRequest.Marshal(b, m, deterministic)
}
func (m *ResumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeRequest.Merge(m, src)
}
func (m *ResumeRequest) XXX_Size() int {
	return xxx_messageInfo_ResumeRequest.Size(m)
}
func (m *ResumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeRequest proto.InternalMessageInfo

func (m *ResumeRequest) GetSeed() int32 {
	if m != nil {
		return m.Seed
	}
	return 0
}

func init() {
	proto.RegisterType((*Number)(nil), "numbers.Number")
	proto.RegisterType((*BeginRequest)(nil), "numbers.BeginRequest")
	proto.RegisterType((*ResumeRequest)(nil), "numbers.ResumeRequest")
}

func init() { proto.RegisterFile("numbers.proto", fileDescriptor_874cc518860a4f8b) }

var fileDescriptor_874cc518860a4f8b = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x2b, 0xcd, 0x4d,
	0x4a, 0x2d, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xe4, 0xb8,
	0xd8, 0xfc, 0xc0, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xd6, 0x20, 0x08, 0x47, 0x89, 0x8f, 0x8b, 0xc7, 0x29, 0x35, 0x3d, 0x33, 0x2f, 0x28,
	0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x49, 0x99, 0x8b, 0x37, 0x28, 0xb5, 0xb8, 0x34, 0x37, 0x15,
	0x2a, 0x20, 0x24, 0xc4, 0xc5, 0x52, 0x9c, 0x9a, 0x9a, 0x02, 0xd5, 0x05, 0x66, 0x1b, 0x95, 0x73,
	0x71, 0xba, 0xa7, 0xe6, 0xa5, 0x16, 0x25, 0x96, 0xe4, 0x17, 0x09, 0x19, 0x73, 0xb1, 0x82, 0x4d,
	0x10, 0x12, 0xd5, 0x83, 0xb9, 0x01, 0xd9, 0x44, 0x29, 0x7e, 0xb8, 0x30, 0xc4, 0x21, 0x4a, 0x0c,
	0x06, 0x8c, 0x42, 0xa6, 0x5c, 0x6c, 0x10, 0x6b, 0x84, 0xc4, 0xe0, 0xd2, 0x28, 0xf6, 0x62, 0xd5,
	0x96, 0xc4, 0x06, 0xf6, 0x9d, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xfb, 0x7c, 0xb7, 0xee,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GeneratorClient is the client API for Generator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeneratorClient interface {
	Begin(ctx context.Context, in *BeginRequest, opts ...grpc.CallOption) (Generator_BeginClient, error)
	Resume(ctx context.Context, in *ResumeRequest, opts ...grpc.CallOption) (Generator_ResumeClient, error)
}

type generatorClient struct {
	cc *grpc.ClientConn
}

func NewGeneratorClient(cc *grpc.ClientConn) GeneratorClient {
	return &generatorClient{cc}
}

func (c *generatorClient) Begin(ctx context.Context, in *BeginRequest, opts ...grpc.CallOption) (Generator_BeginClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Generator_serviceDesc.Streams[0], "/numbers.Generator/Begin", opts...)
	if err != nil {
		return nil, err
	}
	x := &generatorBeginClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Generator_BeginClient interface {
	Recv() (*Number, error)
	grpc.ClientStream
}

type generatorBeginClient struct {
	grpc.ClientStream
}

func (x *generatorBeginClient) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *generatorClient) Resume(ctx context.Context, in *ResumeRequest, opts ...grpc.CallOption) (Generator_ResumeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Generator_serviceDesc.Streams[1], "/numbers.Generator/Resume", opts...)
	if err != nil {
		return nil, err
	}
	x := &generatorResumeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Generator_ResumeClient interface {
	Recv() (*Number, error)
	grpc.ClientStream
}

type generatorResumeClient struct {
	grpc.ClientStream
}

func (x *generatorResumeClient) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GeneratorServer is the server API for Generator service.
type GeneratorServer interface {
	Begin(*BeginRequest, Generator_BeginServer) error
	Resume(*ResumeRequest, Generator_ResumeServer) error
}

func RegisterGeneratorServer(s *grpc.Server, srv GeneratorServer) {
	s.RegisterService(&_Generator_serviceDesc, srv)
}

func _Generator_Begin_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BeginRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GeneratorServer).Begin(m, &generatorBeginServer{stream})
}

type Generator_BeginServer interface {
	Send(*Number) error
	grpc.ServerStream
}

type generatorBeginServer struct {
	grpc.ServerStream
}

func (x *generatorBeginServer) Send(m *Number) error {
	return x.ServerStream.SendMsg(m)
}

func _Generator_Resume_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ResumeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GeneratorServer).Resume(m, &generatorResumeServer{stream})
}

type Generator_ResumeServer interface {
	Send(*Number) error
	grpc.ServerStream
}

type generatorResumeServer struct {
	grpc.ServerStream
}

func (x *generatorResumeServer) Send(m *Number) error {
	return x.ServerStream.SendMsg(m)
}

var _Generator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "numbers.Generator",
	HandlerType: (*GeneratorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Begin",
			Handler:       _Generator_Begin_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Resume",
			Handler:       _Generator_Resume_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "numbers.proto",
}
