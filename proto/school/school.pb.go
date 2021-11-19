// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/school/school.proto

package omo_msp_school

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

type HonorInfo struct {
	Uid                  string        `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string        `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Parent               string        `protobuf:"bytes,4,opt,name=parent,proto3" json:"parent,omitempty"`
	Children             []*HonorInfo  `protobuf:"bytes,5,rep,name=children,proto3" json:"children,omitempty"`
	Bries                []*HonorBrief `protobuf:"bytes,6,rep,name=bries,proto3" json:"bries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *HonorInfo) Reset()         { *m = HonorInfo{} }
func (m *HonorInfo) String() string { return proto.CompactTextString(m) }
func (*HonorInfo) ProtoMessage()    {}
func (*HonorInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_26e8684ccc1feb24, []int{0}
}

func (m *HonorInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HonorInfo.Unmarshal(m, b)
}
func (m *HonorInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HonorInfo.Marshal(b, m, deterministic)
}
func (m *HonorInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HonorInfo.Merge(m, src)
}
func (m *HonorInfo) XXX_Size() int {
	return xxx_messageInfo_HonorInfo.Size(m)
}
func (m *HonorInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HonorInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HonorInfo proto.InternalMessageInfo

func (m *HonorInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *HonorInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HonorInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *HonorInfo) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *HonorInfo) GetChildren() []*HonorInfo {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *HonorInfo) GetBries() []*HonorBrief {
	if m != nil {
		return m.Bries
	}
	return nil
}

type HonorBrief struct {
	Year                 string   `protobuf:"bytes,1,opt,name=year,proto3" json:"year,omitempty"`
	Count                uint32   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Entities             []string `protobuf:"bytes,3,rep,name=entities,proto3" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HonorBrief) Reset()         { *m = HonorBrief{} }
func (m *HonorBrief) String() string { return proto.CompactTextString(m) }
func (*HonorBrief) ProtoMessage()    {}
func (*HonorBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_26e8684ccc1feb24, []int{1}
}

func (m *HonorBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HonorBrief.Unmarshal(m, b)
}
func (m *HonorBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HonorBrief.Marshal(b, m, deterministic)
}
func (m *HonorBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HonorBrief.Merge(m, src)
}
func (m *HonorBrief) XXX_Size() int {
	return xxx_messageInfo_HonorBrief.Size(m)
}
func (m *HonorBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_HonorBrief.DiscardUnknown(m)
}

var xxx_messageInfo_HonorBrief proto.InternalMessageInfo

func (m *HonorBrief) GetYear() string {
	if m != nil {
		return m.Year
	}
	return ""
}

func (m *HonorBrief) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *HonorBrief) GetEntities() []string {
	if m != nil {
		return m.Entities
	}
	return nil
}

type SubjectInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubjectInfo) Reset()         { *m = SubjectInfo{} }
func (m *SubjectInfo) String() string { return proto.CompactTextString(m) }
func (*SubjectInfo) ProtoMessage()    {}
func (*SubjectInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_26e8684ccc1feb24, []int{2}
}

