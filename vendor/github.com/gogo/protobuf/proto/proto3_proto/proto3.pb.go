// Code generated by protoc-gen-gogo.
// source: proto3_proto/proto3.proto
// DO NOT EDIT!

/*
Package proto3_proto is a generated protocol buffer package.

It is generated from these files:
	proto3_proto/proto3.proto

It has these top-level messages:
	Message
	Nested
	MessageWithMap
*/
package proto3_proto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import testdata "github.com/gogo/protobuf/proto/testdata"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Message_Humour int32

const (
	Message_UNKNOWN     Message_Humour = 0
	Message_PUNS        Message_Humour = 1
	Message_SLAPSTICK   Message_Humour = 2
	Message_BILL_BAILEY Message_Humour = 3
)

var Message_Humour_name = map[int32]string{
	0: "UNKNOWN",
	1: "PUNS",
	2: "SLAPSTICK",
	3: "BILL_BAILEY",
}
var Message_Humour_value = map[string]int32{
	"UNKNOWN":     0,
	"PUNS":        1,
	"SLAPSTICK":   2,
	"BILL_BAILEY": 3,
}

func (x Message_Humour) String() string {
	return proto.EnumName(Message_Humour_name, int32(x))
}
func (Message_Humour) EnumDescriptor() ([]byte, []int) { return fileDescriptorProto3, []int{0, 0} }

