// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: im.proto

package im

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
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

// Client API for Im service

type ImService interface {
	PublishMessage(ctx context.Context, in *PublishMessageRequest, opts ...client.CallOption) (*PublishMessageResponse, error)
}

type imService struct {
	c    client.Client
	name string
}

func NewImService(name string, c client.Client) ImService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "im"
	}
	return &imService{
		c:    c,
		name: name,
	}
}

func (c *imService) PublishMessage(ctx context.Context, in *PublishMessageRequest, opts ...client.CallOption) (*PublishMessageResponse, error) {
	req := c.c.NewRequest(c.name, "Im.PublishMessage", in)
	out := new(PublishMessageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Im service

type ImHandler interface {
	PublishMessage(context.Context, *PublishMessageRequest, *PublishMessageResponse) error
}

func RegisterImHandler(s server.Server, hdlr ImHandler, opts ...server.HandlerOption) error {
	type im interface {
		PublishMessage(ctx context.Context, in *PublishMessageRequest, out *PublishMessageResponse) error
	}
	type Im struct {
		im
	}
	h := &imHandler{hdlr}
	return s.Handle(s.NewHandler(&Im{h}, opts...))
}

type imHandler struct {
	ImHandler
}

func (h *imHandler) PublishMessage(ctx context.Context, in *PublishMessageRequest, out *PublishMessageResponse) error {
	return h.ImHandler.PublishMessage(ctx, in, out)
}
