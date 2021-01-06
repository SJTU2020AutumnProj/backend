// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/courseclass/courseclass.proto

//服务名

package courseclass

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

// Api Endpoints for CourseClassService service

func NewCourseClassServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CourseClassService service

type CourseClassService interface {
	//rpc AddCourseClass (CourseClass) returns (EditResponse) {}
	DeleteCourseClass(ctx context.Context, in *CourseID, opts ...client.CallOption) (*EditResponse, error)
	UpdateCourseClass(ctx context.Context, in *CourseClass, opts ...client.CallOption) (*EditResponse, error)
	SearchCourseClass(ctx context.Context, in *CourseID, opts ...client.CallOption) (*SearchCourseClassResponse, error)
	SearchCourseClasses(ctx context.Context, in *CourseIDArray, opts ...client.CallOption) (*SearchCourseClassesResponse, error)
	NewCourse(ctx context.Context, in *NewCourseMessage, opts ...client.CallOption) (*NewCourseResponse, error)
	AddTake(ctx context.Context, in *Take, opts ...client.CallOption) (*EditResponse, error)
	DeleteTake(ctx context.Context, in *UserCourse, opts ...client.CallOption) (*EditResponse, error)
	DeleteTakeByUser(ctx context.Context, in *UserID, opts ...client.CallOption) (*EditResponse, error)
	DeleteTakeByCourseClass(ctx context.Context, in *CourseID, opts ...client.CallOption) (*EditResponse, error)
	SearchTakeByUser(ctx context.Context, in *UserID, opts ...client.CallOption) (*SearchTakeByUserResponse, error)
	SearchTakeByCourse(ctx context.Context, in *CourseID, opts ...client.CallOption) (*SearchTakeByCourseResponse, error)
	SearchUserNotInCourse(ctx context.Context, in *CourseID, opts ...client.CallOption) (*SearchUserNotInCourseResponse, error)
}

type courseClassService struct {
	c    client.Client
	name string
}

func NewCourseClassService(name string, c client.Client) CourseClassService {
	return &courseClassService{
		c:    c,
		name: name,
	}
}

