// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderserver.proto

package orderserver

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CreateOrderRequ struct {
	OrderTitle           string   `protobuf:"bytes,1,opt,name=orderTitle,proto3" json:"orderTitle,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderRequ) Reset()         { *m = CreateOrderRequ{} }
func (m *CreateOrderRequ) String() string { return proto.CompactTextString(m) }
func (*CreateOrderRequ) ProtoMessage()    {}
func (*CreateOrderRequ) Descriptor() ([]byte, []int) {
	return fileDescriptor_df21fb255d4dca2e, []int{0}
}

func (m *CreateOrderRequ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderRequ.Unmarshal(m, b)
}
func (m *CreateOrderRequ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderRequ.Marshal(b, m, deterministic)
}
func (m *CreateOrderRequ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderRequ.Merge(m, src)
}
func (m *CreateOrderRequ) XXX_Size() int {
	return xxx_messageInfo_CreateOrderRequ.Size(m)
}
func (m *CreateOrderRequ) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderRequ.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderRequ proto.InternalMessageInfo

func (m *CreateOrderRequ) GetOrderTitle() string {
	if m != nil {
		return m.OrderTitle
	}
	return ""
}

func (m *CreateOrderRequ) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CreateOrderResp struct {
	OrderAmount          float64  `protobuf:"fixed64,1,opt,name=orderAmount,proto3" json:"orderAmount,omitempty"`
	OrderNumber          int64    `protobuf:"varint,2,opt,name=orderNumber,proto3" json:"orderNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderResp) Reset()         { *m = CreateOrderResp{} }
func (m *CreateOrderResp) String() string { return proto.CompactTextString(m) }
func (*CreateOrderResp) ProtoMessage()    {}
func (*CreateOrderResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_df21fb255d4dca2e, []int{1}
}

func (m *CreateOrderResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderResp.Unmarshal(m, b)
}
func (m *CreateOrderResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderResp.Marshal(b, m, deterministic)
}
func (m *CreateOrderResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderResp.Merge(m, src)
}
func (m *CreateOrderResp) XXX_Size() int {
	return xxx_messageInfo_CreateOrderResp.Size(m)
}
func (m *CreateOrderResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderResp proto.InternalMessageInfo

func (m *CreateOrderResp) GetOrderAmount() float64 {
	if m != nil {
		return m.OrderAmount
	}
	return 0
}

func (m *CreateOrderResp) GetOrderNumber() int64 {
	if m != nil {
		return m.OrderNumber
	}
	return 0
}

type GetOrderRequ struct {
	OrderNumber          int64    `protobuf:"varint,1,opt,name=orderNumber,proto3" json:"orderNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderRequ) Reset()         { *m = GetOrderRequ{} }
func (m *GetOrderRequ) String() string { return proto.CompactTextString(m) }
func (*GetOrderRequ) ProtoMessage()    {}
func (*GetOrderRequ) Descriptor() ([]byte, []int) {
	return fileDescriptor_df21fb255d4dca2e, []int{2}
}

func (m *GetOrderRequ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderRequ.Unmarshal(m, b)
}
func (m *GetOrderRequ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderRequ.Marshal(b, m, deterministic)
}
func (m *GetOrderRequ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderRequ.Merge(m, src)
}
func (m *GetOrderRequ) XXX_Size() int {
	return xxx_messageInfo_GetOrderRequ.Size(m)
}
func (m *GetOrderRequ) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderRequ.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderRequ proto.InternalMessageInfo

func (m *GetOrderRequ) GetOrderNumber() int64 {
	if m != nil {
		return m.OrderNumber
	}
	return 0
}

type GetOrderResp struct {
	OrderNumber          int64    `protobuf:"varint,1,opt,name=orderNumber,proto3" json:"orderNumber,omitempty"`
	OrderTitle           string   `protobuf:"bytes,2,opt,name=orderTitle,proto3" json:"orderTitle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderResp) Reset()         { *m = GetOrderResp{} }
func (m *GetOrderResp) String() string { return proto.CompactTextString(m) }
func (*GetOrderResp) ProtoMessage()    {}
func (*GetOrderResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_df21fb255d4dca2e, []int{3}
}

func (m *GetOrderResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderResp.Unmarshal(m, b)
}
func (m *GetOrderResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderResp.Marshal(b, m, deterministic)
}
func (m *GetOrderResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderResp.Merge(m, src)
}
func (m *GetOrderResp) XXX_Size() int {
	return xxx_messageInfo_GetOrderResp.Size(m)
}
func (m *GetOrderResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderResp proto.InternalMessageInfo

func (m *GetOrderResp) GetOrderNumber() int64 {
	if m != nil {
		return m.OrderNumber
	}
	return 0
}

func (m *GetOrderResp) GetOrderTitle() string {
	if m != nil {
		return m.OrderTitle
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateOrderRequ)(nil), "proto.CreateOrderRequ")
	proto.RegisterType((*CreateOrderResp)(nil), "proto.CreateOrderResp")
	proto.RegisterType((*GetOrderRequ)(nil), "proto.GetOrderRequ")
	proto.RegisterType((*GetOrderResp)(nil), "proto.GetOrderResp")
}

func init() {
	proto.RegisterFile("orderserver.proto", fileDescriptor_df21fb255d4dca2e)
}

var fileDescriptor_df21fb255d4dca2e = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x2f, 0x4a, 0x49,
	0x2d, 0x2a, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05,
	0x53, 0x4a, 0x9e, 0x5c, 0xfc, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9, 0xfe, 0x20, 0x15, 0x41, 0xa9,
	0x85, 0xa5, 0x42, 0x72, 0x5c, 0x5c, 0x60, 0xe5, 0x21, 0x99, 0x25, 0x39, 0xa9, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x48, 0x22, 0x42, 0x62, 0x5c, 0x6c, 0xa5, 0xc5, 0xa9, 0x45, 0x9e, 0x29,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x50, 0x9e, 0x52, 0x28, 0x9a, 0x51, 0xc5, 0x05, 0x42,
	0x0a, 0x5c, 0xdc, 0x60, 0x8d, 0x8e, 0xb9, 0xf9, 0xa5, 0x79, 0x25, 0x60, 0xb3, 0x18, 0x83, 0x90,
	0x85, 0xe0, 0x2a, 0xfc, 0x4a, 0x73, 0x93, 0x52, 0x8b, 0xa0, 0x26, 0x22, 0x0b, 0x29, 0x19, 0x70,
	0xf1, 0xb8, 0xa7, 0x96, 0x20, 0x9c, 0x87, 0xa6, 0x83, 0x11, 0x53, 0x47, 0x00, 0xb2, 0x0e, 0x24,
	0x57, 0xe0, 0xd4, 0x81, 0xe6, 0x65, 0x26, 0x74, 0x2f, 0x1b, 0xb5, 0x31, 0x72, 0x71, 0x83, 0xcd,
	0x0b, 0x06, 0x07, 0xa1, 0x90, 0x3d, 0x17, 0x37, 0x92, 0x57, 0x85, 0xc4, 0x20, 0x61, 0xaa, 0x87,
	0x16, 0x92, 0x52, 0x58, 0xc5, 0x8b, 0x0b, 0x94, 0x18, 0x84, 0xcc, 0xb8, 0x38, 0x60, 0x4e, 0x14,
	0x12, 0x86, 0xaa, 0x42, 0xf6, 0xa5, 0x14, 0xa6, 0x20, 0x48, 0x5f, 0x12, 0x1b, 0x58, 0xd4, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xf1, 0x36, 0x9a, 0xd1, 0x01, 0x00, 0x00,
}
