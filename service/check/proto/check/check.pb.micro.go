// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/check/check.proto

//服务名

package check

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for CheckService service

func NewCheckServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CheckService service

type CheckService interface {
	CreateCheck(ctx context.Context, in *CreateCheckParam, opts ...client.CallOption) (*CreateCheckResponse, error)
	DeleteCheck(ctx context.Context, in *CheckID, opts ...client.CallOption) (*DeleteCheckResponse, error)
	UpdateCheck(ctx context.Context, in *UpdateCheckParam, opts ...client.CallOption) (*UpdateCheckResponse, error)
	SearchCheckByID(ctx context.Context, in *CheckID, opts ...client.CallOption) (*SearchCheckByIDResponse, error)
}

type checkService struct {
	c    client.Client
	name string
}

func NewCheckService(name string, c client.Client) CheckService {
	return &checkService{
		c:    c,
		name: name,
	}
}

func (c *checkService) CreateCheck(ctx context.Context, in *CreateCheckParam, opts ...client.CallOption) (*CreateCheckResponse, error) {
	req := c.c.NewRequest(c.name, "CheckService.CreateCheck", in)
	out := new(CreateCheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkService) DeleteCheck(ctx context.Context, in *CheckID, opts ...client.CallOption) (*DeleteCheckResponse, error) {
	req := c.c.NewRequest(c.name, "CheckService.DeleteCheck", in)
	out := new(DeleteCheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkService) UpdateCheck(ctx context.Context, in *UpdateCheckParam, opts ...client.CallOption) (*UpdateCheckResponse, error) {
	req := c.c.NewRequest(c.name, "CheckService.UpdateCheck", in)
	out := new(UpdateCheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkService) SearchCheckByID(ctx context.Context, in *CheckID, opts ...client.CallOption) (*SearchCheckByIDResponse, error) {
	req := c.c.NewRequest(c.name, "CheckService.SearchCheckByID", in)
	out := new(SearchCheckByIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CheckService service

type CheckServiceHandler interface {
	CreateCheck(context.Context, *CreateCheckParam, *CreateCheckResponse) error
	DeleteCheck(context.Context, *CheckID, *DeleteCheckResponse) error
	UpdateCheck(context.Context, *UpdateCheckParam, *UpdateCheckResponse) error
	SearchCheckByID(context.Context, *CheckID, *SearchCheckByIDResponse) error
}

func RegisterCheckServiceHandler(s server.Server, hdlr CheckServiceHandler, opts ...server.HandlerOption) error {
	type checkService interface {
		CreateCheck(ctx context.Context, in *CreateCheckParam, out *CreateCheckResponse) error
		DeleteCheck(ctx context.Context, in *CheckID, out *DeleteCheckResponse) error
		UpdateCheck(ctx context.Context, in *UpdateCheckParam, out *UpdateCheckResponse) error
		SearchCheckByID(ctx context.Context, in *CheckID, out *SearchCheckByIDResponse) error
	}
	type CheckService struct {
		checkService
	}
	h := &checkServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CheckService{h}, opts...))
}

type checkServiceHandler struct {
	CheckServiceHandler
}

func (h *checkServiceHandler) CreateCheck(ctx context.Context, in *CreateCheckParam, out *CreateCheckResponse) error {
	return h.CheckServiceHandler.CreateCheck(ctx, in, out)
}

func (h *checkServiceHandler) DeleteCheck(ctx context.Context, in *CheckID, out *DeleteCheckResponse) error {
	return h.CheckServiceHandler.DeleteCheck(ctx, in, out)
}

func (h *checkServiceHandler) UpdateCheck(ctx context.Context, in *UpdateCheckParam, out *UpdateCheckResponse) error {
	return h.CheckServiceHandler.UpdateCheck(ctx, in, out)
}

func (h *checkServiceHandler) SearchCheckByID(ctx context.Context, in *CheckID, out *SearchCheckByIDResponse) error {
	return h.CheckServiceHandler.SearchCheckByID(ctx, in, out)
}
