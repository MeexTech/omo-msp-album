// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/album/collective.proto

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

type CollectiveInfo struct {
	Uid                  string        `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64        `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64         `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64         `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string        `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string        `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string        `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string        `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	Cover                string        `protobuf:"bytes,9,opt,name=cover,proto3" json:"cover,omitempty"`
	Owner                string        `protobuf:"bytes,10,opt,name=owner,proto3" json:"owner,omitempty"`
	Size                 uint64        `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	Limit                uint32        `protobuf:"varint,12,opt,name=limit,proto3" json:"limit,omitempty"`
	Star                 uint32        `protobuf:"varint,13,opt,name=star,proto3" json:"star,omitempty"`
	Status               uint32        `protobuf:"varint,14,opt,name=status,proto3" json:"status,omitempty"`
	Permission           uint32        `protobuf:"varint,15,opt,name=permission,proto3" json:"permission,omitempty"`
	Type                 uint32        `protobuf:"varint,16,opt,name=type,proto3" json:"type,omitempty"`
	Duration             *DurationInfo `protobuf:"bytes,17,opt,name=duration,proto3" json:"duration,omitempty"`
	Assets               []string      `protobuf:"bytes,21,rep,name=assets,proto3" json:"assets,omitempty"`
	Tags                 []string      `protobuf:"bytes,22,rep,name=tags,proto3" json:"tags,omitempty"`
	Medias               []*MediaInfo  `protobuf:"bytes,31,rep,name=medias,proto3" json:"medias,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CollectiveInfo) Reset()         { *m = CollectiveInfo{} }
func (m *CollectiveInfo) String() string { return proto.CompactTextString(m) }
func (*CollectiveInfo) ProtoMessage()    {}
func (*CollectiveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_81657e36c0b4b126, []int{0}
}

func (m *CollectiveInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectiveInfo.Unmarshal(m, b)
}
func (m *CollectiveInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectiveInfo.Marshal(b, m, deterministic)
}
func (m *CollectiveInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectiveInfo.Merge(m, src)
}
func (m *CollectiveInfo) XXX_Size() int {
	return xxx_messageInfo_CollectiveInfo.Size(m)
}
func (m *CollectiveInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectiveInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CollectiveInfo proto.InternalMessageInfo

func (m *CollectiveInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CollectiveInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CollectiveInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *CollectiveInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *CollectiveInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *CollectiveInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *CollectiveInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CollectiveInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *CollectiveInfo) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *CollectiveInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *CollectiveInfo) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *CollectiveInfo) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *CollectiveInfo) GetStar() uint32 {
	if m != nil {
		return m.Star
	}
	return 0
}

func (m *CollectiveInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CollectiveInfo) GetPermission() uint32 {
	if m != nil {
		return m.Permission
	}
	return 0
}

func (m *CollectiveInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *CollectiveInfo) GetDuration() *DurationInfo {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *CollectiveInfo) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *CollectiveInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *CollectiveInfo) GetMedias() []*MediaInfo {
	if m != nil {
		return m.Medias
	}
	return nil
}

type ReqCollectiveAdd struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string        `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string        `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Owner                string        `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Type                 uint32        `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Duration             *DurationInfo `protobuf:"bytes,11,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReqCollectiveAdd) Reset()         { *m = ReqCollectiveAdd{} }
func (m *ReqCollectiveAdd) String() string { return proto.CompactTextString(m) }
func (*ReqCollectiveAdd) ProtoMessage()    {}
func (*ReqCollectiveAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_81657e36c0b4b126, []int{1}
}

func (m *ReqCollectiveAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqCollectiveAdd.Unmarshal(m, b)
}
func (m *ReqCollectiveAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqCollectiveAdd.Marshal(b, m, deterministic)
}
func (m *ReqCollectiveAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqCollectiveAdd.Merge(m, src)
}
func (m *ReqCollectiveAdd) XXX_Size() int {
	return xxx_messageInfo_ReqCollectiveAdd.Size(m)
}
func (m *ReqCollectiveAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqCollectiveAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqCollectiveAdd proto.InternalMessageInfo

func (m *ReqCollectiveAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqCollectiveAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqCollectiveAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqCollectiveAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqCollectiveAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqCollectiveAdd) GetDuration() *DurationInfo {
	if m != nil {
		return m.Duration
	}
	return nil
}

type ReqCollectiveUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Cover                string   `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqCollectiveUpdate) Reset()         { *m = ReqCollectiveUpdate{} }
func (m *ReqCollectiveUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqCollectiveUpdate) ProtoMessage()    {}
func (*ReqCollectiveUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_81657e36c0b4b126, []int{2}
}

