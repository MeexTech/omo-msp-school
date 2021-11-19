// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/school/student.proto

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

type CustodianInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Identify             string   `protobuf:"bytes,2,opt,name=identify,proto3" json:"identify,omitempty"`
	Phones               []string `protobuf:"bytes,3,rep,name=phones,proto3" json:"phones,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustodianInfo) Reset()         { *m = CustodianInfo{} }
func (m *CustodianInfo) String() string { return proto.CompactTextString(m) }
func (*CustodianInfo) ProtoMessage()    {}
func (*CustodianInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f522acf3f8a748b, []int{0}
}

func (m *CustodianInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustodianInfo.Unmarshal(m, b)
}
func (m *CustodianInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustodianInfo.Marshal(b, m, deterministic)
}
func (m *CustodianInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustodianInfo.Merge(m, src)
}
func (m *CustodianInfo) XXX_Size() int {
	return xxx_messageInfo_CustodianInfo.Size(m)
}
func (m *CustodianInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CustodianInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CustodianInfo proto.InternalMessageInfo

func (m *CustodianInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustodianInfo) GetIdentify() string {
	if m != nil {
		return m.Identify
	}
	return ""
}

func (m *CustodianInfo) GetPhones() []string {
	if m != nil {
		return m.Phones
	}
	return nil
}

type StudentInfo struct {
	Id                   uint64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string           `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string           `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Created              uint64           `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              uint64           `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string           `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string           `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Sn                   string           `protobuf:"bytes,8,opt,name=sn,proto3" json:"sn,omitempty"`
	Card                 string           `protobuf:"bytes,9,opt,name=card,proto3" json:"card,omitempty"`
	Sex                  uint32           `protobuf:"varint,10,opt,name=sex,proto3" json:"sex,omitempty"`
	Enrol                string           `protobuf:"bytes,11,opt,name=enrol,proto3" json:"enrol,omitempty"`
	Entity               string           `protobuf:"bytes,12,opt,name=entity,proto3" json:"entity,omitempty"`
	Sid                  string           `protobuf:"bytes,13,opt,name=sid,proto3" json:"sid,omitempty"`
	Custodians           []*CustodianInfo `protobuf:"bytes,14,rep,name=custodians,proto3" json:"custodians,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *StudentInfo) Reset()         { *m = StudentInfo{} }
func (m *StudentInfo) String() string { return proto.CompactTextString(m) }
func (*StudentInfo) ProtoMessage()    {}
func (*StudentInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f522acf3f8a748b, []int{1}
}

func (m *StudentInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StudentInfo.Unmarshal(m, b)
}
func (m *StudentInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StudentInfo.Marshal(b, m, deterministic)
}
func (m *StudentInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StudentInfo.Merge(m, src)
}
func (m *StudentInfo) XXX_Size() int {
	return xxx_messageInfo_StudentInfo.Size(m)
}
func (m *StudentInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StudentInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StudentInfo proto.InternalMessageInfo

func (m *StudentInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StudentInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *StudentInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StudentInfo) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *StudentInfo) GetUpdated() uint64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *StudentInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *StudentInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *StudentInfo) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *StudentInfo) GetCard() string {
	if m != nil {
		return m.Card
	}
	return ""
}

func (m *StudentInfo) GetSex() uint32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *StudentInfo) GetEnrol() string {
	if m != nil {
		return m.Enrol
	}
	return ""
}

func (m *StudentInfo) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *StudentInfo) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *StudentInfo) GetCustodians() []*CustodianInfo {
	if m != nil {
		return m.Custodians
	}
	return nil
}

