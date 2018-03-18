// Code generated by protoc-gen-go. DO NOT EDIT.
// source: staffmana.proto

/*
Package Staffmanarpc is a generated protocol buffer package.

It is generated from these files:
	staffmana.proto

It has these top-level messages:
	AddStaffReq
	AddStaffResp
	DelStaffReq
	DelStaffResp
	UpdateStaffReq
	UpdateStaffResp
	SelectAllStaffReq
	SelectAllStaffResp
	SelectStaffByIdReq
	SelectStaffByIdResp
*/
package Staffmanarpc

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

type AddStaffReq struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Passwd   string `protobuf:"bytes,2,opt,name=passwd" json:"passwd,omitempty"`
	Position int64  `protobuf:"varint,3,opt,name=position" json:"position,omitempty"`
}

func (m *AddStaffReq) Reset()                    { *m = AddStaffReq{} }
func (m *AddStaffReq) String() string            { return proto.CompactTextString(m) }
func (*AddStaffReq) ProtoMessage()               {}
func (*AddStaffReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddStaffReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddStaffReq) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

func (m *AddStaffReq) GetPosition() int64 {
	if m != nil {
		return m.Position
	}
	return 0
}

type AddStaffResp struct {
	Success bool `protobuf:"varint,1,opt,name=Success,json=success" json:"Success,omitempty"`
}

