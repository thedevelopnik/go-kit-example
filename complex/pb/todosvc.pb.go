// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todosvc.proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetTodosRequest struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodosRequest) Reset()         { *m = GetTodosRequest{} }
func (m *GetTodosRequest) String() string { return proto.CompactTextString(m) }
func (*GetTodosRequest) ProtoMessage()    {}
func (*GetTodosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4fe0a8e5ec7537d, []int{0}
}

func (m *GetTodosRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodosRequest.Unmarshal(m, b)
}
func (m *GetTodosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodosRequest.Marshal(b, m, deterministic)
}
func (m *GetTodosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodosRequest.Merge(m, src)
}
func (m *GetTodosRequest) XXX_Size() int {
	return xxx_messageInfo_GetTodosRequest.Size(m)
}
func (m *GetTodosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodosRequest proto.InternalMessageInfo

func (m *GetTodosRequest) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type GetTodosResponse struct {
	Todos                []*Todo  `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodosResponse) Reset()         { *m = GetTodosResponse{} }
func (m *GetTodosResponse) String() string { return proto.CompactTextString(m) }
func (*GetTodosResponse) ProtoMessage()    {}
func (*GetTodosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4fe0a8e5ec7537d, []int{1}
}

func (m *GetTodosResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodosResponse.Unmarshal(m, b)
}
func (m *GetTodosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodosResponse.Marshal(b, m, deterministic)
}
func (m *GetTodosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodosResponse.Merge(m, src)
}
func (m *GetTodosResponse) XXX_Size() int {
	return xxx_messageInfo_GetTodosResponse.Size(m)
}
func (m *GetTodosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodosResponse proto.InternalMessageInfo

func (m *GetTodosResponse) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

type Todo struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4fe0a8e5ec7537d, []int{2}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Todo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*GetTodosRequest)(nil), "pb.GetTodosRequest")
	proto.RegisterType((*GetTodosResponse)(nil), "pb.GetTodosResponse")
	proto.RegisterType((*Todo)(nil), "pb.Todo")
}

func init() { proto.RegisterFile("todosvc.proto", fileDescriptor_d4fe0a8e5ec7537d) }

var fileDescriptor_d4fe0a8e5ec7537d = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0xc9, 0x4f, 0xc9,
	0x2f, 0x2e, 0x4b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe4,
	0xe2, 0x77, 0x4f, 0x2d, 0x09, 0x01, 0x89, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89,
	0x71, 0xb1, 0x95, 0x16, 0xa7, 0x16, 0x79, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x04, 0x41,
	0x79, 0x4a, 0x46, 0x5c, 0x02, 0x08, 0xa5, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x72, 0x5c,
	0xac, 0x60, 0x33, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x38, 0xf4, 0x0a, 0x92, 0xf4, 0x40,
	0x2a, 0x82, 0x20, 0xc2, 0x4a, 0x5a, 0x5c, 0x2c, 0x20, 0xae, 0x10, 0x1f, 0x17, 0x53, 0x26, 0xcc,
	0x3c, 0xa6, 0xcc, 0x14, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x30, 0x3b, 0x89, 0x0d, 0xec, 0x2a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5e, 0xdd, 0xed, 0x34, 0xa6, 0x00, 0x00, 0x00,
}
