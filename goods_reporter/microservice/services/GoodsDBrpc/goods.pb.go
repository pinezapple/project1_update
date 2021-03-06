// Code generated by protoc-gen-go. DO NOT EDIT.
// source: goods.proto

/*
Package GoodsDBrpc is a generated protocol buffer package.

It is generated from these files:
	goods.proto

It has these top-level messages:
	AddReq
	AddResp
	UpdateReq
	UpdateResp
	DelReq
	DelResp
	SelectAllReq
	SelectAllResp
	SelectByIDReq
	SelectByIDResp
	SelectByNameReq
	SelectByNameResp
	SelectByGoodsPriceReq
	SelectByGoodsPriceResp
	IfExistReq
	IfExistResp
*/
package GoodsDBrpc

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

type AddReq struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Quantity int64  `protobuf:"varint,2,opt,name=quantity" json:"quantity,omitempty"`
	Price    int64  `protobuf:"varint,3,opt,name=price" json:"price,omitempty"`
}

func (m *AddReq) Reset()                    { *m = AddReq{} }
func (m *AddReq) String() string            { return proto.CompactTextString(m) }
func (*AddReq) ProtoMessage()               {}
func (*AddReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddReq) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *AddReq) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type AddResp struct {
	Success bool `protobuf:"varint,1,opt,name=Success,json=success" json:"Success,omitempty"`
}

func (m *AddResp) Reset()                    { *m = AddResp{} }
func (m *AddResp) String() string            { return proto.CompactTextString(m) }
func (*AddResp) ProtoMessage()               {}
func (*AddResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type UpdateReq struct {
	Id       int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Quantity int64  `protobuf:"varint,3,opt,name=quantity" json:"quantity,omitempty"`
	Price    int64  `protobuf:"varint,4,opt,name=price" json:"price,omitempty"`
}

func (m *UpdateReq) Reset()                    { *m = UpdateReq{} }
func (m *UpdateReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateReq) ProtoMessage()               {}
func (*UpdateReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateReq) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *UpdateReq) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type UpdateResp struct {
	Success bool `protobuf:"varint,1,opt,name=Success,json=success" json:"Success,omitempty"`
}