func (m *SubjectInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubjectInfo.Unmarshal(m, b)
}
func (m *SubjectInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubjectInfo.Marshal(b, m, deterministic)
}
func (m *SubjectInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubjectInfo.Merge(m, src)
}
func (m *SubjectInfo) XXX_Size() int {
	return xxx_messageInfo_SubjectInfo.Size(m)
}
func (m *SubjectInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SubjectInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SubjectInfo proto.InternalMessageInfo

func (m *SubjectInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *SubjectInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SubjectInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

type SchoolInfo struct {
	Id                   uint64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string         `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Scene                string         `protobuf:"bytes,3,opt,name=scene,proto3" json:"scene,omitempty"`
	Created              uint64         `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              uint64         `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string         `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string         `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string         `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Grade                uint32         `protobuf:"varint,9,opt,name=grade,proto3" json:"grade,omitempty"`
	Entity               string         `protobuf:"bytes,10,opt,name=entity,proto3" json:"entity,omitempty"`
	Honors               []*HonorInfo   `protobuf:"bytes,11,rep,name=honors,proto3" json:"honors,omitempty"`
	Respects             []*HonorInfo   `protobuf:"bytes,12,rep,name=respects,proto3" json:"respects,omitempty"`
	Subjects             []*SubjectInfo `protobuf:"bytes,13,rep,name=subjects,proto3" json:"subjects,omitempty"`
	Teachers             []string       `protobuf:"bytes,14,rep,name=teachers,proto3" json:"teachers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SchoolInfo) Reset()         { *m = SchoolInfo{} }
func (m *SchoolInfo) String() string { return proto.CompactTextString(m) }
func (*SchoolInfo) ProtoMessage()    {}
func (*SchoolInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_26e8684ccc1feb24, []int{3}
}

func (m *SchoolInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchoolInfo.Unmarshal(m, b)
}
func (m *SchoolInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchoolInfo.Marshal(b, m, deterministic)
}
func (m *SchoolInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchoolInfo.Merge(m, src)
}
func (m *SchoolInfo) XXX_Size() int {
	return xxx_messageInfo_SchoolInfo.Size(m)
}
func (m *SchoolInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SchoolInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SchoolInfo proto.InternalMessageInfo

func (m *SchoolInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SchoolInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *SchoolInfo) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *SchoolInfo) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *SchoolInfo) GetUpdated() uint64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *SchoolInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SchoolInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *SchoolInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SchoolInfo) GetGrade() uint32 {
	if m != nil {
		return m.Grade
	}
	return 0
}

func (m *SchoolInfo) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *SchoolInfo) GetHonors() []*HonorInfo {
	if m != nil {
		return m.Honors
	}
	return nil
}

func (m *SchoolInfo) GetRespects() []*HonorInfo {
	if m != nil {
		return m.Respects
	}
	return nil
}

func (m *SchoolInfo) GetSubjects() []*SubjectInfo {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *SchoolInfo) GetTeachers() []string {
	if m != nil {
		return m.Teachers
	}
	return nil
}

type ReqSchoolAdd struct {
	Scene                string   `protobuf:"bytes,1,opt,name=scene,proto3" json:"scene,omitempty"`
	Grade                uint32   `protobuf:"varint,2,opt,name=grade,proto3" json:"grade,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Cover                string   `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Entity               string   `protobuf:"bytes,5,opt,name=entity,proto3" json:"entity,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSchoolAdd) Reset()         { *m = ReqSchoolAdd{} }
func (m *ReqSchoolAdd) String() string { return proto.CompactTextString(m) }
func (*ReqSchoolAdd) ProtoMessage()    {}
func (*ReqSchoolAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_26e8684ccc1feb24, []int{4}
}

func (m *ReqSchoolAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSchoolAdd.Unmarshal(m, b)
}
func (m *ReqSchoolAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSchoolAdd.Marshal(b, m, deterministic)
}
func (m *ReqSchoolAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSchoolAdd.Merge(m, src)
}
func (m *ReqSchoolAdd) XXX_Size() int {
	return xxx_messageInfo_ReqSchoolAdd.Size(m)
}
func (m *ReqSchoolAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSchoolAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSchoolAdd proto.InternalMessageInfo

func (m *ReqSchoolAdd) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ReqSchoolAdd) GetGrade() uint32 {
	if m != nil {
		return m.Grade
	}
	return 0
}

func (m *ReqSchoolAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqSchoolAdd) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqSchoolAdd) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *ReqSchoolAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReqSchoolSubject struct {
	Option               OptionType `protobuf:"varint,1,opt,name=option,proto3,enum=omo.msp.school.OptionType" json:"option,omitempty"`
	Uid                  string     `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Scene                string     `protobuf:"bytes,3,opt,name=scene,proto3" json:"scene,omitempty"`
	Name                 string     `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string     `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string     `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReqSchoolSubject) Reset()         { *m = ReqSchoolSubject{} }
func (m *ReqSchoolSubject) String() string { return proto.CompactTextString(m) }
func (*ReqSchoolSubject) ProtoMessage()    {}
func (*ReqSchoolSubject) Descriptor() ([]byte, []int) {
	return fileDescriptor_26e8684ccc1feb24, []int{5}
}

func (m *ReqSchoolSubject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSchoolSubject.Unmarshal(m, b)
}
func (m *ReqSchoolSubject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSchoolSubject.Marshal(b, m, deterministic)
}
func (m *ReqSchoolSubject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSchoolSubject.Merge(m, src)
}
func (m *ReqSchoolSubject) XXX_Size() int {
	return xxx_messageInfo_ReqSchoolSubject.Size(m)
}
func (m *ReqSchoolSubject) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSchoolSubject.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSchoolSubject proto.InternalMessageInfo

func (m *ReqSchoolSubject) GetOption() OptionType {
	if m != nil {
		return m.Option
	}
	return OptionType_Add
}

func (m *ReqSchoolSubject) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqSchoolSubject) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ReqSchoolSubject) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqSchoolSubject) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqSchoolSubject) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReqSchoolUpdate struct {
	Scene                string   `protobuf:"bytes,1,opt,name=scene,proto3" json:"scene,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSchoolUpdate) Reset()         { *m = ReqSchoolUpdate{} }
func (m *ReqSchoolUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqSchoolUpdate) ProtoMessage()    {}
func (*ReqSchoolUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_26e8684ccc1feb24, []int{6}
}

func (m *ReqSchoolUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSchoolUpdate.Unmarshal(m, b)
}
func (m *ReqSchoolUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSchoolUpdate.Marshal(b, m, deterministic)
}
func (m *ReqSchoolUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSchoolUpdate.Merge(m, src)
}
func (m *ReqSchoolUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqSchoolUpdate.Size(m)
}
func (m *ReqSchoolUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSchoolUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSchoolUpdate proto.InternalMessageInfo

func (m *ReqSchoolUpdate) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ReqSchoolUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqSchoolUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqSchoolUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplySchoolInfo struct {
	Info                 *SchoolInfo  `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Role                 uint32       `protobuf:"varint,2,opt,name=role,proto3" json:"role,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplySchoolInfo) Reset()         { *m = ReplySchoolInfo{} }
func (m *ReplySchoolInfo) String() string { return proto.CompactTextString(m) }
func (*ReplySchoolInfo) ProtoMessage()    {}
func (*ReplySchoolInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_26e8684ccc1feb24, []int{7}
}

func (m *ReplySchoolInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplySchoolInfo.Unmarshal(m, b)
}
func (m *ReplySchoolInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplySchoolInfo.Marshal(b, m, deterministic)
}
func (m *ReplySchoolInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplySchoolInfo.Merge(m, src)
}
func (m *ReplySchoolInfo) XXX_Size() int {
	return xxx_messageInfo_ReplySchoolInfo.Size(m)
}
func (m *ReplySchoolInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplySchoolInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplySchoolInfo proto.InternalMessageInfo

func (m *ReplySchoolInfo) GetInfo() *SchoolInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplySchoolInfo) GetRole() uint32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *ReplySchoolInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplySchoolHonors struct {
	Scene                string       `protobuf:"bytes,1,opt,name=scene,proto3" json:"scene,omitempty"`
	Type                 uint32       `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Honors               []*HonorInfo `protobuf:"bytes,3,rep,name=honors,proto3" json:"honors,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplySchoolHonors) Reset()         { *m = ReplySchoolHonors{} }
func (m *ReplySchoolHonors) String() string { return proto.CompactTextString(m) }
func (*ReplySchoolHonors) ProtoMessage()    {}
func (*ReplySchoolHonors) Descriptor() ([]byte, []int) {
	return fileDescriptor_26e8684ccc1feb24, []int{8}
}

func (m *ReplySchoolHonors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplySchoolHonors.Unmarshal(m, b)
}
func (m *ReplySchoolHonors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplySchoolHonors.Marshal(b, m, deterministic)
}
func (m *ReplySchoolHonors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplySchoolHonors.Merge(m, src)
}
func (m *ReplySchoolHonors) XXX_Size() int {
	return xxx_messageInfo_ReplySchoolHonors.Size(m)
}
func (m *ReplySchoolHonors) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplySchoolHonors.DiscardUnknown(m)
}

var xxx_messageInfo_ReplySchoolHonors proto.InternalMessageInfo

func (m *ReplySchoolHonors) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ReplySchoolHonors) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReplySchoolHonors) GetHonors() []*HonorInfo {
	if m != nil {
		return m.Honors
	}
	return nil
}

func (m *ReplySchoolHonors) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplySchoolSubjects struct {
	Scene                string         `protobuf:"bytes,1,opt,name=scene,proto3" json:"scene,omitempty"`
	Subjects             []*SubjectInfo `protobuf:"bytes,2,rep,name=subjects,proto3" json:"subjects,omitempty"`
	Status               *ReplyStatus   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplySchoolSubjects) Reset()         { *m = ReplySchoolSubjects{} }
func (m *ReplySchoolSubjects) String() string { return proto.CompactTextString(m) }
func (*ReplySchoolSubjects) ProtoMessage()    {}
func (*ReplySchoolSubjects) Descriptor() ([]byte, []int) {
	return fileDescriptor_26e8684ccc1feb24, []int{9}
}

func (m *ReplySchoolSubjects) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplySchoolSubjects.Unmarshal(m, b)
}
func (m *ReplySchoolSubjects) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplySchoolSubjects.Marshal(b, m, deterministic)
}
func (m *ReplySchoolSubjects) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplySchoolSubjects.Merge(m, src)
}
func (m *ReplySchoolSubjects) XXX_Size() int {
	return xxx_messageInfo_ReplySchoolSubjects.Size(m)
}
func (m *ReplySchoolSubjects) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplySchoolSubjects.DiscardUnknown(m)
}

var xxx_messageInfo_ReplySchoolSubjects proto.InternalMessageInfo

func (m *ReplySchoolSubjects) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ReplySchoolSubjects) GetSubjects() []*SubjectInfo {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *ReplySchoolSubjects) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplySchoolList struct {
	Total                uint32        `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Page                 uint32        `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32        `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*SchoolInfo `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus  `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplySchoolList) Reset()         { *m = ReplySchoolList{} }
