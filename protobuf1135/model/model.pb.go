// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model.proto

package model

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Status int32

const (
	A Status = 0
	B Status = 1
	C Status = 2
)

var Status_name = map[int32]string{
	0: "A",
	1: "B",
	2: "C",
}

var Status_value = map[string]int32{
	"A": 0,
	"B": 1,
	"C": 2,
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}

func init() {
	proto.RegisterEnum("model.Status", Status_name, Status_value)
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xb4, 0x64, 0xb9, 0xd8, 0x82,
	0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x85, 0x58, 0xb9, 0x18, 0x1d, 0x05, 0x18, 0x40, 0x94, 0x93, 0x00,
	0x23, 0x88, 0x72, 0x16, 0x60, 0x72, 0x32, 0xb9, 0xf0, 0x50, 0x8e, 0xe1, 0xc6, 0x43, 0x39, 0x86,
	0x0f, 0x0f, 0xe5, 0x18, 0x1b, 0x1e, 0xc9, 0x31, 0xae, 0x78, 0x24, 0xc7, 0x78, 0xe2, 0x91, 0x1c,
	0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xbe, 0x78, 0x24, 0xc7, 0xf0, 0xe1, 0x91,
	0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x90, 0xc4,
	0x06, 0xb6, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x4d, 0xeb, 0xa8, 0x71, 0x00, 0x00,
	0x00,
}

func (x Status) String() string {
	s, ok := Status_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
