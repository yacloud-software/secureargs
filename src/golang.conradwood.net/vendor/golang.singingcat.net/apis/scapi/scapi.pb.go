// Code generated by protoc-gen-go.
// source: golang.singingcat.net/apis/scapi/scapi.proto
// DO NOT EDIT!

/*
Package scapi is a generated protocol buffer package.

It is generated from these files:
	golang.singingcat.net/apis/scapi/scapi.proto

It has these top-level messages:
	PingRequest
	PingResponse
	ModuleNameUpdate
	SensorConfigUpdateRequest
	UpdateUserAppVersionRequest
	UpdateUserAppRepoRequest
	Version
	UserAppVersions
*/
package scapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import scweb "golang.singingcat.net/apis/scweb"
import scfunctions "golang.singingcat.net/apis/scfunctions"
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

// comment: message pingrequest
type PingRequest struct {
	// comment: payload
	Payload string `protobuf:"bytes,2,opt,name=Payload" json:"Payload,omitempty"`
	// comment: sequencenumber
	SequenceNumber uint32 `protobuf:"varint,1,opt,name=SequenceNumber" json:"SequenceNumber,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *PingRequest) GetSequenceNumber() uint32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

// comment: message pingresponse
type PingResponse struct {
	// comment: field pingresponse.response
	Response *PingRequest `protobuf:"bytes,1,opt,name=Response" json:"Response,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingResponse) GetResponse() *PingRequest {
	if m != nil {
		return m.Response
	}
	return nil
}

type ModuleNameUpdate struct {
	ModuleID uint64 `protobuf:"varint,1,opt,name=ModuleID" json:"ModuleID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
}

func (m *ModuleNameUpdate) Reset()                    { *m = ModuleNameUpdate{} }
func (m *ModuleNameUpdate) String() string            { return proto.CompactTextString(m) }
func (*ModuleNameUpdate) ProtoMessage()               {}
func (*ModuleNameUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ModuleNameUpdate) GetModuleID() uint64 {
	if m != nil {
		return m.ModuleID
	}
	return 0
}

func (m *ModuleNameUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SensorConfigUpdateRequest struct {
	SensorDefID       uint64 `protobuf:"varint,1,opt,name=SensorDefID" json:"SensorDefID,omitempty"`
	Enable            bool   `protobuf:"varint,2,opt,name=Enable" json:"Enable,omitempty"`
	ReportingInterval uint32 `protobuf:"varint,3,opt,name=ReportingInterval" json:"ReportingInterval,omitempty"`
	PollingInterval   uint32 `protobuf:"varint,4,opt,name=PollingInterval" json:"PollingInterval,omitempty"`
	FriendlyName      string `protobuf:"bytes,5,opt,name=FriendlyName" json:"FriendlyName,omitempty"`
}

func (m *SensorConfigUpdateRequest) Reset()                    { *m = SensorConfigUpdateRequest{} }
func (m *SensorConfigUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*SensorConfigUpdateRequest) ProtoMessage()               {}
func (*SensorConfigUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SensorConfigUpdateRequest) GetSensorDefID() uint64 {
	if m != nil {
		return m.SensorDefID
	}
	return 0
}

func (m *SensorConfigUpdateRequest) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *SensorConfigUpdateRequest) GetReportingInterval() uint32 {
	if m != nil {
		return m.ReportingInterval
	}
	return 0
}

func (m *SensorConfigUpdateRequest) GetPollingInterval() uint32 {
	if m != nil {
		return m.PollingInterval
	}
	return 0
}

func (m *SensorConfigUpdateRequest) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

type UpdateUserAppVersionRequest struct {
	ModuleID uint64 `protobuf:"varint,1,opt,name=ModuleID" json:"ModuleID,omitempty"`
	Build    uint64 `protobuf:"varint,2,opt,name=Build" json:"Build,omitempty"`
}

func (m *UpdateUserAppVersionRequest) Reset()                    { *m = UpdateUserAppVersionRequest{} }
func (m *UpdateUserAppVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserAppVersionRequest) ProtoMessage()               {}
func (*UpdateUserAppVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateUserAppVersionRequest) GetModuleID() uint64 {
	if m != nil {
		return m.ModuleID
	}
	return 0
}

func (m *UpdateUserAppVersionRequest) GetBuild() uint64 {
	if m != nil {
		return m.Build
	}
	return 0
}

type UpdateUserAppRepoRequest struct {
	NewValue string `protobuf:"bytes,1,opt,name=NewValue" json:"NewValue,omitempty"`
	ModuleID uint64 `protobuf:"varint,2,opt,name=ModuleID" json:"ModuleID,omitempty"`
}

func (m *UpdateUserAppRepoRequest) Reset()                    { *m = UpdateUserAppRepoRequest{} }
func (m *UpdateUserAppRepoRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserAppRepoRequest) ProtoMessage()               {}
func (*UpdateUserAppRepoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateUserAppRepoRequest) GetNewValue() string {
	if m != nil {
		return m.NewValue
	}
	return ""
}

func (m *UpdateUserAppRepoRequest) GetModuleID() uint64 {
	if m != nil {
		return m.ModuleID
	}
	return 0
}

type Version struct {
	ID       string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Selected bool   `protobuf:"varint,3,opt,name=Selected" json:"Selected,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Version) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Version) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Version) GetSelected() bool {
	if m != nil {
		return m.Selected
	}
	return false
}