type Message struct {
	Name         string                           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hilarity     Message_Humour                   `protobuf:"varint,2,opt,name=hilarity,proto3,enum=proto3_proto.Message_Humour" json:"hilarity,omitempty"`
	HeightInCm   uint32                           `protobuf:"varint,3,opt,name=height_in_cm,json=heightInCm,proto3" json:"height_in_cm,omitempty"`
	Data         []byte                           `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	ResultCount  int64                            `protobuf:"varint,7,opt,name=result_count,json=resultCount,proto3" json:"result_count,omitempty"`
	TrueScotsman bool                             `protobuf:"varint,8,opt,name=true_scotsman,json=trueScotsman,proto3" json:"true_scotsman,omitempty"`
	Score        float32                          `protobuf:"fixed32,9,opt,name=score,proto3" json:"score,omitempty"`
	Key          []uint64                         `protobuf:"varint,5,rep,name=key" json:"key,omitempty"`
	Nested       *Nested                          `protobuf:"bytes,6,opt,name=nested" json:"nested,omitempty"`
	RFunny       []Message_Humour                 `protobuf:"varint,16,rep,name=r_funny,json=rFunny,enum=proto3_proto.Message_Humour" json:"r_funny,omitempty"`
	Terrain      map[string]*Nested               `protobuf:"bytes,10,rep,name=terrain" json:"terrain,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	Proto2Field  *testdata.SubDefaults            `protobuf:"bytes,11,opt,name=proto2_field,json=proto2Field" json:"proto2_field,omitempty"`
	Proto2Value  map[string]*testdata.SubDefaults `protobuf:"bytes,13,rep,name=proto2_value,json=proto2Value" json:"proto2_value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptorProto3, []int{0} }

func (m *Message) GetNested() *Nested {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *Message) GetTerrain() map[string]*Nested {
	if m != nil {
		return m.Terrain
	}
	return nil
}

func (m *Message) GetProto2Field() *testdata.SubDefaults {
	if m != nil {
		return m.Proto2Field
	}
	return nil
}

func (m *Message) GetProto2Value() map[string]*testdata.SubDefaults {
	if m != nil {
		return m.Proto2Value
	}
	return nil
}

type Nested struct {
	Bunny string `protobuf:"bytes,1,opt,name=bunny,proto3" json:"bunny,omitempty"`
	Cute  bool   `protobuf:"varint,2,opt,name=cute,proto3" json:"cute,omitempty"`
}

func (m *Nested) Reset()                    { *m = Nested{} }
func (m *Nested) String() string            { return proto.CompactTextString(m) }
func (*Nested) ProtoMessage()               {}
func (*Nested) Descriptor() ([]byte, []int) { return fileDescriptorProto3, []int{1} }

type MessageWithMap struct {
	ByteMapping map[bool][]byte `protobuf:"bytes,1,rep,name=byte_mapping,json=byteMapping" json:"byte_mapping,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *MessageWithMap) Reset()                    { *m = MessageWithMap{} }
func (m *MessageWithMap) String() string            { return proto.CompactTextString(m) }
func (*MessageWithMap) ProtoMessage()               {}
func (*MessageWithMap) Descriptor() ([]byte, []int) { return fileDescriptorProto3, []int{2} }

func (m *MessageWithMap) GetByteMapping() map[bool][]byte {
	if m != nil {
		return m.ByteMapping
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "proto3_proto.Message")
	proto.RegisterType((*Nested)(nil), "proto3_proto.Nested")
	proto.RegisterType((*MessageWithMap)(nil), "proto3_proto.MessageWithMap")
	proto.RegisterEnum("proto3_proto.Message_Humour", Message_Humour_name, Message_Humour_value)
}

func init() { proto.RegisterFile("proto3_proto/proto3.proto", fileDescriptorProto3) }

var fileDescriptorProto3 = []byte{
	// 585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x52, 0xef, 0x6a, 0xdb, 0x3e,
	0x14, 0xfd, 0x29, 0x4e, 0x9d, 0xf4, 0xda, 0xe9, 0xcf, 0x88, 0x0e, 0xb4, 0xb0, 0x0f, 0x5a, 0x06,
	0xc3, 0xec, 0x8f, 0x0b, 0x19, 0x83, 0x32, 0xc6, 0x46, 0xdb, 0xb5, 0x2c, 0x34, 0xcd, 0x82, 0xd2,
	0xae, 0xec, 0x93, 0xb1, 0x53, 0xc5, 0x31, 0x8b, 0xed, 0x60, 0x4b, 0x03, 0x3f, 0xc5, 0xde, 0x61,
	0x4f, 0x3a, 0x2c, 0x39, 0xad, 0x5b, 0xb2, 0xed, 0x93, 0xaf, 0x8e, 0xcf, 0xbd, 0xe7, 0xe8, 0x5c,
	0xc1, 0xe3, 0x75, 0x9e, 0x89, 0xec, 0x8d, 0xaf, 0x3e, 0x07, 0xfa, 0xe0, 0xa9, 0x0f, 0xb6, 0x9b,
	0xbf, 0xfa, 0xc3, 0x28, 0x16, 0x4b, 0x19, 0x7a, 0xf3, 0x2c, 0x39, 0x88, 0xb2, 0xa8, 0xe6, 0x86,
	0x72, 0xa1, 0x8b, 0x03, 0xc1, 0x0b, 0x71, 0x13, 0x88, 0x40, 0x15, 0x7a, 0xc2, 0xe0, 0xa7, 0x09,
	0x9d, 0x0b, 0x5e, 0x14, 0x41, 0xc4, 0x31, 0x86, 0x76, 0x1a, 0x24, 0x9c, 0x20, 0x8a, 0xdc, 0x5d,
	0xa6, 0x6a, 0x7c, 0x08, 0xdd, 0x65, 0xbc, 0x0a, 0xf2, 0x58, 0x94, 0xa4, 0x45, 0x91, 0xbb, 0x37,
	0x7c, 0xe2, 0x35, 0x45, 0xbd, 0xba, 0xd9, 0xfb, 0x2c, 0x93, 0x4c, 0xe6, 0xec, 0x96, 0x8d, 0x29,
	0xd8, 0x4b, 0x1e, 0x47, 0x4b, 0xe1, 0xc7, 0xa9, 0x3f, 0x4f, 0x88, 0x41, 0x91, 0xdb, 0x63, 0xa0,
	0xb1, 0x51, 0x7a, 0x92, 0x54, 0x7a, 0x95, 0x1d, 0xd2, 0xa6, 0xc8, 0xb5, 0x99, 0xaa, 0xf1, 0x53,
	0xb0, 0x73, 0x5e, 0xc8, 0x95, 0xf0, 0xe7, 0x99, 0x4c, 0x05, 0xe9, 0x50, 0xe4, 0x1a, 0xcc, 0xd2,
	0xd8, 0x49, 0x05, 0xe1, 0x67, 0xd0, 0x13, 0xb9, 0xe4, 0x7e, 0x31, 0xcf, 0x44, 0x91, 0x04, 0x29,
	0xe9, 0x52, 0xe4, 0x76, 0x99, 0x5d, 0x81, 0xb3, 0x1a, 0xc3, 0xfb, 0xb0, 0x53, 0xcc, 0xb3, 0x9c,
	0x93, 0x5d, 0x8a, 0xdc, 0x16, 0xd3, 0x07, 0xec, 0x80, 0xf1, 0x9d, 0x97, 0x64, 0x87, 0x1a, 0x6e,
	0x9b, 0x55, 0x25, 0x7e, 0x05, 0x66, 0xca, 0x0b, 0xc1, 0x6f, 0x88, 0x49, 0x91, 0x6b, 0x0d, 0xf7,
	0xef, 0xdf, 0x6e, 0xa2, 0xfe, 0xb1, 0x9a, 0x83, 0xdf, 0x42, 0x27, 0xf7, 0x17, 0x32, 0x4d, 0x4b,
	0xe2, 0x50, 0xe3, 0x9f, 0x61, 0x98, 0xf9, 0x59, 0xc5, 0xc5, 0xef, 0xa1, 0x23, 0x78, 0x9e, 0x07,
	0x71, 0x4a, 0x80, 0x1a, 0xae, 0x35, 0x1c, 0x6c, 0x6f, 0xbb, 0xd4, 0xa4, 0xd3, 0x54, 0xe4, 0x25,
	0xdb, 0xb4, 0xe0, 0x43, 0xd0, 0x6b, 0x1e, 0xfa, 0x8b, 0x98, 0xaf, 0x6e, 0x88, 0xa5, 0x8c, 0x3e,
	0xf2, 0x36, 0xeb, 0xf4, 0x66, 0x32, 0xfc, 0xc4, 0x17, 0x81, 0x5c, 0x89, 0x82, 0x59, 0x9a, 0x7a,
	0x56, 0x31, 0xf1, 0xe8, 0xb6, 0xf3, 0x47, 0xb0, 0x92, 0x9c, 0xf4, 0x94, 0xf8, 0xf3, 0xed, 0xe2,
	0x53, 0xc5, 0xfc, 0x5a, 0x11, 0xb5, 0x81, 0x7a, 0x94, 0x42, 0xfa, 0x53, 0xb0, 0x9b, 0xee, 0x36,
	0x49, 0xea, 0xa7, 0xa2, 0x92, 0x7c, 0x01, 0x3b, 0x5a, 0xa5, 0xf5, 0x97, 0x20, 0x35, 0xe5, 0x5d,
	0xeb, 0x10, 0xf5, 0xaf, 0xc0, 0x79, 0x28, 0xb9, 0x65, 0xea, 0xcb, 0xfb, 0x53, 0xff, 0x70, 0xeb,
	0xbb, 0xb1, 0x83, 0x8f, 0x60, 0xea, 0xf4, 0xb1, 0x05, 0x9d, 0xab, 0xc9, 0xf9, 0xe4, 0xcb, 0xf5,
	0xc4, 0xf9, 0x0f, 0x77, 0xa1, 0x3d, 0xbd, 0x9a, 0xcc, 0x1c, 0x84, 0x7b, 0xb0, 0x3b, 0x1b, 0x1f,
	0x4d, 0x67, 0x97, 0xa3, 0x93, 0x73, 0xa7, 0x85, 0xff, 0x07, 0xeb, 0x78, 0x34, 0x1e, 0xfb, 0xc7,
	0x47, 0xa3, 0xf1, 0xe9, 0x37, 0xc7, 0x18, 0x0c, 0xc1, 0xd4, 0x66, 0xab, 0x37, 0x14, 0xaa, 0x5d,
	0x6b, 0x3f, 0xfa, 0x50, 0xbd, 0xda, 0xb9, 0x14, 0xda, 0x50, 0x97, 0xa9, 0x7a, 0xf0, 0x0b, 0xc1,
	0x5e, 0x9d, 0xe3, 0x75, 0x2c, 0x96, 0x17, 0xc1, 0x1a, 0x4f, 0xc1, 0x0e, 0x4b, 0xc1, 0xfd, 0x24,
	0x58, 0xaf, 0xe3, 0x34, 0x22, 0x48, 0x65, 0xff, 0x7a, 0x6b, 0xf6, 0x75, 0x8f, 0x77, 0x5c, 0x0a,
	0x7e, 0xa1, 0xf9, 0xf5, 0x0a, 0xc2, 0x3b, 0xa4, 0xff, 0x01, 0x9c, 0x87, 0x84, 0x66, 0x60, 0x5d,
	0x1d, 0xd8, 0x7e, 0x33, 0x30, 0xbb, 0x91, 0x4c, 0x68, 0x6a, 0xe9, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x75, 0x74, 0x87, 0xb4, 0x50, 0x04, 0x00, 0x00,
}
