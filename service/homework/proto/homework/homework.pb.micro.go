// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/homework/homework.proto

//服务名

package homework

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

// Api Endpoints for HomeworkService service

func NewHomeworkServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for HomeworkService service

type HomeworkService interface {
	AssignHomework(ctx context.Context, in *AssignHomeworkParam, opts ...client.CallOption) (*AssignHomeworkResponse, error)
	DeleteHomework(ctx context.Context, in *HomeworkID, opts ...client.CallOption) (*DeleteHomeworkResponse, error)
	UpdateHomework(ctx context.Context, in *HomeworkInfo, opts ...client.CallOption) (*UpdateHomeworkResponse, error)
	SearchHomework(ctx context.Context, in *HomeworkID, opts ...client.CallOption) (*SearchHomeworkResponse, error)
	SearchHomeworkByUserID(ctx context.Context, in *UserID, opts ...client.CallOption) (*SearchHomeworkByUserIDResponse, error)
}

type homeworkService struct {
	c    client.Client
	name string
}

func NewHomeworkService(name string, c client.Client) HomeworkService {
	return &homeworkService{
		c:    c,
		name: name,
	}
}

func (c *homeworkService) AssignHomework(ctx context.Context, in *AssignHomeworkParam, opts ...client.CallOption) (*AssignHomeworkResponse, error) {
	req := c.c.NewRequest(c.name, "HomeworkService.AssignHomework", in)
	out := new(AssignHomeworkResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeworkService) DeleteHomework(ctx context.Context, in *HomeworkID, opts ...client.CallOption) (*DeleteHomeworkResponse, error) {
	req := c.c.NewRequest(c.name, "HomeworkService.DeleteHomework", in)
	out := new(DeleteHomeworkResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeworkService) UpdateHomework(ctx context.Context, in *HomeworkInfo, opts ...client.CallOption) (*UpdateHomeworkResponse, error) {
	req := c.c.NewRequest(c.name, "HomeworkService.UpdateHomework", in)
	out := new(UpdateHomeworkResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeworkService) SearchHomework(ctx context.Context, in *HomeworkID, opts ...client.CallOption) (*SearchHomeworkResponse, error) {
	req := c.c.NewRequest(c.name, "HomeworkService.SearchHomework", in)
	out := new(SearchHomeworkResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeworkService) SearchHomeworkByUserID(ctx context.Context, in *UserID, opts ...client.CallOption) (*SearchHomeworkByUserIDResponse, error) {
	req := c.c.NewRequest(c.name, "HomeworkService.SearchHomeworkByUserID", in)
	out := new(SearchHomeworkByUserIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HomeworkService service

type HomeworkServiceHandler interface {
	AssignHomework(context.Context, *AssignHomeworkParam, *AssignHomeworkResponse) error
	DeleteHomework(context.Context, *HomeworkID, *DeleteHomeworkResponse) error
	UpdateHomework(context.Context, *HomeworkInfo, *UpdateHomeworkResponse) error
	SearchHomework(context.Context, *HomeworkID, *SearchHomeworkResponse) error
	SearchHomeworkByUserID(context.Context, *UserID, *SearchHomeworkByUserIDResponse) error
}

func RegisterHomeworkServiceHandler(s server.Server, hdlr HomeworkServiceHandler, opts ...server.HandlerOption) error {
	type homeworkService interface {
		AssignHomework(ctx context.Context, in *AssignHomeworkParam, out *AssignHomeworkResponse) error
		DeleteHomework(ctx context.Context, in *HomeworkID, out *DeleteHomeworkResponse) error
		UpdateHomework(ctx context.Context, in *HomeworkInfo, out *UpdateHomeworkResponse) error
		SearchHomework(ctx context.Context, in *HomeworkID, out *SearchHomeworkResponse) error
		SearchHomeworkByUserID(ctx context.Context, in *UserID, out *SearchHomeworkByUserIDResponse) error
	}
	type HomeworkService struct {
		homeworkService
	}
	h := &homeworkServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&HomeworkService{h}, opts...))
}

type homeworkServiceHandler struct {
	HomeworkServiceHandler
}

func (h *homeworkServiceHandler) AssignHomework(ctx context.Context, in *AssignHomeworkParam, out *AssignHomeworkResponse) error {
	return h.HomeworkServiceHandler.AssignHomework(ctx, in, out)
}

func (h *homeworkServiceHandler) DeleteHomework(ctx context.Context, in *HomeworkID, out *DeleteHomeworkResponse) error {
	return h.HomeworkServiceHandler.DeleteHomework(ctx, in, out)
}

func (h *homeworkServiceHandler) UpdateHomework(ctx context.Context, in *HomeworkInfo, out *UpdateHomeworkResponse) error {
	return h.HomeworkServiceHandler.UpdateHomework(ctx, in, out)
}

func (h *homeworkServiceHandler) SearchHomework(ctx context.Context, in *HomeworkID, out *SearchHomeworkResponse) error {
	return h.HomeworkServiceHandler.SearchHomework(ctx, in, out)
}

func (h *homeworkServiceHandler) SearchHomeworkByUserID(ctx context.Context, in *UserID, out *SearchHomeworkByUserIDResponse) error {
	return h.HomeworkServiceHandler.SearchHomeworkByUserID(ctx, in, out)
}
