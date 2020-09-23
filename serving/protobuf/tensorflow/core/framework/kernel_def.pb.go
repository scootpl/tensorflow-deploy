// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/kernel_def.proto

package framework // import "github.com/grupawp/tensorflow-deploy/serving/protobuf/tensorflow/core/framework"

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

type KernelDef struct {
	// Must match the name of an Op.
	Op string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	// Type of device this kernel runs on.
	DeviceType string                      `protobuf:"bytes,2,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	Constraint []*KernelDef_AttrConstraint `protobuf:"bytes,3,rep,name=constraint,proto3" json:"constraint,omitempty"`
	// Names of the Op's input_/output_args that reside in host memory
	// instead of device memory.
	HostMemoryArg []string `protobuf:"bytes,4,rep,name=host_memory_arg,json=hostMemoryArg,proto3" json:"host_memory_arg,omitempty"`
	// This allows experimental kernels to be registered for an op that
	// won't be used unless the user specifies a "_kernel" attr with
	// value matching this.
	Label string `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	// Prioritization of kernel amongst different devices. By default we assume
	// priority is 0. The higher the priority the better. By default (i.e. if
	// this is not set), we prefer GPU kernels over CPU.
	Priority             int32    `protobuf:"varint,6,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KernelDef) Reset()         { *m = KernelDef{} }
func (m *KernelDef) String() string { return proto.CompactTextString(m) }
func (*KernelDef) ProtoMessage()    {}
func (*KernelDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_kernel_def_b23c0cb3dcd8b347, []int{0}
}
func (m *KernelDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelDef.Unmarshal(m, b)
}
func (m *KernelDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelDef.Marshal(b, m, deterministic)
}
func (dst *KernelDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelDef.Merge(dst, src)
}
func (m *KernelDef) XXX_Size() int {
	return xxx_messageInfo_KernelDef.Size(m)
}
func (m *KernelDef) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelDef.DiscardUnknown(m)
}

var xxx_messageInfo_KernelDef proto.InternalMessageInfo

func (m *KernelDef) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *KernelDef) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *KernelDef) GetConstraint() []*KernelDef_AttrConstraint {
	if m != nil {
		return m.Constraint
	}
	return nil
}

func (m *KernelDef) GetHostMemoryArg() []string {
	if m != nil {
		return m.HostMemoryArg
	}
	return nil
}

func (m *KernelDef) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *KernelDef) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

type KernelDef_AttrConstraint struct {
	// Name of an attr from the Op.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of values that this kernel supports for this attr.
	// Like OpDef.AttrDef.allowed_values, except for kernels instead of Ops.
	AllowedValues        *AttrValue `protobuf:"bytes,2,opt,name=allowed_values,json=allowedValues,proto3" json:"allowed_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *KernelDef_AttrConstraint) Reset()         { *m = KernelDef_AttrConstraint{} }
func (m *KernelDef_AttrConstraint) String() string { return proto.CompactTextString(m) }
func (*KernelDef_AttrConstraint) ProtoMessage()    {}
func (*KernelDef_AttrConstraint) Descriptor() ([]byte, []int) {
	return fileDescriptor_kernel_def_b23c0cb3dcd8b347, []int{0, 0}
}
func (m *KernelDef_AttrConstraint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelDef_AttrConstraint.Unmarshal(m, b)
}
func (m *KernelDef_AttrConstraint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelDef_AttrConstraint.Marshal(b, m, deterministic)
}
func (dst *KernelDef_AttrConstraint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelDef_AttrConstraint.Merge(dst, src)
}
func (m *KernelDef_AttrConstraint) XXX_Size() int {
	return xxx_messageInfo_KernelDef_AttrConstraint.Size(m)
}
func (m *KernelDef_AttrConstraint) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelDef_AttrConstraint.DiscardUnknown(m)
}

var xxx_messageInfo_KernelDef_AttrConstraint proto.InternalMessageInfo

func (m *KernelDef_AttrConstraint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KernelDef_AttrConstraint) GetAllowedValues() *AttrValue {
	if m != nil {
		return m.AllowedValues
	}
	return nil
}

