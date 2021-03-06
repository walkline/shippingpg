// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/grpc/PortDomainService.proto

package pb

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

type PortID struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PortID) Reset()         { *m = PortID{} }
func (m *PortID) String() string { return proto.CompactTextString(m) }
func (*PortID) ProtoMessage()    {}
func (*PortID) Descriptor() ([]byte, []int) {
	return fileDescriptor_6213773b07236b36, []int{0}
}

func (m *PortID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortID.Unmarshal(m, b)
}
func (m *PortID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortID.Marshal(b, m, deterministic)
}
func (m *PortID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortID.Merge(m, src)
}
func (m *PortID) XXX_Size() int {
	return xxx_messageInfo_PortID.Size(m)
}
func (m *PortID) XXX_DiscardUnknown() {
	xxx_messageInfo_PortID.DiscardUnknown(m)
}

var xxx_messageInfo_PortID proto.InternalMessageInfo

func (m *PortID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Geo2DPoint struct {
	Lati                 string   `protobuf:"bytes,1,opt,name=lati,proto3" json:"lati,omitempty"`
	Longi                string   `protobuf:"bytes,2,opt,name=longi,proto3" json:"longi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Geo2DPoint) Reset()         { *m = Geo2DPoint{} }
func (m *Geo2DPoint) String() string { return proto.CompactTextString(m) }
func (*Geo2DPoint) ProtoMessage()    {}
func (*Geo2DPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_6213773b07236b36, []int{1}
}

func (m *Geo2DPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Geo2DPoint.Unmarshal(m, b)
}
func (m *Geo2DPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Geo2DPoint.Marshal(b, m, deterministic)
}
func (m *Geo2DPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Geo2DPoint.Merge(m, src)
}
func (m *Geo2DPoint) XXX_Size() int {
	return xxx_messageInfo_Geo2DPoint.Size(m)
}
func (m *Geo2DPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Geo2DPoint.DiscardUnknown(m)
}

var xxx_messageInfo_Geo2DPoint proto.InternalMessageInfo

func (m *Geo2DPoint) GetLati() string {
	if m != nil {
		return m.Lati
	}
	return ""
}

func (m *Geo2DPoint) GetLongi() string {
	if m != nil {
		return m.Longi
	}
	return ""
}

type Port struct {
	ID                   string      `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Country              string      `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	City                 string      `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Province             string      `protobuf:"bytes,5,opt,name=province,proto3" json:"province,omitempty"`
	Timezone             string      `protobuf:"bytes,6,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Code                 string      `protobuf:"bytes,7,opt,name=code,proto3" json:"code,omitempty"`
	Coordinates          *Geo2DPoint `protobuf:"bytes,8,opt,name=coordinates,proto3" json:"coordinates,omitempty"`
	Regions              []string    `protobuf:"bytes,9,rep,name=regions,proto3" json:"regions,omitempty"`
	Alias                []string    `protobuf:"bytes,10,rep,name=alias,proto3" json:"alias,omitempty"`
	Unlocs               []string    `protobuf:"bytes,11,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_6213773b07236b36, []int{2}
}

func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (m *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(m, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Port) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Port) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Port) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *Port) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *Port) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Port) GetCoordinates() *Geo2DPoint {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

func (m *Port) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *Port) GetAlias() []string {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *Port) GetUnlocs() []string {
	if m != nil {
		return m.Unlocs
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_6213773b07236b36, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PortID)(nil), "PortID")
	proto.RegisterType((*Geo2DPoint)(nil), "Geo2DPoint")
	proto.RegisterType((*Port)(nil), "Port")
	proto.RegisterType((*Empty)(nil), "Empty")
}

func init() { proto.RegisterFile("proto/grpc/PortDomainService.proto", fileDescriptor_6213773b07236b36) }

