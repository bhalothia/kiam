// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package kiam is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	GetPodRoleRequest
	Role
	GetRoleCredentialsRequest
	Credentials
*/
package kiam

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

type GetPodRoleRequest struct {
	Ip string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
}

func (m *GetPodRoleRequest) Reset()                    { *m = GetPodRoleRequest{} }
func (m *GetPodRoleRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPodRoleRequest) ProtoMessage()               {}
func (*GetPodRoleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetPodRoleRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type Role struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Role) Reset()                    { *m = Role{} }
func (m *Role) String() string            { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()               {}
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetRoleCredentialsRequest struct {
	Role *Role `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
}

func (m *GetRoleCredentialsRequest) Reset()                    { *m = GetRoleCredentialsRequest{} }
func (m *GetRoleCredentialsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRoleCredentialsRequest) ProtoMessage()               {}
func (*GetRoleCredentialsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetRoleCredentialsRequest) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type Credentials struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *Credentials) Reset()                    { *m = Credentials{} }
func (m *Credentials) String() string            { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()               {}
func (*Credentials) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Credentials) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPodRoleRequest)(nil), "kiam.GetPodRoleRequest")
	proto.RegisterType((*Role)(nil), "kiam.Role")
	proto.RegisterType((*GetRoleCredentialsRequest)(nil), "kiam.GetRoleCredentialsRequest")
	proto.RegisterType((*Credentials)(nil), "kiam.Credentials")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KiamService service

type KiamServiceClient interface {
	GetPodRole(ctx context.Context, in *GetPodRoleRequest, opts ...grpc.CallOption) (*Role, error)
	GetRoleCredentials(ctx context.Context, in *GetRoleCredentialsRequest, opts ...grpc.CallOption) (*Credentials, error)
}

type kiamServiceClient struct {
	cc *grpc.ClientConn
}

func NewKiamServiceClient(cc *grpc.ClientConn) KiamServiceClient {
	return &kiamServiceClient{cc}
}

func (c *kiamServiceClient) GetPodRole(ctx context.Context, in *GetPodRoleRequest, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := grpc.Invoke(ctx, "/kiam.KiamService/GetPodRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kiamServiceClient) GetRoleCredentials(ctx context.Context, in *GetRoleCredentialsRequest, opts ...grpc.CallOption) (*Credentials, error) {
	out := new(Credentials)
	err := grpc.Invoke(ctx, "/kiam.KiamService/GetRoleCredentials", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KiamService service

type KiamServiceServer interface {
	GetPodRole(context.Context, *GetPodRoleRequest) (*Role, error)
	GetRoleCredentials(context.Context, *GetRoleCredentialsRequest) (*Credentials, error)
}

func RegisterKiamServiceServer(s *grpc.Server, srv KiamServiceServer) {
	s.RegisterService(&_KiamService_serviceDesc, srv)
}

func _KiamService_GetPodRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPodRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KiamServiceServer).GetPodRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kiam.KiamService/GetPodRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KiamServiceServer).GetPodRole(ctx, req.(*GetPodRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KiamService_GetRoleCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KiamServiceServer).GetRoleCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kiam.KiamService/GetRoleCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KiamServiceServer).GetRoleCredentials(ctx, req.(*GetRoleCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KiamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kiam.KiamService",
	HandlerType: (*KiamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPodRole",
			Handler:    _KiamService_GetPodRole_Handler,
		},
		{
			MethodName: "GetRoleCredentials",
			Handler:    _KiamService_GetRoleCredentials_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0xce, 0x4c, 0xcc, 0x55,
	0x52, 0xe6, 0x12, 0x74, 0x4f, 0x2d, 0x09, 0xc8, 0x4f, 0x09, 0xca, 0xcf, 0x49, 0x0d, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x2c, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x62, 0xca, 0x2c, 0x50, 0x92, 0xe2, 0x62, 0x01, 0x49, 0x0b, 0x09, 0x71, 0xb1, 0xe4, 0x25,
	0xe6, 0xa6, 0x42, 0x65, 0xc0, 0x6c, 0x25, 0x6b, 0x2e, 0x49, 0xf7, 0xd4, 0x12, 0x90, 0xb4, 0x73,
	0x51, 0x6a, 0x4a, 0x6a, 0x5e, 0x49, 0x66, 0x62, 0x4e, 0x31, 0xcc, 0x20, 0x39, 0x2e, 0x96, 0xa2,
	0xfc, 0x1c, 0x88, 0x06, 0x6e, 0x23, 0x2e, 0x3d, 0x90, 0x95, 0x7a, 0x60, 0x9b, 0xc0, 0xe2, 0x4a,
	0xca, 0x5c, 0xdc, 0x48, 0xba, 0x84, 0x44, 0xb8, 0x58, 0x4b, 0xf2, 0xb3, 0x53, 0xf3, 0xa0, 0x16,
	0x40, 0x38, 0x46, 0x7d, 0x8c, 0x5c, 0xdc, 0xde, 0x99, 0x89, 0xb9, 0xc1, 0x10, 0xe7, 0x0b, 0x19,
	0x73, 0x71, 0x21, 0x9c, 0x2c, 0x24, 0x0e, 0x31, 0x14, 0xc3, 0x13, 0x52, 0x48, 0xb6, 0x29, 0x31,
	0x08, 0x79, 0x71, 0x09, 0x61, 0x3a, 0x53, 0x48, 0x1e, 0xae, 0x19, 0xbb, 0x07, 0xa4, 0x04, 0x21,
	0x0a, 0x90, 0x64, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0x01, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0xce, 0xb7, 0x69, 0xf8, 0x51, 0x01, 0x00, 0x00,
}