type UserAppVersions struct {
	Versions []*Version `protobuf:"bytes,1,rep,name=Versions" json:"Versions,omitempty"`
}

func (m *UserAppVersions) Reset()                    { *m = UserAppVersions{} }
func (m *UserAppVersions) String() string            { return proto.CompactTextString(m) }
func (*UserAppVersions) ProtoMessage()               {}
func (*UserAppVersions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UserAppVersions) GetVersions() []*Version {
	if m != nil {
		return m.Versions
	}
	return nil
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "scapi.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "scapi.PingResponse")
	proto.RegisterType((*ModuleNameUpdate)(nil), "scapi.ModuleNameUpdate")
	proto.RegisterType((*SensorConfigUpdateRequest)(nil), "scapi.SensorConfigUpdateRequest")
	proto.RegisterType((*UpdateUserAppVersionRequest)(nil), "scapi.UpdateUserAppVersionRequest")
	proto.RegisterType((*UpdateUserAppRepoRequest)(nil), "scapi.UpdateUserAppRepoRequest")
	proto.RegisterType((*Version)(nil), "scapi.Version")
	proto.RegisterType((*UserAppVersions)(nil), "scapi.UserAppVersions")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SCApiService service

type SCApiServiceClient interface {
	// just ping the api
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// update name of a module
	UpdateModuleName(ctx context.Context, in *ModuleNameUpdate, opts ...grpc.CallOption) (*common.Void, error)
	// update friendly name of sensor
	UpdateSensorName(ctx context.Context, in *scweb.SensorNameUpdate, opts ...grpc.CallOption) (*common.Void, error)
	// set a config flag
	SetConfigFlag(ctx context.Context, in *scfunctions.SetFlagRequest, opts ...grpc.CallOption) (*common.Void, error)
	// update sensor configuration
	UpdateSensorConfiguration(ctx context.Context, in *SensorConfigUpdateRequest, opts ...grpc.CallOption) (*common.Void, error)
	// update the user app repo
	UpdateUserAppRepo(ctx context.Context, in *UpdateUserAppRepoRequest, opts ...grpc.CallOption) (*UserAppVersions, error)
	// update the user app version
	UpdateUserAppVersion(ctx context.Context, in *UpdateUserAppVersionRequest, opts ...grpc.CallOption) (*common.Void, error)
}

type sCApiServiceClient struct {
	cc *grpc.ClientConn
}

func NewSCApiServiceClient(cc *grpc.ClientConn) SCApiServiceClient {
	return &sCApiServiceClient{cc}
}

func (c *sCApiServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/scapi.SCApiService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCApiServiceClient) UpdateModuleName(ctx context.Context, in *ModuleNameUpdate, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/scapi.SCApiService/UpdateModuleName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCApiServiceClient) UpdateSensorName(ctx context.Context, in *scweb.SensorNameUpdate, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/scapi.SCApiService/UpdateSensorName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCApiServiceClient) SetConfigFlag(ctx context.Context, in *scfunctions.SetFlagRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/scapi.SCApiService/SetConfigFlag", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCApiServiceClient) UpdateSensorConfiguration(ctx context.Context, in *SensorConfigUpdateRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/scapi.SCApiService/UpdateSensorConfiguration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCApiServiceClient) UpdateUserAppRepo(ctx context.Context, in *UpdateUserAppRepoRequest, opts ...grpc.CallOption) (*UserAppVersions, error) {
	out := new(UserAppVersions)
	err := grpc.Invoke(ctx, "/scapi.SCApiService/UpdateUserAppRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCApiServiceClient) UpdateUserAppVersion(ctx context.Context, in *UpdateUserAppVersionRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/scapi.SCApiService/UpdateUserAppVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SCApiService service

type SCApiServiceServer interface {
	// just ping the api
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// update name of a module
	UpdateModuleName(context.Context, *ModuleNameUpdate) (*common.Void, error)
	// update friendly name of sensor
	UpdateSensorName(context.Context, *scweb.SensorNameUpdate) (*common.Void, error)
	// set a config flag
	SetConfigFlag(context.Context, *scfunctions.SetFlagRequest) (*common.Void, error)
	// update sensor configuration
	UpdateSensorConfiguration(context.Context, *SensorConfigUpdateRequest) (*common.Void, error)
	// update the user app repo
	UpdateUserAppRepo(context.Context, *UpdateUserAppRepoRequest) (*UserAppVersions, error)
	// update the user app version
	UpdateUserAppVersion(context.Context, *UpdateUserAppVersionRequest) (*common.Void, error)
}

func RegisterSCApiServiceServer(s *grpc.Server, srv SCApiServiceServer) {
	s.RegisterService(&_SCApiService_serviceDesc, srv)
}

func _SCApiService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCApiServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scapi.SCApiService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCApiServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCApiService_UpdateModuleName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModuleNameUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCApiServiceServer).UpdateModuleName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scapi.SCApiService/UpdateModuleName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCApiServiceServer).UpdateModuleName(ctx, req.(*ModuleNameUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCApiService_UpdateSensorName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(scweb.SensorNameUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCApiServiceServer).UpdateSensorName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scapi.SCApiService/UpdateSensorName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCApiServiceServer).UpdateSensorName(ctx, req.(*scweb.SensorNameUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCApiService_SetConfigFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(scfunctions.SetFlagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCApiServiceServer).SetConfigFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scapi.SCApiService/SetConfigFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCApiServiceServer).SetConfigFlag(ctx, req.(*scfunctions.SetFlagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCApiService_UpdateSensorConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SensorConfigUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCApiServiceServer).UpdateSensorConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scapi.SCApiService/UpdateSensorConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCApiServiceServer).UpdateSensorConfiguration(ctx, req.(*SensorConfigUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCApiService_UpdateUserAppRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserAppRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCApiServiceServer).UpdateUserAppRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scapi.SCApiService/UpdateUserAppRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCApiServiceServer).UpdateUserAppRepo(ctx, req.(*UpdateUserAppRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCApiService_UpdateUserAppVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserAppVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCApiServiceServer).UpdateUserAppVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scapi.SCApiService/UpdateUserAppVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCApiServiceServer).UpdateUserAppVersion(ctx, req.(*UpdateUserAppVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SCApiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scapi.SCApiService",
	HandlerType: (*SCApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SCApiService_Ping_Handler,
		},
		{
			MethodName: "UpdateModuleName",
			Handler:    _SCApiService_UpdateModuleName_Handler,
		},
		{
			MethodName: "UpdateSensorName",
			Handler:    _SCApiService_UpdateSensorName_Handler,
		},
		{
			MethodName: "SetConfigFlag",
			Handler:    _SCApiService_SetConfigFlag_Handler,
		},
		{
			MethodName: "UpdateSensorConfiguration",
			Handler:    _SCApiService_UpdateSensorConfiguration_Handler,
		},
		{
			MethodName: "UpdateUserAppRepo",
			Handler:    _SCApiService_UpdateUserAppRepo_Handler,
		},
		{
			MethodName: "UpdateUserAppVersion",
			Handler:    _SCApiService_UpdateUserAppVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.singingcat.net/apis/scapi/scapi.proto",
}

func init() { proto.RegisterFile("golang.singingcat.net/apis/scapi/scapi.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x54, 0xdf, 0x4f, 0xdb, 0x30,
	0x10, 0x56, 0x4b, 0x81, 0x72, 0x2d, 0xbf, 0x3c, 0xc4, 0x4a, 0x78, 0x58, 0x94, 0x87, 0xa9, 0x9a,
	0x50, 0x90, 0xd8, 0xcb, 0x36, 0x69, 0x93, 0xf8, 0x31, 0x34, 0xb4, 0x0d, 0x90, 0x23, 0x78, 0x77,
	0x93, 0x23, 0xb2, 0xe4, 0xda, 0x59, 0xe2, 0x80, 0xf8, 0x33, 0xf7, 0xb2, 0xbf, 0x67, 0x4a, 0xec,
	0xa4, 0x49, 0x29, 0x7d, 0x49, 0x7c, 0xe7, 0xfb, 0x3e, 0xdf, 0x9d, 0x3f, 0x1f, 0x1c, 0xc5, 0x4a,
	0x30, 0x19, 0xfb, 0x19, 0x97, 0x31, 0x97, 0x71, 0xc8, 0xb4, 0x2f, 0x51, 0x1f, 0xb3, 0x84, 0x67,
	0xc7, 0x59, 0xc8, 0x12, 0x6e, 0xbe, 0x7e, 0x92, 0x2a, 0xad, 0xc8, 0x6a, 0x69, 0x38, 0xcb, 0x41,
	0x4f, 0x38, 0x31, 0x5f, 0x03, 0x72, 0x3e, 0x2d, 0x8d, 0x7e, 0xc8, 0x65, 0xa8, 0xb9, 0x92, 0xad,
	0xb5, 0x45, 0xfa, 0x16, 0x19, 0x2a, 0x99, 0xb2, 0xe8, 0x49, 0xa9, 0x68, 0x86, 0x0c, 0xd5, 0x74,
	0xaa, 0xa4, 0xfd, 0x99, 0x78, 0xef, 0x06, 0x06, 0xb7, 0x5c, 0xc6, 0x14, 0xff, 0xe4, 0x98, 0x69,
	0x32, 0x82, 0xf5, 0x5b, 0xf6, 0x2c, 0x14, 0x8b, 0x46, 0x5d, 0xb7, 0x33, 0xde, 0xa0, 0x95, 0x49,
	0xde, 0xc3, 0x56, 0x50, 0x04, 0xc9, 0x10, 0xaf, 0xf3, 0xe9, 0x04, 0xd3, 0x51, 0xc7, 0xed, 0x8c,
	0x37, 0xe9, 0x9c, 0xd7, 0xfb, 0x06, 0x43, 0x43, 0x98, 0x25, 0x4a, 0x66, 0x48, 0x7c, 0xe8, 0x57,
	0xeb, 0x12, 0x31, 0x38, 0x21, 0xbe, 0xe9, 0x4f, 0xe3, 0x5c, 0x5a, 0xc7, 0x78, 0x67, 0xb0, 0xf3,
	0x5b, 0x45, 0xb9, 0xc0, 0x6b, 0x36, 0xc5, 0xbb, 0x24, 0x62, 0x1a, 0x89, 0x03, 0x7d, 0xe3, 0xbb,
	0xba, 0x28, 0x39, 0x7a, 0xb4, 0xb6, 0x09, 0x81, 0x5e, 0x11, 0x69, 0xd3, 0x2d, 0xd7, 0xde, 0xdf,
	0x0e, 0x1c, 0x04, 0x28, 0x33, 0x95, 0x9e, 0x2b, 0xf9, 0xc0, 0x63, 0x43, 0x53, 0xd5, 0xe8, 0xc2,
	0xc0, 0x6c, 0x5e, 0xe0, 0x43, 0x4d, 0xd8, 0x74, 0x91, 0x7d, 0x58, 0xfb, 0x2e, 0xd9, 0x44, 0x18,
	0xd6, 0x3e, 0xb5, 0x16, 0x39, 0x82, 0x5d, 0x8a, 0x89, 0x4a, 0x35, 0x97, 0xf1, 0x95, 0xd4, 0x98,
	0x3e, 0x32, 0x31, 0x5a, 0x29, 0xdb, 0xf0, 0x72, 0x83, 0x8c, 0x61, 0xfb, 0x56, 0x09, 0xd1, 0x8c,
	0xed, 0x95, 0xb1, 0xf3, 0x6e, 0xe2, 0xc1, 0xf0, 0x32, 0xe5, 0x28, 0x23, 0xf1, 0x5c, 0xd6, 0xb2,
	0x5a, 0xd6, 0xd2, 0xf2, 0x79, 0x37, 0x70, 0x68, 0xca, 0xb8, 0xcb, 0x30, 0x3d, 0x4d, 0x92, 0x7b,
	0x4c, 0x33, 0xae, 0x64, 0x55, 0xd4, 0xb2, 0x16, 0xed, 0xc1, 0xea, 0x59, 0xce, 0x85, 0xb9, 0xd2,
	0x1e, 0x35, 0x86, 0x47, 0x61, 0xd4, 0x22, 0x2c, 0x0a, 0x68, 0xb0, 0x5d, 0xe3, 0xd3, 0x3d, 0x13,
	0xb9, 0xb9, 0xb4, 0x0d, 0x5a, 0xdb, 0xad, 0x93, 0xba, 0xed, 0x93, 0xbc, 0x2b, 0x58, 0xb7, 0x79,
	0x91, 0x2d, 0xe8, 0xda, 0x54, 0x36, 0x68, 0x77, 0xf1, 0x3d, 0x15, 0x54, 0x01, 0x0a, 0x0c, 0x35,
	0x46, 0x65, 0x1b, 0xfb, 0xb4, 0xb6, 0xbd, 0xaf, 0xb0, 0xdd, 0xae, 0x34, 0x23, 0x1f, 0xa0, 0x5f,
	0xad, 0x47, 0x1d, 0x77, 0x65, 0x3c, 0x38, 0xd9, 0xb2, 0x52, 0xaa, 0x9a, 0x51, 0xef, 0x9f, 0xfc,
	0x5b, 0x81, 0x61, 0x70, 0x7e, 0x9a, 0xf0, 0x00, 0xd3, 0x47, 0x1e, 0x22, 0x39, 0x86, 0x5e, 0x21,
	0x38, 0xb2, 0x40, 0x7d, 0xce, 0x9b, 0x96, 0xcf, 0x0a, 0xf7, 0x33, 0xec, 0x98, 0xfe, 0xcc, 0xe4,
	0x48, 0xde, 0xda, 0xc0, 0x79, 0x85, 0x3a, 0x43, 0xdf, 0xbe, 0xaa, 0x7b, 0xc5, 0xa3, 0x19, 0xd4,
	0x88, 0xaa, 0x86, 0x16, 0x0f, 0x7c, 0xe6, 0x5a, 0x08, 0xfd, 0x02, 0x9b, 0x01, 0x6a, 0x23, 0xdb,
	0x4b, 0xc1, 0x62, 0x72, 0xe8, 0x37, 0x1f, 0x79, 0x80, 0xba, 0xf0, 0x56, 0x89, 0xb7, 0xb1, 0x3f,
	0xe1, 0xa0, 0x79, 0xac, 0x21, 0xc9, 0x53, 0x56, 0x40, 0x89, 0x6b, 0x53, 0x7f, 0xf5, 0x5d, 0xcc,
	0x91, 0xfd, 0x82, 0xdd, 0x17, 0xf2, 0x20, 0xef, 0x2c, 0xc9, 0x6b, 0xc2, 0x71, 0xf6, 0xab, 0x80,
	0xb9, 0xab, 0xfb, 0x01, 0x7b, 0x8b, 0xd4, 0x4b, 0xbc, 0x45, 0x84, 0x6d, 0x69, 0xb7, 0xf3, 0x3a,
	0xf3, 0xc0, 0x95, 0xa8, 0x9b, 0xf3, 0xcd, 0x4e, 0xbc, 0x62, 0xc4, 0x19, 0xb6, 0xc9, 0x5a, 0x39,
	0xdb, 0x3e, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x72, 0xf3, 0x51, 0x15, 0xaa, 0x05, 0x00, 0x00,
}
