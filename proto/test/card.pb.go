// Code generated by protoc-gen-go. DO NOT EDIT.
// source: card.proto

/*
Package test is a generated protocol buffer package.

It is generated from these files:
	card.proto
	test.proto

It has these top-level messages:
	Card
	Req
	RespBody
	Resp
*/
package test

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Card struct {
	Id               *uint32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Level            *uint32 `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	Amount           *uint32 `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	State            *uint32 `protobuf:"varint,4,opt,name=state" json:"state,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Card) Reset()                    { *m = Card{} }
func (m *Card) String() string            { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()               {}
func (*Card) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Card) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Card) GetLevel() uint32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *Card) GetAmount() uint32 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

func (m *Card) GetState() uint32 {
	if m != nil && m.State != nil {
		return *m.State
	}
	return 0
}

func init() {
	proto.RegisterType((*Card)(nil), "proto.card")
}

func init() { proto.RegisterFile("card.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4e, 0x2c, 0x4a,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x0e, 0x5c, 0x2c, 0x20, 0x41,
	0x21, 0x2e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x26, 0x0d, 0x5e, 0x21, 0x5e, 0x2e, 0xd6,
	0x9c, 0xd4, 0xb2, 0xd4, 0x1c, 0x09, 0x26, 0x30, 0x97, 0x8f, 0x8b, 0x2d, 0x31, 0x37, 0xbf, 0x34,
	0xaf, 0x44, 0x82, 0x59, 0x81, 0x11, 0x22, 0x5d, 0x5c, 0x92, 0x58, 0x92, 0x2a, 0xc1, 0x02, 0xe2,
	0x3a, 0x19, 0x44, 0xe9, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7,
	0x24, 0x96, 0x65, 0xa6, 0x64, 0xa4, 0xe6, 0x15, 0x55, 0xa6, 0x27, 0xe6, 0xeb, 0xbb, 0xe7, 0xe7,
	0x24, 0xe6, 0xa5, 0x87, 0xa4, 0x16, 0x97, 0xe8, 0x83, 0xad, 0xd3, 0x2f, 0x49, 0x2d, 0x2e, 0x01,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x7f, 0xf4, 0xf9, 0x87, 0x00, 0x00, 0x00,
}