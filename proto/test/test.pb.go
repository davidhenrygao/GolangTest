// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package test is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
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

type CMD int32

const (
	CMD_LOGIN CMD = 1
)

var CMD_name = map[int32]string{
	1: "LOGIN",
}
var CMD_value = map[string]int32{
	"LOGIN": 1,
}

func (x CMD) Enum() *CMD {
	p := new(CMD)
	*p = x
	return p
}
func (x CMD) String() string {
	return proto.EnumName(CMD_name, int32(x))
}
func (x *CMD) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CMD_value, data, "CMD")
	if err != nil {
		return err
	}
	*x = CMD(value)
	return nil
}
func (CMD) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RetCode int32

const (
	RetCode_SUCCESS        RetCode = 0
	RetCode_PROTOCOL_ERROR RetCode = 1
	RetCode_WRONG_PASSWORD RetCode = 2
)

var RetCode_name = map[int32]string{
	0: "SUCCESS",
	1: "PROTOCOL_ERROR",
	2: "WRONG_PASSWORD",
}
var RetCode_value = map[string]int32{
	"SUCCESS":        0,
	"PROTOCOL_ERROR": 1,
	"WRONG_PASSWORD": 2,
}

func (x RetCode) Enum() *RetCode {
	p := new(RetCode)
	*p = x
	return p
}
func (x RetCode) String() string {
	return proto.EnumName(RetCode_name, int32(x))
}
func (x *RetCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RetCode_value, data, "RetCode")
	if err != nil {
		return err
	}
	*x = RetCode(value)
	return nil
}
func (RetCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Req struct {
	Account          *string `protobuf:"bytes,1,req,name=account" json:"account,omitempty"`
	Password         *string `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
	Token            *string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	GameserverId     *int32  `protobuf:"varint,4,opt,name=gameserver_id" json:"gameserver_id,omitempty"`
	Msg              *string `protobuf:"bytes,5,opt,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Req) Reset()                    { *m = Req{} }
func (m *Req) String() string            { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()               {}
func (*Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Req) GetAccount() string {
	if m != nil && m.Account != nil {
		return *m.Account
	}
	return ""
}

func (m *Req) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *Req) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *Req) GetGameserverId() int32 {
	if m != nil && m.GameserverId != nil {
		return *m.GameserverId
	}
	return 0
}

func (m *Req) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

type RespBody struct {
	Secrect          *string `protobuf:"bytes,1,opt,name=secrect" json:"secrect,omitempty"`
	MsgBack          *string `protobuf:"bytes,2,opt,name=msg_back" json:"msg_back,omitempty"`
	Options          []int32 `protobuf:"varint,3,rep,name=options" json:"options,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RespBody) Reset()                    { *m = RespBody{} }
func (m *RespBody) String() string            { return proto.CompactTextString(m) }
func (*RespBody) ProtoMessage()               {}
func (*RespBody) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RespBody) GetSecrect() string {
	if m != nil && m.Secrect != nil {
		return *m.Secrect
	}
	return ""
}

func (m *RespBody) GetMsgBack() string {
	if m != nil && m.MsgBack != nil {
		return *m.MsgBack
	}
	return ""
}

func (m *RespBody) GetOptions() []int32 {
	if m != nil {
		return m.Options
	}
	return nil
}

type Resp struct {
	Rc               *int32    `protobuf:"varint,1,req,name=rc" json:"rc,omitempty"`
	RespBody         *RespBody `protobuf:"bytes,2,opt,name=resp_body" json:"resp_body,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Resp) Reset()                    { *m = Resp{} }
func (m *Resp) String() string            { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()               {}
func (*Resp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Resp) GetRc() int32 {
	if m != nil && m.Rc != nil {
		return *m.Rc
	}
	return 0
}

func (m *Resp) GetRespBody() *RespBody {
	if m != nil {
		return m.RespBody
	}
	return nil
}

func init() {
	proto.RegisterType((*Req)(nil), "test.Req")
	proto.RegisterType((*RespBody)(nil), "test.RespBody")
	proto.RegisterType((*Resp)(nil), "test.Resp")
	proto.RegisterEnum("test.CMD", CMD_name, CMD_value)
	proto.RegisterEnum("test.RetCode", RetCode_name, RetCode_value)
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xc1, 0x6b, 0xfa, 0x30,
	0x14, 0x80, 0x7f, 0x6d, 0x0c, 0xb5, 0x4f, 0xd4, 0x10, 0xf8, 0x41, 0x8e, 0x9d, 0xa7, 0xe2, 0xc1,
	0xc3, 0x60, 0x57, 0x61, 0xab, 0x22, 0x03, 0x67, 0x24, 0xdd, 0xf0, 0xb0, 0x43, 0xa9, 0xe9, 0x43,
	0x44, 0x6c, 0xba, 0x24, 0xdb, 0xf0, 0xbf, 0x1f, 0xe9, 0xd8, 0xed, 0xf1, 0xf1, 0x3e, 0xbe, 0xf7,
	0x00, 0x3c, 0x3a, 0xbf, 0xe8, 0xac, 0xf1, 0x86, 0x0f, 0xc2, 0x3c, 0x7b, 0x07, 0xa2, 0xf0, 0x83,
	0x4f, 0x21, 0xa9, 0xb5, 0x36, 0x9f, 0xad, 0x17, 0x51, 0x16, 0xe7, 0x29, 0x67, 0x30, 0xec, 0x6a,
	0xe7, 0xbe, 0x8d, 0x6d, 0x44, 0xdc, 0x93, 0x31, 0x50, 0x6f, 0x2e, 0xd8, 0x0a, 0x92, 0x45, 0x79,
	0xca, 0xff, 0xc3, 0xf8, 0x54, 0x5f, 0xd1, 0xa1, 0xfd, 0x42, 0x5b, 0x9d, 0x1b, 0x31, 0xc8, 0xa2,
	0x9c, 0xf2, 0x11, 0x90, 0xab, 0x3b, 0x09, 0x1a, 0x76, 0x66, 0x4b, 0x18, 0x2a, 0x74, 0xdd, 0x93,
	0x69, 0x6e, 0xa1, 0xe0, 0x50, 0x5b, 0xd4, 0xa1, 0x10, 0xfd, 0x16, 0xae, 0xee, 0x54, 0x1d, 0x6b,
	0x7d, 0x11, 0x71, 0x4f, 0xa6, 0x90, 0x98, 0xce, 0x9f, 0x4d, 0xeb, 0x04, 0xc9, 0x48, 0x4e, 0x67,
	0x0f, 0x30, 0x08, 0x3e, 0x07, 0x88, 0xad, 0xee, 0x0f, 0xa3, 0xfc, 0x0e, 0x52, 0x8b, 0xae, 0xab,
	0x8e, 0xa6, 0xb9, 0xf5, 0xde, 0xe8, 0x7e, 0xb2, 0xe8, 0xdf, 0xfa, 0x4b, 0xcd, 0x19, 0x90, 0xe2,
	0x65, 0xc5, 0x53, 0xa0, 0x5b, 0xb9, 0x79, 0xde, 0xb1, 0x68, 0xbe, 0x84, 0x44, 0xa1, 0x2f, 0x4c,
	0x83, 0x7c, 0x04, 0x49, 0xf9, 0x56, 0x14, 0xeb, 0xb2, 0x64, 0xff, 0x38, 0x87, 0xc9, 0x5e, 0xc9,
	0x57, 0x59, 0xc8, 0x6d, 0xb5, 0x56, 0x4a, 0x2a, 0x16, 0x05, 0x76, 0x50, 0x72, 0xb7, 0xa9, 0xf6,
	0x8f, 0x65, 0x79, 0x90, 0x6a, 0xc5, 0xe2, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0x3b, 0xd8,
	0xd4, 0x38, 0x01, 0x00, 0x00,
}