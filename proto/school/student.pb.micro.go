// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/school/student.proto

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

// Client API for StudentService service

type StudentService interface {
	AddOne(ctx context.Context, in *ReqStudentAdd, opts ...client.CallOption) (*ReplyStudentInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyStudentInfo, error)
	GetByFilter(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplyStudentList, error)
	GetList(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplyStudentList, error)
	GetArray(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyStudentList, error)
	GetStatistic(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateOne(ctx context.Context, in *ReqStudentUpdate, opts ...client.CallOption) (*ReplyStudentInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	AddBatch(ctx context.Context, in *ReqStudentBatch, opts ...client.CallOption) (*ReplyStudentList, error)
	BindEntity(ctx context.Context, in *ReqStudentBind, opts ...client.CallOption) (*ReplyStudentInfo, error)
	UpdateCustodian(ctx context.Context, in *ReqStudentCustodian, opts ...client.CallOption) (*ReplyStudentInfo, error)
	UpdateTags(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyList, error)
	UpdateStatus(ctx context.Context, in *RequestState, opts ...client.CallOption) (*ReplyInfo, error)
}

type studentService struct {
	c    client.Client
	name string
}

func NewStudentService(name string, c client.Client) StudentService {
	return &studentService{
		c:    c,
		name: name,
	}
}

func (c *studentService) AddOne(ctx context.Context, in *ReqStudentAdd, opts ...client.CallOption) (*ReplyStudentInfo, error) {
	req := c.c.NewRequest(c.name, "StudentService.AddOne", in)
	out := new(ReplyStudentInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyStudentInfo, error) {
	req := c.c.NewRequest(c.name, "StudentService.GetOne", in)
	out := new(ReplyStudentInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentService) GetByFilter(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplyStudentList, error) {
	req := c.c.NewRequest(c.name, "StudentService.GetByFilter", in)
	out := new(ReplyStudentList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentService) GetList(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplyStudentList, error) {
	req := c.c.NewRequest(c.name, "StudentService.GetList", in)
	out := new(ReplyStudentList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentService) GetArray(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyStudentList, error) {
	req := c.c.NewRequest(c.name, "StudentService.GetArray", in)
	out := new(ReplyStudentList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentService) GetStatistic(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "StudentService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentService) UpdateOne(ctx context.Context, in *ReqStudentUpdate, opts ...client.CallOption) (*ReplyStudentInfo, error) {
	req := c.c.NewRequest(c.name, "StudentService.UpdateOne", in)
	out := new(ReplyStudentInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "StudentService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentService) AddBatch(ctx context.Context, in *ReqStudentBatch, opts ...client.CallOption) (*ReplyStudentList, error) {
	req := c.c.NewRequest(c.name, "StudentService.AddBatch", in)
	out := new(ReplyStudentList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentService) BindEntity(ctx context.Context, in *ReqStudentBind, opts ...client.CallOption) (*ReplyStudentInfo, error) {
	req := c.c.NewRequest(c.name, "StudentService.BindEntity", in)
	out := new(ReplyStudentInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentService) UpdateCustodian(ctx context.Context, in *ReqStudentCustodian, opts ...client.CallOption) (*ReplyStudentInfo, error) {
	req := c.c.NewRequest(c.name, "StudentService.UpdateCustodian", in)
	out := new(ReplyStudentInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentService) UpdateTags(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyList, error) {
	req := c.c.NewRequest(c.name, "StudentService.UpdateTags", in)
	out := new(ReplyList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentService) UpdateStatus(ctx context.Context, in *RequestState, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "StudentService.UpdateStatus", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StudentService service

type StudentServiceHandler interface {
	AddOne(context.Context, *ReqStudentAdd, *ReplyStudentInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyStudentInfo) error
	GetByFilter(context.Context, *RequestPage, *ReplyStudentList) error
	GetList(context.Context, *RequestPage, *ReplyStudentList) error
	GetArray(context.Context, *RequestList, *ReplyStudentList) error
	GetStatistic(context.Context, *RequestPage, *ReplyStatistic) error
	UpdateOne(context.Context, *ReqStudentUpdate, *ReplyStudentInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	AddBatch(context.Context, *ReqStudentBatch, *ReplyStudentList) error
	BindEntity(context.Context, *ReqStudentBind, *ReplyStudentInfo) error
	UpdateCustodian(context.Context, *ReqStudentCustodian, *ReplyStudentInfo) error
	UpdateTags(context.Context, *RequestList, *ReplyList) error
	UpdateStatus(context.Context, *RequestState, *ReplyInfo) error
}

func RegisterStudentServiceHandler(s server.Server, hdlr StudentServiceHandler, opts ...server.HandlerOption) error {
	type studentService interface {
		AddOne(ctx context.Context, in *ReqStudentAdd, out *ReplyStudentInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyStudentInfo) error
		GetByFilter(ctx context.Context, in *RequestPage, out *ReplyStudentList) error
		GetList(ctx context.Context, in *RequestPage, out *ReplyStudentList) error
		GetArray(ctx context.Context, in *RequestList, out *ReplyStudentList) error
		GetStatistic(ctx context.Context, in *RequestPage, out *ReplyStatistic) error
		UpdateOne(ctx context.Context, in *ReqStudentUpdate, out *ReplyStudentInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		AddBatch(ctx context.Context, in *ReqStudentBatch, out *ReplyStudentList) error
		BindEntity(ctx context.Context, in *ReqStudentBind, out *ReplyStudentInfo) error
		UpdateCustodian(ctx context.Context, in *ReqStudentCustodian, out *ReplyStudentInfo) error
		UpdateTags(ctx context.Context, in *RequestList, out *ReplyList) error
		UpdateStatus(ctx context.Context, in *RequestState, out *ReplyInfo) error
	}
	type StudentService struct {
		studentService
	}
	h := &studentServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&StudentService{h}, opts...))
}

type studentServiceHandler struct {
	StudentServiceHandler
}

func (h *studentServiceHandler) AddOne(ctx context.Context, in *ReqStudentAdd, out *ReplyStudentInfo) error {
	return h.StudentServiceHandler.AddOne(ctx, in, out)
}

func (h *studentServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyStudentInfo) error {
	return h.StudentServiceHandler.GetOne(ctx, in, out)
}

func (h *studentServiceHandler) GetByFilter(ctx context.Context, in *RequestPage, out *ReplyStudentList) error {
	return h.StudentServiceHandler.GetByFilter(ctx, in, out)
}

func (h *studentServiceHandler) GetList(ctx context.Context, in *RequestPage, out *ReplyStudentList) error {
	return h.StudentServiceHandler.GetList(ctx, in, out)
}

func (h *studentServiceHandler) GetArray(ctx context.Context, in *RequestList, out *ReplyStudentList) error {
	return h.StudentServiceHandler.GetArray(ctx, in, out)
}

func (h *studentServiceHandler) GetStatistic(ctx context.Context, in *RequestPage, out *ReplyStatistic) error {
	return h.StudentServiceHandler.GetStatistic(ctx, in, out)
}

func (h *studentServiceHandler) UpdateOne(ctx context.Context, in *ReqStudentUpdate, out *ReplyStudentInfo) error {
	return h.StudentServiceHandler.UpdateOne(ctx, in, out)
}

func (h *studentServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.StudentServiceHandler.RemoveOne(ctx, in, out)
}

func (h *studentServiceHandler) AddBatch(ctx context.Context, in *ReqStudentBatch, out *ReplyStudentList) error {
	return h.StudentServiceHandler.AddBatch(ctx, in, out)
}

func (h *studentServiceHandler) BindEntity(ctx context.Context, in *ReqStudentBind, out *ReplyStudentInfo) error {
	return h.StudentServiceHandler.BindEntity(ctx, in, out)
}

func (h *studentServiceHandler) UpdateCustodian(ctx context.Context, in *ReqStudentCustodian, out *ReplyStudentInfo) error {
	return h.StudentServiceHandler.UpdateCustodian(ctx, in, out)
}

func (h *studentServiceHandler) UpdateTags(ctx context.Context, in *RequestList, out *ReplyList) error {
	return h.StudentServiceHandler.UpdateTags(ctx, in, out)
}

func (h *studentServiceHandler) UpdateStatus(ctx context.Context, in *RequestState, out *ReplyInfo) error {
	return h.StudentServiceHandler.UpdateStatus(ctx, in, out)
}