type ReqStudentAdd struct {
	Owner                string           `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name                 string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Sn                   string           `protobuf:"bytes,5,opt,name=sn,proto3" json:"sn,omitempty"`
	Card                 string           `protobuf:"bytes,6,opt,name=card,proto3" json:"card,omitempty"`
	Sex                  uint32           `protobuf:"varint,7,opt,name=sex,proto3" json:"sex,omitempty"`
	Operator             string           `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
	Enrol                string           `protobuf:"bytes,9,opt,name=enrol,proto3" json:"enrol,omitempty"`
	Class                string           `protobuf:"bytes,10,opt,name=class,proto3" json:"class,omitempty"`
	Custodians           []*CustodianInfo `protobuf:"bytes,11,rep,name=custodians,proto3" json:"custodians,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReqStudentAdd) Reset()         { *m = ReqStudentAdd{} }
func (m *ReqStudentAdd) String() string { return proto.CompactTextString(m) }
func (*ReqStudentAdd) ProtoMessage()    {}
func (*ReqStudentAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f522acf3f8a748b, []int{2}
}

func (m *ReqStudentAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqStudentAdd.Unmarshal(m, b)
}
func (m *ReqStudentAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqStudentAdd.Marshal(b, m, deterministic)
}
func (m *ReqStudentAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqStudentAdd.Merge(m, src)
}
func (m *ReqStudentAdd) XXX_Size() int {
	return xxx_messageInfo_ReqStudentAdd.Size(m)
}
func (m *ReqStudentAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqStudentAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqStudentAdd proto.InternalMessageInfo

func (m *ReqStudentAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqStudentAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqStudentAdd) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *ReqStudentAdd) GetCard() string {
	if m != nil {
		return m.Card
	}
	return ""
}

func (m *ReqStudentAdd) GetSex() uint32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *ReqStudentAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqStudentAdd) GetEnrol() string {
	if m != nil {
		return m.Enrol
	}
	return ""
}

func (m *ReqStudentAdd) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *ReqStudentAdd) GetCustodians() []*CustodianInfo {
	if m != nil {
		return m.Custodians
	}
	return nil
}

type ReqStudentUpdate struct {
	Owner                string           `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Uid                  string           `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string           `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Card                 string           `protobuf:"bytes,4,opt,name=card,proto3" json:"card,omitempty"`
	Sex                  uint32           `protobuf:"varint,5,opt,name=sex,proto3" json:"sex,omitempty"`
	Sn                   string           `protobuf:"bytes,6,opt,name=sn,proto3" json:"sn,omitempty"`
	Operator             string           `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Enrol                string           `protobuf:"bytes,9,opt,name=enrol,proto3" json:"enrol,omitempty"`
	Custodians           []*CustodianInfo `protobuf:"bytes,10,rep,name=custodians,proto3" json:"custodians,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReqStudentUpdate) Reset()         { *m = ReqStudentUpdate{} }
func (m *ReqStudentUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqStudentUpdate) ProtoMessage()    {}
func (*ReqStudentUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f522acf3f8a748b, []int{3}
}

func (m *ReqStudentUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqStudentUpdate.Unmarshal(m, b)
}
func (m *ReqStudentUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqStudentUpdate.Marshal(b, m, deterministic)
}
func (m *ReqStudentUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqStudentUpdate.Merge(m, src)
}
func (m *ReqStudentUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqStudentUpdate.Size(m)
}
func (m *ReqStudentUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqStudentUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqStudentUpdate proto.InternalMessageInfo

func (m *ReqStudentUpdate) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqStudentUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqStudentUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqStudentUpdate) GetCard() string {
	if m != nil {
		return m.Card
	}
	return ""
}

func (m *ReqStudentUpdate) GetSex() uint32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *ReqStudentUpdate) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *ReqStudentUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqStudentUpdate) GetEnrol() string {
	if m != nil {
		return m.Enrol
	}
	return ""
}

func (m *ReqStudentUpdate) GetCustodians() []*CustodianInfo {
	if m != nil {
		return m.Custodians
	}
	return nil
}

type ReqStudentBind struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Entity               string   `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	Uid                  string   `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Card                 string   `protobuf:"bytes,4,opt,name=card,proto3" json:"card,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqStudentBind) Reset()         { *m = ReqStudentBind{} }
func (m *ReqStudentBind) String() string { return proto.CompactTextString(m) }
func (*ReqStudentBind) ProtoMessage()    {}
func (*ReqStudentBind) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f522acf3f8a748b, []int{4}
}

func (m *ReqStudentBind) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqStudentBind.Unmarshal(m, b)
}
func (m *ReqStudentBind) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqStudentBind.Marshal(b, m, deterministic)
}
func (m *ReqStudentBind) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqStudentBind.Merge(m, src)
}
func (m *ReqStudentBind) XXX_Size() int {
	return xxx_messageInfo_ReqStudentBind.Size(m)
}
func (m *ReqStudentBind) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqStudentBind.DiscardUnknown(m)
}

var xxx_messageInfo_ReqStudentBind proto.InternalMessageInfo

func (m *ReqStudentBind) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqStudentBind) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *ReqStudentBind) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqStudentBind) GetCard() string {
	if m != nil {
		return m.Card
	}
	return ""
}

func (m *ReqStudentBind) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReqStudentBatch struct {
	Owner                string           `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Operator             string           `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	List                 []*ReqStudentAdd `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReqStudentBatch) Reset()         { *m = ReqStudentBatch{} }
