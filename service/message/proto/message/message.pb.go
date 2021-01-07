//声明proto本版

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/message/message.proto

//服务名

package message

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

type GetMessageByUserIDResponse_Status int32

const (
	GetMessageByUserIDResponse_SUCCESS GetMessageByUserIDResponse_Status = 0
	GetMessageByUserIDResponse_ERROR   GetMessageByUserIDResponse_Status = -1
)

// Enum value maps for GetMessageByUserIDResponse_Status.
var (
	GetMessageByUserIDResponse_Status_name = map[int32]string{
		0:  "SUCCESS",
		-1: "ERROR",
	}
	GetMessageByUserIDResponse_Status_value = map[string]int32{
		"SUCCESS": 0,
		"ERROR":   -1,
	}
)

func (x GetMessageByUserIDResponse_Status) Enum() *GetMessageByUserIDResponse_Status {
	p := new(GetMessageByUserIDResponse_Status)
	*p = x
	return p
}

func (x GetMessageByUserIDResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetMessageByUserIDResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_message_message_proto_enumTypes[0].Descriptor()
}

func (GetMessageByUserIDResponse_Status) Type() protoreflect.EnumType {
	return &file_proto_message_message_proto_enumTypes[0]
}

func (x GetMessageByUserIDResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetMessageByUserIDResponse_Status.Descriptor instead.
func (GetMessageByUserIDResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{2, 0}
}

type GetMessageByCourseIDResponse_Status int32

const (
	GetMessageByCourseIDResponse_SUCCESS GetMessageByCourseIDResponse_Status = 0
	GetMessageByCourseIDResponse_ERROR   GetMessageByCourseIDResponse_Status = -1
)

// Enum value maps for GetMessageByCourseIDResponse_Status.
var (
	GetMessageByCourseIDResponse_Status_name = map[int32]string{
		0:  "SUCCESS",
		-1: "ERROR",
	}
	GetMessageByCourseIDResponse_Status_value = map[string]int32{
		"SUCCESS": 0,
		"ERROR":   -1,
	}
)

func (x GetMessageByCourseIDResponse_Status) Enum() *GetMessageByCourseIDResponse_Status {
	p := new(GetMessageByCourseIDResponse_Status)
	*p = x
	return p
}

func (x GetMessageByCourseIDResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetMessageByCourseIDResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_message_message_proto_enumTypes[1].Descriptor()
}

func (GetMessageByCourseIDResponse_Status) Type() protoreflect.EnumType {
	return &file_proto_message_message_proto_enumTypes[1]
}

func (x GetMessageByCourseIDResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetMessageByCourseIDResponse_Status.Descriptor instead.
func (GetMessageByCourseIDResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{3, 0}
}

type GetMessageByUserIDParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int32 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetMessageByUserIDParam) Reset() {
	*x = GetMessageByUserIDParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageByUserIDParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageByUserIDParam) ProtoMessage() {}

func (x *GetMessageByUserIDParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageByUserIDParam.ProtoReflect.Descriptor instead.
func (*GetMessageByUserIDParam) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{0}
}

func (x *GetMessageByUserIDParam) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type GetMessageByCourseIDParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseID int32 `protobuf:"varint,1,opt,name=courseID,proto3" json:"courseID,omitempty"`
}

func (x *GetMessageByCourseIDParam) Reset() {
	*x = GetMessageByCourseIDParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageByCourseIDParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageByCourseIDParam) ProtoMessage() {}

func (x *GetMessageByCourseIDParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageByCourseIDParam.ProtoReflect.Descriptor instead.
func (*GetMessageByCourseIDParam) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{1}
}

func (x *GetMessageByCourseIDParam) GetCourseID() int32 {
	if x != nil {
		return x.CourseID
	}
	return 0
}

type GetMessageByUserIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status GetMessageByUserIDResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=go.micro.service.message.GetMessageByUserIDResponse_Status" json:"status,omitempty"`
	Msg    string                            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data   []*Message                        `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetMessageByUserIDResponse) Reset() {
	*x = GetMessageByUserIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageByUserIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageByUserIDResponse) ProtoMessage() {}

func (x *GetMessageByUserIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageByUserIDResponse.ProtoReflect.Descriptor instead.
func (*GetMessageByUserIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{2}
}

func (x *GetMessageByUserIDResponse) GetStatus() GetMessageByUserIDResponse_Status {
	if x != nil {
		return x.Status
	}
	return GetMessageByUserIDResponse_SUCCESS
}

func (x *GetMessageByUserIDResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetMessageByUserIDResponse) GetData() []*Message {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetMessageByCourseIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status GetMessageByCourseIDResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=go.micro.service.message.GetMessageByCourseIDResponse_Status" json:"status,omitempty"`
	Msg    string                              `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data   []*Message                          `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetMessageByCourseIDResponse) Reset() {
	*x = GetMessageByCourseIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageByCourseIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageByCourseIDResponse) ProtoMessage() {}