func (m *ReplySchoolList) String() string { return proto.CompactTextString(m) }
func (*ReplySchoolList) ProtoMessage()    {}
func (*ReplySchoolList) Descriptor() ([]byte, []int) {
	return fileDescriptor_26e8684ccc1feb24, []int{10}
}

func (m *ReplySchoolList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplySchoolList.Unmarshal(m, b)
}
func (m *ReplySchoolList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplySchoolList.Marshal(b, m, deterministic)
}
func (m *ReplySchoolList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplySchoolList.Merge(m, src)
}
func (m *ReplySchoolList) XXX_Size() int {
	return xxx_messageInfo_ReplySchoolList.Size(m)
}
func (m *ReplySchoolList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplySchoolList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplySchoolList proto.InternalMessageInfo

func (m *ReplySchoolList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplySchoolList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplySchoolList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplySchoolList) GetList() []*SchoolInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplySchoolList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*HonorInfo)(nil), "omo.msp.school.HonorInfo")
	proto.RegisterType((*HonorBrief)(nil), "omo.msp.school.HonorBrief")
	proto.RegisterType((*SubjectInfo)(nil), "omo.msp.school.SubjectInfo")
	proto.RegisterType((*SchoolInfo)(nil), "omo.msp.school.SchoolInfo")
	proto.RegisterType((*ReqSchoolAdd)(nil), "omo.msp.school.ReqSchoolAdd")
	proto.RegisterType((*ReqSchoolSubject)(nil), "omo.msp.school.ReqSchoolSubject")
	proto.RegisterType((*ReqSchoolUpdate)(nil), "omo.msp.school.ReqSchoolUpdate")
	proto.RegisterType((*ReplySchoolInfo)(nil), "omo.msp.school.ReplySchoolInfo")
	proto.RegisterType((*ReplySchoolHonors)(nil), "omo.msp.school.ReplySchoolHonors")
	proto.RegisterType((*ReplySchoolSubjects)(nil), "omo.msp.school.ReplySchoolSubjects")
	proto.RegisterType((*ReplySchoolList)(nil), "omo.msp.school.ReplySchoolList")
}

