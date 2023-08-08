// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/album/sheet.proto

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

type SheetInfo struct {
	Uid                  string     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64     `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64      `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64      `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string     `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string     `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string     `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string     `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	Owner                string     `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
	Area                 string     `protobuf:"bytes,10,opt,name=area,proto3" json:"area,omitempty"`
	Time                 int64      `protobuf:"varint,11,opt,name=time,proto3" json:"time,omitempty"`
	Pages                []*PairInt `protobuf:"bytes,21,rep,name=pages,proto3" json:"pages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SheetInfo) Reset()         { *m = SheetInfo{} }
func (m *SheetInfo) String() string { return proto.CompactTextString(m) }
func (*SheetInfo) ProtoMessage()    {}
func (*SheetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_16583bcb620c2e7c, []int{0}
}

func (m *SheetInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SheetInfo.Unmarshal(m, b)
}
func (m *SheetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SheetInfo.Marshal(b, m, deterministic)
}
func (m *SheetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SheetInfo.Merge(m, src)
}
func (m *SheetInfo) XXX_Size() int {
	return xxx_messageInfo_SheetInfo.Size(m)
}
func (m *SheetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SheetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SheetInfo proto.InternalMessageInfo

func (m *SheetInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *SheetInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SheetInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *SheetInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *SheetInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SheetInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *SheetInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SheetInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *SheetInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *SheetInfo) GetArea() string {
	if m != nil {
		return m.Area
	}
	return ""
}

func (m *SheetInfo) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *SheetInfo) GetPages() []*PairInt {
	if m != nil {
		return m.Pages
	}
	return nil
}

type ReqSheetAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Owner                string   `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSheetAdd) Reset()         { *m = ReqSheetAdd{} }
func (m *ReqSheetAdd) String() string { return proto.CompactTextString(m) }
func (*ReqSheetAdd) ProtoMessage()    {}
func (*ReqSheetAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_16583bcb620c2e7c, []int{1}
}

