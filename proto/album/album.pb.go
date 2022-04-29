// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/album/album.proto

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

type AlbumInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	Location             string   `protobuf:"bytes,9,opt,name=location,proto3" json:"location,omitempty"`
	Passwords            string   `protobuf:"bytes,10,opt,name=passwords,proto3" json:"passwords,omitempty"`
	Cover                string   `protobuf:"bytes,11,opt,name=cover,proto3" json:"cover,omitempty"`
	Type                 uint32   `protobuf:"varint,12,opt,name=type,proto3" json:"type,omitempty"`
	Status               uint32   `protobuf:"varint,13,opt,name=status,proto3" json:"status,omitempty"`
	Style                uint32   `protobuf:"varint,14,opt,name=style,proto3" json:"style,omitempty"`
	Size                 uint64   `protobuf:"varint,15,opt,name=size,proto3" json:"size,omitempty"`
	Limit                uint64   `protobuf:"varint,16,opt,name=limit,proto3" json:"limit,omitempty"`
	Star                 uint32   `protobuf:"varint,17,opt,name=star,proto3" json:"star,omitempty"`
	Assets               []string `protobuf:"bytes,19,rep,name=assets,proto3" json:"assets,omitempty"`
	Tags                 []string `protobuf:"bytes,20,rep,name=tags,proto3" json:"tags,omitempty"`
	Targets              []string `protobuf:"bytes,21,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlbumInfo) Reset()         { *m = AlbumInfo{} }
func (m *AlbumInfo) String() string { return proto.CompactTextString(m) }
func (*AlbumInfo) ProtoMessage()    {}
func (*AlbumInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9099fd9bfb5dac32, []int{0}
}

func (m *AlbumInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlbumInfo.Unmarshal(m, b)
}
func (m *AlbumInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlbumInfo.Marshal(b, m, deterministic)
}
func (m *AlbumInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlbumInfo.Merge(m, src)
}
func (m *AlbumInfo) XXX_Size() int {
	return xxx_messageInfo_AlbumInfo.Size(m)
}
func (m *AlbumInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AlbumInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AlbumInfo proto.InternalMessageInfo

func (m *AlbumInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *AlbumInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AlbumInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *AlbumInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *AlbumInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *AlbumInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *AlbumInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AlbumInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *AlbumInfo) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *AlbumInfo) GetPasswords() string {
	if m != nil {
		return m.Passwords
	}
	return ""
}

func (m *AlbumInfo) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *AlbumInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *AlbumInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AlbumInfo) GetStyle() uint32 {
	if m != nil {
		return m.Style
	}
	return 0
}

func (m *AlbumInfo) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *AlbumInfo) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *AlbumInfo) GetStar() uint32 {
	if m != nil {
		return m.Star
	}
	return 0
}

func (m *AlbumInfo) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *AlbumInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *AlbumInfo) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ReqAlbumAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	Location             string   `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Type                 uint32   `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Style                uint32   `protobuf:"varint,6,opt,name=style,proto3" json:"style,omitempty"`
	Targets              []string `protobuf:"bytes,7,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqAlbumAdd) Reset()         { *m = ReqAlbumAdd{} }
func (m *ReqAlbumAdd) String() string { return proto.CompactTextString(m) }
func (*ReqAlbumAdd) ProtoMessage()    {}
func (*ReqAlbumAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_9099fd9bfb5dac32, []int{1}
}

func (m *ReqAlbumAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqAlbumAdd.Unmarshal(m, b)
}
func (m *ReqAlbumAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqAlbumAdd.Marshal(b, m, deterministic)
}
func (m *ReqAlbumAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqAlbumAdd.Merge(m, src)
}
func (m *ReqAlbumAdd) XXX_Size() int {
	return xxx_messageInfo_ReqAlbumAdd.Size(m)
}
func (m *ReqAlbumAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqAlbumAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqAlbumAdd proto.InternalMessageInfo

func (m *ReqAlbumAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqAlbumAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqAlbumAdd) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *ReqAlbumAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqAlbumAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqAlbumAdd) GetStyle() uint32 {
	if m != nil {
		return m.Style
	}
	return 0
}

