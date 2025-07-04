// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/album/common.proto

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

type PairInt struct {
	Key                  uint32   `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PairInt) Reset()         { *m = PairInt{} }
func (m *PairInt) String() string { return proto.CompactTextString(m) }
func (*PairInt) ProtoMessage()    {}
func (*PairInt) Descriptor() ([]byte, []int) {
	return fileDescriptor_24d55da2c5adfb30, []int{0}
}

func (m *PairInt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PairInt.Unmarshal(m, b)
}
func (m *PairInt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PairInt.Marshal(b, m, deterministic)
}
func (m *PairInt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairInt.Merge(m, src)
}
func (m *PairInt) XXX_Size() int {
	return xxx_messageInfo_PairInt.Size(m)
}
func (m *PairInt) XXX_DiscardUnknown() {
	xxx_messageInfo_PairInt.DiscardUnknown(m)
}

var xxx_messageInfo_PairInt proto.InternalMessageInfo

func (m *PairInt) GetKey() uint32 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *PairInt) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type MediaInfo struct {
	Cover                string   `protobuf:"bytes,1,opt,name=cover,proto3" json:"cover,omitempty"`
	Asset                string   `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	Type                 uint32   `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MediaInfo) Reset()         { *m = MediaInfo{} }
func (m *MediaInfo) String() string { return proto.CompactTextString(m) }
func (*MediaInfo) ProtoMessage()    {}
func (*MediaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_24d55da2c5adfb30, []int{1}
}

func (m *MediaInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaInfo.Unmarshal(m, b)
}
func (m *MediaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaInfo.Marshal(b, m, deterministic)
}
func (m *MediaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaInfo.Merge(m, src)
}
func (m *MediaInfo) XXX_Size() int {
	return xxx_messageInfo_MediaInfo.Size(m)
}
func (m *MediaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MediaInfo proto.InternalMessageInfo

func (m *MediaInfo) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *MediaInfo) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *MediaInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type ContactInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Remark               string   `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactInfo) Reset()         { *m = ContactInfo{} }
func (m *ContactInfo) String() string { return proto.CompactTextString(m) }
func (*ContactInfo) ProtoMessage()    {}
func (*ContactInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_24d55da2c5adfb30, []int{2}
}

