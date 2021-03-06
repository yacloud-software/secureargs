// Code generated by protoc-gen-go.
// source: protos/golang.conradwood.net/apis/secureargs/secureargs.proto
// DO NOT EDIT!

/*
Package secureargs is a generated protocol buffer package.

It is generated from these files:
	protos/golang.conradwood.net/apis/secureargs/secureargs.proto

It has these top-level messages:
	GetArgsRequest
	ArgsResponse
	SetArgRequest
	Arg
*/
package secureargs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "golang.conradwood.net/apis/common"

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

type GetArgsRequest struct {
	RepositoryID uint64 `protobuf:"varint,1,opt,name=RepositoryID" json:"RepositoryID,omitempty"`
}

func (m *GetArgsRequest) Reset()                    { *m = GetArgsRequest{} }
func (m *GetArgsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetArgsRequest) ProtoMessage()               {}
func (*GetArgsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetArgsRequest) GetRepositoryID() uint64 {
	if m != nil {
		return m.RepositoryID
	}
	return 0
}

type ArgsResponse struct {
	Args map[string]string `protobuf:"bytes,1,rep,name=Args" json:"Args,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ArgsResponse) Reset()                    { *m = ArgsResponse{} }
func (m *ArgsResponse) String() string            { return proto.CompactTextString(m) }
func (*ArgsResponse) ProtoMessage()               {}
func (*ArgsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ArgsResponse) GetArgs() map[string]string {
	if m != nil {
		return m.Args
	}
	return nil
}

type SetArgRequest struct {
	RepositoryID uint64 `protobuf:"varint,1,opt,name=RepositoryID" json:"RepositoryID,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Value        string `protobuf:"bytes,3,opt,name=Value" json:"Value,omitempty"`
}

func (m *SetArgRequest) Reset()                    { *m = SetArgRequest{} }
func (m *SetArgRequest) String() string            { return proto.CompactTextString(m) }
func (*SetArgRequest) ProtoMessage()               {}
func (*SetArgRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SetArgRequest) GetRepositoryID() uint64 {
	if m != nil {
		return m.RepositoryID
	}
	return 0
}

func (m *SetArgRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SetArgRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// not very normalized database....
type Arg struct {
	ID uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	// string Repository = 2; // obsolete, use ID instead
	Name         string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Value        string `protobuf:"bytes,4,opt,name=Value" json:"Value,omitempty"`
	RepositoryID uint64 `protobuf:"varint,5,opt,name=RepositoryID" json:"RepositoryID,omitempty"`
}

func (m *Arg) Reset()                    { *m = Arg{} }
func (m *Arg) String() string            { return proto.CompactTextString(m) }
func (*Arg) ProtoMessage()               {}
func (*Arg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Arg) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Arg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Arg) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Arg) GetRepositoryID() uint64 {
	if m != nil {
		return m.RepositoryID
	}
	return 0
}