func (c *courseClassService) DeleteCourseClass(ctx context.Context, in *CourseID, opts ...client.CallOption) (*EditResponse, error) {
	req := c.c.NewRequest(c.name, "CourseClassService.DeleteCourseClass", in)
	out := new(EditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClassService) UpdateCourseClass(ctx context.Context, in *CourseClass, opts ...client.CallOption) (*EditResponse, error) {
	req := c.c.NewRequest(c.name, "CourseClassService.UpdateCourseClass", in)
	out := new(EditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClassService) SearchCourseClass(ctx context.Context, in *CourseID, opts ...client.CallOption) (*SearchCourseClassResponse, error) {
	req := c.c.NewRequest(c.name, "CourseClassService.SearchCourseClass", in)
	out := new(SearchCourseClassResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClassService) SearchCourseClasses(ctx context.Context, in *CourseIDArray, opts ...client.CallOption) (*SearchCourseClassesResponse, error) {
	req := c.c.NewRequest(c.name, "CourseClassService.SearchCourseClasses", in)
	out := new(SearchCourseClassesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClassService) NewCourse(ctx context.Context, in *NewCourseMessage, opts ...client.CallOption) (*NewCourseResponse, error) {
	req := c.c.NewRequest(c.name, "CourseClassService.NewCourse", in)
	out := new(NewCourseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClassService) AddTake(ctx context.Context, in *Take, opts ...client.CallOption) (*EditResponse, error) {
	req := c.c.NewRequest(c.name, "CourseClassService.AddTake", in)
	out := new(EditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClassService) DeleteTake(ctx context.Context, in *UserCourse, opts ...client.CallOption) (*EditResponse, error) {
	req := c.c.NewRequest(c.name, "CourseClassService.DeleteTake", in)
	out := new(EditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClassService) DeleteTakeByUser(ctx context.Context, in *UserID, opts ...client.CallOption) (*EditResponse, error) {
	req := c.c.NewRequest(c.name, "CourseClassService.DeleteTakeByUser", in)
	out := new(EditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClassService) DeleteTakeByCourseClass(ctx context.Context, in *CourseID, opts ...client.CallOption) (*EditResponse, error) {
	req := c.c.NewRequest(c.name, "CourseClassService.DeleteTakeByCourseClass", in)
	out := new(EditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClassService) SearchTakeByUser(ctx context.Context, in *UserID, opts ...client.CallOption) (*SearchTakeByUserResponse, error) {
	req := c.c.NewRequest(c.name, "CourseClassService.SearchTakeByUser", in)
	out := new(SearchTakeByUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClassService) SearchTakeByCourse(ctx context.Context, in *CourseID, opts ...client.CallOption) (*SearchTakeByCourseResponse, error) {
	req := c.c.NewRequest(c.name, "CourseClassService.SearchTakeByCourse", in)
	out := new(SearchTakeByCourseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClassService) SearchUserNotInCourse(ctx context.Context, in *CourseID, opts ...client.CallOption) (*SearchUserNotInCourseResponse, error) {
	req := c.c.NewRequest(c.name, "CourseClassService.SearchUserNotInCourse", in)
	out := new(SearchUserNotInCourseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CourseClassService service

type CourseClassServiceHandler interface {
	//rpc AddCourseClass (CourseClass) returns (EditResponse) {}
	DeleteCourseClass(context.Context, *CourseID, *EditResponse) error
	UpdateCourseClass(context.Context, *CourseClass, *EditResponse) error
	SearchCourseClass(context.Context, *CourseID, *SearchCourseClassResponse) error
	SearchCourseClasses(context.Context, *CourseIDArray, *SearchCourseClassesResponse) error
	NewCourse(context.Context, *NewCourseMessage, *NewCourseResponse) error
	AddTake(context.Context, *Take, *EditResponse) error
	DeleteTake(context.Context, *UserCourse, *EditResponse) error
	DeleteTakeByUser(context.Context, *UserID, *EditResponse) error
	DeleteTakeByCourseClass(context.Context, *CourseID, *EditResponse) error
	SearchTakeByUser(context.Context, *UserID, *SearchTakeByUserResponse) error
	SearchTakeByCourse(context.Context, *CourseID, *SearchTakeByCourseResponse) error
	SearchUserNotInCourse(context.Context, *CourseID, *SearchUserNotInCourseResponse) error
}

func RegisterCourseClassServiceHandler(s server.Server, hdlr CourseClassServiceHandler, opts ...server.HandlerOption) error {
	type courseClassService interface {
		DeleteCourseClass(ctx context.Context, in *CourseID, out *EditResponse) error
		UpdateCourseClass(ctx context.Context, in *CourseClass, out *EditResponse) error
		SearchCourseClass(ctx context.Context, in *CourseID, out *SearchCourseClassResponse) error
		SearchCourseClasses(ctx context.Context, in *CourseIDArray, out *SearchCourseClassesResponse) error
		NewCourse(ctx context.Context, in *NewCourseMessage, out *NewCourseResponse) error
		AddTake(ctx context.Context, in *Take, out *EditResponse) error
		DeleteTake(ctx context.Context, in *UserCourse, out *EditResponse) error
		DeleteTakeByUser(ctx context.Context, in *UserID, out *EditResponse) error
		DeleteTakeByCourseClass(ctx context.Context, in *CourseID, out *EditResponse) error
		SearchTakeByUser(ctx context.Context, in *UserID, out *SearchTakeByUserResponse) error
		SearchTakeByCourse(ctx context.Context, in *CourseID, out *SearchTakeByCourseResponse) error
		SearchUserNotInCourse(ctx context.Context, in *CourseID, out *SearchUserNotInCourseResponse) error
	}
	type CourseClassService struct {
		courseClassService
	}
	h := &courseClassServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CourseClassService{h}, opts...))
}

type courseClassServiceHandler struct {
	CourseClassServiceHandler
}

func (h *courseClassServiceHandler) DeleteCourseClass(ctx context.Context, in *CourseID, out *EditResponse) error {
	return h.CourseClassServiceHandler.DeleteCourseClass(ctx, in, out)
}

func (h *courseClassServiceHandler) UpdateCourseClass(ctx context.Context, in *CourseClass, out *EditResponse) error {
	return h.CourseClassServiceHandler.UpdateCourseClass(ctx, in, out)
}

func (h *courseClassServiceHandler) SearchCourseClass(ctx context.Context, in *CourseID, out *SearchCourseClassResponse) error {
	return h.CourseClassServiceHandler.SearchCourseClass(ctx, in, out)
}

func (h *courseClassServiceHandler) SearchCourseClasses(ctx context.Context, in *CourseIDArray, out *SearchCourseClassesResponse) error {
	return h.CourseClassServiceHandler.SearchCourseClasses(ctx, in, out)
}

func (h *courseClassServiceHandler) NewCourse(ctx context.Context, in *NewCourseMessage, out *NewCourseResponse) error {
	return h.CourseClassServiceHandler.NewCourse(ctx, in, out)
}

func (h *courseClassServiceHandler) AddTake(ctx context.Context, in *Take, out *EditResponse) error {
	return h.CourseClassServiceHandler.AddTake(ctx, in, out)
}

func (h *courseClassServiceHandler) DeleteTake(ctx context.Context, in *UserCourse, out *EditResponse) error {
	return h.CourseClassServiceHandler.DeleteTake(ctx, in, out)
}

func (h *courseClassServiceHandler) DeleteTakeByUser(ctx context.Context, in *UserID, out *EditResponse) error {
	return h.CourseClassServiceHandler.DeleteTakeByUser(ctx, in, out)
}

func (h *courseClassServiceHandler) DeleteTakeByCourseClass(ctx context.Context, in *CourseID, out *EditResponse) error {
	return h.CourseClassServiceHandler.DeleteTakeByCourseClass(ctx, in, out)
}

func (h *courseClassServiceHandler) SearchTakeByUser(ctx context.Context, in *UserID, out *SearchTakeByUserResponse) error {
	return h.CourseClassServiceHandler.SearchTakeByUser(ctx, in, out)
}

func (h *courseClassServiceHandler) SearchTakeByCourse(ctx context.Context, in *CourseID, out *SearchTakeByCourseResponse) error {
	return h.CourseClassServiceHandler.SearchTakeByCourse(ctx, in, out)
}

func (h *courseClassServiceHandler) SearchUserNotInCourse(ctx context.Context, in *CourseID, out *SearchUserNotInCourseResponse) error {
	return h.CourseClassServiceHandler.SearchUserNotInCourse(ctx, in, out)
}
