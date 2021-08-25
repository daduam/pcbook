// Code generated by protoc-gen-go. DO NOT EDIT.
// source: laptop_message.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Laptop struct {
	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand    string     `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Name     string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Cpu      *CPU       `protobuf:"bytes,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Ram      *Memory    `protobuf:"bytes,5,opt,name=ram,proto3" json:"ram,omitempty"`
	Gpus     []*GPU     `protobuf:"bytes,6,rep,name=gpus,proto3" json:"gpus,omitempty"`
	Storages []*Storage `protobuf:"bytes,7,rep,name=storages,proto3" json:"storages,omitempty"`
	Screen   *Screen    `protobuf:"bytes,8,opt,name=screen,proto3" json:"screen,omitempty"`
	Keyboard *Keyboard  `protobuf:"bytes,9,opt,name=keyboard,proto3" json:"keyboard,omitempty"`
	// Types that are valid to be assigned to Weight:
	//	*Laptop_WeightKg
	//	*Laptop_WeightLb
	Weight               isLaptop_Weight      `protobuf_oneof:"weight"`
	PriceUsd             float64              `protobuf:"fixed64,12,opt,name=price_usd,json=priceUsd,proto3" json:"price_usd,omitempty"`
	ReleaseYear          uint32               `protobuf:"varint,13,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Laptop) Reset()         { *m = Laptop{} }
func (m *Laptop) String() string { return proto.CompactTextString(m) }
func (*Laptop) ProtoMessage()    {}
func (*Laptop) Descriptor() ([]byte, []int) {
	return fileDescriptor_07a3824d46f4b731, []int{0}
}

func (m *Laptop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Laptop.Unmarshal(m, b)
}
func (m *Laptop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Laptop.Marshal(b, m, deterministic)
}
func (m *Laptop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Laptop.Merge(m, src)
}
func (m *Laptop) XXX_Size() int {
	return xxx_messageInfo_Laptop.Size(m)
}
func (m *Laptop) XXX_DiscardUnknown() {
	xxx_messageInfo_Laptop.DiscardUnknown(m)
}

var xxx_messageInfo_Laptop proto.InternalMessageInfo

func (m *Laptop) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Laptop) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Laptop) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Laptop) GetCpu() *CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *Laptop) GetRam() *Memory {
	if m != nil {
		return m.Ram
	}
	return nil
}

func (m *Laptop) GetGpus() []*GPU {
	if m != nil {
		return m.Gpus
	}
	return nil
}

func (m *Laptop) GetStorages() []*Storage {
	if m != nil {
		return m.Storages
	}
	return nil
}

func (m *Laptop) GetScreen() *Screen {
	if m != nil {
		return m.Screen
	}
	return nil
}

func (m *Laptop) GetKeyboard() *Keyboard {
	if m != nil {
		return m.Keyboard
	}
	return nil
}

type isLaptop_Weight interface {
	isLaptop_Weight()
}

type Laptop_WeightKg struct {
	WeightKg float64 `protobuf:"fixed64,10,opt,name=weight_kg,json=weightKg,proto3,oneof"`
}

type Laptop_WeightLb struct {
	WeightLb float64 `protobuf:"fixed64,11,opt,name=weight_lb,json=weightLb,proto3,oneof"`
}

func (*Laptop_WeightKg) isLaptop_Weight() {}

func (*Laptop_WeightLb) isLaptop_Weight() {}

func (m *Laptop) GetWeight() isLaptop_Weight {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *Laptop) GetWeightKg() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightKg); ok {
		return x.WeightKg
	}
	return 0
}

func (m *Laptop) GetWeightLb() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightLb); ok {
		return x.WeightLb
	}
	return 0
}

func (m *Laptop) GetPriceUsd() float64 {
	if m != nil {
		return m.PriceUsd
	}
	return 0
}

func (m *Laptop) GetReleaseYear() uint32 {
	if m != nil {
		return m.ReleaseYear
	}
	return 0
}

func (m *Laptop) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Laptop) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Laptop_WeightKg)(nil),
		(*Laptop_WeightLb)(nil),
	}
}

func init() {
	proto.RegisterType((*Laptop)(nil), "pcbook.Laptop")
}

func init() { proto.RegisterFile("laptop_message.proto", fileDescriptor_07a3824d46f4b731) }

var fileDescriptor_07a3824d46f4b731 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x49, 0xd3, 0x85, 0xe4, 0x75, 0x2b, 0xc8, 0x2a, 0x60, 0x15, 0x4d, 0x0b, 0x1c, 0x50,
	0x25, 0x50, 0x26, 0xc1, 0x89, 0x23, 0xe3, 0x00, 0xd2, 0x86, 0x34, 0x79, 0xf4, 0x00, 0x97, 0xc8,
	0x89, 0x1f, 0x26, 0x6a, 0x53, 0x5b, 0xb6, 0x23, 0xd4, 0x7f, 0x99, 0xbf, 0x62, 0xaa, 0x9d, 0x4c,
	0x6b, 0x6e, 0x79, 0x9f, 0xef, 0x47, 0xf1, 0xfb, 0x01, 0x8b, 0x2d, 0xd7, 0x4e, 0xe9, 0xb2, 0x45,
	0x6b, 0xb9, 0xc4, 0x42, 0x1b, 0xe5, 0x14, 0x49, 0x74, 0x5d, 0x29, 0xb5, 0x59, 0xbe, 0xd2, 0x46,
	0xd5, 0x68, 0xad, 0x32, 0xc7, 0xc2, 0x72, 0xd1, 0x62, 0xab, 0xcc, 0x7e, 0x44, 0x5f, 0x58, 0xa7,
	0x0c, 0x97, 0x38, 0x96, 0x6d, 0x6d, 0x10, 0x77, 0x23, 0xfa, 0x72, 0x83, 0xfb, 0x4a, 0x71, 0x23,
	0x46, 0xfc, 0x42, 0x2a, 0x25, 0xb7, 0x78, 0xe9, 0xab, 0xaa, 0xfb, 0x73, 0xe9, 0x9a, 0x16, 0xad,
	0xe3, 0xad, 0x0e, 0xc2, 0xdb, 0xff, 0x31, 0x24, 0x37, 0xbe, 0x6b, 0x32, 0x87, 0x49, 0x23, 0x68,
	0x94, 0x47, 0xab, 0x8c, 0x4d, 0x1a, 0x41, 0x16, 0x70, 0x52, 0x19, 0xbe, 0x13, 0x74, 0xe2, 0x51,
	0x28, 0x08, 0x81, 0xe9, 0x8e, 0xb7, 0x48, 0x63, 0x0f, 0xfd, 0x37, 0x39, 0x87, 0xb8, 0xd6, 0x1d,
	0x9d, 0xe6, 0xd1, 0x6a, 0xf6, 0x71, 0x56, 0x84, 0x79, 0x8b, 0xaf, 0xb7, 0x6b, 0x76, 0xe0, 0x24,
	0x87, 0xd8, 0xf0, 0x96, 0x9e, 0xf8, 0x78, 0x3e, 0xc4, 0x3f, 0xfc, 0xd0, 0xec, 0x10, 0x91, 0x0b,
	0x98, 0x4a, 0xdd, 0x59, 0x9a, 0xe4, 0xf1, 0xe3, 0x3f, 0x7c, 0xbb, 0x5d, 0x33, 0x1f, 0x90, 0xf7,
	0x90, 0xf6, 0xeb, 0xb0, 0xf4, 0xa9, 0x97, 0x9e, 0x0d, 0xd2, 0x5d, 0xe0, 0xec, 0x41, 0x20, 0xef,
	0x20, 0x09, 0x4b, 0xa2, 0xe9, 0xf1, 0x93, 0x77, 0x9e, 0xb2, 0x3e, 0x25, 0x1f, 0x20, 0x1d, 0xd6,
	0x46, 0x33, 0x6f, 0x3e, 0x1f, 0xcc, 0xeb, 0x9e, 0xb3, 0x07, 0x83, 0x9c, 0x43, 0xf6, 0x0f, 0x1b,
	0xf9, 0xd7, 0x95, 0x1b, 0x49, 0x21, 0x8f, 0x56, 0xd1, 0xf7, 0x27, 0x2c, 0x0d, 0xe8, 0x5a, 0x3e,
	0x8a, 0xb7, 0x15, 0x9d, 0x1d, 0xc7, 0x37, 0x15, 0x79, 0x0d, 0x99, 0x36, 0x4d, 0x8d, 0x65, 0x67,
	0x05, 0x3d, 0x3d, 0xc4, 0x2c, 0xf5, 0x60, 0x6d, 0x05, 0x79, 0x03, 0xa7, 0x06, 0xb7, 0xc8, 0x2d,
	0x96, 0x7b, 0xe4, 0x86, 0x9e, 0xe5, 0xd1, 0xea, 0x8c, 0xcd, 0x7a, 0xf6, 0x0b, 0xb9, 0x21, 0x9f,
	0x01, 0x3a, 0x2d, 0xb8, 0x43, 0x51, 0x72, 0x47, 0xe7, 0xbe, 0xdb, 0x65, 0x11, 0xae, 0x5b, 0x0c,
	0xd7, 0x2d, 0x7e, 0x0e, 0xd7, 0x65, 0x59, 0x6f, 0x7f, 0x71, 0x57, 0x29, 0x24, 0xa1, 0x8d, 0xab,
	0xe9, 0xef, 0x89, 0xae, 0xaa, 0xc4, 0xeb, 0x9f, 0xee, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x1e,
	0x79, 0x27, 0xae, 0x02, 0x00, 0x00,
}