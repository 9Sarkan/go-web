// Code generated by protoc-gen-go. DO NOT EDIT.
// source: persion.proto

package protofiles

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

type Person_PhoneType int32

const (
	Person_Mobile Person_PhoneType = 0
	Person_Work   Person_PhoneType = 1
	Person_Home   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "Mobile",
	1: "Work",
	2: "Home",
}

var Person_PhoneType_value = map[string]int32{
	"Mobile": 0,
	"Work":   1,
	"Home":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a7f5642358fd2963, []int{0, 0}
}

type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=Phones,proto3" json:"Phones,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7f5642358fd2963, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type Person_PhoneNumber struct {
	Phone                string           `protobuf:"bytes,1,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=Type,proto3,enum=protofiles.Person_PhoneType" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7f5642358fd2963, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_Mobile
}

type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7f5642358fd2963, []int{1}
}

func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterEnum("protofiles.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Person)(nil), "protofiles.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "protofiles.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "protofiles.AddressBook")
}

func init() { proto.RegisterFile("persion.proto", fileDescriptor_a7f5642358fd2963) }

var fileDescriptor_a7f5642358fd2963 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x4d, 0xb7, 0x1b, 0xdc, 0x29, 0x2e, 0x65, 0xf0, 0x50, 0x44, 0xa4, 0xf4, 0x54, 0x14,
	0x8a, 0xac, 0x20, 0x78, 0x54, 0x10, 0xdc, 0x83, 0xcb, 0x12, 0x14, 0xcf, 0x5b, 0x3a, 0x62, 0x69,
	0xd3, 0x09, 0x89, 0x1e, 0xfc, 0xf3, 0x22, 0x49, 0x8b, 0x0a, 0xb2, 0xa7, 0xcc, 0x9b, 0x7c, 0xbc,
	0xf7, 0x06, 0x8e, 0x0c, 0x59, 0xd7, 0xf2, 0x50, 0x19, 0xcb, 0xef, 0x8c, 0x10, 0x9e, 0xd7, 0xb6,
	0x27, 0x57, 0x7c, 0x09, 0x90, 0x5b, 0xb2, 0x8e, 0x07, 0x44, 0x88, 0x37, 0x3b, 0x4d, 0x99, 0xc8,
	0x45, 0xb9, 0x50, 0x61, 0xc6, 0x25, 0x44, 0xeb, 0x26, 0x8b, 0x72, 0x51, 0xce, 0x55, 0xb4, 0x6e,
	0xf0, 0x18, 0xe6, 0xf7, 0x7a, 0xd7, 0xf6, 0xd9, 0x2c, 0x40, 0xa3, 0xc0, 0x6b, 0x90, 0xdb, 0x37,
	0x1e, 0xc8, 0x65, 0x71, 0x3e, 0x2b, 0x93, 0xd5, 0x59, 0xf5, 0x9b, 0x50, 0x8d, 0xee, 0x55, 0x00,
	0x36, 0x1f, 0xba, 0x26, 0xab, 0x26, 0xfa, 0xe4, 0x19, 0x92, 0x3f, 0x6b, 0x6f, 0x1e, 0xe4, 0xd4,
	0x60, 0x14, 0x78, 0x09, 0xf1, 0xd3, 0xa7, 0xa1, 0x50, 0x62, 0xb9, 0x3a, 0xdd, 0x67, 0xed, 0x19,
	0x15, 0xc8, 0xe2, 0x02, 0x16, 0x3f, 0x2b, 0x04, 0x90, 0x8f, 0x5c, 0xb7, 0x3d, 0xa5, 0x07, 0x78,
	0x08, 0xf1, 0x0b, 0xdb, 0x2e, 0x15, 0x7e, 0x7a, 0x60, 0x4d, 0x69, 0x54, 0xdc, 0x40, 0x72, 0xdb,
	0x34, 0x96, 0x9c, 0xbb, 0x63, 0xee, 0xf0, 0x1c, 0xa4, 0x21, 0x36, 0xbd, 0x2f, 0xe1, 0x4f, 0xc1,
	0xff, 0x79, 0x6a, 0x22, 0x6a, 0x19, 0xbe, 0xae, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x67,
	0xcd, 0x0b, 0x5f, 0x01, 0x00, 0x00,
}