func (m *ReqStudentBatch) String() string { return proto.CompactTextString(m) }
func (*ReqStudentBatch) ProtoMessage()    {}
func (*ReqStudentBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f522acf3f8a748b, []int{5}
}

func (m *ReqStudentBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqStudentBatch.Unmarshal(m, b)
}
func (m *ReqStudentBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqStudentBatch.Marshal(b, m, deterministic)
}
func (m *ReqStudentBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqStudentBatch.Merge(m, src)
}
func (m *ReqStudentBatch) XXX_Size() int {
	return xxx_messageInfo_ReqStudentBatch.Size(m)
}
func (m *ReqStudentBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqStudentBatch.DiscardUnknown(m)
}

var xxx_messageInfo_ReqStudentBatch proto.InternalMessageInfo

func (m *ReqStudentBatch) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqStudentBatch) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqStudentBatch) GetList() []*ReqStudentAdd {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyStudentInfo struct {
	Info                 *StudentInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyStudentInfo) Reset()         { *m = ReplyStudentInfo{} }
func (m *ReplyStudentInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyStudentInfo) ProtoMessage()    {}
func (*ReplyStudentInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f522acf3f8a748b, []int{6}
}

func (m *ReplyStudentInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStudentInfo.Unmarshal(m, b)
}
func (m *ReplyStudentInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStudentInfo.Marshal(b, m, deterministic)
}
func (m *ReplyStudentInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStudentInfo.Merge(m, src)
}
func (m *ReplyStudentInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyStudentInfo.Size(m)
}
func (m *ReplyStudentInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStudentInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStudentInfo proto.InternalMessageInfo

func (m *ReplyStudentInfo) GetInfo() *StudentInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyStudentInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyStudentList struct {
	Total                uint32         `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Page                 uint32         `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32         `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*StudentInfo `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyStudentList) Reset()         { *m = ReplyStudentList{} }
func (m *ReplyStudentList) String() string { return proto.CompactTextString(m) }
func (*ReplyStudentList) ProtoMessage()    {}
func (*ReplyStudentList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f522acf3f8a748b, []int{7}
}

func (m *ReplyStudentList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStudentList.Unmarshal(m, b)
}
func (m *ReplyStudentList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStudentList.Marshal(b, m, deterministic)
}
func (m *ReplyStudentList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStudentList.Merge(m, src)
}
func (m *ReplyStudentList) XXX_Size() int {
	return xxx_messageInfo_ReplyStudentList.Size(m)
}
func (m *ReplyStudentList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStudentList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStudentList proto.InternalMessageInfo

func (m *ReplyStudentList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyStudentList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyStudentList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyStudentList) GetList() []*StudentInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyStudentList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*CustodianInfo)(nil), "omo.msp.school.CustodianInfo")
	proto.RegisterType((*StudentInfo)(nil), "omo.msp.school.StudentInfo")
	proto.RegisterType((*ReqStudentAdd)(nil), "omo.msp.school.ReqStudentAdd")
	proto.RegisterType((*ReqStudentUpdate)(nil), "omo.msp.school.ReqStudentUpdate")
	proto.RegisterType((*ReqStudentBind)(nil), "omo.msp.school.ReqStudentBind")
	proto.RegisterType((*ReqStudentBatch)(nil), "omo.msp.school.ReqStudentBatch")
	proto.RegisterType((*ReplyStudentInfo)(nil), "omo.msp.school.ReplyStudentInfo")
	proto.RegisterType((*ReplyStudentList)(nil), "omo.msp.school.ReplyStudentList")
}

func init() {
	proto.RegisterFile("proto/school/student.proto", fileDescriptor_1f522acf3f8a748b)
}