func init() {
	proto.RegisterType((*GetArgsRequest)(nil), "secureargs.GetArgsRequest")
	proto.RegisterType((*ArgsResponse)(nil), "secureargs.ArgsResponse")
	proto.RegisterType((*SetArgRequest)(nil), "secureargs.SetArgRequest")
	proto.RegisterType((*Arg)(nil), "secureargs.Arg")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SecureArgsService service

type SecureArgsServiceClient interface {
	// get all arguments for a given repository
	GetArgs(ctx context.Context, in *GetArgsRequest, opts ...grpc.CallOption) (*ArgsResponse, error)
	// set an arg
	SetArg(ctx context.Context, in *SetArgRequest, opts ...grpc.CallOption) (*common.Void, error)
}

type secureArgsServiceClient struct {
	cc *grpc.ClientConn
}

func NewSecureArgsServiceClient(cc *grpc.ClientConn) SecureArgsServiceClient {
	return &secureArgsServiceClient{cc}
}

func (c *secureArgsServiceClient) GetArgs(ctx context.Context, in *GetArgsRequest, opts ...grpc.CallOption) (*ArgsResponse, error) {
	out := new(ArgsResponse)
	err := grpc.Invoke(ctx, "/secureargs.SecureArgsService/GetArgs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secureArgsServiceClient) SetArg(ctx context.Context, in *SetArgRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/secureargs.SecureArgsService/SetArg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SecureArgsService service

type SecureArgsServiceServer interface {
	// get all arguments for a given repository
	GetArgs(context.Context, *GetArgsRequest) (*ArgsResponse, error)
	// set an arg
	SetArg(context.Context, *SetArgRequest) (*common.Void, error)
}

func RegisterSecureArgsServiceServer(s *grpc.Server, srv SecureArgsServiceServer) {
	s.RegisterService(&_SecureArgsService_serviceDesc, srv)
}

func _SecureArgsService_GetArgs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArgsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureArgsServiceServer).GetArgs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/secureargs.SecureArgsService/GetArgs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureArgsServiceServer).GetArgs(ctx, req.(*GetArgsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecureArgsService_SetArg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetArgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureArgsServiceServer).SetArg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/secureargs.SecureArgsService/SetArg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureArgsServiceServer).SetArg(ctx, req.(*SetArgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecureArgsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "secureargs.SecureArgsService",
	HandlerType: (*SecureArgsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArgs",
			Handler:    _SecureArgsService_GetArgs_Handler,
		},
		{
			MethodName: "SetArg",
			Handler:    _SecureArgsService_SetArg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/golang.conradwood.net/apis/secureargs/secureargs.proto",
}

func init() {
	proto.RegisterFile("protos/golang.conradwood.net/apis/secureargs/secureargs.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x25, 0x1f, 0xad, 0x74, 0xac, 0x45, 0x17, 0x0f, 0x31, 0xa7, 0x12, 0x10, 0x7b, 0x4a, 0xb1,
	0x8a, 0x8a, 0x20, 0x52, 0x51, 0xc4, 0x8b, 0x87, 0x14, 0x7a, 0xf3, 0xb0, 0xa6, 0x43, 0x08, 0xb6,
	0x3b, 0x71, 0x77, 0x5b, 0xe9, 0xc9, 0xa3, 0x7f, 0x5b, 0xb2, 0x89, 0x75, 0x4b, 0x55, 0xf0, 0xb4,
	0x33, 0xc3, 0xcc, 0x7b, 0x8f, 0xb7, 0x0f, 0xae, 0x0a, 0x49, 0x9a, 0x54, 0x3f, 0xa3, 0x29, 0x17,
	0x59, 0x9c, 0x92, 0x90, 0x7c, 0xf2, 0x46, 0x34, 0x89, 0x05, 0xea, 0x3e, 0x2f, 0x72, 0xd5, 0x57,
	0x98, 0xce, 0x25, 0x72, 0x99, 0xd9, 0x65, 0x6c, 0xee, 0x18, 0x7c, 0x4f, 0xc2, 0xf8, 0x0f, 0x8c,
	0x94, 0x66, 0x33, 0x12, 0xf5, 0x53, 0xdd, 0x46, 0xa7, 0xd0, 0xb9, 0x47, 0x3d, 0x94, 0x99, 0x4a,
	0xf0, 0x75, 0x8e, 0x4a, 0xb3, 0x08, 0xda, 0x09, 0x16, 0xa4, 0x72, 0x4d, 0x72, 0xf9, 0x70, 0x1b,
	0x38, 0x5d, 0xa7, 0xe7, 0x27, 0x6b, 0xb3, 0xe8, 0x1d, 0xda, 0xd5, 0x89, 0x2a, 0x48, 0x28, 0x64,
	0x67, 0xe0, 0x97, 0x7d, 0xe0, 0x74, 0xbd, 0xde, 0xf6, 0x20, 0x8a, 0x2d, 0x89, 0xf6, 0x9e, 0x69,
	0xee, 0x84, 0x96, 0xcb, 0xc4, 0xec, 0x87, 0xe7, 0xd0, 0x5a, 0x8d, 0xd8, 0x2e, 0x78, 0x2f, 0xb8,
	0x34, 0x7c, 0xad, 0xa4, 0x2c, 0xd9, 0x3e, 0x34, 0x16, 0x7c, 0x3a, 0xc7, 0xc0, 0x35, 0xb3, 0xaa,
	0xb9, 0x74, 0x2f, 0x9c, 0xe8, 0x09, 0x76, 0x46, 0x46, 0xf6, 0x3f, 0x54, 0x33, 0x06, 0xfe, 0x23,
	0x9f, 0x7d, 0xa1, 0x99, 0xba, 0xa4, 0x18, 0x1b, 0x0a, 0xaf, 0xa2, 0x30, 0x4d, 0x94, 0x82, 0x37,
	0x94, 0x19, 0xeb, 0x80, 0xbb, 0x82, 0x72, 0x2d, 0x00, 0xef, 0x27, 0x00, 0xdf, 0x02, 0xd8, 0x90,
	0xd3, 0xd8, 0x94, 0x33, 0xf8, 0x70, 0x60, 0x6f, 0x64, 0x8c, 0x2a, 0x3d, 0x18, 0xa1, 0x5c, 0xe4,
	0x29, 0xb2, 0x6b, 0xd8, 0xaa, 0x3f, 0x84, 0x85, 0xb6, 0x8f, 0xeb, 0xbf, 0x14, 0x06, 0xbf, 0x79,
	0xcc, 0x8e, 0xa1, 0x59, 0x59, 0xc3, 0x0e, 0xec, 0x9d, 0x35, 0xbb, 0xc2, 0x76, 0x5c, 0xa7, 0x60,
	0x4c, 0xf9, 0xe4, 0xe6, 0x08, 0x0e, 0x05, 0x6a, 0x3b, 0x33, 0x75, 0x8a, 0xca, 0xd8, 0x58, 0x20,
	0xcf, 0x4d, 0x13, 0x9a, 0x93, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xd3, 0x37, 0x46, 0xb1,
	0x02, 0x00, 0x00,
}
