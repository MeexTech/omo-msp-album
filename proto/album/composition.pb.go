// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/album/composition.proto

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

type CompositionInfo struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64      `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64       `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64       `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string      `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string      `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string      `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string      `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	Owner                string      `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
	Aspect               uint32      `protobuf:"varint,10,opt,name=aspect,proto3" json:"aspect,omitempty"`
	Slots                []*SlotInfo `protobuf:"bytes,21,rep,name=slots,proto3" json:"slots,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CompositionInfo) Reset()         { *m = CompositionInfo{} }
func (m *CompositionInfo) String() string { return proto.CompactTextString(m) }
func (*CompositionInfo) ProtoMessage()    {}
func (*CompositionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e97c75056c972412, []int{0}
}

func (m *CompositionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompositionInfo.Unmarshal(m, b)
}
func (m *CompositionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompositionInfo.Marshal(b, m, deterministic)
}
func (m *CompositionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompositionInfo.Merge(m, src)
}
func (m *CompositionInfo) XXX_Size() int {
	return xxx_messageInfo_CompositionInfo.Size(m)
}
func (m *CompositionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CompositionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CompositionInfo proto.InternalMessageInfo

func (m *CompositionInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CompositionInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CompositionInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *CompositionInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *CompositionInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *CompositionInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *CompositionInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CompositionInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *CompositionInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *CompositionInfo) GetAspect() uint32 {
	if m != nil {
		return m.Aspect
	}
	return 0
}

func (m *CompositionInfo) GetSlots() []*SlotInfo {
	if m != nil {
		return m.Slots
	}
	return nil
}

type SlotInfo struct {
	Index                uint32   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	X                    int32    `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	Width                uint32   `protobuf:"varint,4,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32   `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty"`
	Type                 uint32   `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SlotInfo) Reset()         { *m = SlotInfo{} }
func (m *SlotInfo) String() string { return proto.CompactTextString(m) }
func (*SlotInfo) ProtoMessage()    {}
func (*SlotInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e97c75056c972412, []int{1}
}

func (m *SlotInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlotInfo.Unmarshal(m, b)
}
func (m *SlotInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlotInfo.Marshal(b, m, deterministic)
}
func (m *SlotInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlotInfo.Merge(m, src)
}
func (m *SlotInfo) XXX_Size() int {
	return xxx_messageInfo_SlotInfo.Size(m)
}
func (m *SlotInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SlotInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SlotInfo proto.InternalMessageInfo

func (m *SlotInfo) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *SlotInfo) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *SlotInfo) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *SlotInfo) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *SlotInfo) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *SlotInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type ReqCompositionAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Aspect               uint32   `protobuf:"varint,4,opt,name=aspect,proto3" json:"aspect,omitempty"`
	Owner                string   `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqCompositionAdd) Reset()         { *m = ReqCompositionAdd{} }
func (m *ReqCompositionAdd) String() string { return proto.CompactTextString(m) }
func (*ReqCompositionAdd) ProtoMessage()    {}
func (*ReqCompositionAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_e97c75056c972412, []int{2}
}