var fileDescriptor_1f522acf3f8a748b = []byte{
	// 731 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xc1, 0x6e, 0xdb, 0x38,
	0x14, 0x8c, 0x65, 0x59, 0x8e, 0x9f, 0xd7, 0xde, 0x40, 0x08, 0x16, 0x8a, 0x83, 0xec, 0x1a, 0x3a,
	0xf9, 0xe4, 0x60, 0x93, 0x63, 0xd1, 0x43, 0x52, 0x04, 0x41, 0xda, 0xa2, 0x4d, 0x19, 0x14, 0x05,
	0x7a, 0x09, 0x14, 0x91, 0x89, 0x09, 0x58, 0xa4, 0x22, 0xd2, 0x69, 0x72, 0xeb, 0x27, 0xf5, 0xdc,
	0xcf, 0xe9, 0x3f, 0xb4, 0xe7, 0x82, 0x8f, 0xb2, 0x2d, 0xd9, 0x96, 0xeb, 0xdc, 0x38, 0xe4, 0xe3,
	0x70, 0x66, 0x1e, 0x45, 0x08, 0x7a, 0x69, 0x26, 0xb5, 0x3c, 0x54, 0xf1, 0x48, 0xca, 0xf1, 0xa1,
	0xd2, 0x13, 0xca, 0x84, 0x1e, 0xe2, 0xa4, 0xdf, 0x95, 0x89, 0x1c, 0x26, 0x2a, 0x1d, 0xda, 0xd5,
	0xde, 0x5e, 0xa9, 0x36, 0x96, 0x49, 0x22, 0x85, 0x2d, 0x0d, 0x3f, 0x41, 0xe7, 0xd5, 0x44, 0x69,
	0x49, 0x79, 0x24, 0x2e, 0xc4, 0xad, 0xf4, 0x7d, 0x70, 0x45, 0x94, 0xb0, 0xa0, 0xd6, 0xaf, 0x0d,
	0x5a, 0x04, 0xc7, 0x7e, 0x0f, 0xb6, 0xb9, 0xa1, 0xe7, 0xb7, 0x4f, 0x81, 0x83, 0xf3, 0x33, 0xec,
	0xff, 0x03, 0x5e, 0x3a, 0x92, 0x82, 0xa9, 0xa0, 0xde, 0xaf, 0x0f, 0x5a, 0x24, 0x47, 0xe1, 0x0f,
	0x07, 0xda, 0x57, 0x56, 0x15, 0xf2, 0x76, 0xc1, 0xe1, 0x14, 0x59, 0x5d, 0xe2, 0x70, 0xea, 0xef,
	0x40, 0x7d, 0xc2, 0x69, 0x4e, 0x67, 0x86, 0xb3, 0x93, 0xeb, 0x85, 0x93, 0x03, 0x68, 0xc6, 0x19,
	0x8b, 0x34, 0xa3, 0x81, 0x8b, 0x5b, 0xa7, 0xd0, 0xac, 0x4c, 0x52, 0x8a, 0x2b, 0x0d, 0xbb, 0x92,
	0xc3, 0xd9, 0x1e, 0x99, 0x05, 0x1e, 0x52, 0x4d, 0xa1, 0xf1, 0x21, 0x53, 0x96, 0xe1, 0x52, 0xd3,
	0xfa, 0x98, 0x62, 0xa3, 0x4f, 0x89, 0x60, 0x1b, 0x67, 0x1d, 0x25, 0x8c, 0x9a, 0x38, 0xca, 0x68,
	0xd0, 0xb2, 0x6a, 0xcc, 0xd8, 0x68, 0x56, 0xec, 0x31, 0x80, 0x7e, 0x6d, 0xd0, 0x21, 0x66, 0xe8,
	0xef, 0x42, 0x83, 0x89, 0x4c, 0x8e, 0x83, 0x36, 0x96, 0x59, 0x60, 0x32, 0x31, 0xe9, 0xe8, 0xa7,
	0xe0, 0x2f, 0x9c, 0xce, 0x11, 0xee, 0xe7, 0x34, 0xe8, 0x58, 0xcf, 0x8a, 0x53, 0xff, 0x25, 0x40,
	0x3c, 0x8d, 0x5f, 0x05, 0xdd, 0x7e, 0x7d, 0xd0, 0x3e, 0x3a, 0x18, 0x96, 0xdb, 0x37, 0x2c, 0x35,
	0x88, 0x14, 0x36, 0x84, 0xbf, 0x6a, 0xd0, 0x21, 0xec, 0x3e, 0xcf, 0xf9, 0x84, 0x52, 0x23, 0x48,
	0x7e, 0x11, 0x2c, 0xcb, 0xfb, 0x67, 0xc1, 0x2c, 0x5a, 0xa7, 0x10, 0xad, 0x35, 0xdc, 0x58, 0x32,
	0xec, 0x2d, 0x1b, 0x6e, 0xce, 0x0d, 0x17, 0x23, 0xdc, 0x5e, 0x88, 0x70, 0x16, 0x46, 0xab, 0x18,
	0xc6, 0x2e, 0x34, 0xe2, 0x71, 0xa4, 0x14, 0xc6, 0xd6, 0x22, 0x16, 0x2c, 0x18, 0x6f, 0x3f, 0xd7,
	0xf8, 0xcf, 0x1a, 0xec, 0xcc, 0x8d, 0x7f, 0xc4, 0xce, 0x57, 0x78, 0xdf, 0xec, 0xa2, 0x4d, 0xdd,
	0xbb, 0xcb, 0xee, 0x1b, 0x73, 0xf7, 0x36, 0x33, 0x6f, 0x96, 0xd9, 0xba, 0x0b, 0xb5, 0x3a, 0x8d,
	0xb2, 0x6f, 0x78, 0xae, 0xef, 0xaf, 0x35, 0xe8, 0xce, 0x7d, 0x9f, 0x72, 0x51, 0xd5, 0xf1, 0xf9,
	0x15, 0x74, 0x16, 0xaf, 0xa0, 0x49, 0xa3, 0x5e, 0x4a, 0x63, 0xc9, 0x79, 0xd1, 0x57, 0xa3, 0xec,
	0x2b, 0x7c, 0x80, 0xbf, 0x0b, 0x0a, 0x22, 0x1d, 0x8f, 0x2a, 0x24, 0x14, 0x49, 0x9c, 0x85, 0x70,
	0xfe, 0x07, 0x77, 0xcc, 0x95, 0xc6, 0x37, 0x63, 0x45, 0x00, 0xa5, 0x3b, 0x4d, 0xb0, 0x34, 0x7c,
	0x34, 0x1d, 0x4f, 0xc7, 0x4f, 0xc5, 0x47, 0xe5, 0x10, 0x5c, 0x2e, 0x6e, 0x25, 0x9e, 0xdb, 0x3e,
	0xda, 0x5f, 0xa4, 0x29, 0x94, 0x12, 0x2c, 0xf4, 0x8f, 0xc1, 0x53, 0x3a, 0xd2, 0x13, 0x85, 0x8a,
	0x56, 0x6c, 0xc9, 0x8f, 0x30, 0x25, 0x24, 0x2f, 0x0d, 0xbf, 0xd7, 0xca, 0x47, 0xbf, 0xe5, 0x4a,
	0x1b, 0xcf, 0x5a, 0xea, 0x68, 0x8c, 0x67, 0x77, 0x88, 0x05, 0x26, 0xcc, 0x34, 0xba, 0xb3, 0x1f,
	0x5a, 0x87, 0xe0, 0xd8, 0xb4, 0x42, 0x4c, 0x92, 0x1b, 0x96, 0x61, 0xea, 0x1d, 0x92, 0x23, 0x23,
	0x1e, 0x33, 0x70, 0x31, 0x83, 0xf5, 0xe2, 0x4d, 0x61, 0x41, 0x7c, 0x63, 0x63, 0xf1, 0x47, 0xdf,
	0x5c, 0xe8, 0xe6, 0x54, 0x57, 0x2c, 0x7b, 0xe0, 0x31, 0xf3, 0xdf, 0x80, 0x77, 0x42, 0xe9, 0x7b,
	0xc1, 0xfc, 0xf5, 0xc1, 0xf7, 0xfa, 0x15, 0x07, 0xcc, 0x84, 0x85, 0x5b, 0xfe, 0x05, 0x78, 0xe7,
	0x4c, 0x1b, 0xb2, 0x15, 0x72, 0xee, 0x27, 0x4c, 0x61, 0xe1, 0x46, 0x54, 0xaf, 0xa1, 0x79, 0xce,
	0x6c, 0xba, 0x55, 0x5c, 0x97, 0xd1, 0x1d, 0x5b, 0xcf, 0x65, 0xb6, 0x87, 0x5b, 0xfe, 0x07, 0x68,
	0xd9, 0x57, 0xc1, 0x28, 0xeb, 0x57, 0xdb, 0xb4, 0x45, 0x1b, 0xc9, 0x3b, 0x83, 0x16, 0x61, 0x89,
	0x7c, 0x60, 0x7f, 0x34, 0xbb, 0xb7, 0x92, 0x2d, 0xa7, 0x79, 0x07, 0xcd, 0x13, 0x4a, 0x4f, 0xa3,
	0x78, 0xe4, 0xff, 0x57, 0xad, 0x0b, 0x3f, 0xac, 0x8d, 0x9c, 0x5e, 0x02, 0x98, 0x77, 0xe0, 0xcc,
	0x7e, 0xdf, 0xff, 0xae, 0xa1, 0xe4, 0x62, 0xa3, 0x96, 0x9e, 0x1e, 0x7c, 0xde, 0x2f, 0xfe, 0x30,
	0xbc, 0x90, 0x89, 0xbc, 0x4e, 0x54, 0x7a, 0x6d, 0xe1, 0x8d, 0x87, 0x8b, 0xc7, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x60, 0x96, 0x5f, 0xc7, 0x82, 0x08, 0x00, 0x00,
}