func (m *ReqCollectiveUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqCollectiveUpdate.Unmarshal(m, b)
}
func (m *ReqCollectiveUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqCollectiveUpdate.Marshal(b, m, deterministic)
}
func (m *ReqCollectiveUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqCollectiveUpdate.Merge(m, src)
}
func (m *ReqCollectiveUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqCollectiveUpdate.Size(m)
}
func (m *ReqCollectiveUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqCollectiveUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqCollectiveUpdate proto.InternalMessageInfo

func (m *ReqCollectiveUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqCollectiveUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqCollectiveUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqCollectiveUpdate) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqCollectiveUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyCollectiveInfo struct {
	Status               *ReplyStatus    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *CollectiveInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReplyCollectiveInfo) Reset()         { *m = ReplyCollectiveInfo{} }
func (m *ReplyCollectiveInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyCollectiveInfo) ProtoMessage()    {}
func (*ReplyCollectiveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_81657e36c0b4b126, []int{3}
}

func (m *ReplyCollectiveInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyCollectiveInfo.Unmarshal(m, b)
}
func (m *ReplyCollectiveInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyCollectiveInfo.Marshal(b, m, deterministic)
}
func (m *ReplyCollectiveInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyCollectiveInfo.Merge(m, src)
}
func (m *ReplyCollectiveInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyCollectiveInfo.Size(m)
}
func (m *ReplyCollectiveInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyCollectiveInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyCollectiveInfo proto.InternalMessageInfo

func (m *ReplyCollectiveInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyCollectiveInfo) GetInfo() *CollectiveInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyCollectiveList struct {
	Status               *ReplyStatus      `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*CollectiveInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReplyCollectiveList) Reset()         { *m = ReplyCollectiveList{} }
func (m *ReplyCollectiveList) String() string { return proto.CompactTextString(m) }
func (*ReplyCollectiveList) ProtoMessage()    {}
func (*ReplyCollectiveList) Descriptor() ([]byte, []int) {
	return fileDescriptor_81657e36c0b4b126, []int{4}
}

func (m *ReplyCollectiveList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyCollectiveList.Unmarshal(m, b)
}
func (m *ReplyCollectiveList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyCollectiveList.Marshal(b, m, deterministic)
}
func (m *ReplyCollectiveList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyCollectiveList.Merge(m, src)
}
func (m *ReplyCollectiveList) XXX_Size() int {
	return xxx_messageInfo_ReplyCollectiveList.Size(m)
}
func (m *ReplyCollectiveList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyCollectiveList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyCollectiveList proto.InternalMessageInfo

func (m *ReplyCollectiveList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyCollectiveList) GetList() []*CollectiveInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*CollectiveInfo)(nil), "omo.msp.album.CollectiveInfo")
	proto.RegisterType((*ReqCollectiveAdd)(nil), "omo.msp.album.ReqCollectiveAdd")
	proto.RegisterType((*ReqCollectiveUpdate)(nil), "omo.msp.album.ReqCollectiveUpdate")
	proto.RegisterType((*ReplyCollectiveInfo)(nil), "omo.msp.album.ReplyCollectiveInfo")
	proto.RegisterType((*ReplyCollectiveList)(nil), "omo.msp.album.ReplyCollectiveList")
}

func init() {
	proto.RegisterFile("proto/album/collective.proto", fileDescriptor_81657e36c0b4b126)
}

var fileDescriptor_81657e36c0b4b126 = []byte{
	// 703 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0xc6, 0xf9, 0x31, 0x70, 0x42, 0x02, 0xcc, 0xbd, 0x17, 0x8d, 0x72, 0xa1, 0x44, 0x5e, 0x65,
	0x15, 0xda, 0x74, 0xd1, 0x75, 0xa0, 0x25, 0x02, 0x81, 0x90, 0x1c, 0xb1, 0xe9, 0xce, 0xd8, 0x03,
	0x1d, 0xd5, 0xf6, 0x98, 0x99, 0x71, 0x2a, 0xaa, 0x6e, 0xfb, 0x08, 0x7d, 0x9c, 0xbe, 0x49, 0x1f,
	0xa6, 0x9a, 0x63, 0x63, 0xe2, 0xe0, 0x10, 0x5a, 0x76, 0x73, 0xfe, 0xbe, 0xf3, 0xf3, 0x7d, 0x89,
	0x61, 0x37, 0x91, 0x42, 0x8b, 0x03, 0x2f, 0xbc, 0x4a, 0xa3, 0x03, 0x5f, 0x84, 0x21, 0xf3, 0x35,
	0x9f, 0xb2, 0x01, 0xba, 0x49, 0x5b, 0x44, 0x62, 0x10, 0xa9, 0x64, 0x80, 0xf1, 0x2e, 0x2d, 0x27,
	0x47, 0x91, 0x88, 0xb3, 0x44, 0xe7, 0x47, 0x03, 0x3a, 0x47, 0x45, 0xf5, 0x49, 0x7c, 0x2d, 0xc8,
	0x16, 0xd4, 0x53, 0x1e, 0x50, 0xab, 0x67, 0xf5, 0xd7, 0x5d, 0xf3, 0x24, 0x1d, 0xa8, 0xf1, 0x80,
	0xd6, 0x7a, 0x56, 0xbf, 0xe1, 0xd6, 0x78, 0x40, 0x28, 0xac, 0xfa, 0x92, 0x79, 0x9a, 0x05, 0xb4,
	0xde, 0xb3, 0xfa, 0x75, 0xf7, 0xde, 0x34, 0x91, 0x34, 0x09, 0x30, 0xd2, 0xc8, 0x22, 0xb9, 0x59,
	0xd4, 0x08, 0x49, 0x9b, 0x88, 0x7c, 0x6f, 0x92, 0x2e, 0xac, 0x89, 0x84, 0x49, 0x0c, 0xd9, 0x18,
	0x2a, 0x6c, 0x42, 0xa0, 0x11, 0x7b, 0x11, 0xa3, 0xab, 0xe8, 0xc7, 0x37, 0xd9, 0x01, 0x5b, 0xb2,
	0xc8, 0x93, 0x9f, 0xe9, 0x1a, 0x7a, 0x73, 0x8b, 0xfc, 0x0b, 0x4d, 0x5f, 0x4c, 0x99, 0xa4, 0xeb,
	0xe8, 0xce, 0x0c, 0xe3, 0x15, 0x5f, 0x62, 0x26, 0x29, 0x64, 0x5e, 0x34, 0x0c, 0xae, 0xe2, 0x5f,
	0x19, 0x6d, 0xe1, 0x4e, 0xf8, 0x36, 0x99, 0x21, 0x8f, 0xb8, 0xa6, 0x1b, 0x3d, 0xab, 0xdf, 0x76,
	0x33, 0x03, 0x33, 0xb5, 0x27, 0x69, 0x1b, 0x9d, 0xf8, 0x36, 0x13, 0x28, 0xed, 0xe9, 0x54, 0xd1,
	0x0e, 0x7a, 0x73, 0x8b, 0xbc, 0x02, 0x48, 0x98, 0x8c, 0xb8, 0x52, 0x5c, 0xc4, 0x74, 0x13, 0x63,
	0x33, 0x1e, 0x83, 0xa5, 0xef, 0x12, 0x46, 0xb7, 0x32, 0x2c, 0xf3, 0x26, 0xef, 0x60, 0x2d, 0x48,
	0xa5, 0xa7, 0x4d, 0xc5, 0x76, 0xcf, 0xea, 0xb7, 0x86, 0xff, 0x0f, 0x4a, 0xe4, 0x0d, 0xde, 0xe7,
	0x61, 0x43, 0x8e, 0x5b, 0x24, 0x9b, 0x21, 0x3c, 0xa5, 0x98, 0x56, 0xf4, 0xbf, 0x5e, 0xdd, 0x9c,
	0x21, 0xb3, 0xb0, 0x89, 0x77, 0xa3, 0xe8, 0x0e, 0x7a, 0xf1, 0x4d, 0x5e, 0x83, 0x1d, 0xb1, 0x80,
	0x7b, 0x8a, 0xee, 0xf7, 0xea, 0xfd, 0xd6, 0x90, 0xce, 0xb5, 0x38, 0x37, 0x41, 0xc4, 0xcf, 0xf3,
	0x9c, 0x9f, 0x16, 0x6c, 0xb9, 0xec, 0xf6, 0x41, 0x1a, 0xa3, 0x20, 0x28, 0xd8, 0xb0, 0x2a, 0xd9,
	0xa8, 0x95, 0xd8, 0x98, 0x65, 0xb5, 0x3e, 0xc7, 0x6a, 0xc1, 0x49, 0x63, 0x8e, 0x13, 0xbc, 0x4e,
	0x73, 0xc1, 0x75, 0x5a, 0x7f, 0x70, 0x1d, 0xe7, 0xbb, 0x05, 0xff, 0x94, 0xe6, 0xbf, 0x44, 0x1d,
	0x56, 0x88, 0xfb, 0x7e, 0xa9, 0x5a, 0xe5, 0x52, 0xf5, 0x6a, 0x89, 0x35, 0x66, 0x25, 0x36, 0xbb,
	0x6a, 0xb3, 0xbc, 0xaa, 0xf3, 0xcd, 0x8c, 0x91, 0x84, 0x77, 0x73, 0xbf, 0xb1, 0x61, 0xa1, 0x20,
	0x0b, 0xb7, 0xea, 0xce, 0x6d, 0x85, 0x35, 0x13, 0xcc, 0x28, 0xd4, 0xf5, 0x06, 0x1a, 0x3c, 0xbe,
	0x16, 0x38, 0x68, 0x6b, 0xb8, 0x37, 0x57, 0x51, 0x6e, 0xe0, 0x62, 0x6a, 0x45, 0xf7, 0x33, 0xae,
	0xf4, 0xdf, 0x76, 0x0f, 0xb9, 0xd2, 0xb4, 0x86, 0x02, 0x5a, 0xd6, 0xdd, 0xa4, 0x0e, 0x7f, 0xd9,
	0xb0, 0xfd, 0x10, 0x98, 0x30, 0x39, 0xe5, 0x3e, 0x23, 0x17, 0x60, 0x8f, 0x82, 0xe0, 0x22, 0x66,
	0x64, 0xff, 0x51, 0xdb, 0xb2, 0xde, 0xba, 0x4e, 0xd5, 0x5c, 0xe5, 0x56, 0xce, 0x0a, 0x39, 0x03,
	0xc8, 0xc8, 0x3d, 0xf4, 0x14, 0x23, 0xce, 0x53, 0xa0, 0x59, 0x5e, 0x97, 0x56, 0xe1, 0xe6, 0x68,
	0x47, 0xb0, 0xee, 0xb2, 0x48, 0x4c, 0x99, 0x99, 0xf0, 0xf1, 0x61, 0x6e, 0x53, 0xa6, 0xb4, 0x49,
	0x7d, 0x12, 0xe4, 0x14, 0xec, 0x31, 0xd3, 0xcb, 0x10, 0x9e, 0xb7, 0xde, 0x29, 0xd8, 0x13, 0xe6,
	0x49, 0xff, 0xd3, 0x4b, 0xb0, 0x0c, 0xed, 0xce, 0x0a, 0xb9, 0x84, 0xcd, 0x31, 0xd3, 0xc6, 0x38,
	0xbc, 0x3b, 0xe6, 0xa1, 0x66, 0x92, 0xec, 0x56, 0x83, 0x66, 0xd1, 0x67, 0xc2, 0x9e, 0xc3, 0xc6,
	0x98, 0x69, 0x23, 0x18, 0xae, 0x34, 0xf7, 0x97, 0x60, 0xee, 0x2d, 0x52, 0x1b, 0x16, 0xe3, 0xc6,
	0x9d, 0x9c, 0xd0, 0x25, 0x43, 0x3e, 0x83, 0xce, 0x13, 0xd8, 0xc8, 0xb2, 0x32, 0x39, 0x93, 0xbd,
	0x45, 0x37, 0xd4, 0xc7, 0xa1, 0x77, 0xf3, 0x24, 0xd4, 0x07, 0x68, 0x8d, 0x92, 0x84, 0xc5, 0xc1,
	0xc8, 0xfc, 0xd1, 0x2e, 0x62, 0xc3, 0x9c, 0xa4, 0x1a, 0x26, 0x3f, 0xd6, 0x18, 0xda, 0x93, 0xf4,
	0x4a, 0x4b, 0xcf, 0xd7, 0x2f, 0x02, 0x3a, 0x6c, 0x7f, 0x6c, 0xcd, 0x7c, 0xd6, 0xaf, 0x6c, 0x34,
	0xde, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x82, 0x16, 0x4d, 0x19, 0x08, 0x00, 0x00,
}