func (m *ReqCompositionAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqCompositionAdd.Unmarshal(m, b)
}
func (m *ReqCompositionAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqCompositionAdd.Marshal(b, m, deterministic)
}
func (m *ReqCompositionAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqCompositionAdd.Merge(m, src)
}
func (m *ReqCompositionAdd) XXX_Size() int {
	return xxx_messageInfo_ReqCompositionAdd.Size(m)
}
func (m *ReqCompositionAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqCompositionAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqCompositionAdd proto.InternalMessageInfo

func (m *ReqCompositionAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqCompositionAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqCompositionAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqCompositionAdd) GetAspect() uint32 {
	if m != nil {
		return m.Aspect
	}
	return 0
}

func (m *ReqCompositionAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type ReqCompositionUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Aspect               uint32   `protobuf:"varint,5,opt,name=aspect,proto3" json:"aspect,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqCompositionUpdate) Reset()         { *m = ReqCompositionUpdate{} }
func (m *ReqCompositionUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqCompositionUpdate) ProtoMessage()    {}
func (*ReqCompositionUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_e97c75056c972412, []int{3}
}

func (m *ReqCompositionUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqCompositionUpdate.Unmarshal(m, b)
}
func (m *ReqCompositionUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqCompositionUpdate.Marshal(b, m, deterministic)
}
func (m *ReqCompositionUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqCompositionUpdate.Merge(m, src)
}
func (m *ReqCompositionUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqCompositionUpdate.Size(m)
}
func (m *ReqCompositionUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqCompositionUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqCompositionUpdate proto.InternalMessageInfo

func (m *ReqCompositionUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqCompositionUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqCompositionUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqCompositionUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqCompositionUpdate) GetAspect() uint32 {
	if m != nil {
		return m.Aspect
	}
	return 0
}

type ReqCompositionSlots struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string      `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Slots                []*SlotInfo `protobuf:"bytes,3,rep,name=slots,proto3" json:"slots,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReqCompositionSlots) Reset()         { *m = ReqCompositionSlots{} }
func (m *ReqCompositionSlots) String() string { return proto.CompactTextString(m) }
func (*ReqCompositionSlots) ProtoMessage()    {}
func (*ReqCompositionSlots) Descriptor() ([]byte, []int) {
	return fileDescriptor_e97c75056c972412, []int{4}
}

func (m *ReqCompositionSlots) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqCompositionSlots.Unmarshal(m, b)
}
func (m *ReqCompositionSlots) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqCompositionSlots.Marshal(b, m, deterministic)
}
func (m *ReqCompositionSlots) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqCompositionSlots.Merge(m, src)
}
func (m *ReqCompositionSlots) XXX_Size() int {
	return xxx_messageInfo_ReqCompositionSlots.Size(m)
}
func (m *ReqCompositionSlots) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqCompositionSlots.DiscardUnknown(m)
}

var xxx_messageInfo_ReqCompositionSlots proto.InternalMessageInfo

func (m *ReqCompositionSlots) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqCompositionSlots) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqCompositionSlots) GetSlots() []*SlotInfo {
	if m != nil {
		return m.Slots
	}
	return nil
}

type ReplyCompositionSlots struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*SlotInfo  `protobuf:"bytes,11,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyCompositionSlots) Reset()         { *m = ReplyCompositionSlots{} }
func (m *ReplyCompositionSlots) String() string { return proto.CompactTextString(m) }
func (*ReplyCompositionSlots) ProtoMessage()    {}
func (*ReplyCompositionSlots) Descriptor() ([]byte, []int) {
	return fileDescriptor_e97c75056c972412, []int{5}
}

func (m *ReplyCompositionSlots) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyCompositionSlots.Unmarshal(m, b)
}
func (m *ReplyCompositionSlots) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyCompositionSlots.Marshal(b, m, deterministic)
}
func (m *ReplyCompositionSlots) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyCompositionSlots.Merge(m, src)
}
func (m *ReplyCompositionSlots) XXX_Size() int {
	return xxx_messageInfo_ReplyCompositionSlots.Size(m)
}
func (m *ReplyCompositionSlots) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyCompositionSlots.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyCompositionSlots proto.InternalMessageInfo