func (m *AddStaffResp) Reset()                    { *m = AddStaffResp{} }
func (m *AddStaffResp) String() string            { return proto.CompactTextString(m) }
func (*AddStaffResp) ProtoMessage()               {}
func (*AddStaffResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddStaffResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type DelStaffReq struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DelStaffReq) Reset()                    { *m = DelStaffReq{} }
func (m *DelStaffReq) String() string            { return proto.CompactTextString(m) }
func (*DelStaffReq) ProtoMessage()               {}
func (*DelStaffReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DelStaffReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DelStaffResp struct {
	Success bool `protobuf:"varint,1,opt,name=Success,json=success" json:"Success,omitempty"`
}

func (m *DelStaffResp) Reset()                    { *m = DelStaffResp{} }
func (m *DelStaffResp) String() string            { return proto.CompactTextString(m) }
func (*DelStaffResp) ProtoMessage()               {}
func (*DelStaffResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DelStaffResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type UpdateStaffReq struct {
	Id       int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Passwd   string `protobuf:"bytes,3,opt,name=passwd" json:"passwd,omitempty"`
	Position int64  `protobuf:"varint,4,opt,name=position" json:"position,omitempty"`
}

func (m *UpdateStaffReq) Reset()                    { *m = UpdateStaffReq{} }
func (m *UpdateStaffReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateStaffReq) ProtoMessage()               {}
func (*UpdateStaffReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateStaffReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateStaffReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateStaffReq) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

func (m *UpdateStaffReq) GetPosition() int64 {
	if m != nil {
		return m.Position
	}
	return 0
}

type UpdateStaffResp struct {
	Success bool `protobuf:"varint,1,opt,name=Success,json=success" json:"Success,omitempty"`
}

func (m *UpdateStaffResp) Reset()                    { *m = UpdateStaffResp{} }
func (m *UpdateStaffResp) String() string            { return proto.CompactTextString(m) }
func (*UpdateStaffResp) ProtoMessage()               {}
func (*UpdateStaffResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateStaffResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type SelectAllStaffReq struct {
}

func (m *SelectAllStaffReq) Reset()                    { *m = SelectAllStaffReq{} }
func (m *SelectAllStaffReq) String() string            { return proto.CompactTextString(m) }
func (*SelectAllStaffReq) ProtoMessage()               {}
func (*SelectAllStaffReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type SelectAllStaffResp struct {
	Payload []byte `protobuf:"bytes,1,opt,name=Payload,json=payload,proto3" json:"Payload,omitempty"`
}

func (m *SelectAllStaffResp) Reset()                    { *m = SelectAllStaffResp{} }
func (m *SelectAllStaffResp) String() string            { return proto.CompactTextString(m) }
func (*SelectAllStaffResp) ProtoMessage()               {}
func (*SelectAllStaffResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SelectAllStaffResp) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type SelectStaffByIdReq struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *SelectStaffByIdReq) Reset()                    { *m = SelectStaffByIdReq{} }
func (m *SelectStaffByIdReq) String() string            { return proto.CompactTextString(m) }
func (*SelectStaffByIdReq) ProtoMessage()               {}
func (*SelectStaffByIdReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SelectStaffByIdReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SelectStaffByIdResp struct {
	Payload []byte `protobuf:"bytes,1,opt,name=Payload,json=payload,proto3" json:"Payload,omitempty"`
}

func (m *SelectStaffByIdResp) Reset()                    { *m = SelectStaffByIdResp{} }
func (m *SelectStaffByIdResp) String() string            { return proto.CompactTextString(m) }
func (*SelectStaffByIdResp) ProtoMessage()               {}
func (*SelectStaffByIdResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SelectStaffByIdResp) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*AddStaffReq)(nil), "Staffmanarpc.AddStaffReq")
	proto.RegisterType((*AddStaffResp)(nil), "Staffmanarpc.AddStaffResp")
	proto.RegisterType((*DelStaffReq)(nil), "Staffmanarpc.DelStaffReq")
	proto.RegisterType((*DelStaffResp)(nil), "Staffmanarpc.DelStaffResp")
	proto.RegisterType((*UpdateStaffReq)(nil), "Staffmanarpc.UpdateStaffReq")
	proto.RegisterType((*UpdateStaffResp)(nil), "Staffmanarpc.UpdateStaffResp")
	proto.RegisterType((*SelectAllStaffReq)(nil), "Staffmanarpc.SelectAllStaffReq")
	proto.RegisterType((*SelectAllStaffResp)(nil), "Staffmanarpc.SelectAllStaffResp")
	proto.RegisterType((*SelectStaffByIdReq)(nil), "Staffmanarpc.SelectStaffByIdReq")
	proto.RegisterType((*SelectStaffByIdResp)(nil), "Staffmanarpc.SelectStaffByIdResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Staffmana service

type StaffmanaClient interface {
	AddStaff(ctx context.Context, in *AddStaffReq, opts ...grpc.CallOption) (*AddStaffResp, error)
	DelStaff(ctx context.Context, in *DelStaffReq, opts ...grpc.CallOption) (*DelStaffResp, error)
	UpdateStaff(ctx context.Context, in *UpdateStaffReq, opts ...grpc.CallOption) (*UpdateStaffResp, error)
	SelectAllStaff(ctx context.Context, in *SelectAllStaffReq, opts ...grpc.CallOption) (*SelectAllStaffResp, error)
	SelectStaffById(ctx context.Context, in *SelectStaffByIdReq, opts ...grpc.CallOption) (*SelectStaffByIdResp, error)
}

type staffmanaClient struct {
	cc *grpc.ClientConn
}

func NewStaffmanaClient(cc *grpc.ClientConn) StaffmanaClient {
	return &staffmanaClient{cc}
}

func (c *staffmanaClient) AddStaff(ctx context.Context, in *AddStaffReq, opts ...grpc.CallOption) (*AddStaffResp, error) {
	out := new(AddStaffResp)
	err := grpc.Invoke(ctx, "/Staffmanarpc.staffmana/AddStaff", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffmanaClient) DelStaff(ctx context.Context, in *DelStaffReq, opts ...grpc.CallOption) (*DelStaffResp, error) {
	out := new(DelStaffResp)
	err := grpc.Invoke(ctx, "/Staffmanarpc.staffmana/DelStaff", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffmanaClient) UpdateStaff(ctx context.Context, in *UpdateStaffReq, opts ...grpc.CallOption) (*UpdateStaffResp, error) {
	out := new(UpdateStaffResp)
	err := grpc.Invoke(ctx, "/Staffmanarpc.staffmana/UpdateStaff", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffmanaClient) SelectAllStaff(ctx context.Context, in *SelectAllStaffReq, opts ...grpc.CallOption) (*SelectAllStaffResp, error) {
	out := new(SelectAllStaffResp)
	err := grpc.Invoke(ctx, "/Staffmanarpc.staffmana/SelectAllStaff", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffmanaClient) SelectStaffById(ctx context.Context, in *SelectStaffByIdReq, opts ...grpc.CallOption) (*SelectStaffByIdResp, error) {
	out := new(SelectStaffByIdResp)
	err := grpc.Invoke(ctx, "/Staffmanarpc.staffmana/SelectStaffById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Staffmana service

type StaffmanaServer interface {
	AddStaff(context.Context, *AddStaffReq) (*AddStaffResp, error)
	DelStaff(context.Context, *DelStaffReq) (*DelStaffResp, error)
	UpdateStaff(context.Context, *UpdateStaffReq) (*UpdateStaffResp, error)
	SelectAllStaff(context.Context, *SelectAllStaffReq) (*SelectAllStaffResp, error)
	SelectStaffById(context.Context, *SelectStaffByIdReq) (*SelectStaffByIdResp, error)
}

func RegisterStaffmanaServer(s *grpc.Server, srv StaffmanaServer) {
	s.RegisterService(&_Staffmana_serviceDesc, srv)
}

func _Staffmana_AddStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStaffReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffmanaServer).AddStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Staffmanarpc.staffmana/AddStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffmanaServer).AddStaff(ctx, req.(*AddStaffReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staffmana_DelStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelStaffReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffmanaServer).DelStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Staffmanarpc.staffmana/DelStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffmanaServer).DelStaff(ctx, req.(*DelStaffReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staffmana_UpdateStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStaffReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffmanaServer).UpdateStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Staffmanarpc.staffmana/UpdateStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffmanaServer).UpdateStaff(ctx, req.(*UpdateStaffReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staffmana_SelectAllStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectAllStaffReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffmanaServer).SelectAllStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Staffmanarpc.staffmana/SelectAllStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffmanaServer).SelectAllStaff(ctx, req.(*SelectAllStaffReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staffmana_SelectStaffById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectStaffByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffmanaServer).SelectStaffById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Staffmanarpc.staffmana/SelectStaffById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffmanaServer).SelectStaffById(ctx, req.(*SelectStaffByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Staffmana_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Staffmanarpc.staffmana",
	HandlerType: (*StaffmanaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddStaff",
			Handler:    _Staffmana_AddStaff_Handler,
		},
		{
			MethodName: "DelStaff",
			Handler:    _Staffmana_DelStaff_Handler,
		},
		{
			MethodName: "UpdateStaff",
			Handler:    _Staffmana_UpdateStaff_Handler,
		},
		{
			MethodName: "SelectAllStaff",
			Handler:    _Staffmana_SelectAllStaff_Handler,
		},
		{
			MethodName: "SelectStaffById",
			Handler:    _Staffmana_SelectStaffById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staffmana.proto",
}

func init() { proto.RegisterFile("staffmana.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x69, 0x52, 0xfa, 0x67, 0x5a, 0x5a, 0x9c, 0x82, 0xc4, 0x60, 0xb1, 0x06, 0x0f, 0x05,
	0x21, 0x82, 0x3e, 0x41, 0xc5, 0x8b, 0x9e, 0x64, 0x43, 0xbd, 0xaf, 0xd9, 0x2d, 0x06, 0xd2, 0x66,
	0xed, 0x44, 0xa4, 0xaf, 0xe0, 0x53, 0x4b, 0x37, 0x26, 0xcd, 0xa6, 0x4d, 0xbc, 0x75, 0x66, 0xbe,
	0xfd, 0x75, 0x66, 0xbe, 0x09, 0x8c, 0x29, 0xe5, 0xab, 0xd5, 0x9a, 0x6f, 0xb8, 0xaf, 0xb6, 0x49,
	0x9a, 0xe0, 0x30, 0xc8, 0x13, 0x5b, 0x15, 0x7a, 0x4b, 0x18, 0x2c, 0x84, 0xd0, 0x29, 0x26, 0x3f,
	0x11, 0xa1, 0xbd, 0xe1, 0x6b, 0xe9, 0xb4, 0x66, 0xad, 0x79, 0x9f, 0xe9, 0xdf, 0x78, 0x0e, 0x1d,
	0xc5, 0x89, 0xbe, 0x85, 0x63, 0xe9, 0xec, 0x5f, 0x84, 0x2e, 0xf4, 0x54, 0x42, 0x51, 0x1a, 0x25,
	0x1b, 0xc7, 0x9e, 0xb5, 0xe6, 0x36, 0x2b, 0x62, 0x6f, 0x0e, 0xc3, 0x03, 0x96, 0x14, 0x3a, 0xd0,
	0x0d, 0xbe, 0xc2, 0x50, 0x12, 0x69, 0x74, 0x8f, 0x75, 0x29, 0x0b, 0xbd, 0x29, 0x0c, 0x9e, 0x64,
	0x5c, 0x34, 0x30, 0x02, 0x2b, 0x12, 0x5a, 0x63, 0x33, 0x2b, 0x12, 0x7b, 0xd0, 0xa1, 0xdc, 0x08,
	0xfa, 0x80, 0xd1, 0x52, 0x09, 0x9e, 0xca, 0x3a, 0x56, 0x31, 0x9c, 0x75, 0x72, 0x38, 0xbb, 0x76,
	0xb8, 0x76, 0x65, 0xb8, 0x5b, 0x18, 0x1b, 0xff, 0xd4, 0xd8, 0xd6, 0x04, 0xce, 0x02, 0x19, 0xcb,
	0x30, 0x5d, 0xc4, 0xc5, 0x94, 0x9e, 0x0f, 0x58, 0x4d, 0x66, 0x90, 0x57, 0xbe, 0x8b, 0x13, 0x9e,
	0x35, 0x3d, 0x64, 0x5d, 0x95, 0x85, 0xde, 0x4d, 0xae, 0xd7, 0xe2, 0xc7, 0xdd, 0xb3, 0x38, 0xb5,
	0xab, 0x3b, 0x98, 0x1c, 0xa9, 0x9a, 0xb0, 0xf7, 0x3f, 0x36, 0xf4, 0x8b, 0xf3, 0xc0, 0x05, 0xf4,
	0x72, 0xcf, 0xf0, 0xc2, 0x2f, 0x5f, 0x89, 0x5f, 0x3a, 0x11, 0xd7, 0xad, 0x2b, 0x91, 0xda, 0x23,
	0x72, 0xb7, 0xaa, 0x88, 0x92, 0xc9, 0x55, 0x84, 0x61, 0xf0, 0x0b, 0x0c, 0x4a, 0xcb, 0xc5, 0x4b,
	0x53, 0x6a, 0x3a, 0xec, 0x4e, 0x1b, 0xaa, 0xa4, 0x30, 0x80, 0x91, 0xb9, 0x66, 0xbc, 0x32, 0x1f,
	0x1c, 0x39, 0xe3, 0xce, 0x9a, 0x05, 0xa4, 0xf0, 0x0d, 0xc6, 0x95, 0x2d, 0xe3, 0xc9, 0x47, 0x65,
	0xab, 0xdc, 0xeb, 0x7f, 0x14, 0xa4, 0xde, 0x3b, 0xfa, 0xf3, 0x7c, 0xf8, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xc1, 0x34, 0xcf, 0x91, 0xb1, 0x03, 0x00, 0x00,
}
