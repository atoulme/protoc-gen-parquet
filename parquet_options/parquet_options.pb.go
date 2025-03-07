// Code generated by protoc-gen-go. DO NOT EDIT.
// source: parquet_options/parquet_options.proto

package parquet_options

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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

type TimestampType int32

const (
	TimestampType_UNSPECIFIED      TimestampType = 0
	TimestampType_TIMESTAMP_MILLIS TimestampType = 1
	TimestampType_TIMESTAMP_MICROS TimestampType = 2
	TimestampType_TIMESTAMP_NANOS  TimestampType = 3
)

var TimestampType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "TIMESTAMP_MILLIS",
	2: "TIMESTAMP_MICROS",
	3: "TIMESTAMP_NANOS",
}

var TimestampType_value = map[string]int32{
	"UNSPECIFIED":      0,
	"TIMESTAMP_MILLIS": 1,
	"TIMESTAMP_MICROS": 2,
	"TIMESTAMP_NANOS":  3,
}

func (x TimestampType) String() string {
	return proto.EnumName(TimestampType_name, int32(x))
}

func (TimestampType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_306afa5d4177b3d5, []int{0}
}

type MessageOptions struct {
	// As long as table_name is not blank,
	// a schema is generated for top-level messages in each file.
	TableName            string   `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageOptions) Reset()         { *m = MessageOptions{} }
func (m *MessageOptions) String() string { return proto.CompactTextString(m) }
func (*MessageOptions) ProtoMessage()    {}
func (*MessageOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_306afa5d4177b3d5, []int{0}
}

func (m *MessageOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageOptions.Unmarshal(m, b)
}
func (m *MessageOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageOptions.Marshal(b, m, deterministic)
}
func (m *MessageOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageOptions.Merge(m, src)
}
func (m *MessageOptions) XXX_Size() int {
	return xxx_messageInfo_MessageOptions.Size(m)
}
func (m *MessageOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageOptions.DiscardUnknown(m)
}

var xxx_messageInfo_MessageOptions proto.InternalMessageInfo

func (m *MessageOptions) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

type FieldOptions struct {
	TimestampType        TimestampType `protobuf:"varint,1,opt,name=timestamp_type,json=timestampType,proto3,enum=parquet_options.TimestampType" json:"timestamp_type,omitempty"`
	IsJson               bool          `protobuf:"varint,2,opt,name=is_json,json=isJson,proto3" json:"is_json,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FieldOptions) Reset()         { *m = FieldOptions{} }
func (m *FieldOptions) String() string { return proto.CompactTextString(m) }
func (*FieldOptions) ProtoMessage()    {}
func (*FieldOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_306afa5d4177b3d5, []int{1}
}

func (m *FieldOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldOptions.Unmarshal(m, b)
}
func (m *FieldOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldOptions.Marshal(b, m, deterministic)
}
func (m *FieldOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldOptions.Merge(m, src)
}
func (m *FieldOptions) XXX_Size() int {
	return xxx_messageInfo_FieldOptions.Size(m)
}
func (m *FieldOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldOptions.DiscardUnknown(m)
}

var xxx_messageInfo_FieldOptions proto.InternalMessageInfo

func (m *FieldOptions) GetTimestampType() TimestampType {
	if m != nil {
		return m.TimestampType
	}
	return TimestampType_UNSPECIFIED
}

func (m *FieldOptions) GetIsJson() bool {
	if m != nil {
		return m.IsJson
	}
	return false
}

var E_MessageOpts = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.MessageOptions)(nil),
	ExtensionType: (*MessageOptions)(nil),
	Field:         1023,
	Name:          "parquet_options.message_opts",
	Tag:           "bytes,1023,opt,name=message_opts",
	Filename:      "parquet_options/parquet_options.proto",
}

var E_FieldOpts = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.FieldOptions)(nil),
	ExtensionType: (*FieldOptions)(nil),
	Field:         1023,
	Name:          "parquet_options.field_opts",
	Tag:           "bytes,1023,opt,name=field_opts",
	Filename:      "parquet_options/parquet_options.proto",
}

func init() {
	proto.RegisterEnum("parquet_options.TimestampType", TimestampType_name, TimestampType_value)
	proto.RegisterType((*MessageOptions)(nil), "parquet_options.MessageOptions")
	proto.RegisterType((*FieldOptions)(nil), "parquet_options.FieldOptions")
	proto.RegisterExtension(E_MessageOpts)
	proto.RegisterExtension(E_FieldOpts)
}

func init() {
	proto.RegisterFile("parquet_options/parquet_options.proto", fileDescriptor_306afa5d4177b3d5)
}