func (m *ReplyCompositionSlots) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyCompositionSlots) GetList() []*SlotInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyCompositionInfo struct {
	Status               *ReplyStatus     `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *CompositionInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReplyCompositionInfo) Reset()         { *m = ReplyCompositionInfo{} }
func (m *ReplyCompositionInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyCompositionInfo) ProtoMessage()    {}
func (*ReplyCompositionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e97c75056c972412, []int{6}
}

func (m *ReplyCompositionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyCompositionInfo.Unmarshal(m, b)
}
func (m *ReplyCompositionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyCompositionInfo.Marshal(b, m, deterministic)
}
func (m *ReplyCompositionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyCompositionInfo.Merge(m, src)
}
func (m *ReplyCompositionInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyCompositionInfo.Size(m)
}
func (m *ReplyCompositionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyCompositionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyCompositionInfo proto.InternalMessageInfo

func (m *ReplyCompositionInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyCompositionInfo) GetInfo() *CompositionInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyCompositionList struct {
	Status               *ReplyStatus       `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*CompositionInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ReplyCompositionList) Reset()         { *m = ReplyCompositionList{} }
func (m *ReplyCompositionList) String() string { return proto.CompactTextString(m) }
func (*ReplyCompositionList) ProtoMessage()    {}
func (*ReplyCompositionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_e97c75056c972412, []int{7}
}

func (m *ReplyCompositionList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyCompositionList.Unmarshal(m, b)
}
func (m *ReplyCompositionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyCompositionList.Marshal(b, m, deterministic)
}
func (m *ReplyCompositionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyCompositionList.Merge(m, src)
}
func (m *ReplyCompositionList) XXX_Size() int {
	return xxx_messageInfo_ReplyCompositionList.Size(m)
}
func (m *ReplyCompositionList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyCompositionList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyCompositionList proto.InternalMessageInfo

func (m *ReplyCompositionList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyCompositionList) GetList() []*CompositionInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*CompositionInfo)(nil), "omo.msp.album.CompositionInfo")
	proto.RegisterType((*SlotInfo)(nil), "omo.msp.album.SlotInfo")
	proto.RegisterType((*ReqCompositionAdd)(nil), "omo.msp.album.ReqCompositionAdd")
	proto.RegisterType((*ReqCompositionUpdate)(nil), "omo.msp.album.ReqCompositionUpdate")
	proto.RegisterType((*ReqCompositionSlots)(nil), "omo.msp.album.ReqCompositionSlots")
	proto.RegisterType((*ReplyCompositionSlots)(nil), "omo.msp.album.ReplyCompositionSlots")
	proto.RegisterType((*ReplyCompositionInfo)(nil), "omo.msp.album.ReplyCompositionInfo")
	proto.RegisterType((*ReplyCompositionList)(nil), "omo.msp.album.ReplyCompositionList")
}

func init() {
	proto.RegisterFile("proto/album/composition.proto", fileDescriptor_e97c75056c972412)
}

