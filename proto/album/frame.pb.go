// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/album/frame.proto

package album

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

type FrameInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	Asset                string   `protobuf:"bytes,9,opt,name=asset,proto3" json:"asset,omitempty"`
	Width                uint32   `protobuf:"varint,10,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32   `protobuf:"varint,11,opt,name=height,proto3" json:"height,omitempty"`
	Owner                string   `protobuf:"bytes,12,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrameInfo) Reset()         { *m = FrameInfo{} }
func (m *FrameInfo) String() string { return proto.CompactTextString(m) }
func (*FrameInfo) ProtoMessage()    {}
func (*FrameInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1080b1b1ac4037e, []int{0}
}

func (m *FrameInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrameInfo.Unmarshal(m, b)
}
func (m *FrameInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrameInfo.Marshal(b, m, deterministic)
}
func (m *FrameInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrameInfo.Merge(m, src)
}
func (m *FrameInfo) XXX_Size() int {
	return xxx_messageInfo_FrameInfo.Size(m)
}
func (m *FrameInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FrameInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FrameInfo proto.InternalMessageInfo

func (m *FrameInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *FrameInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FrameInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *FrameInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *FrameInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *FrameInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *FrameInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FrameInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *FrameInfo) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *FrameInfo) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *FrameInfo) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *FrameInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type ReqFrameAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Asset                string   `protobuf:"bytes,4,opt,name=asset,proto3" json:"asset,omitempty"`
	Width                uint32   `protobuf:"varint,5,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32   `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty"`
	Owner                string   `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqFrameAdd) Reset()         { *m = ReqFrameAdd{} }
func (m *ReqFrameAdd) String() string { return proto.CompactTextString(m) }
func (*ReqFrameAdd) ProtoMessage()    {}
func (*ReqFrameAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1080b1b1ac4037e, []int{1}
}

func (m *ReqFrameAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqFrameAdd.Unmarshal(m, b)
}
func (m *ReqFrameAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqFrameAdd.Marshal(b, m, deterministic)
}
func (m *ReqFrameAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqFrameAdd.Merge(m, src)
}
func (m *ReqFrameAdd) XXX_Size() int {
	return xxx_messageInfo_ReqFrameAdd.Size(m)
}
func (m *ReqFrameAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqFrameAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqFrameAdd proto.InternalMessageInfo

func (m *ReqFrameAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqFrameAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqFrameAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqFrameAdd) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *ReqFrameAdd) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ReqFrameAdd) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ReqFrameAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type ReqFrameUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqFrameUpdate) Reset()         { *m = ReqFrameUpdate{} }
func (m *ReqFrameUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqFrameUpdate) ProtoMessage()    {}
func (*ReqFrameUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1080b1b1ac4037e, []int{2}
}

func (m *ReqFrameUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqFrameUpdate.Unmarshal(m, b)
}
func (m *ReqFrameUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqFrameUpdate.Marshal(b, m, deterministic)
}
func (m *ReqFrameUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqFrameUpdate.Merge(m, src)
}
func (m *ReqFrameUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqFrameUpdate.Size(m)
}
func (m *ReqFrameUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqFrameUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqFrameUpdate proto.InternalMessageInfo

func (m *ReqFrameUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqFrameUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqFrameUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqFrameUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyFrameInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *FrameInfo   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyFrameInfo) Reset()         { *m = ReplyFrameInfo{} }
func (m *ReplyFrameInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyFrameInfo) ProtoMessage()    {}
func (*ReplyFrameInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1080b1b1ac4037e, []int{3}
}

func (m *ReplyFrameInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyFrameInfo.Unmarshal(m, b)
}
func (m *ReplyFrameInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyFrameInfo.Marshal(b, m, deterministic)
}
func (m *ReplyFrameInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyFrameInfo.Merge(m, src)
}
func (m *ReplyFrameInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyFrameInfo.Size(m)
}
func (m *ReplyFrameInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyFrameInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyFrameInfo proto.InternalMessageInfo

func (m *ReplyFrameInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyFrameInfo) GetInfo() *FrameInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyFrameList struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*FrameInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyFrameList) Reset()         { *m = ReplyFrameList{} }
func (m *ReplyFrameList) String() string { return proto.CompactTextString(m) }
func (*ReplyFrameList) ProtoMessage()    {}
func (*ReplyFrameList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1080b1b1ac4037e, []int{4}
}

func (m *ReplyFrameList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyFrameList.Unmarshal(m, b)
}
func (m *ReplyFrameList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyFrameList.Marshal(b, m, deterministic)
}
func (m *ReplyFrameList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyFrameList.Merge(m, src)
}
func (m *ReplyFrameList) XXX_Size() int {
	return xxx_messageInfo_ReplyFrameList.Size(m)
}
func (m *ReplyFrameList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyFrameList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyFrameList proto.InternalMessageInfo

func (m *ReplyFrameList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyFrameList) GetList() []*FrameInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*FrameInfo)(nil), "omo.msp.album.FrameInfo")
	proto.RegisterType((*ReqFrameAdd)(nil), "omo.msp.album.ReqFrameAdd")
	proto.RegisterType((*ReqFrameUpdate)(nil), "omo.msp.album.ReqFrameUpdate")
	proto.RegisterType((*ReplyFrameInfo)(nil), "omo.msp.album.ReplyFrameInfo")
	proto.RegisterType((*ReplyFrameList)(nil), "omo.msp.album.ReplyFrameList")
}

func init() {
	proto.RegisterFile("proto/album/frame.proto", fileDescriptor_f1080b1b1ac4037e)
}

var fileDescriptor_f1080b1b1ac4037e = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x8d, 0x7f, 0xe2, 0x34, 0x93, 0x9f, 0xef, 0xd3, 0x0a, 0xc1, 0x2a, 0xa2, 0x52, 0xe4, 0xab,
	0x5c, 0xa0, 0x54, 0x0a, 0x4f, 0xd0, 0xa0, 0x34, 0x02, 0x81, 0x40, 0x8e, 0xb8, 0xe1, 0xce, 0x8d,
	0x27, 0x64, 0x21, 0xf6, 0x86, 0xdd, 0x4d, 0xab, 0x3e, 0x0c, 0xcf, 0xc0, 0x1b, 0xf1, 0x2c, 0x68,
	0xc7, 0x8e, 0x49, 0x8a, 0x53, 0x4b, 0xb9, 0xdb, 0x33, 0x73, 0xe6, 0xcc, 0x19, 0xcf, 0xc8, 0xf0,
	0x62, 0xab, 0xa4, 0x91, 0x57, 0xf1, 0xe6, 0x76, 0x97, 0x5e, 0xad, 0x54, 0x9c, 0xe2, 0x98, 0x22,
	0xac, 0x27, 0x53, 0x39, 0x4e, 0xf5, 0x76, 0x4c, 0xa9, 0x01, 0x3f, 0xe4, 0x2d, 0x65, 0x9a, 0xca,
	0x2c, 0x27, 0x86, 0x3f, 0x5d, 0x68, 0xdf, 0xd8, 0xc2, 0xb7, 0xd9, 0x4a, 0xb2, 0xff, 0xc1, 0xdb,
	0x89, 0x84, 0x3b, 0x43, 0x67, 0xd4, 0x8e, 0xec, 0x93, 0xf5, 0xc1, 0x15, 0x09, 0x77, 0x87, 0xce,
	0xc8, 0x8f, 0x5c, 0x91, 0x30, 0x0e, 0xad, 0xa5, 0xc2, 0xd8, 0x60, 0xc2, 0xbd, 0xa1, 0x33, 0xf2,
	0xa2, 0x3d, 0xb4, 0x99, 0xdd, 0x36, 0xa1, 0x8c, 0x9f, 0x67, 0x0a, 0x58, 0xd6, 0x48, 0xc5, 0x9b,
	0xa4, 0xbc, 0x87, 0x6c, 0x00, 0x17, 0x72, 0x8b, 0x8a, 0x52, 0x01, 0xa5, 0x4a, 0xcc, 0x18, 0xf8,
	0x59, 0x9c, 0x22, 0x6f, 0x51, 0x9c, 0xde, 0xec, 0x39, 0x04, 0x0a, 0xd3, 0x58, 0x7d, 0xe7, 0x17,
	0x14, 0x2d, 0x10, 0x7b, 0x06, 0xcd, 0x58, 0x6b, 0x34, 0xbc, 0x4d, 0xe1, 0x1c, 0xd8, 0xe8, 0xbd,
	0x48, 0xcc, 0x9a, 0xc3, 0xd0, 0x19, 0xf5, 0xa2, 0x1c, 0x58, 0x8d, 0x35, 0x8a, 0xaf, 0x6b, 0xc3,
	0x3b, 0x14, 0x2e, 0x90, 0x65, 0xcb, 0xfb, 0x0c, 0x15, 0xef, 0xe6, 0x1a, 0x04, 0xc2, 0x5f, 0x0e,
	0x74, 0x22, 0xfc, 0x41, 0x9f, 0xe8, 0x3a, 0x49, 0x4a, 0x57, 0x4e, 0xa5, 0x2b, 0xf7, 0xc8, 0xd5,
	0xe1, 0x74, 0xde, 0xa3, 0xe9, 0x4a, 0xc7, 0x7e, 0xa5, 0xe3, 0x66, 0xb5, 0xe3, 0xa0, 0xda, 0x71,
	0xeb, 0xd0, 0xf1, 0x37, 0xe8, 0xef, 0x0d, 0x7f, 0xa6, 0x05, 0x54, 0x6c, 0x75, 0x3f, 0x85, 0x5b,
	0x39, 0x85, 0x77, 0x72, 0x0a, 0xff, 0x78, 0x8a, 0x50, 0xd9, 0x5e, 0xdb, 0xcd, 0xc3, 0xdf, 0x0b,
	0x9a, 0x40, 0xa0, 0x4d, 0x6c, 0x76, 0x9a, 0xda, 0x75, 0x26, 0x83, 0xf1, 0xd1, 0x25, 0x8e, 0x89,
	0xbe, 0x20, 0x46, 0x54, 0x30, 0xd9, 0x2b, 0xf0, 0x45, 0xb6, 0x92, 0xe4, 0xa6, 0x33, 0xe1, 0x8f,
	0x2a, 0x4a, 0xed, 0x88, 0x58, 0xc7, 0x3d, 0xdf, 0x0b, 0x6d, 0xce, 0xed, 0xb9, 0x11, 0xda, 0x70,
	0x77, 0xe8, 0x3d, 0xdd, 0xd3, 0xb2, 0x26, 0xbf, 0x7d, 0xe8, 0x52, 0x6c, 0x81, 0xea, 0x4e, 0x2c,
	0x91, 0xcd, 0x20, 0xb8, 0x4e, 0x92, 0x8f, 0x19, 0xb2, 0x7f, 0x9b, 0x95, 0xc7, 0x32, 0xb8, 0xac,
	0x32, 0x52, 0x6a, 0x87, 0x0d, 0x36, 0x07, 0xc8, 0x77, 0x34, 0x8d, 0x35, 0xb2, 0xcb, 0x13, 0x52,
	0x39, 0x65, 0xc0, 0xab, 0xd4, 0x0a, 0xa1, 0x37, 0xd0, 0x8e, 0x30, 0x95, 0x77, 0x78, 0xc2, 0xd2,
	0x0e, 0xb5, 0xb1, 0xd4, 0x27, 0x45, 0x66, 0x10, 0xcc, 0xd1, 0xd4, 0x29, 0xd4, 0x0e, 0x35, 0x83,
	0x60, 0x81, 0xb1, 0x5a, 0xae, 0xcf, 0x94, 0xb1, 0x3b, 0x0d, 0x1b, 0xec, 0x13, 0xfc, 0x37, 0x47,
	0x63, 0xc1, 0xf4, 0xe1, 0x46, 0x6c, 0x0c, 0x2a, 0xf6, 0xb2, 0x5a, 0x2f, 0xcf, 0xd6, 0x2b, 0x7e,
	0x80, 0xee, 0x1c, 0x8d, 0x3d, 0x04, 0xa1, 0x8d, 0x58, 0x9e, 0x23, 0x57, 0x16, 0x87, 0x0d, 0xf6,
	0x0e, 0xfa, 0xc5, 0xf2, 0x6a, 0xfc, 0xd5, 0xef, 0x6f, 0xda, 0xfb, 0xd2, 0x39, 0xf8, 0x45, 0xdf,
	0x06, 0x04, 0x5e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x8e, 0xde, 0x1e, 0xe0, 0x05, 0x00,
	0x00,
}
