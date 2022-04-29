// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/album/photocopy.proto

package album

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for PhotocopyService service

type PhotocopyService interface {
	AddOne(ctx context.Context, in *ReqPhotocopyAdd, opts ...client.CallOption) (*ReplyPhotocopyInfo, error)
	UpdateBase(ctx context.Context, in *ReqPhotocopyUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyPhotocopyInfo, error)
	Search(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyPhotocopyList, error)
	GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyPhotocopyList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	AppendSlot(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyList, error)
	SubtractSlot(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyList, error)
}

type photocopyService struct {
	c    client.Client
	name string
}

func NewPhotocopyService(name string, c client.Client) PhotocopyService {
	return &photocopyService{
		c:    c,
		name: name,
	}
}

func (c *photocopyService) AddOne(ctx context.Context, in *ReqPhotocopyAdd, opts ...client.CallOption) (*ReplyPhotocopyInfo, error) {
	req := c.c.NewRequest(c.name, "PhotocopyService.AddOne", in)
	out := new(ReplyPhotocopyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photocopyService) UpdateBase(ctx context.Context, in *ReqPhotocopyUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "PhotocopyService.UpdateBase", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photocopyService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "PhotocopyService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photocopyService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyPhotocopyInfo, error) {
	req := c.c.NewRequest(c.name, "PhotocopyService.GetOne", in)
	out := new(ReplyPhotocopyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photocopyService) Search(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyPhotocopyList, error) {
	req := c.c.NewRequest(c.name, "PhotocopyService.Search", in)
	out := new(ReplyPhotocopyList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photocopyService) GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyPhotocopyList, error) {
	req := c.c.NewRequest(c.name, "PhotocopyService.GetListByFilter", in)
	out := new(ReplyPhotocopyList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photocopyService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "PhotocopyService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photocopyService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "PhotocopyService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photocopyService) AppendSlot(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyList, error) {
	req := c.c.NewRequest(c.name, "PhotocopyService.AppendSlot", in)
	out := new(ReplyList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photocopyService) SubtractSlot(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyList, error) {
	req := c.c.NewRequest(c.name, "PhotocopyService.SubtractSlot", in)
	out := new(ReplyList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PhotocopyService service

type PhotocopyServiceHandler interface {
	AddOne(context.Context, *ReqPhotocopyAdd, *ReplyPhotocopyInfo) error
	UpdateBase(context.Context, *ReqPhotocopyUpdate, *ReplyInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyPhotocopyInfo) error
	Search(context.Context, *RequestInfo, *ReplyPhotocopyList) error
	GetListByFilter(context.Context, *RequestFilter, *ReplyPhotocopyList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
	AppendSlot(context.Context, *RequestList, *ReplyList) error
	SubtractSlot(context.Context, *RequestList, *ReplyList) error
}

func RegisterPhotocopyServiceHandler(s server.Server, hdlr PhotocopyServiceHandler, opts ...server.HandlerOption) error {
	type photocopyService interface {
		AddOne(ctx context.Context, in *ReqPhotocopyAdd, out *ReplyPhotocopyInfo) error
		UpdateBase(ctx context.Context, in *ReqPhotocopyUpdate, out *ReplyInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyPhotocopyInfo) error
		Search(ctx context.Context, in *RequestInfo, out *ReplyPhotocopyList) error
		GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyPhotocopyList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
		AppendSlot(ctx context.Context, in *RequestList, out *ReplyList) error
		SubtractSlot(ctx context.Context, in *RequestList, out *ReplyList) error
	}
	type PhotocopyService struct {
		photocopyService
	}
	h := &photocopyServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PhotocopyService{h}, opts...))
}

type photocopyServiceHandler struct {
	PhotocopyServiceHandler
}

func (h *photocopyServiceHandler) AddOne(ctx context.Context, in *ReqPhotocopyAdd, out *ReplyPhotocopyInfo) error {
	return h.PhotocopyServiceHandler.AddOne(ctx, in, out)
}

func (h *photocopyServiceHandler) UpdateBase(ctx context.Context, in *ReqPhotocopyUpdate, out *ReplyInfo) error {
	return h.PhotocopyServiceHandler.UpdateBase(ctx, in, out)
}

func (h *photocopyServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.PhotocopyServiceHandler.RemoveOne(ctx, in, out)
}

func (h *photocopyServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyPhotocopyInfo) error {
	return h.PhotocopyServiceHandler.GetOne(ctx, in, out)
}

func (h *photocopyServiceHandler) Search(ctx context.Context, in *RequestInfo, out *ReplyPhotocopyList) error {
	return h.PhotocopyServiceHandler.Search(ctx, in, out)
}

func (h *photocopyServiceHandler) GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyPhotocopyList) error {
	return h.PhotocopyServiceHandler.GetListByFilter(ctx, in, out)
}

func (h *photocopyServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.PhotocopyServiceHandler.GetStatistic(ctx, in, out)
}

func (h *photocopyServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.PhotocopyServiceHandler.UpdateByFilter(ctx, in, out)
}

func (h *photocopyServiceHandler) AppendSlot(ctx context.Context, in *RequestList, out *ReplyList) error {
	return h.PhotocopyServiceHandler.AppendSlot(ctx, in, out)
}

func (h *photocopyServiceHandler) SubtractSlot(ctx context.Context, in *RequestList, out *ReplyList) error {
	return h.PhotocopyServiceHandler.SubtractSlot(ctx, in, out)
}