var fileDescriptor_e97c75056c972412 = []byte{
	// 672 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0x6d, 0xd2, 0x24, 0x6b, 0x6f, 0x97, 0xdf, 0x0f, 0xcc, 0x06, 0x51, 0xc5, 0x50, 0x95, 0xf1,
	0x50, 0x09, 0xd1, 0x49, 0xe5, 0x13, 0x6c, 0x43, 0x4c, 0x68, 0x4c, 0x80, 0xab, 0xbd, 0xf0, 0x96,
	0x35, 0x1e, 0xb5, 0x68, 0xe2, 0x2c, 0x76, 0xb7, 0xf5, 0x65, 0xcf, 0xf0, 0x45, 0xf8, 0x42, 0x7c,
	0x21, 0xe4, 0x9b, 0xb4, 0x4b, 0xb3, 0xf4, 0x8f, 0xc6, 0x9b, 0x8f, 0x7d, 0x7d, 0xce, 0xf1, 0xb9,
	0x37, 0x2d, 0xec, 0x25, 0xa9, 0x50, 0xe2, 0x20, 0x18, 0x5f, 0x4c, 0xa2, 0x83, 0xa1, 0x88, 0x12,
	0x21, 0xb9, 0xe2, 0x22, 0xee, 0xe1, 0x3e, 0x71, 0x45, 0x24, 0x7a, 0x91, 0x4c, 0x7a, 0x58, 0xd0,
	0xf6, 0x4a, 0xd5, 0xd1, 0xac, 0xd0, 0xff, 0x6d, 0xc2, 0xff, 0xc7, 0xf7, 0xd7, 0x3f, 0xc6, 0x97,
	0x82, 0x3c, 0x81, 0xfa, 0x84, 0x87, 0x9e, 0xd1, 0x31, 0xba, 0x4d, 0xaa, 0x97, 0xe4, 0x3f, 0x30,
	0x79, 0xe8, 0x99, 0x1d, 0xa3, 0x6b, 0x51, 0x93, 0x87, 0xc4, 0x83, 0xad, 0x61, 0xca, 0x02, 0xc5,
	0x42, 0xaf, 0xde, 0x31, 0xba, 0x75, 0x3a, 0x83, 0xfa, 0x64, 0x92, 0x84, 0x78, 0x62, 0x65, 0x27,
	0x39, 0x9c, 0xdf, 0x11, 0xa9, 0x67, 0x23, 0xf3, 0x0c, 0x92, 0x36, 0x34, 0x44, 0xc2, 0x52, 0x3c,
	0x72, 0xf0, 0x68, 0x8e, 0x09, 0x01, 0x2b, 0x0e, 0x22, 0xe6, 0x6d, 0xe1, 0x3e, 0xae, 0xc9, 0x73,
	0x70, 0x52, 0x16, 0x05, 0xe9, 0x0f, 0xaf, 0x81, 0xbb, 0x39, 0x22, 0x3b, 0x60, 0x8b, 0x9b, 0x98,
	0xa5, 0x5e, 0x13, 0xb7, 0x33, 0xa0, 0xab, 0x03, 0x99, 0xb0, 0xa1, 0xf2, 0xa0, 0x63, 0x74, 0x5d,
	0x9a, 0x23, 0xf2, 0x16, 0x6c, 0x39, 0x16, 0x4a, 0x7a, 0xbb, 0x9d, 0x7a, 0xb7, 0xd5, 0x7f, 0xd1,
	0x5b, 0x88, 0xac, 0x37, 0x18, 0x0b, 0xa5, 0xd3, 0xa0, 0x59, 0x95, 0x7f, 0x07, 0x8d, 0xd9, 0x96,
	0x16, 0xe2, 0x71, 0xc8, 0x6e, 0x31, 0x22, 0x97, 0x66, 0x80, 0x6c, 0x83, 0x71, 0x8b, 0x19, 0xd9,
	0xd4, 0x40, 0x34, 0xc5, 0x70, 0x6c, 0x6a, 0x4c, 0xf5, 0x8d, 0x1b, 0x1e, 0xaa, 0x11, 0x86, 0xe2,
	0xd2, 0x0c, 0x68, 0x6b, 0x23, 0xc6, 0xbf, 0x8f, 0x14, 0x26, 0xe2, 0xd2, 0x1c, 0xe9, 0x47, 0xab,
	0x69, 0xc2, 0x30, 0x0c, 0x97, 0xe2, 0xda, 0xff, 0x65, 0xc0, 0x53, 0xca, 0xae, 0x0a, 0xbd, 0x3a,
	0x0c, 0xc3, 0x79, 0x3c, 0x46, 0x65, 0x3c, 0xe6, 0x42, 0x3c, 0xc5, 0x98, 0xeb, 0xa5, 0x98, 0xef,
	0x43, 0xb2, 0x16, 0x42, 0x9a, 0x47, 0x6a, 0x17, 0x22, 0xf5, 0x7f, 0x1a, 0xb0, 0xb3, 0xe8, 0xe5,
	0x1c, 0x9b, 0x5c, 0x31, 0x39, 0x33, 0x83, 0x66, 0xa5, 0xc1, 0xfa, 0x52, 0x83, 0xd6, 0x52, 0x83,
	0x76, 0xd1, 0xa0, 0x9f, 0xc2, 0xb3, 0x45, 0x27, 0xba, 0x49, 0xb2, 0xc2, 0x48, 0x91, 0xdc, 0x2c,
	0x91, 0xcf, 0x47, 0xa1, 0xbe, 0xd1, 0x28, 0xdc, 0xc2, 0x2e, 0x65, 0xc9, 0x78, 0xfa, 0x40, 0xb5,
	0x0f, 0x8e, 0x54, 0x81, 0x9a, 0x48, 0x14, 0x6e, 0xf5, 0xdb, 0x25, 0x22, 0xbc, 0x35, 0xc0, 0x0a,
	0x9a, 0x57, 0x92, 0x37, 0x60, 0x8d, 0xb9, 0x54, 0x5e, 0x6b, 0xb5, 0x34, 0x16, 0xf9, 0x77, 0x3a,
	0xf7, 0x45, 0x65, 0x1c, 0xc8, 0xc7, 0x08, 0xf7, 0xc1, 0xe2, 0xf1, 0xa5, 0xc0, 0x30, 0x5a, 0xfd,
	0x57, 0xa5, 0x1b, 0x25, 0x05, 0x8a, 0xb5, 0x55, 0xfa, 0x9f, 0xb8, 0x54, 0x8f, 0xd5, 0xc7, 0x87,
	0x9b, 0xf8, 0xf0, 0xb5, 0xfa, 0xba, 0xb6, 0xff, 0xc7, 0x06, 0x52, 0x4c, 0x9d, 0xa5, 0xd7, 0x7c,
	0xc8, 0xc8, 0x57, 0x70, 0x0e, 0xc3, 0xf0, 0x73, 0xcc, 0x48, 0xe7, 0x81, 0x70, 0xe9, 0x8b, 0x69,
	0xef, 0x57, 0x59, 0x2b, 0xa9, 0xf9, 0x35, 0x72, 0x06, 0x90, 0xcd, 0xf4, 0x51, 0x20, 0x19, 0xd9,
	0x5f, 0x49, 0x9b, 0x15, 0xb6, 0xbd, 0x2a, 0xe6, 0x9c, 0xee, 0x18, 0x9a, 0x94, 0x45, 0xe2, 0x9a,
	0x69, 0x93, 0x0f, 0xd3, 0xb9, 0x9a, 0x30, 0x89, 0x7d, 0x5e, 0x49, 0x72, 0x0a, 0xce, 0x09, 0x53,
	0xeb, 0x18, 0x36, 0x7c, 0xe0, 0x29, 0x38, 0x03, 0x16, 0xa4, 0xc3, 0xd1, 0x3f, 0x91, 0xe9, 0xee,
	0xfb, 0x35, 0xf2, 0x05, 0x9a, 0x27, 0x4c, 0x69, 0x70, 0x34, 0x25, 0x2f, 0xab, 0xf9, 0x3e, 0xf0,
	0xb1, 0x62, 0xe9, 0xa6, 0x8c, 0x67, 0xb0, 0x7d, 0xc2, 0x94, 0x1e, 0x19, 0x2e, 0x15, 0x1f, 0xae,
	0x21, 0xdd, 0x5b, 0x36, 0x6f, 0x78, 0xd9, 0xaf, 0x91, 0xf7, 0xd0, 0xc8, 0xdb, 0xb9, 0xd4, 0xdf,
	0x06, 0x5d, 0x3c, 0x87, 0xad, 0x01, 0x53, 0xfa, 0x9b, 0x24, 0xfe, 0xca, 0x89, 0xc0, 0x9f, 0x83,
	0xf6, 0xeb, 0x35, 0x4f, 0xc5, 0x2a, 0xbf, 0x76, 0xe4, 0x7e, 0x6b, 0x15, 0xfe, 0x9f, 0x2f, 0x1c,
	0x04, 0xef, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x95, 0x59, 0xb9, 0x1a, 0xe3, 0x07, 0x00, 0x00,
}