// A collection of KernelDefs
type KernelList struct {
	Kernel               []*KernelDef `protobuf:"bytes,1,rep,name=kernel,proto3" json:"kernel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *KernelList) Reset()         { *m = KernelList{} }
func (m *KernelList) String() string { return proto.CompactTextString(m) }
func (*KernelList) ProtoMessage()    {}
func (*KernelList) Descriptor() ([]byte, []int) {
	return fileDescriptor_kernel_def_b23c0cb3dcd8b347, []int{1}
}
func (m *KernelList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelList.Unmarshal(m, b)
}
func (m *KernelList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelList.Marshal(b, m, deterministic)
}
func (dst *KernelList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelList.Merge(dst, src)
}
func (m *KernelList) XXX_Size() int {
	return xxx_messageInfo_KernelList.Size(m)
}
func (m *KernelList) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelList.DiscardUnknown(m)
}

var xxx_messageInfo_KernelList proto.InternalMessageInfo

func (m *KernelList) GetKernel() []*KernelDef {
	if m != nil {
		return m.Kernel
	}
	return nil
}

func init() {
	proto.RegisterType((*KernelDef)(nil), "tensorflow.KernelDef")
	proto.RegisterType((*KernelDef_AttrConstraint)(nil), "tensorflow.KernelDef.AttrConstraint")
	proto.RegisterType((*KernelList)(nil), "tensorflow.KernelList")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/kernel_def.proto", fileDescriptor_kernel_def_b23c0cb3dcd8b347)
}

var fileDescriptor_kernel_def_b23c0cb3dcd8b347 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x8a, 0xdb, 0x40,
	0x0c, 0xc6, 0xb1, 0xf3, 0x87, 0x46, 0x21, 0x09, 0x0c, 0x0d, 0x0c, 0xb9, 0xd4, 0x84, 0x52, 0x4c,
	0xa1, 0x36, 0xa4, 0xc7, 0xb6, 0x87, 0xa4, 0xb9, 0xb5, 0x85, 0x60, 0xca, 0x1e, 0xf6, 0x62, 0xc6,
	0x8e, 0xec, 0x98, 0xd8, 0x96, 0x19, 0x4f, 0x12, 0xfc, 0x76, 0xfb, 0x58, 0x7b, 0x5c, 0x3c, 0x0e,
	0x8e, 0x77, 0x09, 0x7b, 0x1b, 0x49, 0x9f, 0xa4, 0xf9, 0xf4, 0x83, 0xaf, 0x0a, 0xf3, 0x92, 0x64,
	0x94, 0xd2, 0xc5, 0x0d, 0x49, 0xa2, 0x1b, 0x49, 0x91, 0xe1, 0x85, 0xe4, 0xd1, 0x3d, 0xa2, 0xcc,
	0x31, 0xf5, 0xf7, 0x18, 0x39, 0x85, 0x24, 0x45, 0x0c, 0x6e, 0xda, 0xc5, 0x3b, 0x7d, 0x42, 0x29,
	0xe9, 0x9f, 0x45, 0x7a, 0xc2, 0xa6, 0x6f, 0xf9, 0x64, 0xc2, 0xe8, 0x8f, 0x1e, 0xb6, 0xc5, 0x88,
	0x4d, 0xc1, 0xa4, 0x82, 0x1b, 0x96, 0x61, 0x8f, 0x3c, 0x93, 0x0a, 0xf6, 0x09, 0xc6, 0x7b, 0x3c,
	0x27, 0x21, 0xfa, 0xaa, 0x2a, 0x90, 0x9b, 0xba, 0x00, 0x4d, 0xea, 0x7f, 0x55, 0x20, 0xdb, 0x02,
	0x84, 0x94, 0x97, 0x4a, 0x8a, 0x24, 0x57, 0xbc, 0x67, 0xf5, 0xec, 0xf1, 0xea, 0xb3, 0x73, 0xdb,
	0xef, 0xb4, 0xb3, 0x9d, 0xb5, 0x52, 0xf2, 0x77, 0xab, 0xf5, 0x3a, 0x7d, 0xec, 0x0b, 0xcc, 0x0e,
	0x54, 0x2a, 0x3f, 0xc3, 0x8c, 0x64, 0xe5, 0x0b, 0x19, 0xf3, 0xbe, 0xd5, 0xb3, 0x47, 0xde, 0xa4,
	0x4e, 0xff, 0xd3, 0xd9, 0xb5, 0x8c, 0xd9, 0x47, 0x18, 0xa4, 0x22, 0xc0, 0x94, 0x0f, 0xf4, 0x47,
	0x9a, 0x80, 0x2d, 0xe0, 0x43, 0x21, 0x13, 0x92, 0x89, 0xaa, 0xf8, 0xd0, 0x32, 0xec, 0x81, 0xd7,
	0xc6, 0x8b, 0x00, 0xa6, 0xaf, 0xf7, 0x32, 0x06, 0xfd, 0x5c, 0x64, 0x78, 0x35, 0xa9, 0xdf, 0xec,
	0x27, 0x4c, 0x45, 0x9a, 0xd2, 0x05, 0xf7, 0xcd, 0x6d, 0x4a, 0xed, 0x74, 0xbc, 0x9a, 0x77, 0x9d,
	0xd4, 0x73, 0x1e, 0xea, 0xaa, 0x37, 0xb9, 0x8a, 0x75, 0x54, 0x2e, 0x7f, 0x00, 0x34, 0x2e, 0xff,
	0x26, 0xa5, 0x62, 0xdf, 0x60, 0xd8, 0xc0, 0xe1, 0x86, 0xbe, 0xc6, 0xfc, 0xee, 0x35, 0xbc, 0xab,
	0x68, 0x43, 0xc0, 0x49, 0xc6, 0x5d, 0x4d, 0x0b, 0x6b, 0x33, 0x6b, 0xe5, 0xbb, 0x9a, 0x55, 0xb9,
	0x33, 0x1e, 0x7f, 0xc5, 0x89, 0x3a, 0x9c, 0x02, 0x27, 0xa4, 0xcc, 0xed, 0x50, 0xbe, 0xff, 0x8c,
	0xe9, 0x0d, 0xfe, 0x67, 0xc3, 0x08, 0x86, 0x9a, 0xfb, 0xf7, 0x97, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x4e, 0x89, 0x26, 0xb6, 0x5d, 0x02, 0x00, 0x00,
}