func (m *ReqSheetAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSheetAdd.Unmarshal(m, b)
}
func (m *ReqSheetAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSheetAdd.Marshal(b, m, deterministic)
}
func (m *ReqSheetAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSheetAdd.Merge(m, src)
}
func (m *ReqSheetAdd) XXX_Size() int {
	return xxx_messageInfo_ReqSheetAdd.Size(m)
}
func (m *ReqSheetAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSheetAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSheetAdd proto.InternalMessageInfo

func (m *ReqSheetAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqSheetAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqSheetAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqSheetAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type ReqSheetUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSheetUpdate) Reset()         { *m = ReqSheetUpdate{} }
func (m *ReqSheetUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqSheetUpdate) ProtoMessage()    {}
func (*ReqSheetUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_16583bcb620c2e7c, []int{2}
}

func (m *ReqSheetUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSheetUpdate.Unmarshal(m, b)
}
func (m *ReqSheetUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSheetUpdate.Marshal(b, m, deterministic)
}
func (m *ReqSheetUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSheetUpdate.Merge(m, src)
}
func (m *ReqSheetUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqSheetUpdate.Size(m)
}
func (m *ReqSheetUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSheetUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSheetUpdate proto.InternalMessageInfo

func (m *ReqSheetUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqSheetUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqSheetUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqSheetUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReqSheetPages struct {
	Uid                  string     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string     `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	List                 []*PairInt `protobuf:"bytes,11,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReqSheetPages) Reset()         { *m = ReqSheetPages{} }
func (m *ReqSheetPages) String() string { return proto.CompactTextString(m) }
func (*ReqSheetPages) ProtoMessage()    {}
func (*ReqSheetPages) Descriptor() ([]byte, []int) {
	return fileDescriptor_16583bcb620c2e7c, []int{3}
}

func (m *ReqSheetPages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSheetPages.Unmarshal(m, b)
}
func (m *ReqSheetPages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSheetPages.Marshal(b, m, deterministic)
}
func (m *ReqSheetPages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSheetPages.Merge(m, src)
}
func (m *ReqSheetPages) XXX_Size() int {
	return xxx_messageInfo_ReqSheetPages.Size(m)
}
func (m *ReqSheetPages) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSheetPages.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSheetPages proto.InternalMessageInfo

func (m *ReqSheetPages) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqSheetPages) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqSheetPages) GetList() []*PairInt {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplySheetPages struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*PairList  `protobuf:"bytes,11,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplySheetPages) Reset()         { *m = ReplySheetPages{} }
func (m *ReplySheetPages) String() string { return proto.CompactTextString(m) }
func (*ReplySheetPages) ProtoMessage()    {}
func (*ReplySheetPages) Descriptor() ([]byte, []int) {
	return fileDescriptor_16583bcb620c2e7c, []int{4}
}

func (m *ReplySheetPages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplySheetPages.Unmarshal(m, b)
}
func (m *ReplySheetPages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplySheetPages.Marshal(b, m, deterministic)
}
func (m *ReplySheetPages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplySheetPages.Merge(m, src)
}
func (m *ReplySheetPages) XXX_Size() int {
	return xxx_messageInfo_ReplySheetPages.Size(m)
}
func (m *ReplySheetPages) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplySheetPages.DiscardUnknown(m)
}

var xxx_messageInfo_ReplySheetPages proto.InternalMessageInfo

func (m *ReplySheetPages) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplySheetPages) GetList() []*PairList {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplySheetInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *SheetInfo   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplySheetInfo) Reset()         { *m = ReplySheetInfo{} }
func (m *ReplySheetInfo) String() string { return proto.CompactTextString(m) }
func (*ReplySheetInfo) ProtoMessage()    {}
func (*ReplySheetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_16583bcb620c2e7c, []int{5}
}

func (m *ReplySheetInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplySheetInfo.Unmarshal(m, b)
}
func (m *ReplySheetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplySheetInfo.Marshal(b, m, deterministic)
}
func (m *ReplySheetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplySheetInfo.Merge(m, src)
}
func (m *ReplySheetInfo) XXX_Size() int {
	return xxx_messageInfo_ReplySheetInfo.Size(m)
}
func (m *ReplySheetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplySheetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplySheetInfo proto.InternalMessageInfo

func (m *ReplySheetInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplySheetInfo) GetInfo() *SheetInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplySheetList struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*SheetInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplySheetList) Reset()         { *m = ReplySheetList{} }
func (m *ReplySheetList) String() string { return proto.CompactTextString(m) }
func (*ReplySheetList) ProtoMessage()    {}
func (*ReplySheetList) Descriptor() ([]byte, []int) {
	return fileDescriptor_16583bcb620c2e7c, []int{6}
}

func (m *ReplySheetList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplySheetList.Unmarshal(m, b)
}
func (m *ReplySheetList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplySheetList.Marshal(b, m, deterministic)
}
func (m *ReplySheetList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplySheetList.Merge(m, src)
}
func (m *ReplySheetList) XXX_Size() int {
	return xxx_messageInfo_ReplySheetList.Size(m)
}
func (m *ReplySheetList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplySheetList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplySheetList proto.InternalMessageInfo

func (m *ReplySheetList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplySheetList) GetList() []*SheetInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*SheetInfo)(nil), "omo.msp.album.SheetInfo")
	proto.RegisterType((*ReqSheetAdd)(nil), "omo.msp.album.ReqSheetAdd")
	proto.RegisterType((*ReqSheetUpdate)(nil), "omo.msp.album.ReqSheetUpdate")
	proto.RegisterType((*ReqSheetPages)(nil), "omo.msp.album.ReqSheetPages")
	proto.RegisterType((*ReplySheetPages)(nil), "omo.msp.album.ReplySheetPages")
	proto.RegisterType((*ReplySheetInfo)(nil), "omo.msp.album.ReplySheetInfo")
	proto.RegisterType((*ReplySheetList)(nil), "omo.msp.album.ReplySheetList")
}

func init() {
	proto.RegisterFile("proto/album/sheet.proto", fileDescriptor_16583bcb620c2e7c)
}

var fileDescriptor_16583bcb620c2e7c = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xd1, 0x6e, 0xda, 0x4a,
	0x10, 0x05, 0xdb, 0x38, 0x61, 0x1c, 0xb8, 0x57, 0xab, 0x36, 0x59, 0x59, 0x49, 0x85, 0xfc, 0x84,
	0xda, 0x88, 0x48, 0xf4, 0x0b, 0x42, 0x4b, 0x51, 0xaa, 0x46, 0x8d, 0x8c, 0xfa, 0xd2, 0x37, 0x07,
	0x4f, 0x9a, 0x6d, 0xb0, 0xd7, 0x5d, 0x2f, 0xa9, 0xf2, 0x39, 0xfd, 0x82, 0xfe, 0x62, 0xe5, 0x31,
	0x76, 0x80, 0xda, 0x50, 0xf1, 0xb6, 0x33, 0x73, 0xf6, 0x9c, 0x33, 0x33, 0x6b, 0x19, 0x4e, 0x12,
	0x25, 0xb5, 0xbc, 0x08, 0xe6, 0xb7, 0x8b, 0xe8, 0x22, 0xbd, 0x47, 0xd4, 0x03, 0xca, 0xb0, 0x8e,
	0x8c, 0xe4, 0x20, 0x4a, 0x93, 0x01, 0x95, 0x5c, 0xbe, 0x8a, 0x9b, 0xc9, 0x28, 0x92, 0x71, 0x0e,
	0xf4, 0x7e, 0x1b, 0xd0, 0x9e, 0x66, 0x17, 0xaf, 0xe2, 0x3b, 0xc9, 0xfe, 0x07, 0x73, 0x21, 0x42,
	0xde, 0xec, 0x35, 0xfb, 0x6d, 0x3f, 0x3b, 0xb2, 0x2e, 0x18, 0x22, 0xe4, 0x46, 0xaf, 0xd9, 0xb7,
	0x7c, 0x43, 0x84, 0x8c, 0xc3, 0xc1, 0x4c, 0x61, 0xa0, 0x31, 0xe4, 0x66, 0xaf, 0xd9, 0x37, 0xfd,
	0x22, 0xcc, 0x2a, 0x8b, 0x24, 0xa4, 0x8a, 0x95, 0x57, 0x96, 0x61, 0x79, 0x47, 0x2a, 0xde, 0x22,
	0xe6, 0x22, 0x64, 0x2e, 0x1c, 0xca, 0x04, 0x15, 0x95, 0x6c, 0x2a, 0x95, 0x31, 0x63, 0x60, 0xc5,
	0x41, 0x84, 0xfc, 0x80, 0xf2, 0x74, 0x66, 0xc7, 0x60, 0x2b, 0x8c, 0x02, 0xf5, 0xc0, 0x0f, 0x29,
	0xbb, 0x8c, 0xd8, 0x0b, 0x68, 0xc9, 0x9f, 0x31, 0x2a, 0xde, 0xa6, 0x74, 0x1e, 0x64, 0x0c, 0x81,
	0xc2, 0x80, 0x43, 0xce, 0x90, 0x9d, 0xb3, 0x9c, 0x16, 0x11, 0x72, 0x87, 0x2c, 0xd2, 0x99, 0x9d,
	0x43, 0x2b, 0x09, 0xbe, 0x61, 0xca, 0x5f, 0xf6, 0xcc, 0xbe, 0x33, 0x3c, 0x1e, 0xac, 0x0d, 0x6f,
	0x70, 0x13, 0x08, 0x75, 0x15, 0x6b, 0x3f, 0x07, 0x79, 0x0f, 0xe0, 0xf8, 0xf8, 0x83, 0x66, 0x76,
	0x19, 0x86, 0xa5, 0xcd, 0x66, 0xa5, 0x4d, 0x63, 0xcd, 0xe6, 0x6a, 0xbb, 0xe6, 0x46, 0xbb, 0x65,
	0x0b, 0xd6, 0x4a, 0x0b, 0xde, 0x77, 0xe8, 0x16, 0x62, 0x5f, 0x68, 0x9a, 0x15, 0x2b, 0x2a, 0x1c,
	0x18, 0x95, 0x0e, 0xcc, 0x5a, 0x07, 0xd6, 0xba, 0x03, 0x4f, 0x40, 0xa7, 0xd0, 0xba, 0xc9, 0x3a,
	0xad, 0x90, 0xda, 0x72, 0x9d, 0xbd, 0x06, 0x6b, 0x2e, 0x52, 0xcd, 0x9d, 0xad, 0x43, 0x24, 0x8c,
	0xa7, 0xe0, 0x3f, 0x1f, 0x93, 0xf9, 0xd3, 0x8a, 0xd8, 0x10, 0xec, 0x54, 0x07, 0x7a, 0x91, 0x92,
	0x9e, 0x33, 0x74, 0x37, 0x08, 0x72, 0x3c, 0x21, 0xfc, 0x25, 0x92, 0xbd, 0x59, 0x93, 0x3c, 0xa9,
	0x90, 0xfc, 0x24, 0xd2, 0x67, 0xcd, 0xee, 0xb3, 0x26, 0xbd, 0xf6, 0x7d, 0x24, 0xcf, 0xc1, 0x12,
	0xf1, 0x9d, 0xa4, 0x61, 0x3b, 0x43, 0xbe, 0x71, 0xa3, 0xe4, 0xf6, 0x09, 0xb5, 0xae, 0x99, 0x79,
	0xd9, 0x57, 0x93, 0xda, 0x34, 0xa8, 0xcd, 0x2d, 0x9a, 0x19, 0x6a, 0xf8, 0xab, 0x05, 0x47, 0x94,
	0x9b, 0xa2, 0x7a, 0x14, 0x33, 0x64, 0x63, 0xb0, 0x2f, 0xc3, 0xf0, 0x73, 0x8c, 0xec, 0x6f, 0xb1,
	0xf2, 0x1d, 0xbb, 0x67, 0x95, 0x46, 0x0a, 0x6e, 0xaf, 0xc1, 0x26, 0x00, 0xf9, 0x13, 0x1c, 0x05,
	0x29, 0xb2, 0xb3, 0x1a, 0xaa, 0x1c, 0xe2, 0xf2, 0x2a, 0xb6, 0x25, 0xd1, 0x3b, 0x68, 0xfb, 0x18,
	0xc9, 0x47, 0xac, 0xb1, 0xb4, 0xc0, 0x94, 0x34, 0xb7, 0x92, 0x8c, 0xc1, 0x9e, 0xa0, 0xde, 0xc5,
	0xb0, 0xb3, 0xa9, 0x31, 0xd8, 0x53, 0x0c, 0xd4, 0xec, 0x7e, 0x4f, 0x9a, 0x6c, 0xa7, 0x5e, 0x83,
	0x7d, 0x84, 0xf6, 0x24, 0x0f, 0x46, 0x4f, 0xec, 0xb4, 0x9a, 0xe9, 0x83, 0x98, 0x6b, 0x54, 0xbb,
	0xb9, 0xae, 0xe1, 0x68, 0x82, 0x3a, 0x7b, 0x02, 0x22, 0xd5, 0x62, 0xb6, 0x17, 0x5d, 0x71, 0xd9,
	0x6b, 0xb0, 0xf7, 0x70, 0xb8, 0x5c, 0x5b, 0xad, 0xb3, 0x7f, 0xd8, 0xd9, 0x35, 0x38, 0x39, 0x2a,
	0xff, 0x58, 0x4f, 0x6b, 0xb6, 0x4f, 0x55, 0xf7, 0x55, 0x6d, 0x8b, 0x54, 0xf7, 0x1a, 0xa3, 0xce,
	0x57, 0x67, 0xe5, 0x8f, 0x74, 0x6b, 0x53, 0xf0, 0xf6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbe,
	0xbf, 0xaf, 0x6e, 0xcf, 0x06, 0x00, 0x00,
}