var fileDescriptor_306afa5d4177b3d5 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xef, 0x4b, 0xc2, 0x40,
	0x18, 0x6e, 0x06, 0x9a, 0xaf, 0xbf, 0xc6, 0x0a, 0x92, 0xc0, 0x1c, 0x42, 0x20, 0x81, 0x1b, 0x18,
	0x11, 0xf8, 0xcd, 0x6c, 0xc2, 0xc2, 0x4d, 0xd9, 0xd6, 0x97, 0x08, 0xc6, 0x4d, 0xcf, 0xb5, 0xd8,
	0xed, 0xd6, 0xee, 0xf6, 0xc1, 0xbf, 0xbe, 0xd0, 0xb9, 0xd2, 0xf9, 0xf1, 0x79, 0xee, 0x79, 0xee,
	0x79, 0xef, 0xb9, 0x17, 0xee, 0x62, 0x94, 0x7c, 0xa7, 0x98, 0xbb, 0x34, 0xe6, 0x01, 0x8d, 0x98,
	0x5a, 0xc0, 0x4a, 0x9c, 0x50, 0x4e, 0xa5, 0x56, 0x81, 0xbe, 0x91, 0x7d, 0x4a, 0xfd, 0x10, 0xab,
	0xbb, 0x63, 0x2f, 0x5d, 0xab, 0x2b, 0xcc, 0x96, 0x49, 0x10, 0x73, 0x9a, 0x64, 0x96, 0x9e, 0x0a,
	0x4d, 0x03, 0x33, 0x86, 0x7c, 0x3c, 0xcf, 0x3c, 0x52, 0x07, 0x80, 0x23, 0x2f, 0xc4, 0x6e, 0x84,
	0x08, 0x6e, 0x0b, 0xb2, 0xd0, 0xaf, 0x5a, 0xd5, 0x1d, 0x63, 0x22, 0x82, 0x7b, 0x11, 0xd4, 0xa7,
	0x01, 0x0e, 0x57, 0xb9, 0x5c, 0x83, 0x26, 0x0f, 0x08, 0x66, 0x1c, 0x91, 0xd8, 0xe5, 0x9b, 0x38,
	0xb3, 0x34, 0x87, 0xb7, 0x4a, 0x71, 0x46, 0x27, 0x97, 0x39, 0x9b, 0x18, 0x5b, 0x0d, 0x7e, 0x08,
	0xa5, 0x6b, 0xa8, 0x04, 0xcc, 0xfd, 0x62, 0x34, 0x6a, 0x97, 0x64, 0xa1, 0x7f, 0x61, 0x95, 0x03,
	0xf6, 0xca, 0x68, 0x74, 0x8f, 0xa0, 0x71, 0x64, 0x94, 0x5a, 0x50, 0x7b, 0x33, 0xed, 0x85, 0x36,
	0xd1, 0xa7, 0xba, 0xf6, 0x22, 0x9e, 0x49, 0x57, 0x20, 0x3a, 0xba, 0xa1, 0xd9, 0xce, 0xd8, 0x58,
	0xb8, 0x86, 0x3e, 0x9b, 0xe9, 0xb6, 0x28, 0x14, 0xd9, 0x89, 0x35, 0xb7, 0xc5, 0x92, 0x74, 0x09,
	0xad, 0x7f, 0xd6, 0x1c, 0x9b, 0x73, 0x5b, 0x3c, 0x1f, 0xad, 0xa0, 0x4e, 0xb2, 0x0e, 0xb6, 0xb3,
	0x32, 0xa9, 0xab, 0x64, 0xb5, 0x29, 0x79, 0x6d, 0xca, 0x71, 0x45, 0xed, 0x9f, 0x8a, 0x2c, 0xf4,
	0x6b, 0xc3, 0xee, 0xc9, 0x13, 0x8f, 0x75, 0x56, 0x8d, 0xfc, 0x61, 0x36, 0xfa, 0x00, 0x58, 0x6f,
	0x8b, 0xcb, 0x32, 0x3a, 0x27, 0x19, 0x87, 0xad, 0xe6, 0x09, 0x9d, 0x93, 0x84, 0x43, 0x95, 0x55,
	0x5d, 0xef, 0x11, 0x7b, 0x7e, 0x7a, 0x7f, 0xf4, 0x03, 0xfe, 0x99, 0x7a, 0xca, 0x92, 0x12, 0x15,
	0x71, 0x9a, 0x86, 0x64, 0xff, 0xef, 0xcb, 0x81, 0x8f, 0xa3, 0xc1, 0xfe, 0xa2, 0xe2, 0xe6, 0x78,
	0xe5, 0x9d, 0xe6, 0xe1, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x16, 0x91, 0x0e, 0x63, 0x02, 0x00,
	0x00,
}
