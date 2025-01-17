// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package zpd_proto

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

type Status struct {
	//code = 1 means success
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type MessageResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageResponse) Reset()         { *m = MessageResponse{} }
func (m *MessageResponse) String() string { return proto.CompactTextString(m) }
func (*MessageResponse) ProtoMessage()    {}
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *MessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageResponse.Unmarshal(m, b)
}
func (m *MessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageResponse.Marshal(b, m, deterministic)
}
func (m *MessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageResponse.Merge(m, src)
}
func (m *MessageResponse) XXX_Size() int {
	return xxx_messageInfo_MessageResponse.Size(m)
}
func (m *MessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessageResponse proto.InternalMessageInfo

func (m *MessageResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "zpd.common.proto.Status")
	proto.RegisterType((*MessageResponse)(nil), "zpd.common.proto.MessageResponse")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xa8, 0x2a, 0x48, 0xd1, 0x43, 0x16,
	0x51, 0x32, 0xe2, 0x62, 0x0b, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x16, 0x12, 0xe2, 0x62, 0x49, 0xce,
	0x4f, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0xb3, 0x85, 0x44, 0xb8, 0x58, 0x53,
	0x8b, 0x8a, 0xf2, 0x8b, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x25, 0x67, 0x2e,
	0x7e, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0xd4, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54,
	0x21, 0x03, 0x2e, 0xb6, 0x62, 0xb0, 0x31, 0x60, 0xed, 0xdc, 0x46, 0x12, 0x7a, 0xe8, 0x36, 0xe9,
	0x41, 0xac, 0x09, 0x82, 0xaa, 0x73, 0xe2, 0x0d, 0x60, 0x8c, 0xe2, 0x04, 0x29, 0x02, 0xcb, 0x26,
	0xb1, 0x81, 0x29, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x52, 0x05, 0x73, 0xa7, 0xb0, 0x00,
	0x00, 0x00,
}