var fileDescriptor_6213773b07236b36 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x41, 0x6b, 0xe3, 0x30,
	0x10, 0x85, 0x13, 0x27, 0xb6, 0x93, 0x31, 0x2c, 0xac, 0x58, 0x96, 0x21, 0xf4, 0x10, 0x74, 0xca,
	0xa5, 0x0e, 0xa4, 0xd0, 0x4b, 0x6f, 0xc1, 0x6d, 0x31, 0xbd, 0x84, 0xe4, 0xd6, 0x9b, 0x23, 0x8b,
	0x20, 0xb0, 0x35, 0x46, 0x56, 0x02, 0xee, 0x5f, 0xe9, 0x9f, 0x2d, 0x92, 0x9d, 0xb6, 0xb4, 0xb7,
	0x79, 0xef, 0x7b, 0x3c, 0x46, 0x1a, 0xe0, 0x8d, 0x21, 0x4b, 0xeb, 0x93, 0x69, 0xc4, 0x7a, 0x47,
	0xc6, 0x66, 0x54, 0x17, 0x4a, 0x1f, 0xa4, 0xb9, 0x28, 0x21, 0x53, 0x0f, 0x39, 0x42, 0xe4, 0x50,
	0x9e, 0xb1, 0x3f, 0x10, 0xe4, 0x19, 0x8e, 0x97, 0xe3, 0xd5, 0x7c, 0x1f, 0xe4, 0x19, 0xbf, 0x07,
	0x78, 0x96, 0xb4, 0xc9, 0x76, 0xa4, 0xb4, 0x65, 0x0c, 0xa6, 0x55, 0x61, 0xd5, 0xc0, 0xfd, 0xcc,
	0xfe, 0x41, 0x58, 0x91, 0x3e, 0x29, 0x0c, 0xbc, 0xd9, 0x0b, 0xfe, 0x1e, 0xc0, 0xd4, 0x55, 0xfe,
	0x2c, 0x74, 0x15, 0xba, 0xa8, 0xe5, 0x90, 0xf6, 0x33, 0x43, 0x88, 0x05, 0x9d, 0xb5, 0x35, 0x1d,
	0x4e, 0xbc, 0x7d, 0x95, 0x2e, 0x2d, 0x94, 0xed, 0x70, 0xda, 0xa7, 0xdd, 0xcc, 0x16, 0x30, 0x6b,
	0x0c, 0x5d, 0x94, 0x16, 0x12, 0x43, 0xef, 0x7f, 0x6a, 0xc7, 0xac, 0xaa, 0xe5, 0x1b, 0x69, 0x89,
	0x51, 0xcf, 0xae, 0xda, 0x77, 0x51, 0x29, 0x31, 0x1e, 0xba, 0xa8, 0x94, 0xec, 0x16, 0x12, 0x41,
	0x64, 0x4a, 0xa5, 0x0b, 0x2b, 0x5b, 0x9c, 0x2d, 0xc7, 0xab, 0x64, 0x93, 0xa4, 0x5f, 0x4f, 0xde,
	0x7f, 0xe7, 0x6e, 0x51, 0x23, 0x4f, 0x8a, 0x74, 0x8b, 0xf3, 0xe5, 0xc4, 0x2d, 0x3a, 0x48, 0xf7,
	0x0b, 0x45, 0xa5, 0x8a, 0x16, 0xc1, 0xfb, 0xbd, 0x60, 0xff, 0x21, 0x3a, 0xeb, 0x8a, 0x44, 0x8b,
	0x89, 0xb7, 0x07, 0xc5, 0x63, 0x08, 0x1f, 0xeb, 0xc6, 0x76, 0x9b, 0x17, 0xf8, 0xfb, 0xeb, 0x26,
	0x0c, 0x21, 0x3c, 0x58, 0x32, 0x92, 0x85, 0xa9, 0x83, 0x8b, 0x28, 0xf5, 0x61, 0x3e, 0x62, 0x37,
	0x30, 0x7b, 0x52, 0xba, 0xdc, 0x76, 0x79, 0xc6, 0xe2, 0xb4, 0x3f, 0xd9, 0xa2, 0x4f, 0xf1, 0xd1,
	0x36, 0x7e, 0x0d, 0x9b, 0xe3, 0x43, 0x73, 0x3c, 0x46, 0xfe, 0xaa, 0x77, 0x1f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x5b, 0x57, 0xf4, 0xb1, 0xfb, 0x01, 0x00, 0x00,
}
