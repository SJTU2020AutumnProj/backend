// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/answer/answer.proto

//服务名

package answer

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

// Api Endpoints for AnswerService service

func NewAnswerServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AnswerService service

type AnswerService interface {
	CreateAnswer(ctx context.Context, in *CreateAnswerParam, opts ...client.CallOption) (*CreateAnswerResponse, error)
	DeleteAnswer(ctx context.Context, in *AnswerID, opts ...client.CallOption) (*DeleteAnswerResponse, error)
	UpdateAnswer(ctx context.Context, in *AnswerInfo, opts ...client.CallOption) (*UpdateAnswerResponse, error)
	SearchAnswer(ctx context.Context, in *AnswerID, opts ...client.CallOption) (*SearchAnswerResponse, error)
	SearchAnswerByStudentID(ctx context.Context, in *StudentID, opts ...client.CallOption) (*SearchAnswerByStudentIDResponse, error)
	SearchAnswerByHomeworkID(ctx context.Context, in *HomeworkID, opts ...client.CallOption) (*SearchAnswerByHomeworkIDResponse, error)
	SearchAnswerByStudentIDAndHomeworkID(ctx context.Context, in *StudentIDAndHomeworkID, opts ...client.CallOption) (*SearchAnswerByStudentIDAndHomeworkIDResponse, error)
}

type answerService struct {
	c    client.Client
	name string
}

func NewAnswerService(name string, c client.Client) AnswerService {
	return &answerService{
		c:    c,
		name: name,
	}
}

func (c *answerService) CreateAnswer(ctx context.Context, in *CreateAnswerParam, opts ...client.CallOption) (*CreateAnswerResponse, error) {
	req := c.c.NewRequest(c.name, "AnswerService.CreateAnswer", in)
	out := new(CreateAnswerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerService) DeleteAnswer(ctx context.Context, in *AnswerID, opts ...client.CallOption) (*DeleteAnswerResponse, error) {
	req := c.c.NewRequest(c.name, "AnswerService.DeleteAnswer", in)
	out := new(DeleteAnswerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerService) UpdateAnswer(ctx context.Context, in *AnswerInfo, opts ...client.CallOption) (*UpdateAnswerResponse, error) {
	req := c.c.NewRequest(c.name, "AnswerService.UpdateAnswer", in)
	out := new(UpdateAnswerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerService) SearchAnswer(ctx context.Context, in *AnswerID, opts ...client.CallOption) (*SearchAnswerResponse, error) {
	req := c.c.NewRequest(c.name, "AnswerService.SearchAnswer", in)
	out := new(SearchAnswerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerService) SearchAnswerByStudentID(ctx context.Context, in *StudentID, opts ...client.CallOption) (*SearchAnswerByStudentIDResponse, error) {
	req := c.c.NewRequest(c.name, "AnswerService.SearchAnswerByStudentID", in)
	out := new(SearchAnswerByStudentIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerService) SearchAnswerByHomeworkID(ctx context.Context, in *HomeworkID, opts ...client.CallOption) (*SearchAnswerByHomeworkIDResponse, error) {
	req := c.c.NewRequest(c.name, "AnswerService.SearchAnswerByHomeworkID", in)
	out := new(SearchAnswerByHomeworkIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerService) SearchAnswerByStudentIDAndHomeworkID(ctx context.Context, in *StudentIDAndHomeworkID, opts ...client.CallOption) (*SearchAnswerByStudentIDAndHomeworkIDResponse, error) {
	req := c.c.NewRequest(c.name, "AnswerService.SearchAnswerByStudentIDAndHomeworkID", in)
	out := new(SearchAnswerByStudentIDAndHomeworkIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AnswerService service

type AnswerServiceHandler interface {
	CreateAnswer(context.Context, *CreateAnswerParam, *CreateAnswerResponse) error
	DeleteAnswer(context.Context, *AnswerID, *DeleteAnswerResponse) error
	UpdateAnswer(context.Context, *AnswerInfo, *UpdateAnswerResponse) error
	SearchAnswer(context.Context, *AnswerID, *SearchAnswerResponse) error
	SearchAnswerByStudentID(context.Context, *StudentID, *SearchAnswerByStudentIDResponse) error
	SearchAnswerByHomeworkID(context.Context, *HomeworkID, *SearchAnswerByHomeworkIDResponse) error
	SearchAnswerByStudentIDAndHomeworkID(context.Context, *StudentIDAndHomeworkID, *SearchAnswerByStudentIDAndHomeworkIDResponse) error
}

func RegisterAnswerServiceHandler(s server.Server, hdlr AnswerServiceHandler, opts ...server.HandlerOption) error {
	type answerService interface {
		CreateAnswer(ctx context.Context, in *CreateAnswerParam, out *CreateAnswerResponse) error
		DeleteAnswer(ctx context.Context, in *AnswerID, out *DeleteAnswerResponse) error
		UpdateAnswer(ctx context.Context, in *AnswerInfo, out *UpdateAnswerResponse) error
		SearchAnswer(ctx context.Context, in *AnswerID, out *SearchAnswerResponse) error
		SearchAnswerByStudentID(ctx context.Context, in *StudentID, out *SearchAnswerByStudentIDResponse) error
		SearchAnswerByHomeworkID(ctx context.Context, in *HomeworkID, out *SearchAnswerByHomeworkIDResponse) error
		SearchAnswerByStudentIDAndHomeworkID(ctx context.Context, in *StudentIDAndHomeworkID, out *SearchAnswerByStudentIDAndHomeworkIDResponse) error
	}
	type AnswerService struct {
		answerService
	}
	h := &answerServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AnswerService{h}, opts...))
}

type answerServiceHandler struct {
	AnswerServiceHandler
}

func (h *answerServiceHandler) CreateAnswer(ctx context.Context, in *CreateAnswerParam, out *CreateAnswerResponse) error {
	return h.AnswerServiceHandler.CreateAnswer(ctx, in, out)
}

func (h *answerServiceHandler) DeleteAnswer(ctx context.Context, in *AnswerID, out *DeleteAnswerResponse) error {
	return h.AnswerServiceHandler.DeleteAnswer(ctx, in, out)
}

func (h *answerServiceHandler) UpdateAnswer(ctx context.Context, in *AnswerInfo, out *UpdateAnswerResponse) error {
	return h.AnswerServiceHandler.UpdateAnswer(ctx, in, out)
}

func (h *answerServiceHandler) SearchAnswer(ctx context.Context, in *AnswerID, out *SearchAnswerResponse) error {
	return h.AnswerServiceHandler.SearchAnswer(ctx, in, out)
}

func (h *answerServiceHandler) SearchAnswerByStudentID(ctx context.Context, in *StudentID, out *SearchAnswerByStudentIDResponse) error {
	return h.AnswerServiceHandler.SearchAnswerByStudentID(ctx, in, out)
}

func (h *answerServiceHandler) SearchAnswerByHomeworkID(ctx context.Context, in *HomeworkID, out *SearchAnswerByHomeworkIDResponse) error {
	return h.AnswerServiceHandler.SearchAnswerByHomeworkID(ctx, in, out)
}

func (h *answerServiceHandler) SearchAnswerByStudentIDAndHomeworkID(ctx context.Context, in *StudentIDAndHomeworkID, out *SearchAnswerByStudentIDAndHomeworkIDResponse) error {
	return h.AnswerServiceHandler.SearchAnswerByStudentIDAndHomeworkID(ctx, in, out)
}