func (x *GetMessageByCourseIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageByCourseIDResponse.ProtoReflect.Descriptor instead.
func (*GetMessageByCourseIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{3}
}

func (x *GetMessageByCourseIDResponse) GetStatus() GetMessageByCourseIDResponse_Status {
	if x != nil {
		return x.Status
	}
	return GetMessageByCourseIDResponse_SUCCESS
}

func (x *GetMessageByCourseIDResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetMessageByCourseIDResponse) GetData() []*Message {
	if x != nil {
		return x.Data
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageID   int32  `protobuf:"varint,1,opt,name=messageID,proto3" json:"messageID,omitempty"`
	MessageTime int64  `protobuf:"varint,2,opt,name=messageTime,proto3" json:"messageTime,omitempty"`
	MessageType int32  `protobuf:"varint,3,opt,name=messageType,proto3" json:"messageType,omitempty"`
	UserID      int32  `protobuf:"varint,4,opt,name=userID,proto3" json:"userID,omitempty"`
	CourseID    int32  `protobuf:"varint,5,opt,name=courseID,proto3" json:"courseID,omitempty"`
	Title       string `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	Content     string `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	State       int32  `protobuf:"varint,8,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{4}
}

func (x *Message) GetMessageID() int32 {
	if x != nil {
		return x.MessageID
	}
	return 0
}

func (x *Message) GetMessageTime() int64 {
	if x != nil {
		return x.MessageTime
	}
	return 0
}

func (x *Message) GetMessageType() int32 {
	if x != nil {
		return x.MessageType
	}
	return 0
}

func (x *Message) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Message) GetCourseID() int32 {
	if x != nil {
		return x.CourseID
	}
	return 0
}

func (x *Message) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

var File_proto_message_message_proto protoreflect.FileDescriptor

var file_proto_message_message_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x31, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x37, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x49, 0x44, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x49, 0x44, 0x22, 0xe5, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x53, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x29, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x22, 0xe9, 0x01, 0x0a, 0x1c,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3d, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x29, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x22, 0xe5, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x32,
	0x99, 0x02, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x7f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x31, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x34, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x85, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44, 0x12, 0x33, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x1a, 0x36, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_message_message_proto_rawDescOnce sync.Once
	file_proto_message_message_proto_rawDescData = file_proto_message_message_proto_rawDesc
)

func file_proto_message_message_proto_rawDescGZIP() []byte {
	file_proto_message_message_proto_rawDescOnce.Do(func() {
		file_proto_message_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_message_proto_rawDescData)
	})
	return file_proto_message_message_proto_rawDescData
}

var file_proto_message_message_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_message_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_message_message_proto_goTypes = []interface{}{
	(GetMessageByUserIDResponse_Status)(0),   // 0: go.micro.service.message.GetMessageByUserIDResponse.Status
	(GetMessageByCourseIDResponse_Status)(0), // 1: go.micro.service.message.GetMessageByCourseIDResponse.Status
	(*GetMessageByUserIDParam)(nil),          // 2: go.micro.service.message.GetMessageByUserIDParam
	(*GetMessageByCourseIDParam)(nil),        // 3: go.micro.service.message.GetMessageByCourseIDParam
	(*GetMessageByUserIDResponse)(nil),       // 4: go.micro.service.message.GetMessageByUserIDResponse
	(*GetMessageByCourseIDResponse)(nil),     // 5: go.micro.service.message.GetMessageByCourseIDResponse
	(*Message)(nil),                          // 6: go.micro.service.message.Message
}
var file_proto_message_message_proto_depIdxs = []int32{
	0, // 0: go.micro.service.message.GetMessageByUserIDResponse.status:type_name -> go.micro.service.message.GetMessageByUserIDResponse.Status
	6, // 1: go.micro.service.message.GetMessageByUserIDResponse.data:type_name -> go.micro.service.message.Message
	1, // 2: go.micro.service.message.GetMessageByCourseIDResponse.status:type_name -> go.micro.service.message.GetMessageByCourseIDResponse.Status
	6, // 3: go.micro.service.message.GetMessageByCourseIDResponse.data:type_name -> go.micro.service.message.Message
	2, // 4: go.micro.service.message.MessageService.GetMessageByUserID:input_type -> go.micro.service.message.GetMessageByUserIDParam
	3, // 5: go.micro.service.message.MessageService.GetMessageByCourseID:input_type -> go.micro.service.message.GetMessageByCourseIDParam
	4, // 6: go.micro.service.message.MessageService.GetMessageByUserID:output_type -> go.micro.service.message.GetMessageByUserIDResponse
	5, // 7: go.micro.service.message.MessageService.GetMessageByCourseID:output_type -> go.micro.service.message.GetMessageByCourseIDResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_message_message_proto_init() }
func file_proto_message_message_proto_init() {
	if File_proto_message_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_message_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageByUserIDParam); i {
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
		file_proto_message_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageByCourseIDParam); i {
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
		file_proto_message_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageByUserIDResponse); i {
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
		file_proto_message_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageByCourseIDResponse); i {
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
		file_proto_message_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_proto_message_message_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_message_message_proto_goTypes,
		DependencyIndexes: file_proto_message_message_proto_depIdxs,
		EnumInfos:         file_proto_message_message_proto_enumTypes,
		MessageInfos:      file_proto_message_message_proto_msgTypes,
	}.Build()
	File_proto_message_message_proto = out.File
	file_proto_message_message_proto_rawDesc = nil
	file_proto_message_message_proto_goTypes = nil
	file_proto_message_message_proto_depIdxs = nil
}
