//声明proto本版

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/check/check.proto

//服务名

package check

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateCheckResponse_Status int32

const (
	CreateCheckResponse_SUCCESS CreateCheckResponse_Status = 0
	CreateCheckResponse_ERROR   CreateCheckResponse_Status = -1
)

// Enum value maps for CreateCheckResponse_Status.
var (
	CreateCheckResponse_Status_name = map[int32]string{
		0:  "SUCCESS",
		-1: "ERROR",
	}
	CreateCheckResponse_Status_value = map[string]int32{
		"SUCCESS": 0,
		"ERROR":   -1,
	}
)

func (x CreateCheckResponse_Status) Enum() *CreateCheckResponse_Status {
	p := new(CreateCheckResponse_Status)
	*p = x
	return p
}

func (x CreateCheckResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateCheckResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_check_check_proto_enumTypes[0].Descriptor()
}

func (CreateCheckResponse_Status) Type() protoreflect.EnumType {
	return &file_proto_check_check_proto_enumTypes[0]
}

func (x CreateCheckResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateCheckResponse_Status.Descriptor instead.
func (CreateCheckResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_check_check_proto_rawDescGZIP(), []int{4, 0}
}

type DeleteCheckResponse_Status int32

const (
	DeleteCheckResponse_SUCCESS DeleteCheckResponse_Status = 0
	DeleteCheckResponse_ERROR   DeleteCheckResponse_Status = -1
)

// Enum value maps for DeleteCheckResponse_Status.
var (
	DeleteCheckResponse_Status_name = map[int32]string{
		0:  "SUCCESS",
		-1: "ERROR",
	}
	DeleteCheckResponse_Status_value = map[string]int32{
		"SUCCESS": 0,
		"ERROR":   -1,
	}
)

func (x DeleteCheckResponse_Status) Enum() *DeleteCheckResponse_Status {
	p := new(DeleteCheckResponse_Status)
	*p = x
	return p
}

func (x DeleteCheckResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeleteCheckResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_check_check_proto_enumTypes[1].Descriptor()
}

func (DeleteCheckResponse_Status) Type() protoreflect.EnumType {
	return &file_proto_check_check_proto_enumTypes[1]
}

func (x DeleteCheckResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeleteCheckResponse_Status.Descriptor instead.
func (DeleteCheckResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_check_check_proto_rawDescGZIP(), []int{5, 0}
}

type UpdateCheckResponse_Status int32

const (
	UpdateCheckResponse_SUCCESS UpdateCheckResponse_Status = 0
	UpdateCheckResponse_ERROR   UpdateCheckResponse_Status = -1
)

// Enum value maps for UpdateCheckResponse_Status.
var (
	UpdateCheckResponse_Status_name = map[int32]string{
		0:  "SUCCESS",
		-1: "ERROR",
	}
	UpdateCheckResponse_Status_value = map[string]int32{
		"SUCCESS": 0,
		"ERROR":   -1,
	}
)

func (x UpdateCheckResponse_Status) Enum() *UpdateCheckResponse_Status {
	p := new(UpdateCheckResponse_Status)
	*p = x
	return p
}

func (x UpdateCheckResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateCheckResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_check_check_proto_enumTypes[2].Descriptor()
}

func (UpdateCheckResponse_Status) Type() protoreflect.EnumType {
	return &file_proto_check_check_proto_enumTypes[2]
}

func (x UpdateCheckResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateCheckResponse_Status.Descriptor instead.
func (UpdateCheckResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_check_check_proto_rawDescGZIP(), []int{6, 0}
}

type SearchCheckByIDResponse_Status int32

const (
	SearchCheckByIDResponse_SUCCESS SearchCheckByIDResponse_Status = 0
	SearchCheckByIDResponse_ERROR   SearchCheckByIDResponse_Status = -1
)

// Enum value maps for SearchCheckByIDResponse_Status.
var (
	SearchCheckByIDResponse_Status_name = map[int32]string{
		0:  "SUCCESS",
		-1: "ERROR",
	}
	SearchCheckByIDResponse_Status_value = map[string]int32{
		"SUCCESS": 0,
		"ERROR":   -1,
	}
)

func (x SearchCheckByIDResponse_Status) Enum() *SearchCheckByIDResponse_Status {
	p := new(SearchCheckByIDResponse_Status)
	*p = x
	return p
}

func (x SearchCheckByIDResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchCheckByIDResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_check_check_proto_enumTypes[3].Descriptor()
}

func (SearchCheckByIDResponse_Status) Type() protoreflect.EnumType {
	return &file_proto_check_check_proto_enumTypes[3]
}

func (x SearchCheckByIDResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchCheckByIDResponse_Status.Descriptor instead.
func (SearchCheckByIDResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_check_check_proto_rawDescGZIP(), []int{7, 0}
}

type CheckInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckID     int32  `protobuf:"varint,1,opt,name=checkID,proto3" json:"checkID,omitempty"`
	CheckTime   int64  `protobuf:"varint,2,opt,name=checkTime,proto3" json:"checkTime,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Comment     string `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	Score       int32  `protobuf:"varint,5,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *CheckInfo) Reset() {
	*x = CheckInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_check_check_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInfo) ProtoMessage() {}

func (x *CheckInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_check_check_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInfo.ProtoReflect.Descriptor instead.
func (*CheckInfo) Descriptor() ([]byte, []int) {
	return file_proto_check_check_proto_rawDescGZIP(), []int{0}
}

func (x *CheckInfo) GetCheckID() int32 {
	if x != nil {
		return x.CheckID
	}
	return 0
}

func (x *CheckInfo) GetCheckTime() int64 {
	if x != nil {
		return x.CheckTime
	}
	return 0
}

func (x *CheckInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CheckInfo) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *CheckInfo) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

//Parameter protocol
type CreateCheckParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnswerID    int32  `protobuf:"varint,1,opt,name=answerID,proto3" json:"answerID,omitempty"`
	HomeworkID  int32  `protobuf:"varint,2,opt,name=homeworkID,proto3" json:"homeworkID,omitempty"`
	TeacherID   int32  `protobuf:"varint,3,opt,name=teacherID,proto3" json:"teacherID,omitempty"`
	StudentID   int32  `protobuf:"varint,4,opt,name=studentID,proto3" json:"studentID,omitempty"`
	CheckTime   int64  `protobuf:"varint,5,opt,name=checkTime,proto3" json:"checkTime,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Comment     string `protobuf:"bytes,7,opt,name=comment,proto3" json:"comment,omitempty"`
	Score       int32  `protobuf:"varint,8,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *CreateCheckParam) Reset() {
	*x = CreateCheckParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_check_check_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCheckParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCheckParam) ProtoMessage() {}

func (x *CreateCheckParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_check_check_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCheckParam.ProtoReflect.Descriptor instead.
func (*CreateCheckParam) Descriptor() ([]byte, []int) {
	return file_proto_check_check_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCheckParam) GetAnswerID() int32 {
	if x != nil {
		return x.AnswerID
	}
	return 0
}

func (x *CreateCheckParam) GetHomeworkID() int32 {
	if x != nil {
		return x.HomeworkID
	}
	return 0
}

func (x *CreateCheckParam) GetTeacherID() int32 {
	if x != nil {
		return x.TeacherID
	}
	return 0
}

func (x *CreateCheckParam) GetStudentID() int32 {
	if x != nil {
		return x.StudentID
	}
	return 0
}

func (x *CreateCheckParam) GetCheckTime() int64 {
	if x != nil {
		return x.CheckTime
	}
	return 0
}

func (x *CreateCheckParam) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateCheckParam) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *CreateCheckParam) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type CheckID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckID int32 `protobuf:"varint,1,opt,name=CheckID,proto3" json:"CheckID,omitempty"`
}

func (x *CheckID) Reset() {
	*x = CheckID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_check_check_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckID) ProtoMessage() {}

func (x *CheckID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_check_check_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckID.ProtoReflect.Descriptor instead.
func (*CheckID) Descriptor() ([]byte, []int) {
	return file_proto_check_check_proto_rawDescGZIP(), []int{2}
}

func (x *CheckID) GetCheckID() int32 {
	if x != nil {
		return x.CheckID
	}
	return 0
}

type UpdateCheckParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckID     int32  `protobuf:"varint,1,opt,name=checkID,proto3" json:"checkID,omitempty"`
	AnswerID    int32  `protobuf:"varint,2,opt,name=answerID,proto3" json:"answerID,omitempty"`
	HomeworkID  int32  `protobuf:"varint,3,opt,name=homeworkID,proto3" json:"homeworkID,omitempty"`
	TeacherID   int32  `protobuf:"varint,4,opt,name=teacherID,proto3" json:"teacherID,omitempty"`
	StudentID   int32  `protobuf:"varint,5,opt,name=studentID,proto3" json:"studentID,omitempty"`
	CheckTime   int64  `protobuf:"varint,6,opt,name=checkTime,proto3" json:"checkTime,omitempty"`
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Comment     string `protobuf:"bytes,8,opt,name=comment,proto3" json:"comment,omitempty"`
	Score       int32  `protobuf:"varint,9,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *UpdateCheckParam) Reset() {
	*x = UpdateCheckParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_check_check_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCheckParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCheckParam) ProtoMessage() {}

func (x *UpdateCheckParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_check_check_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCheckParam.ProtoReflect.Descriptor instead.
func (*UpdateCheckParam) Descriptor() ([]byte, []int) {
	return file_proto_check_check_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCheckParam) GetCheckID() int32 {
	if x != nil {
		return x.CheckID
	}
	return 0
}

func (x *UpdateCheckParam) GetAnswerID() int32 {
	if x != nil {
		return x.AnswerID
	}
	return 0
}

func (x *UpdateCheckParam) GetHomeworkID() int32 {
	if x != nil {
		return x.HomeworkID
	}
	return 0
}

func (x *UpdateCheckParam) GetTeacherID() int32 {
	if x != nil {
		return x.TeacherID
	}
	return 0
}

func (x *UpdateCheckParam) GetStudentID() int32 {
	if x != nil {
		return x.StudentID
	}
	return 0
}

func (x *UpdateCheckParam) GetCheckTime() int64 {
	if x != nil {
		return x.CheckTime
	}
	return 0
}

func (x *UpdateCheckParam) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateCheckParam) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *UpdateCheckParam) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

//Response protocol
type CreateCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  CreateCheckResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=go.micro.service.check.CreateCheckResponse_Status" json:"status,omitempty"`
	Msg     string                     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	CheckID int32                      `protobuf:"varint,3,opt,name=CheckID,proto3" json:"CheckID,omitempty"`
}

func (x *CreateCheckResponse) Reset() {
	*x = CreateCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_check_check_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCheckResponse) ProtoMessage() {}

func (x *CreateCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_check_check_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCheckResponse.ProtoReflect.Descriptor instead.
func (*CreateCheckResponse) Descriptor() ([]byte, []int) {
	return file_proto_check_check_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCheckResponse) GetStatus() CreateCheckResponse_Status {
	if x != nil {
		return x.Status
	}
	return CreateCheckResponse_SUCCESS
}

func (x *CreateCheckResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CreateCheckResponse) GetCheckID() int32 {
	if x != nil {
		return x.CheckID
	}
	return 0
}

type DeleteCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status DeleteCheckResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=go.micro.service.check.DeleteCheckResponse_Status" json:"status,omitempty"`
	Msg    string                     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DeleteCheckResponse) Reset() {
	*x = DeleteCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_check_check_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCheckResponse) ProtoMessage() {}

func (x *DeleteCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_check_check_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCheckResponse.ProtoReflect.Descriptor instead.
func (*DeleteCheckResponse) Descriptor() ([]byte, []int) {
	return file_proto_check_check_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCheckResponse) GetStatus() DeleteCheckResponse_Status {
	if x != nil {
		return x.Status
	}
	return DeleteCheckResponse_SUCCESS
}

func (x *DeleteCheckResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UpdateCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status UpdateCheckResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=go.micro.service.check.UpdateCheckResponse_Status" json:"status,omitempty"`
	Msg    string                     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UpdateCheckResponse) Reset() {
	*x = UpdateCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_check_check_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCheckResponse) ProtoMessage() {}

func (x *UpdateCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_check_check_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCheckResponse.ProtoReflect.Descriptor instead.
func (*UpdateCheckResponse) Descriptor() ([]byte, []int) {
	return file_proto_check_check_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCheckResponse) GetStatus() UpdateCheckResponse_Status {
	if x != nil {
		return x.Status
	}
	return UpdateCheckResponse_SUCCESS
}

func (x *UpdateCheckResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SearchCheckByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status SearchCheckByIDResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=go.micro.service.check.SearchCheckByIDResponse_Status" json:"status,omitempty"`
	Msg    string                         `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Check  *CheckInfo                     `protobuf:"bytes,3,opt,name=check,proto3" json:"check,omitempty"`
}

func (x *SearchCheckByIDResponse) Reset() {
	*x = SearchCheckByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_check_check_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchCheckByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchCheckByIDResponse) ProtoMessage() {}

func (x *SearchCheckByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_check_check_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchCheckByIDResponse.ProtoReflect.Descriptor instead.
func (*SearchCheckByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_check_check_proto_rawDescGZIP(), []int{7}
}

func (x *SearchCheckByIDResponse) GetStatus() SearchCheckByIDResponse_Status {
	if x != nil {
		return x.Status
	}
	return SearchCheckByIDResponse_SUCCESS
}

func (x *SearchCheckByIDResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SearchCheckByIDResponse) GetCheck() *CheckInfo {
	if x != nil {
		return x.Check
	}
	return nil
}

var File_proto_check_check_proto protoreflect.FileDescriptor

var file_proto_check_check_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x22, 0x95, 0x01, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0xfa, 0x01, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x6f,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x23, 0x0a, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x44, 0x22, 0x94, 0x02, 0x0a, 0x10,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x22, 0xb8, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x44, 0x22, 0x29, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x22, 0x9e, 0x01,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x29, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x05, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x22, 0x9e,
	0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0x29, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x05, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x22,
	0xdf, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x37, 0x0a,
	0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x22, 0x29, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x01, 0x32, 0xa4, 0x03, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x66, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x12, 0x28, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x2b, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x44, 0x1a, 0x2b, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x28, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x1a, 0x2b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x65, 0x0a, 0x0f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x42, 0x79, 0x49, 0x44, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x49, 0x44, 0x1a, 0x2f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_check_check_proto_rawDescOnce sync.Once
	file_proto_check_check_proto_rawDescData = file_proto_check_check_proto_rawDesc
)

func file_proto_check_check_proto_rawDescGZIP() []byte {
	file_proto_check_check_proto_rawDescOnce.Do(func() {
		file_proto_check_check_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_check_check_proto_rawDescData)
	})
	return file_proto_check_check_proto_rawDescData
}

var file_proto_check_check_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_proto_check_check_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_check_check_proto_goTypes = []interface{}{
	(CreateCheckResponse_Status)(0),     // 0: go.micro.service.check.CreateCheckResponse.Status
	(DeleteCheckResponse_Status)(0),     // 1: go.micro.service.check.DeleteCheckResponse.Status
	(UpdateCheckResponse_Status)(0),     // 2: go.micro.service.check.UpdateCheckResponse.Status
	(SearchCheckByIDResponse_Status)(0), // 3: go.micro.service.check.SearchCheckByIDResponse.Status
	(*CheckInfo)(nil),                   // 4: go.micro.service.check.CheckInfo
	(*CreateCheckParam)(nil),            // 5: go.micro.service.check.CreateCheckParam
	(*CheckID)(nil),                     // 6: go.micro.service.check.CheckID
	(*UpdateCheckParam)(nil),            // 7: go.micro.service.check.UpdateCheckParam
	(*CreateCheckResponse)(nil),         // 8: go.micro.service.check.CreateCheckResponse
	(*DeleteCheckResponse)(nil),         // 9: go.micro.service.check.DeleteCheckResponse
	(*UpdateCheckResponse)(nil),         // 10: go.micro.service.check.UpdateCheckResponse
	(*SearchCheckByIDResponse)(nil),     // 11: go.micro.service.check.SearchCheckByIDResponse
}
var file_proto_check_check_proto_depIdxs = []int32{
	0,  // 0: go.micro.service.check.CreateCheckResponse.status:type_name -> go.micro.service.check.CreateCheckResponse.Status
	1,  // 1: go.micro.service.check.DeleteCheckResponse.status:type_name -> go.micro.service.check.DeleteCheckResponse.Status
	2,  // 2: go.micro.service.check.UpdateCheckResponse.status:type_name -> go.micro.service.check.UpdateCheckResponse.Status
	3,  // 3: go.micro.service.check.SearchCheckByIDResponse.status:type_name -> go.micro.service.check.SearchCheckByIDResponse.Status
	4,  // 4: go.micro.service.check.SearchCheckByIDResponse.check:type_name -> go.micro.service.check.CheckInfo
	5,  // 5: go.micro.service.check.CheckService.CreateCheck:input_type -> go.micro.service.check.CreateCheckParam
	6,  // 6: go.micro.service.check.CheckService.DeleteCheck:input_type -> go.micro.service.check.CheckID
	7,  // 7: go.micro.service.check.CheckService.UpdateCheck:input_type -> go.micro.service.check.UpdateCheckParam
	6,  // 8: go.micro.service.check.CheckService.SearchCheckByID:input_type -> go.micro.service.check.CheckID
	8,  // 9: go.micro.service.check.CheckService.CreateCheck:output_type -> go.micro.service.check.CreateCheckResponse
	9,  // 10: go.micro.service.check.CheckService.DeleteCheck:output_type -> go.micro.service.check.DeleteCheckResponse
	10, // 11: go.micro.service.check.CheckService.UpdateCheck:output_type -> go.micro.service.check.UpdateCheckResponse
	11, // 12: go.micro.service.check.CheckService.SearchCheckByID:output_type -> go.micro.service.check.SearchCheckByIDResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_check_check_proto_init() }
func file_proto_check_check_proto_init() {
	if File_proto_check_check_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_check_check_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_check_check_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCheckParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_check_check_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_check_check_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCheckParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_check_check_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCheckResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_check_check_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCheckResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_check_check_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCheckResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_check_check_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchCheckByIDResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_check_check_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_check_check_proto_goTypes,
		DependencyIndexes: file_proto_check_check_proto_depIdxs,
		EnumInfos:         file_proto_check_check_proto_enumTypes,
		MessageInfos:      file_proto_check_check_proto_msgTypes,
	}.Build()
	File_proto_check_check_proto = out.File
	file_proto_check_check_proto_rawDesc = nil
	file_proto_check_check_proto_goTypes = nil
	file_proto_check_check_proto_depIdxs = nil
}