func (m *ContactInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactInfo.Unmarshal(m, b)
}
func (m *ContactInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactInfo.Marshal(b, m, deterministic)
}
func (m *ContactInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactInfo.Merge(m, src)
}
func (m *ContactInfo) XXX_Size() int {
	return xxx_messageInfo_ContactInfo.Size(m)
}
func (m *ContactInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ContactInfo proto.InternalMessageInfo

func (m *ContactInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContactInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *ContactInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ContactInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

type DurationInfo struct {
	Begin                int64    `protobuf:"varint,1,opt,name=begin,proto3" json:"begin,omitempty"`
	End                  int64    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DurationInfo) Reset()         { *m = DurationInfo{} }
func (m *DurationInfo) String() string { return proto.CompactTextString(m) }
func (*DurationInfo) ProtoMessage()    {}
func (*DurationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_24d55da2c5adfb30, []int{3}
}

func (m *DurationInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DurationInfo.Unmarshal(m, b)
}
func (m *DurationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DurationInfo.Marshal(b, m, deterministic)
}
func (m *DurationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DurationInfo.Merge(m, src)
}
func (m *DurationInfo) XXX_Size() int {
	return xxx_messageInfo_DurationInfo.Size(m)
}
func (m *DurationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DurationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DurationInfo proto.InternalMessageInfo

func (m *DurationInfo) GetBegin() int64 {
	if m != nil {
		return m.Begin
	}
	return 0
}

func (m *DurationInfo) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

type RequestInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Flag                 string   `protobuf:"bytes,2,opt,name=flag,proto3" json:"flag,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	User                 string   `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestInfo) Reset()         { *m = RequestInfo{} }
func (m *RequestInfo) String() string { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()    {}
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_24d55da2c5adfb30, []int{4}
}

func (m *RequestInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestInfo.Unmarshal(m, b)
}
func (m *RequestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestInfo.Marshal(b, m, deterministic)
}
func (m *RequestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInfo.Merge(m, src)
}
func (m *RequestInfo) XXX_Size() int {
	return xxx_messageInfo_RequestInfo.Size(m)
}
func (m *RequestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInfo proto.InternalMessageInfo

func (m *RequestInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestInfo) GetFlag() string {
	if m != nil {
		return m.Flag
	}
	return ""
}

func (m *RequestInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RequestInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RequestInfo) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type RequestFilter struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Field                string   `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Page                 uint32   `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32   `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Values               []string `protobuf:"bytes,7,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestFilter) Reset()         { *m = RequestFilter{} }
func (m *RequestFilter) String() string { return proto.CompactTextString(m) }
func (*RequestFilter) ProtoMessage()    {}
func (*RequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_24d55da2c5adfb30, []int{5}
}

func (m *RequestFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestFilter.Unmarshal(m, b)
}
func (m *RequestFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestFilter.Marshal(b, m, deterministic)
}
func (m *RequestFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestFilter.Merge(m, src)
}
func (m *RequestFilter) XXX_Size() int {
	return xxx_messageInfo_RequestFilter.Size(m)
}
func (m *RequestFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestFilter.DiscardUnknown(m)
}

var xxx_messageInfo_RequestFilter proto.InternalMessageInfo

func (m *RequestFilter) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RequestFilter) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *RequestFilter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RequestFilter) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *RequestFilter) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *RequestFilter) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RequestFilter) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type RequestUpdate struct {
	Owner                string       `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Uid                  string       `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Field                string       `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
	Value                string       `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Operator             string       `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Values               []string     `protobuf:"bytes,6,rep,name=values,proto3" json:"values,omitempty"`
	List                 []*MediaInfo `protobuf:"bytes,7,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RequestUpdate) Reset()         { *m = RequestUpdate{} }
func (m *RequestUpdate) String() string { return proto.CompactTextString(m) }
func (*RequestUpdate) ProtoMessage()    {}
func (*RequestUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_24d55da2c5adfb30, []int{6}
}

func (m *RequestUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestUpdate.Unmarshal(m, b)
}
func (m *RequestUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestUpdate.Marshal(b, m, deterministic)
}
func (m *RequestUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestUpdate.Merge(m, src)
}
func (m *RequestUpdate) XXX_Size() int {
	return xxx_messageInfo_RequestUpdate.Size(m)
}
func (m *RequestUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_RequestUpdate proto.InternalMessageInfo

func (m *RequestUpdate) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RequestUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestUpdate) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *RequestUpdate) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RequestUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RequestUpdate) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *RequestUpdate) GetList() []*MediaInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyStatus struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyStatus) Reset()         { *m = ReplyStatus{} }
func (m *ReplyStatus) String() string { return proto.CompactTextString(m) }
func (*ReplyStatus) ProtoMessage()    {}
func (*ReplyStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_24d55da2c5adfb30, []int{7}
}

func (m *ReplyStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStatus.Unmarshal(m, b)
}
func (m *ReplyStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStatus.Marshal(b, m, deterministic)
}
func (m *ReplyStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStatus.Merge(m, src)
}
func (m *ReplyStatus) XXX_Size() int {
	return xxx_messageInfo_ReplyStatus.Size(m)
}
func (m *ReplyStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStatus proto.InternalMessageInfo

func (m *ReplyStatus) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ReplyStatus) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ReplyInfo struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyInfo) Reset()         { *m = ReplyInfo{} }
func (m *ReplyInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyInfo) ProtoMessage()    {}
func (*ReplyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_24d55da2c5adfb30, []int{8}
}

