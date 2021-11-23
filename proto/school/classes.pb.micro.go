// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/school/classes.proto

package omo_msp_school

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

// Client API for ClassesService service

type ClassesService interface {
	AddOne(ctx context.Context, in *ReqClassAdd, opts ...client.CallOption) (*ReplyClassList, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyClassInfo, error)
	GetList(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyClassList, error)
	UpdateOne(ctx context.Context, in *ReqClassUpdate, opts ...client.CallOption) (*ReplyClassInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	SetMaster(ctx context.Context, in *ReqClassMaster, opts ...client.CallOption) (*ReplyInfo, error)
	AppendStudent(ctx context.Context, in *ReqClassStudent, opts ...client.CallOption) (*ReplyClassStudents, error)
	SubtractStudent(ctx context.Context, in *ReqClassStudent, opts ...client.CallOption) (*ReplyClassStudents, error)
	AppendTeacher(ctx context.Context, in *ReqClassTeacher, opts ...client.CallOption) (*ReplyList, error)
	SubtractTeacher(ctx context.Context, in *ReqClassTeacher, opts ...client.CallOption) (*ReplyList, error)
}

type classesService struct {
	c    client.Client
	name string
}

func NewClassesService(name string, c client.Client) ClassesService {
	return &classesService{
		c:    c,
		name: name,
	}
}

func (c *classesService) AddOne(ctx context.Context, in *ReqClassAdd, opts ...client.CallOption) (*ReplyClassList, error) {
	req := c.c.NewRequest(c.name, "ClassesService.AddOne", in)
	out := new(ReplyClassList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classesService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyClassInfo, error) {
	req := c.c.NewRequest(c.name, "ClassesService.GetOne", in)
	out := new(ReplyClassInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classesService) GetList(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyClassList, error) {
	req := c.c.NewRequest(c.name, "ClassesService.GetList", in)
	out := new(ReplyClassList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classesService) UpdateOne(ctx context.Context, in *ReqClassUpdate, opts ...client.CallOption) (*ReplyClassInfo, error) {
	req := c.c.NewRequest(c.name, "ClassesService.UpdateOne", in)
	out := new(ReplyClassInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classesService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "ClassesService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classesService) SetMaster(ctx context.Context, in *ReqClassMaster, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "ClassesService.SetMaster", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classesService) AppendStudent(ctx context.Context, in *ReqClassStudent, opts ...client.CallOption) (*ReplyClassStudents, error) {
	req := c.c.NewRequest(c.name, "ClassesService.AppendStudent", in)
	out := new(ReplyClassStudents)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classesService) SubtractStudent(ctx context.Context, in *ReqClassStudent, opts ...client.CallOption) (*ReplyClassStudents, error) {
	req := c.c.NewRequest(c.name, "ClassesService.SubtractStudent", in)
	out := new(ReplyClassStudents)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classesService) AppendTeacher(ctx context.Context, in *ReqClassTeacher, opts ...client.CallOption) (*ReplyList, error) {
	req := c.c.NewRequest(c.name, "ClassesService.AppendTeacher", in)
	out := new(ReplyList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classesService) SubtractTeacher(ctx context.Context, in *ReqClassTeacher, opts ...client.CallOption) (*ReplyList, error) {
	req := c.c.NewRequest(c.name, "ClassesService.SubtractTeacher", in)
	out := new(ReplyList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClassesService service

type ClassesServiceHandler interface {
	AddOne(context.Context, *ReqClassAdd, *ReplyClassList) error
	GetOne(context.Context, *RequestInfo, *ReplyClassInfo) error
	GetList(context.Context, *RequestInfo, *ReplyClassList) error
	UpdateOne(context.Context, *ReqClassUpdate, *ReplyClassInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	SetMaster(context.Context, *ReqClassMaster, *ReplyInfo) error
	AppendStudent(context.Context, *ReqClassStudent, *ReplyClassStudents) error
	SubtractStudent(context.Context, *ReqClassStudent, *ReplyClassStudents) error
	AppendTeacher(context.Context, *ReqClassTeacher, *ReplyList) error
	SubtractTeacher(context.Context, *ReqClassTeacher, *ReplyList) error
}

func RegisterClassesServiceHandler(s server.Server, hdlr ClassesServiceHandler, opts ...server.HandlerOption) error {
	type classesService interface {
		AddOne(ctx context.Context, in *ReqClassAdd, out *ReplyClassList) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyClassInfo) error
		GetList(ctx context.Context, in *RequestInfo, out *ReplyClassList) error
		UpdateOne(ctx context.Context, in *ReqClassUpdate, out *ReplyClassInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		SetMaster(ctx context.Context, in *ReqClassMaster, out *ReplyInfo) error
		AppendStudent(ctx context.Context, in *ReqClassStudent, out *ReplyClassStudents) error
		SubtractStudent(ctx context.Context, in *ReqClassStudent, out *ReplyClassStudents) error
		AppendTeacher(ctx context.Context, in *ReqClassTeacher, out *ReplyList) error
		SubtractTeacher(ctx context.Context, in *ReqClassTeacher, out *ReplyList) error
	}
	type ClassesService struct {
		classesService
	}
	h := &classesServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ClassesService{h}, opts...))
}

type classesServiceHandler struct {
	ClassesServiceHandler
}

func (h *classesServiceHandler) AddOne(ctx context.Context, in *ReqClassAdd, out *ReplyClassList) error {
	return h.ClassesServiceHandler.AddOne(ctx, in, out)
}

func (h *classesServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyClassInfo) error {
	return h.ClassesServiceHandler.GetOne(ctx, in, out)
}

func (h *classesServiceHandler) GetList(ctx context.Context, in *RequestInfo, out *ReplyClassList) error {
	return h.ClassesServiceHandler.GetList(ctx, in, out)
}

func (h *classesServiceHandler) UpdateOne(ctx context.Context, in *ReqClassUpdate, out *ReplyClassInfo) error {
	return h.ClassesServiceHandler.UpdateOne(ctx, in, out)
}

func (h *classesServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.ClassesServiceHandler.RemoveOne(ctx, in, out)
}

func (h *classesServiceHandler) SetMaster(ctx context.Context, in *ReqClassMaster, out *ReplyInfo) error {
	return h.ClassesServiceHandler.SetMaster(ctx, in, out)
}

func (h *classesServiceHandler) AppendStudent(ctx context.Context, in *ReqClassStudent, out *ReplyClassStudents) error {
	return h.ClassesServiceHandler.AppendStudent(ctx, in, out)
}

func (h *classesServiceHandler) SubtractStudent(ctx context.Context, in *ReqClassStudent, out *ReplyClassStudents) error {
	return h.ClassesServiceHandler.SubtractStudent(ctx, in, out)
}

func (h *classesServiceHandler) AppendTeacher(ctx context.Context, in *ReqClassTeacher, out *ReplyList) error {
	return h.ClassesServiceHandler.AppendTeacher(ctx, in, out)
}

func (h *classesServiceHandler) SubtractTeacher(ctx context.Context, in *ReqClassTeacher, out *ReplyList) error {
	return h.ClassesServiceHandler.SubtractTeacher(ctx, in, out)
}