func (m *ReqAlbumAdd) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ReqAlbumUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Cover                string   `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqAlbumUpdate) Reset()         { *m = ReqAlbumUpdate{} }
func (m *ReqAlbumUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqAlbumUpdate) ProtoMessage()    {}
func (*ReqAlbumUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_9099fd9bfb5dac32, []int{2}
}

func (m *ReqAlbumUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqAlbumUpdate.Unmarshal(m, b)
}
func (m *ReqAlbumUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqAlbumUpdate.Marshal(b, m, deterministic)
}
func (m *ReqAlbumUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqAlbumUpdate.Merge(m, src)
}
func (m *ReqAlbumUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqAlbumUpdate.Size(m)
}
func (m *ReqAlbumUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqAlbumUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqAlbumUpdate proto.InternalMessageInfo

func (m *ReqAlbumUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqAlbumUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqAlbumUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqAlbumUpdate) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqAlbumUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyAlbumInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *AlbumInfo   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyAlbumInfo) Reset()         { *m = ReplyAlbumInfo{} }
func (m *ReplyAlbumInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyAlbumInfo) ProtoMessage()    {}
func (*ReplyAlbumInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9099fd9bfb5dac32, []int{3}
}

func (m *ReplyAlbumInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyAlbumInfo.Unmarshal(m, b)
}
func (m *ReplyAlbumInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyAlbumInfo.Marshal(b, m, deterministic)
}
func (m *ReplyAlbumInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyAlbumInfo.Merge(m, src)
}
func (m *ReplyAlbumInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyAlbumInfo.Size(m)
}
func (m *ReplyAlbumInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyAlbumInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyAlbumInfo proto.InternalMessageInfo

func (m *ReplyAlbumInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyAlbumInfo) GetInfo() *AlbumInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyAlbumList struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*AlbumInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyAlbumList) Reset()         { *m = ReplyAlbumList{} }
func (m *ReplyAlbumList) String() string { return proto.CompactTextString(m) }
func (*ReplyAlbumList) ProtoMessage()    {}
func (*ReplyAlbumList) Descriptor() ([]byte, []int) {
	return fileDescriptor_9099fd9bfb5dac32, []int{4}
}

func (m *ReplyAlbumList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyAlbumList.Unmarshal(m, b)
}
func (m *ReplyAlbumList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyAlbumList.Marshal(b, m, deterministic)
}
func (m *ReplyAlbumList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyAlbumList.Merge(m, src)
}
func (m *ReplyAlbumList) XXX_Size() int {
	return xxx_messageInfo_ReplyAlbumList.Size(m)
}
func (m *ReplyAlbumList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyAlbumList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyAlbumList proto.InternalMessageInfo

func (m *ReplyAlbumList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyAlbumList) GetList() []*AlbumInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*AlbumInfo)(nil), "omo.msp.album.AlbumInfo")
	proto.RegisterType((*ReqAlbumAdd)(nil), "omo.msp.album.ReqAlbumAdd")
	proto.RegisterType((*ReqAlbumUpdate)(nil), "omo.msp.album.ReqAlbumUpdate")
	proto.RegisterType((*ReplyAlbumInfo)(nil), "omo.msp.album.ReplyAlbumInfo")
	proto.RegisterType((*ReplyAlbumList)(nil), "omo.msp.album.ReplyAlbumList")
}

func init() {
	proto.RegisterFile("proto/album/album.proto", fileDescriptor_9099fd9bfb5dac32)
}

var fileDescriptor_9099fd9bfb5dac32 = []byte{
	// 670 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0xe3, 0xc4, 0x6d, 0x26, 0x3f, 0x2d, 0x4b, 0x81, 0x55, 0xd4, 0x4a, 0x51, 0x4e, 0x39,
	0xa0, 0x20, 0x85, 0x27, 0x48, 0x51, 0x1b, 0x15, 0x81, 0x40, 0x8e, 0xb8, 0x70, 0xdb, 0xda, 0xdb,
	0xb2, 0xc2, 0xf6, 0x9a, 0xdd, 0x4d, 0x51, 0x38, 0xf1, 0x52, 0x1c, 0x78, 0x14, 0xde, 0x06, 0xed,
	0x6c, 0xe2, 0x3a, 0xad, 0xd3, 0xa0, 0x70, 0xa9, 0xe6, 0x9b, 0x99, 0xfd, 0x66, 0xe6, 0x9b, 0x71,
	0x03, 0x2f, 0x72, 0x25, 0x8d, 0x7c, 0xc5, 0x92, 0xab, 0x79, 0xea, 0xfe, 0x8e, 0xd0, 0x43, 0x3a,
	0x32, 0x95, 0xa3, 0x54, 0xe7, 0x23, 0x74, 0xf6, 0x68, 0x39, 0x2f, 0x92, 0x69, 0x2a, 0x33, 0x97,
	0x38, 0xf8, 0xe3, 0x43, 0x73, 0x62, 0xdd, 0x97, 0xd9, 0xb5, 0x24, 0x47, 0xe0, 0xcf, 0x45, 0x4c,
	0xbd, 0xbe, 0x37, 0x6c, 0x86, 0xd6, 0x24, 0x5d, 0xa8, 0x89, 0x98, 0xd6, 0xfa, 0xde, 0xb0, 0x1e,
	0xd6, 0x44, 0x4c, 0x28, 0xec, 0x47, 0x8a, 0x33, 0xc3, 0x63, 0xea, 0xf7, 0xbd, 0xa1, 0x1f, 0xae,
	0xa0, 0x8d, 0xcc, 0xf3, 0x18, 0x23, 0x75, 0x17, 0x59, 0xc2, 0xe2, 0x8d, 0x54, 0xb4, 0x81, 0xcc,
	0x2b, 0x48, 0x7a, 0x70, 0x20, 0x73, 0xae, 0x30, 0x14, 0x60, 0xa8, 0xc0, 0x84, 0x40, 0x3d, 0x63,
	0x29, 0xa7, 0xfb, 0xe8, 0x47, 0x9b, 0x3c, 0x87, 0x40, 0xf1, 0x94, 0xa9, 0xaf, 0xf4, 0x00, 0xbd,
	0x4b, 0x64, 0x79, 0x12, 0x19, 0x31, 0x23, 0x64, 0x46, 0x9b, 0x8e, 0x67, 0x85, 0xc9, 0x09, 0x34,
	0x73, 0xa6, 0xf5, 0x77, 0xa9, 0x62, 0x4d, 0x01, 0x83, 0x77, 0x0e, 0x72, 0x0c, 0x8d, 0x48, 0xde,
	0x72, 0x45, 0x5b, 0x18, 0x71, 0xc0, 0xd6, 0x36, 0x8b, 0x9c, 0xd3, 0x76, 0xdf, 0x1b, 0x76, 0x42,
	0xb4, 0x6d, 0x6d, 0x6d, 0x98, 0x99, 0x6b, 0xda, 0x41, 0xef, 0x12, 0x59, 0x06, 0x6d, 0x16, 0x09,
	0xa7, 0x5d, 0x74, 0x3b, 0x60, 0x19, 0xb4, 0xf8, 0xc1, 0xe9, 0x21, 0x2a, 0x87, 0xb6, 0xcd, 0x4c,
	0x44, 0x2a, 0x0c, 0x3d, 0x42, 0xa7, 0x03, 0x98, 0x69, 0x98, 0xa2, 0x4f, 0x5c, 0x2d, 0x6b, 0xdb,
	0x5a, 0x4c, 0x6b, 0x6e, 0x34, 0x7d, 0xda, 0xf7, 0xed, 0x9c, 0x0e, 0x61, 0x5f, 0xec, 0x46, 0xd3,
	0x63, 0xf4, 0xa2, 0x6d, 0xd5, 0x35, 0x4c, 0xdd, 0xd8, 0xe4, 0x67, 0xe8, 0x5e, 0xc1, 0xc1, 0x2f,
	0x0f, 0x5a, 0x21, 0xff, 0x86, 0xeb, 0x9d, 0xc4, 0x71, 0xa1, 0xa8, 0x57, 0xa9, 0x68, 0x6d, 0xa3,
	0xa2, 0xfe, 0x3d, 0x45, 0xcb, 0x5b, 0xab, 0x3f, 0xdc, 0x1a, 0x2a, 0xd7, 0x28, 0x29, 0x57, 0x28,
	0x14, 0x94, 0x15, 0x2a, 0xf5, 0xbd, 0xbf, 0xde, 0xf7, 0x4f, 0x0f, 0xba, 0xab, 0xbe, 0x3f, 0xe1,
	0x0d, 0x55, 0x1c, 0xe6, 0x6a, 0x98, 0x5a, 0xe5, 0x30, 0xfe, 0xda, 0x30, 0xc5, 0x92, 0xeb, 0xe5,
	0x25, 0x97, 0xc7, 0x68, 0xac, 0x8f, 0x31, 0x50, 0xb6, 0x83, 0x3c, 0x59, 0xdc, 0x7d, 0x1a, 0xe3,
	0x62, 0xfd, 0xb6, 0x89, 0xd6, 0xb8, 0x37, 0x5a, 0xfb, 0xc4, 0x46, 0x98, 0x3e, 0xc3, 0x8c, 0xe2,
	0x34, 0x5e, 0x42, 0x5d, 0x64, 0xd7, 0x12, 0x7b, 0x6c, 0x8d, 0xe9, 0xbd, 0x17, 0x05, 0x77, 0x88,
	0x59, 0xeb, 0x35, 0xdf, 0x09, 0x6d, 0x76, 0xad, 0x99, 0x08, 0x6d, 0x68, 0xad, 0xef, 0x3f, 0x5e,
	0xd3, 0x66, 0x8d, 0x7f, 0x07, 0xd0, 0x46, 0xdf, 0x8c, 0xab, 0x5b, 0x11, 0x71, 0x72, 0x0e, 0xc1,
	0x24, 0x8e, 0x3f, 0x64, 0x9c, 0x3c, 0x2c, 0x56, 0x5c, 0x52, 0xef, 0xb4, 0xaa, 0x91, 0x82, 0x7b,
	0xb0, 0x47, 0xa6, 0x00, 0x6e, 0x73, 0x67, 0x4c, 0x73, 0x72, 0xba, 0x81, 0xca, 0xa5, 0xf4, 0x68,
	0x15, 0xdb, 0x92, 0xe8, 0x0d, 0x34, 0x43, 0x9e, 0xca, 0x5b, 0xbe, 0xa1, 0xa5, 0x39, 0xd7, 0xc6,
	0xa6, 0x3e, 0x4a, 0x72, 0x0e, 0xc1, 0x94, 0x9b, 0x6d, 0x0c, 0x5b, 0x87, 0x3a, 0x87, 0x60, 0xc6,
	0x99, 0x8a, 0xbe, 0xec, 0x48, 0x63, 0x77, 0x3a, 0xd8, 0x23, 0x1f, 0xe1, 0x70, 0xca, 0x8d, 0x05,
	0x67, 0x8b, 0x0b, 0x91, 0x18, 0xae, 0xc8, 0x49, 0x35, 0x9f, 0x8b, 0x6e, 0x67, 0x7c, 0x0f, 0xed,
	0x29, 0x37, 0xf6, 0x10, 0x84, 0x36, 0x22, 0xda, 0x85, 0xae, 0x78, 0x3c, 0xd8, 0x23, 0x6f, 0xa1,
	0xbb, 0x5c, 0xde, 0x96, 0xfe, 0xfe, 0x61, 0x7f, 0x97, 0xd0, 0x76, 0x59, 0xee, 0x4c, 0xab, 0x4e,
	0xc1, 0x29, 0x67, 0x2e, 0x12, 0x76, 0xb3, 0x65, 0x8b, 0xad, 0x49, 0x9e, 0xf3, 0x2c, 0x9e, 0xd8,
	0x7f, 0x86, 0x9b, 0x76, 0x60, 0x25, 0xa9, 0xa6, 0x59, 0x8a, 0x35, 0x85, 0xce, 0x6c, 0x7e, 0x65,
	0x14, 0x8b, 0xcc, 0x7f, 0x11, 0x9d, 0x75, 0x3e, 0xb7, 0x4a, 0x3f, 0xab, 0x57, 0x01, 0x82, 0xd7,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x03, 0xc8, 0x88, 0xd3, 0x94, 0x07, 0x00, 0x00,
}