func (m *ReplyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyInfo.Unmarshal(m, b)
}
func (m *ReplyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyInfo.Marshal(b, m, deterministic)
}
func (m *ReplyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyInfo.Merge(m, src)
}
func (m *ReplyInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyInfo.Size(m)
}
func (m *ReplyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyInfo proto.InternalMessageInfo

func (m *ReplyInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReplyInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type PhotocopySlot struct {
	Page                 uint32   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Index                uint32   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Role                 uint32   `protobuf:"varint,3,opt,name=role,proto3" json:"role,omitempty"`
	Background           string   `protobuf:"bytes,4,opt,name=background,proto3" json:"background,omitempty"`
	Asset                string   `protobuf:"bytes,5,opt,name=asset,proto3" json:"asset,omitempty"`
	Remark               string   `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhotocopySlot) Reset()         { *m = PhotocopySlot{} }
func (m *PhotocopySlot) String() string { return proto.CompactTextString(m) }
func (*PhotocopySlot) ProtoMessage()    {}
func (*PhotocopySlot) Descriptor() ([]byte, []int) {
	return fileDescriptor_24d55da2c5adfb30, []int{9}
}

func (m *PhotocopySlot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhotocopySlot.Unmarshal(m, b)
}
func (m *PhotocopySlot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhotocopySlot.Marshal(b, m, deterministic)
}
func (m *PhotocopySlot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhotocopySlot.Merge(m, src)
}
func (m *PhotocopySlot) XXX_Size() int {
	return xxx_messageInfo_PhotocopySlot.Size(m)
}
func (m *PhotocopySlot) XXX_DiscardUnknown() {
	xxx_messageInfo_PhotocopySlot.DiscardUnknown(m)
}

var xxx_messageInfo_PhotocopySlot proto.InternalMessageInfo

func (m *PhotocopySlot) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *PhotocopySlot) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PhotocopySlot) GetRole() uint32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *PhotocopySlot) GetBackground() string {
	if m != nil {
		return m.Background
	}
	return ""
}

func (m *PhotocopySlot) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *PhotocopySlot) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

type RequestIntFlag struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Flag                 int32    `protobuf:"varint,2,opt,name=flag,proto3" json:"flag,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestIntFlag) Reset()         { *m = RequestIntFlag{} }
func (m *RequestIntFlag) String() string { return proto.CompactTextString(m) }
func (*RequestIntFlag) ProtoMessage()    {}
func (*RequestIntFlag) Descriptor() ([]byte, []int) {
	return fileDescriptor_24d55da2c5adfb30, []int{10}
}

func (m *RequestIntFlag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestIntFlag.Unmarshal(m, b)
}
func (m *RequestIntFlag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestIntFlag.Marshal(b, m, deterministic)
}
func (m *RequestIntFlag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestIntFlag.Merge(m, src)
}
func (m *RequestIntFlag) XXX_Size() int {
	return xxx_messageInfo_RequestIntFlag.Size(m)
}
func (m *RequestIntFlag) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestIntFlag.DiscardUnknown(m)
}

var xxx_messageInfo_RequestIntFlag proto.InternalMessageInfo

func (m *RequestIntFlag) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestIntFlag) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *RequestIntFlag) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type RequestList struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string       `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	List                 []string     `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	Medias               []*MediaInfo `protobuf:"bytes,4,rep,name=medias,proto3" json:"medias,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RequestList) Reset()         { *m = RequestList{} }
func (m *RequestList) String() string { return proto.CompactTextString(m) }
func (*RequestList) ProtoMessage()    {}
func (*RequestList) Descriptor() ([]byte, []int) {
	return fileDescriptor_24d55da2c5adfb30, []int{11}
}

func (m *RequestList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestList.Unmarshal(m, b)
}
func (m *RequestList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestList.Marshal(b, m, deterministic)
}
func (m *RequestList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestList.Merge(m, src)
}
func (m *RequestList) XXX_Size() int {
	return xxx_messageInfo_RequestList.Size(m)
}
func (m *RequestList) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestList.DiscardUnknown(m)
}

var xxx_messageInfo_RequestList proto.InternalMessageInfo

func (m *RequestList) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestList) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RequestList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *RequestList) GetMedias() []*MediaInfo {
	if m != nil {
		return m.Medias
	}
	return nil
}

type ReplyList struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	List                 []string     `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Medias               []*MediaInfo `protobuf:"bytes,5,rep,name=medias,proto3" json:"medias,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyList) Reset()         { *m = ReplyList{} }
func (m *ReplyList) String() string { return proto.CompactTextString(m) }
func (*ReplyList) ProtoMessage()    {}
func (*ReplyList) Descriptor() ([]byte, []int) {
	return fileDescriptor_24d55da2c5adfb30, []int{12}
}

func (m *ReplyList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyList.Unmarshal(m, b)
}
func (m *ReplyList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyList.Marshal(b, m, deterministic)
}
func (m *ReplyList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyList.Merge(m, src)
}
func (m *ReplyList) XXX_Size() int {
	return xxx_messageInfo_ReplyList.Size(m)
}
func (m *ReplyList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyList proto.InternalMessageInfo

func (m *ReplyList) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReplyList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyList) GetMedias() []*MediaInfo {
	if m != nil {
		return m.Medias
	}
	return nil
}

type ReplyStatistic struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Key                  string       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Owner                string       `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Count                uint32       `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyStatistic) Reset()         { *m = ReplyStatistic{} }
func (m *ReplyStatistic) String() string { return proto.CompactTextString(m) }
func (*ReplyStatistic) ProtoMessage()    {}
func (*ReplyStatistic) Descriptor() ([]byte, []int) {
	return fileDescriptor_24d55da2c5adfb30, []int{13}
}

func (m *ReplyStatistic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStatistic.Unmarshal(m, b)
}
func (m *ReplyStatistic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStatistic.Marshal(b, m, deterministic)
}
func (m *ReplyStatistic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStatistic.Merge(m, src)
}
func (m *ReplyStatistic) XXX_Size() int {
	return xxx_messageInfo_ReplyStatistic.Size(m)
}
func (m *ReplyStatistic) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStatistic.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStatistic proto.InternalMessageInfo

func (m *ReplyStatistic) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyStatistic) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReplyStatistic) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReplyStatistic) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*PairInt)(nil), "omo.msp.album.PairInt")
	proto.RegisterType((*MediaInfo)(nil), "omo.msp.album.MediaInfo")
	proto.RegisterType((*ContactInfo)(nil), "omo.msp.album.ContactInfo")
	proto.RegisterType((*DurationInfo)(nil), "omo.msp.album.DurationInfo")
	proto.RegisterType((*RequestInfo)(nil), "omo.msp.album.RequestInfo")
	proto.RegisterType((*RequestFilter)(nil), "omo.msp.album.RequestFilter")
	proto.RegisterType((*RequestUpdate)(nil), "omo.msp.album.RequestUpdate")
	proto.RegisterType((*ReplyStatus)(nil), "omo.msp.album.ReplyStatus")
	proto.RegisterType((*ReplyInfo)(nil), "omo.msp.album.ReplyInfo")
	proto.RegisterType((*PhotocopySlot)(nil), "omo.msp.album.PhotocopySlot")
	proto.RegisterType((*RequestIntFlag)(nil), "omo.msp.album.RequestIntFlag")
	proto.RegisterType((*RequestList)(nil), "omo.msp.album.RequestList")
	proto.RegisterType((*ReplyList)(nil), "omo.msp.album.ReplyList")
	proto.RegisterType((*ReplyStatistic)(nil), "omo.msp.album.ReplyStatistic")
}

func init() {
	proto.RegisterFile("proto/album/common.proto", fileDescriptor_24d55da2c5adfb30)
}

var fileDescriptor_24d55da2c5adfb30 = []byte{
	// 658 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0x56, 0x36, 0x3f, 0x65, 0x67, 0x09, 0xaa, 0xa2, 0xaa, 0x8a, 0x7a, 0x40, 0x55, 0x4e, 0x3d,
	0xa0, 0x2d, 0x14, 0x09, 0xee, 0x80, 0x2a, 0x55, 0x80, 0x54, 0x5c, 0x71, 0xe1, 0xe6, 0x4d, 0xdc,
	0xad, 0xd5, 0xc4, 0x0e, 0xb6, 0xd3, 0xb2, 0x47, 0x78, 0x09, 0x0e, 0x3c, 0x05, 0x0f, 0xc1, 0x7b,
	0x21, 0x4f, 0xbc, 0x89, 0x5b, 0xba, 0xb4, 0xb7, 0xf9, 0x1c, 0x7b, 0xe6, 0x9b, 0xf9, 0xfc, 0x39,
	0x90, 0xb7, 0x4a, 0x1a, 0x79, 0x48, 0xeb, 0x45, 0xd7, 0x1c, 0x96, 0xb2, 0x69, 0xa4, 0x98, 0xe3,
	0x52, 0x96, 0xca, 0x46, 0xce, 0x1b, 0xdd, 0xce, 0xf1, 0x5b, 0xf1, 0x02, 0xb6, 0x4e, 0x29, 0x57,
	0x27, 0xc2, 0x64, 0xdb, 0x10, 0x5e, 0xb2, 0x55, 0x1e, 0xec, 0x07, 0x07, 0x29, 0xb1, 0x61, 0xb6,
	0x03, 0xf1, 0x15, 0xad, 0x3b, 0x96, 0x4f, 0xf6, 0x83, 0x83, 0x29, 0xe9, 0x41, 0xf1, 0x1e, 0xa6,
	0x1f, 0x59, 0xc5, 0xe9, 0x89, 0x38, 0x97, 0x76, 0x4b, 0x29, 0xaf, 0x98, 0xc2, 0x63, 0x53, 0xd2,
	0x03, 0xbb, 0x4a, 0xb5, 0x66, 0x66, 0x7d, 0x10, 0x41, 0x96, 0x41, 0x64, 0x56, 0x2d, 0xcb, 0x43,
	0xac, 0x80, 0x71, 0xc1, 0x61, 0xf6, 0x56, 0x0a, 0x43, 0x4b, 0x83, 0xe9, 0x32, 0x88, 0x04, 0x6d,
	0x98, 0xcb, 0x86, 0xb1, 0x4d, 0xd6, 0x5e, 0x48, 0x31, 0xb0, 0x40, 0x90, 0xe5, 0xb0, 0x45, 0xab,
	0x4a, 0x31, 0xad, 0x31, 0xdf, 0x94, 0xac, 0x61, 0xb6, 0x0b, 0x89, 0x62, 0x0d, 0x55, 0x97, 0x79,
	0x84, 0x1f, 0x1c, 0x2a, 0x5e, 0xc1, 0xe3, 0x77, 0x9d, 0xa2, 0x86, 0x4b, 0xb1, 0xa6, 0xbe, 0x60,
	0x4b, 0x2e, 0xb0, 0x58, 0x48, 0x7a, 0x60, 0xa7, 0xc0, 0x44, 0x85, 0xb5, 0x42, 0x62, 0xc3, 0xe2,
	0x1a, 0x66, 0x84, 0x7d, 0xed, 0x98, 0xee, 0x29, 0x6e, 0x43, 0xd8, 0xf1, 0xca, 0x31, 0xb4, 0xa1,
	0x25, 0x7d, 0x5e, 0xd3, 0xa5, 0xe3, 0x87, 0x71, 0xb6, 0x07, 0x8f, 0x64, 0xcb, 0x14, 0x35, 0x52,
	0x39, 0x7e, 0x03, 0x1e, 0x9a, 0x8c, 0xbc, 0x26, 0x33, 0x88, 0x3a, 0xcd, 0x54, 0x1e, 0xf7, 0x6b,
	0x36, 0x2e, 0x7e, 0x07, 0x90, 0xba, 0xca, 0xc7, 0xbc, 0x36, 0xfd, 0x5c, 0xe5, 0xb5, 0x18, 0xa7,
	0x8d, 0xc0, 0xae, 0x9e, 0x73, 0x56, 0x57, 0xeb, 0x01, 0x21, 0x18, 0xc5, 0x0b, 0x3d, 0xf1, 0x6c,
	0x9d, 0x96, 0x2e, 0xfb, 0xda, 0x29, 0xc1, 0xd8, 0x0e, 0x4c, 0x74, 0xcd, 0xc2, 0x55, 0x4f, 0x89,
	0x43, 0x37, 0x7a, 0x48, 0x6e, 0xf5, 0xb0, 0x0b, 0x09, 0x26, 0xd4, 0xf9, 0xd6, 0x7e, 0x68, 0x87,
	0xdc, 0xa3, 0xe2, 0xcf, 0xc8, 0xf9, 0x73, 0x5b, 0x51, 0xc3, 0x36, 0x70, 0x76, 0x53, 0x9c, 0x8c,
	0x53, 0x1c, 0xba, 0x08, 0xef, 0xec, 0x22, 0xf2, 0xbb, 0xf0, 0x99, 0xc5, 0x1b, 0x99, 0x25, 0x3e,
	0xb3, 0xec, 0x19, 0x44, 0x35, 0xd7, 0x06, 0xf9, 0xce, 0x8e, 0xf2, 0xf9, 0x0d, 0x1f, 0xcc, 0x87,
	0x1b, 0x4d, 0x70, 0x57, 0xf1, 0xda, 0x8a, 0xde, 0xd6, 0xab, 0x33, 0x43, 0x4d, 0xa7, 0xed, 0xd8,
	0x4a, 0x59, 0x31, 0x67, 0x0e, 0x8c, 0x2d, 0x35, 0xa6, 0x94, 0x54, 0xeb, 0xb1, 0x23, 0x28, 0x3e,
	0xc1, 0x14, 0x0f, 0x6e, 0xb8, 0x2b, 0x47, 0x90, 0x68, 0x4c, 0x89, 0xa7, 0x66, 0x47, 0x7b, 0xb7,
	0x78, 0x78, 0x45, 0x89, 0xdb, 0x59, 0xfc, 0x0a, 0x20, 0x3d, 0xbd, 0x90, 0x46, 0x96, 0xb2, 0x5d,
	0x9d, 0xd5, 0xd2, 0x0c, 0x2a, 0x06, 0x9e, 0x8a, 0x3b, 0x10, 0x73, 0x51, 0xb1, 0x6f, 0x98, 0x38,
	0x25, 0x3d, 0xb0, 0x3b, 0x95, 0xac, 0x07, 0xcf, 0xd9, 0x38, 0x7b, 0x0a, 0xb0, 0xa0, 0xe5, 0xe5,
	0x52, 0xc9, 0x4e, 0x54, 0x6e, 0xb0, 0xde, 0xca, 0xe8, 0xde, 0xd8, 0x77, 0xef, 0x68, 0xab, 0xe4,
	0x86, 0xad, 0x08, 0x3c, 0x19, 0xec, 0x61, 0x8e, 0xed, 0xdd, 0xff, 0xbf, 0x43, 0xe2, 0xfb, 0x1d,
	0x52, 0x7c, 0x0f, 0x06, 0xcf, 0x7d, 0xe0, 0xda, 0xdc, 0x91, 0xd1, 0x3f, 0x3d, 0xf9, 0xd7, 0x5f,
	0xa8, 0x74, 0x88, 0xfa, 0x63, 0x9c, 0x3d, 0x87, 0xa4, 0xb1, 0x12, 0xeb, 0x3c, 0xba, 0x47, 0x7f,
	0xb7, 0xaf, 0xf8, 0x19, 0x38, 0x25, 0x37, 0x30, 0x58, 0x57, 0x99, 0x78, 0x55, 0x46, 0x75, 0xc3,
	0x87, 0xaa, 0xeb, 0x31, 0x8b, 0x1f, 0xc8, 0xec, 0x47, 0x60, 0x47, 0xee, 0x32, 0x71, 0x6d, 0x78,
	0xe9, 0x15, 0x0e, 0x1e, 0x5c, 0xd8, 0xbd, 0xf7, 0xce, 0x82, 0xee, 0xbd, 0xef, 0xad, 0x1a, 0xde,
	0x7a, 0x5e, 0x4a, 0xd9, 0x09, 0xe3, 0xde, 0x8c, 0x1e, 0xbc, 0x49, 0xbf, 0xcc, 0xbc, 0x7f, 0xcc,
	0x22, 0x41, 0xf0, 0xf2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0xa4, 0x25, 0x32, 0x79, 0x06,
	0x00, 0x00,
}