func init() {
	proto.RegisterFile("proto/school/school.proto", fileDescriptor_26e8684ccc1feb24)
}

var fileDescriptor_26e8684ccc1feb24 = []byte{
	// 823 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xcd, 0x6e, 0xe3, 0x36,
	0x10, 0x8e, 0xac, 0x9f, 0xd8, 0x93, 0xd8, 0x4d, 0xd9, 0xa0, 0x50, 0x9c, 0x16, 0x31, 0xd4, 0x4b,
	0x4e, 0x6e, 0x9b, 0xa0, 0xe8, 0xa1, 0xa7, 0x04, 0x28, 0x92, 0xa0, 0x05, 0x52, 0x30, 0x2d, 0x50,
	0xf4, 0x12, 0xc8, 0xd2, 0x24, 0x56, 0xd7, 0x12, 0x15, 0x92, 0x0e, 0xe0, 0x57, 0xd8, 0x27, 0xd8,
	0x7d, 0x80, 0x3d, 0xef, 0x6d, 0xb1, 0xcf, 0xb0, 0x6f, 0xb1, 0x6f, 0xb2, 0x20, 0x29, 0xc9, 0x72,
	0x22, 0xc7, 0xce, 0xee, 0xc9, 0xfc, 0xc8, 0xe1, 0xcc, 0x37, 0xdf, 0xcc, 0xd0, 0x82, 0xbd, 0x9c,
	0x33, 0xc9, 0x7e, 0x14, 0xd1, 0x98, 0xb1, 0x49, 0xf1, 0x33, 0xd4, 0x7b, 0xa4, 0xc7, 0x52, 0x36,
	0x4c, 0x45, 0x3e, 0x34, 0xbb, 0xfd, 0x45, 0xd3, 0x88, 0xa5, 0x29, 0xcb, 0x8c, 0x69, 0xf0, 0xc1,
	0x82, 0xce, 0x39, 0xcb, 0x18, 0xbf, 0xc8, 0x6e, 0x18, 0xd9, 0x01, 0x7b, 0x9a, 0xc4, 0xbe, 0x35,
	0xb0, 0x0e, 0x3b, 0x54, 0x2d, 0x09, 0x01, 0x27, 0x0b, 0x53, 0xf4, 0x5b, 0x7a, 0x4b, 0xaf, 0xc9,
	0xb7, 0xe0, 0x71, 0x4c, 0x43, 0xfe, 0xc2, 0xb7, 0xf5, 0x6e, 0x81, 0xd4, 0x7e, 0x1e, 0x72, 0xcc,
	0xa4, 0xef, 0x98, 0x7d, 0x83, 0xc8, 0x2f, 0xd0, 0x8e, 0xc6, 0xc9, 0x24, 0xe6, 0x98, 0xf9, 0xee,
	0xc0, 0x3e, 0xdc, 0x3a, 0xda, 0x1b, 0x2e, 0x32, 0x1c, 0x56, 0x14, 0x68, 0x65, 0x4a, 0x7e, 0x02,
	0x77, 0xc4, 0x13, 0x14, 0xbe, 0xa7, 0xef, 0xf4, 0x1b, 0xef, 0x9c, 0xf2, 0x04, 0x6f, 0xa8, 0x31,
	0x0c, 0x28, 0xc0, 0x7c, 0x53, 0x51, 0x9f, 0x61, 0xc8, 0x8b, 0x6c, 0xf4, 0x9a, 0xec, 0x82, 0x1b,
	0xb1, 0x69, 0x26, 0x75, 0x3e, 0x5d, 0x6a, 0x00, 0xe9, 0x43, 0x1b, 0x33, 0x99, 0x48, 0x15, 0xcc,
	0x1e, 0xd8, 0x87, 0x1d, 0x5a, 0xe1, 0xe0, 0x0f, 0xd8, 0xba, 0x9a, 0x8e, 0xfe, 0xc7, 0x48, 0x7e,
	0xb9, 0x42, 0xc1, 0x5b, 0x1b, 0xe0, 0x4a, 0xb3, 0xd7, 0xce, 0x7a, 0xd0, 0x2a, 0x7c, 0x39, 0xb4,
	0x95, 0xc4, 0xa5, 0xf3, 0xd6, 0xdc, 0xf9, 0x2e, 0xb8, 0x22, 0xc2, 0x0c, 0x0b, 0x3f, 0x06, 0x10,
	0x1f, 0x36, 0x23, 0x8e, 0xa1, 0xc4, 0x58, 0x2b, 0xed, 0xd0, 0x12, 0xaa, 0x93, 0x69, 0x1e, 0xeb,
	0x13, 0xd7, 0x9c, 0x14, 0xb0, 0xba, 0xc3, 0xb8, 0xef, 0x69, 0x5f, 0x25, 0x54, 0xd9, 0xb3, 0x1c,
	0xb9, 0x3e, 0xda, 0xd4, 0x47, 0x15, 0xae, 0x92, 0x6b, 0xd7, 0x92, 0xdb, 0x05, 0xf7, 0x96, 0x87,
	0x31, 0xfa, 0x1d, 0xa3, 0xa1, 0x06, 0x2a, 0x65, 0xad, 0xd9, 0xcc, 0x07, 0x93, 0xb2, 0x41, 0xe4,
	0x67, 0xf0, 0xc6, 0xaa, 0x26, 0xc2, 0xdf, 0x5a, 0x55, 0xfa, 0xc2, 0x50, 0xf5, 0x0b, 0x47, 0x91,
	0x63, 0x24, 0x85, 0xbf, 0xbd, 0xb2, 0x5f, 0x4a, 0x53, 0xf2, 0x2b, 0xb4, 0x85, 0xa9, 0x94, 0xf0,
	0xbb, 0xfa, 0xda, 0xfe, 0xc3, 0x6b, 0xb5, 0x4a, 0xd2, 0xca, 0x58, 0x09, 0x20, 0x31, 0x8c, 0xc6,
	0xc8, 0x85, 0xdf, 0x33, 0xe5, 0x2f, 0x71, 0xf0, 0xca, 0x82, 0x6d, 0x8a, 0x77, 0xa6, 0x68, 0x27,
	0x71, 0xad, 0x22, 0x56, 0xbd, 0x22, 0x95, 0x26, 0xad, 0xba, 0x26, 0xa5, 0x7a, 0xf6, 0xa2, 0x7a,
	0x11, 0xbb, 0x47, 0x5e, 0xcc, 0x88, 0x01, 0x35, 0xf5, 0xdc, 0x05, 0xf5, 0xea, 0xb5, 0xf1, 0x16,
	0x6b, 0x13, 0xbc, 0xb3, 0x60, 0xa7, 0xa2, 0x56, 0x64, 0x46, 0x8e, 0xc0, 0x63, 0xb9, 0x4c, 0x58,
	0xa6, 0xf9, 0xf5, 0x1e, 0x4f, 0xcd, 0xa5, 0x3e, 0xfd, 0x7b, 0x96, 0x23, 0x2d, 0x2c, 0xd7, 0x6e,
	0xbb, 0x32, 0x1d, 0xa7, 0xb1, 0xd3, 0xdd, 0x85, 0xb7, 0xe0, 0x29, 0xe2, 0x0c, 0xbe, 0xaa, 0x78,
	0xff, 0xa3, 0xdb, 0x73, 0x89, 0xaa, 0xcf, 0x79, 0x7c, 0xea, 0x01, 0x9d, 0x07, 0x01, 0x5f, 0x5a,
	0x2a, 0x62, 0x3e, 0x99, 0xd5, 0x66, 0x6f, 0x08, 0x4e, 0x92, 0xdd, 0x30, 0x1d, 0xb0, 0xe1, 0x71,
	0x99, 0x5b, 0x52, 0x6d, 0xa7, 0xb8, 0x70, 0x36, 0x29, 0x0b, 0xac, 0xd7, 0xe4, 0x18, 0x3c, 0x21,
	0x43, 0x39, 0x15, 0x9a, 0x4b, 0x43, 0xbf, 0x99, 0xa0, 0xda, 0x84, 0x16, 0xa6, 0xc1, 0x1b, 0x0b,
	0xbe, 0xae, 0x91, 0x39, 0x37, 0x3d, 0xbf, 0x54, 0x00, 0x39, 0xcb, 0xab, 0xa0, 0x6a, 0x5d, 0x1b,
	0x28, 0x7b, 0xdd, 0x81, 0x9a, 0xf3, 0x74, 0xd6, 0xe7, 0xf9, 0xda, 0x82, 0x6f, 0x6a, 0x3c, 0xaf,
	0xca, 0x69, 0x69, 0x66, 0x5a, 0x1f, 0xbe, 0xd6, 0x73, 0x86, 0xef, 0xb3, 0x34, 0x7c, 0xbf, 0x58,
	0xd0, 0x3f, 0x13, 0x21, 0x15, 0x2f, 0xc9, 0x64, 0x38, 0xd1, 0xbc, 0xba, 0xd4, 0x00, 0xa5, 0x60,
	0x1e, 0xde, 0x56, 0x0a, 0xaa, 0xb5, 0x6a, 0xa1, 0x6c, 0x9a, 0x8e, 0x90, 0xeb, 0x90, 0x5d, 0x5a,
	0x20, 0xd5, 0x12, 0x93, 0x44, 0xa8, 0x7f, 0x2f, 0x7b, 0x55, 0x4b, 0x28, 0xbb, 0x1a, 0x75, 0x77,
	0x6d, 0xea, 0x47, 0x1f, 0x6d, 0xe8, 0x16, 0x8a, 0x22, 0xbf, 0x4f, 0x22, 0x24, 0x17, 0xe0, 0x9d,
	0xc4, 0xf1, 0x65, 0x86, 0xe4, 0xbb, 0xc7, 0x0e, 0xe6, 0x2f, 0x4f, 0xff, 0xa0, 0xd9, 0x7d, 0xc5,
	0x2a, 0xd8, 0x20, 0xe7, 0xe0, 0x9d, 0xa1, 0x54, 0xae, 0x1a, 0xb8, 0xdc, 0x4d, 0x51, 0x68, 0xf5,
	0xd7, 0xf1, 0x74, 0x01, 0x9b, 0x67, 0x28, 0xb5, 0xb0, 0xcb, 0x5c, 0xfd, 0x15, 0xde, 0xe2, 0x93,
	0xae, 0xd4, 0xed, 0x60, 0x83, 0x5c, 0x42, 0xc7, 0x4c, 0xb9, 0xe2, 0x75, 0xb0, 0x34, 0x45, 0x63,
	0xb3, 0x0e, 0xb7, 0xdf, 0xa1, 0x43, 0x31, 0x65, 0xf7, 0xb8, 0x32, 0xd1, 0xbd, 0x46, 0x67, 0x85,
	0x9b, 0x7f, 0xa1, 0x6b, 0x62, 0x96, 0x6f, 0xe7, 0x60, 0x29, 0xb7, 0xc2, 0xa2, 0xff, 0xc3, 0x13,
	0xe4, 0xca, 0x01, 0x09, 0x36, 0x4e, 0xbf, 0xff, 0x6f, 0xbf, 0xfe, 0xc5, 0xf5, 0x1b, 0x4b, 0xd9,
	0x75, 0x2a, 0xf2, 0x6b, 0x03, 0x47, 0x9e, 0x3e, 0x3c, 0xfe, 0x14, 0x00, 0x00, 0xff, 0xff, 0x65,
	0xeb, 0x41, 0xff, 0xc2, 0x09, 0x00, 0x00,
}