func (m *UpdateResp) Reset()                    { *m = UpdateResp{} }
func (m *UpdateResp) String() string            { return proto.CompactTextString(m) }
func (*UpdateResp) ProtoMessage()               {}
func (*UpdateResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type DelReq struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DelReq) Reset()                    { *m = DelReq{} }
func (m *DelReq) String() string            { return proto.CompactTextString(m) }
func (*DelReq) ProtoMessage()               {}
func (*DelReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DelReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DelResp struct {
	Success bool `protobuf:"varint,1,opt,name=Success,json=success" json:"Success,omitempty"`
}

func (m *DelResp) Reset()                    { *m = DelResp{} }
func (m *DelResp) String() string            { return proto.CompactTextString(m) }
func (*DelResp) ProtoMessage()               {}
func (*DelResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DelResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type SelectAllReq struct {
}

func (m *SelectAllReq) Reset()                    { *m = SelectAllReq{} }
func (m *SelectAllReq) String() string            { return proto.CompactTextString(m) }
func (*SelectAllReq) ProtoMessage()               {}
func (*SelectAllReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type SelectAllResp struct {
	Payload []byte `protobuf:"bytes,1,opt,name=Payload,json=payload,proto3" json:"Payload,omitempty"`
}

func (m *SelectAllResp) Reset()                    { *m = SelectAllResp{} }
func (m *SelectAllResp) String() string            { return proto.CompactTextString(m) }
func (*SelectAllResp) ProtoMessage()               {}
func (*SelectAllResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SelectAllResp) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type SelectByIDReq struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *SelectByIDReq) Reset()                    { *m = SelectByIDReq{} }
func (m *SelectByIDReq) String() string            { return proto.CompactTextString(m) }
func (*SelectByIDReq) ProtoMessage()               {}
func (*SelectByIDReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SelectByIDReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SelectByIDResp struct {
	Payload []byte `protobuf:"bytes,1,opt,name=Payload,json=payload,proto3" json:"Payload,omitempty"`
}

func (m *SelectByIDResp) Reset()                    { *m = SelectByIDResp{} }
func (m *SelectByIDResp) String() string            { return proto.CompactTextString(m) }
func (*SelectByIDResp) ProtoMessage()               {}
func (*SelectByIDResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SelectByIDResp) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type SelectByNameReq struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *SelectByNameReq) Reset()                    { *m = SelectByNameReq{} }
func (m *SelectByNameReq) String() string            { return proto.CompactTextString(m) }
func (*SelectByNameReq) ProtoMessage()               {}
func (*SelectByNameReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SelectByNameReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SelectByNameResp struct {
	Payload []byte `protobuf:"bytes,1,opt,name=Payload,json=payload,proto3" json:"Payload,omitempty"`
}

func (m *SelectByNameResp) Reset()                    { *m = SelectByNameResp{} }
func (m *SelectByNameResp) String() string            { return proto.CompactTextString(m) }
func (*SelectByNameResp) ProtoMessage()               {}
func (*SelectByNameResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *SelectByNameResp) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type SelectByGoodsPriceReq struct {
	Price int64 `protobuf:"varint,1,opt,name=price" json:"price,omitempty"`
}

func (m *SelectByGoodsPriceReq) Reset()                    { *m = SelectByGoodsPriceReq{} }
func (m *SelectByGoodsPriceReq) String() string            { return proto.CompactTextString(m) }
func (*SelectByGoodsPriceReq) ProtoMessage()               {}
func (*SelectByGoodsPriceReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *SelectByGoodsPriceReq) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type SelectByGoodsPriceResp struct {
	Payload []byte `protobuf:"bytes,1,opt,name=Payload,json=payload,proto3" json:"Payload,omitempty"`
}

func (m *SelectByGoodsPriceResp) Reset()                    { *m = SelectByGoodsPriceResp{} }
func (m *SelectByGoodsPriceResp) String() string            { return proto.CompactTextString(m) }
func (*SelectByGoodsPriceResp) ProtoMessage()               {}
func (*SelectByGoodsPriceResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SelectByGoodsPriceResp) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type IfExistReq struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *IfExistReq) Reset()                    { *m = IfExistReq{} }
func (m *IfExistReq) String() string            { return proto.CompactTextString(m) }
func (*IfExistReq) ProtoMessage()               {}
func (*IfExistReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *IfExistReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type IfExistResp struct {
	Success bool `protobuf:"varint,1,opt,name=Success,json=success" json:"Success,omitempty"`
}

func (m *IfExistResp) Reset()                    { *m = IfExistResp{} }
func (m *IfExistResp) String() string            { return proto.CompactTextString(m) }
func (*IfExistResp) ProtoMessage()               {}
func (*IfExistResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *IfExistResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*AddReq)(nil), "GoodsDBrpc.AddReq")
	proto.RegisterType((*AddResp)(nil), "GoodsDBrpc.AddResp")
	proto.RegisterType((*UpdateReq)(nil), "GoodsDBrpc.UpdateReq")
	proto.RegisterType((*UpdateResp)(nil), "GoodsDBrpc.UpdateResp")
	proto.RegisterType((*DelReq)(nil), "GoodsDBrpc.DelReq")
	proto.RegisterType((*DelResp)(nil), "GoodsDBrpc.DelResp")
	proto.RegisterType((*SelectAllReq)(nil), "GoodsDBrpc.SelectAllReq")
	proto.RegisterType((*SelectAllResp)(nil), "GoodsDBrpc.SelectAllResp")
	proto.RegisterType((*SelectByIDReq)(nil), "GoodsDBrpc.SelectByIDReq")
	proto.RegisterType((*SelectByIDResp)(nil), "GoodsDBrpc.SelectByIDResp")
	proto.RegisterType((*SelectByNameReq)(nil), "GoodsDBrpc.SelectByNameReq")
	proto.RegisterType((*SelectByNameResp)(nil), "GoodsDBrpc.SelectByNameResp")
	proto.RegisterType((*SelectByGoodsPriceReq)(nil), "GoodsDBrpc.SelectByGoodsPriceReq")
	proto.RegisterType((*SelectByGoodsPriceResp)(nil), "GoodsDBrpc.SelectByGoodsPriceResp")
	proto.RegisterType((*IfExistReq)(nil), "GoodsDBrpc.IfExistReq")
	proto.RegisterType((*IfExistResp)(nil), "GoodsDBrpc.IfExistResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Goods service

type GoodsClient interface {
	AddGoods(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddResp, error)
	UpdateGoods(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error)
	DelGoods(ctx context.Context, in *DelReq, opts ...grpc.CallOption) (*DelResp, error)
	SelectAllGoods(ctx context.Context, in *SelectAllReq, opts ...grpc.CallOption) (*SelectAllResp, error)
	SelectByGoodsID(ctx context.Context, in *SelectByIDReq, opts ...grpc.CallOption) (*SelectByIDResp, error)
	SelectByGoodsName(ctx context.Context, in *SelectByNameReq, opts ...grpc.CallOption) (*SelectByNameResp, error)
	SelectByGoodsPrice(ctx context.Context, in *SelectByGoodsPriceReq, opts ...grpc.CallOption) (*SelectByGoodsPriceResp, error)
	IfGoodsExist(ctx context.Context, in *IfExistReq, opts ...grpc.CallOption) (*IfExistResp, error)
}

type goodsClient struct {
	cc *grpc.ClientConn
}

func NewGoodsClient(cc *grpc.ClientConn) GoodsClient {
	return &goodsClient{cc}
}

func (c *goodsClient) AddGoods(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddResp, error) {
	out := new(AddResp)
	err := grpc.Invoke(ctx, "/GoodsDBrpc.goods/AddGoods", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateGoods(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error) {
	out := new(UpdateResp)
	err := grpc.Invoke(ctx, "/GoodsDBrpc.goods/UpdateGoods", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) DelGoods(ctx context.Context, in *DelReq, opts ...grpc.CallOption) (*DelResp, error) {
	out := new(DelResp)
	err := grpc.Invoke(ctx, "/GoodsDBrpc.goods/DelGoods", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) SelectAllGoods(ctx context.Context, in *SelectAllReq, opts ...grpc.CallOption) (*SelectAllResp, error) {
	out := new(SelectAllResp)
	err := grpc.Invoke(ctx, "/GoodsDBrpc.goods/SelectAllGoods", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) SelectByGoodsID(ctx context.Context, in *SelectByIDReq, opts ...grpc.CallOption) (*SelectByIDResp, error) {
	out := new(SelectByIDResp)
	err := grpc.Invoke(ctx, "/GoodsDBrpc.goods/SelectByGoodsID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) SelectByGoodsName(ctx context.Context, in *SelectByNameReq, opts ...grpc.CallOption) (*SelectByNameResp, error) {
	out := new(SelectByNameResp)
	err := grpc.Invoke(ctx, "/GoodsDBrpc.goods/SelectByGoodsName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) SelectByGoodsPrice(ctx context.Context, in *SelectByGoodsPriceReq, opts ...grpc.CallOption) (*SelectByGoodsPriceResp, error) {
	out := new(SelectByGoodsPriceResp)
	err := grpc.Invoke(ctx, "/GoodsDBrpc.goods/SelectByGoodsPrice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) IfGoodsExist(ctx context.Context, in *IfExistReq, opts ...grpc.CallOption) (*IfExistResp, error) {
	out := new(IfExistResp)
	err := grpc.Invoke(ctx, "/GoodsDBrpc.goods/IfGoodsExist", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Goods service

type GoodsServer interface {
	AddGoods(context.Context, *AddReq) (*AddResp, error)
	UpdateGoods(context.Context, *UpdateReq) (*UpdateResp, error)
	DelGoods(context.Context, *DelReq) (*DelResp, error)
	SelectAllGoods(context.Context, *SelectAllReq) (*SelectAllResp, error)
	SelectByGoodsID(context.Context, *SelectByIDReq) (*SelectByIDResp, error)
	SelectByGoodsName(context.Context, *SelectByNameReq) (*SelectByNameResp, error)
	SelectByGoodsPrice(context.Context, *SelectByGoodsPriceReq) (*SelectByGoodsPriceResp, error)
	IfGoodsExist(context.Context, *IfExistReq) (*IfExistResp, error)
}

func RegisterGoodsServer(s *grpc.Server, srv GoodsServer) {
	s.RegisterService(&_Goods_serviceDesc, srv)
}

func _Goods_AddGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).AddGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GoodsDBrpc.goods/AddGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).AddGoods(ctx, req.(*AddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GoodsDBrpc.goods/UpdateGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateGoods(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_DelGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).DelGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GoodsDBrpc.goods/DelGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).DelGoods(ctx, req.(*DelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_SelectAllGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).SelectAllGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GoodsDBrpc.goods/SelectAllGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).SelectAllGoods(ctx, req.(*SelectAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_SelectByGoodsID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).SelectByGoodsID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GoodsDBrpc.goods/SelectByGoodsID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).SelectByGoodsID(ctx, req.(*SelectByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_SelectByGoodsName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectByNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).SelectByGoodsName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GoodsDBrpc.goods/SelectByGoodsName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).SelectByGoodsName(ctx, req.(*SelectByNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_SelectByGoodsPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectByGoodsPriceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).SelectByGoodsPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GoodsDBrpc.goods/SelectByGoodsPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).SelectByGoodsPrice(ctx, req.(*SelectByGoodsPriceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_IfGoodsExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IfExistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).IfGoodsExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GoodsDBrpc.goods/IfGoodsExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).IfGoodsExist(ctx, req.(*IfExistReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Goods_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GoodsDBrpc.goods",
	HandlerType: (*GoodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddGoods",
			Handler:    _Goods_AddGoods_Handler,
		},
		{
			MethodName: "UpdateGoods",
			Handler:    _Goods_UpdateGoods_Handler,
		},
		{
			MethodName: "DelGoods",
			Handler:    _Goods_DelGoods_Handler,
		},
		{
			MethodName: "SelectAllGoods",
			Handler:    _Goods_SelectAllGoods_Handler,
		},
		{
			MethodName: "SelectByGoodsID",
			Handler:    _Goods_SelectByGoodsID_Handler,
		},
		{
			MethodName: "SelectByGoodsName",
			Handler:    _Goods_SelectByGoodsName_Handler,
		},
		{
			MethodName: "SelectByGoodsPrice",
			Handler:    _Goods_SelectByGoodsPrice_Handler,
		},
		{
			MethodName: "IfGoodsExist",
			Handler:    _Goods_IfGoodsExist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goods.proto",
}

func init() { proto.RegisterFile("goods.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x65, 0x3b, 0x8d, 0xd3, 0x49, 0x08, 0x30, 0xd0, 0x62, 0x96, 0x4a, 0x14, 0x23, 0xa0,
	0x20, 0xc8, 0xa1, 0xbd, 0x22, 0x21, 0x57, 0xae, 0xc0, 0x97, 0xa8, 0x72, 0xc5, 0x89, 0x93, 0xf1,
	0x6e, 0x91, 0x25, 0x27, 0xde, 0x64, 0x1d, 0x89, 0xbc, 0x14, 0xcf, 0x88, 0x76, 0x37, 0x76, 0x6c,
	0x65, 0xed, 0xde, 0x32, 0x3b, 0xff, 0xfc, 0x33, 0x99, 0xf9, 0x64, 0x18, 0xff, 0x29, 0x0a, 0x2a,
	0x66, 0x7c, 0x5d, 0x94, 0x05, 0xc2, 0x77, 0x19, 0x84, 0xd7, 0x6b, 0x9e, 0xfa, 0x73, 0x18, 0x06,
	0x94, 0xc6, 0x6c, 0x85, 0x08, 0x83, 0x65, 0xb2, 0x60, 0x9e, 0x75, 0x6e, 0x5d, 0x1c, 0xc7, 0xea,
	0x37, 0x12, 0x18, 0xad, 0x36, 0xc9, 0xb2, 0xcc, 0xca, 0xad, 0x67, 0x9f, 0x5b, 0x17, 0x4e, 0x5c,
	0xc7, 0xf8, 0x1c, 0x8e, 0xf8, 0x3a, 0x4b, 0x99, 0xe7, 0xa8, 0x84, 0x0e, 0xfc, 0xb7, 0xe0, 0x2a,
	0x3f, 0xc1, 0xd1, 0x03, 0xf7, 0x6e, 0x93, 0xa6, 0x4c, 0x08, 0xe5, 0x39, 0x8a, 0x5d, 0xa1, 0x43,
	0x3f, 0x81, 0xe3, 0x9f, 0x9c, 0x26, 0x25, 0x93, 0x7d, 0xa7, 0x60, 0x67, 0x54, 0x29, 0x9c, 0xd8,
	0xce, 0x68, 0x3d, 0x87, 0xdd, 0x31, 0x87, 0xd3, 0x35, 0xc7, 0xa0, 0x39, 0xc7, 0x7b, 0x80, 0xaa,
	0x45, 0xef, 0x28, 0x1e, 0x0c, 0x43, 0x96, 0x1b, 0xe6, 0x90, 0xff, 0x44, 0x65, 0x7a, 0xcb, 0xa7,
	0x30, 0xb9, 0x63, 0x39, 0x4b, 0xcb, 0x20, 0x97, 0x26, 0xfe, 0x47, 0x78, 0xd4, 0x88, 0x75, 0xe9,
	0x6d, 0xb2, 0xcd, 0x8b, 0x44, 0x5b, 0x4f, 0x62, 0x97, 0xeb, 0xd0, 0x7f, 0x5d, 0x49, 0xaf, 0xb7,
	0x51, 0x68, 0x1a, 0xe0, 0x13, 0x4c, 0x9b, 0x82, 0x5e, 0xb3, 0x77, 0xf0, 0xb8, 0xd2, 0xce, 0x93,
	0x05, 0xeb, 0xb8, 0xa7, 0xff, 0x19, 0x9e, 0xb4, 0x65, 0xbd, 0xa6, 0x5f, 0xe0, 0xa4, 0x52, 0x2b,
	0x62, 0x6e, 0xe5, 0x66, 0xa5, 0x75, 0xbd, 0x72, 0xab, 0xb9, 0xf2, 0x4b, 0x38, 0x35, 0xc9, 0x7b,
	0x5b, 0x9c, 0x01, 0x44, 0xf7, 0x37, 0x7f, 0x33, 0x51, 0x9a, 0x36, 0xf0, 0x01, 0xc6, 0x75, 0xb6,
	0xef, 0x0c, 0x97, 0xff, 0x06, 0x70, 0xa4, 0x08, 0xc7, 0x2b, 0x18, 0x05, 0x94, 0xaa, 0xfe, 0x88,
	0xb3, 0x3d, 0xe8, 0x33, 0x4d, 0x39, 0x79, 0x76, 0xf0, 0x26, 0x38, 0x7e, 0x85, 0xb1, 0x86, 0x45,
	0xd7, 0x9d, 0x34, 0x35, 0x35, 0xa8, 0xe4, 0xd4, 0xf4, 0x2c, 0xb8, 0x6c, 0x19, 0xb2, 0xdc, 0xd0,
	0x52, 0x83, 0xd5, 0x6e, 0x59, 0x21, 0x75, 0x53, 0x1d, 0x37, 0xc8, 0x77, 0xa5, 0x5e, 0x53, 0xd6,
	0x84, 0x8a, 0xbc, 0xec, 0xc8, 0x08, 0x8e, 0x3f, 0xf6, 0x77, 0x57, 0x9a, 0x28, 0x44, 0x83, 0x7a,
	0x47, 0x18, 0x21, 0x5d, 0x29, 0xc1, 0x71, 0x0e, 0x4f, 0x5b, 0x4e, 0x92, 0x0f, 0x7c, 0x65, 0x2a,
	0xd8, 0x01, 0x46, 0xce, 0xba, 0x93, 0x82, 0xe3, 0x2f, 0xc0, 0x43, 0x1a, 0xf0, 0x8d, 0xa9, 0xa6,
	0x05, 0x17, 0xf1, 0x1f, 0x92, 0x08, 0x8e, 0xdf, 0x60, 0x12, 0xdd, 0xab, 0x37, 0x45, 0x07, 0xb6,
	0x4e, 0xb3, 0x07, 0x8a, 0xbc, 0x30, 0xbe, 0x0b, 0xfe, 0x7b, 0xa8, 0xbe, 0x84, 0x57, 0xff, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xda, 0x0b, 0xca, 0xbc, 0x18, 0x05, 0x00, 0x00,
}
