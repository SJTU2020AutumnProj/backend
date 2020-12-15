//声明proto本版

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/auth/auth.proto

//服务名

package auth

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

type LoginResponse_Status int32

const (
	LoginResponse_SUCCESS LoginResponse_Status = 0
	LoginResponse_ERROR   LoginResponse_Status = -1
)

// Enum value maps for LoginResponse_Status.
var (
	LoginResponse_Status_name = map[int32]string{
		0:  "SUCCESS",
		-1: "ERROR",
	}
	LoginResponse_Status_value = map[string]int32{
		"SUCCESS": 0,
		"ERROR":   -1,
	}
)

func (x LoginResponse_Status) Enum() *LoginResponse_Status {
	p := new(LoginResponse_Status)
	*p = x
	return p
}

func (x LoginResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_auth_auth_proto_enumTypes[0].Descriptor()
}

func (LoginResponse_Status) Type() protoreflect.EnumType {
	return &file_proto_auth_auth_proto_enumTypes[0]
}

func (x LoginResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginResponse_Status.Descriptor instead.
func (LoginResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_auth_auth_proto_rawDescGZIP(), []int{3, 0}
}

type CheckAuthResponse_Status int32

const (
	CheckAuthResponse_SUCCESS       CheckAuthResponse_Status = 0
	CheckAuthResponse_INVALID_TOKEN CheckAuthResponse_Status = -1
)

// Enum value maps for CheckAuthResponse_Status.
var (
	CheckAuthResponse_Status_name = map[int32]string{
		0:  "SUCCESS",
		-1: "INVALID_TOKEN",
	}
	CheckAuthResponse_Status_value = map[string]int32{
		"SUCCESS":       0,
		"INVALID_TOKEN": -1,
	}
)

func (x CheckAuthResponse_Status) Enum() *CheckAuthResponse_Status {
	p := new(CheckAuthResponse_Status)
	*p = x
	return p
}

func (x CheckAuthResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CheckAuthResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_auth_auth_proto_enumTypes[1].Descriptor()
}

func (CheckAuthResponse_Status) Type() protoreflect.EnumType {
	return &file_proto_auth_auth_proto_enumTypes[1]
}

func (x CheckAuthResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CheckAuthResponse_Status.Descriptor instead.
func (CheckAuthResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_auth_auth_proto_rawDescGZIP(), []int{5, 0}
}

type LoginParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginParam) Reset() {
	*x = LoginParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginParam) ProtoMessage() {}

func (x *LoginParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginParam.ProtoReflect.Descriptor instead.
func (*LoginParam) Descriptor() ([]byte, []int) {
	return file_proto_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *LoginParam) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *LoginParam) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CheckAuthParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CheckAuthParam) Reset() {
	*x = CheckAuthParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAuthParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAuthParam) ProtoMessage() {}

func (x *CheckAuthParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAuthParam.ProtoReflect.Descriptor instead.
func (*CheckAuthParam) Descriptor() ([]byte, []int) {
	return file_proto_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *CheckAuthParam) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   int32  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	UserType int32  `protobuf:"varint,2,opt,name=userType,proto3" json:"userType,omitempty"`
	UserName string `protobuf:"bytes,3,opt,name=userName,proto3" json:"userName,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	School   string `protobuf:"bytes,5,opt,name=school,proto3" json:"school,omitempty"`
	Id       string `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	Phone    string `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Email    string `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_proto_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *UserData) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UserData) GetUserType() int32 {
	if x != nil {
		return x.UserType
	}
	return 0
}

func (x *UserData) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserData) GetSchool() string {
	if x != nil {
		return x.School
	}
	return ""
}

func (x *UserData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserData) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status LoginResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=go.micro.service.auth.LoginResponse_Status" json:"status,omitempty"`
	Msg    string               `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data   *UserData            `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Token  string               `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetStatus() LoginResponse_Status {
	if x != nil {
		return x.Status
	}
	return LoginResponse_SUCCESS
}

func (x *LoginResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *LoginResponse) GetData() *UserData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   int32  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	UserType int32  `protobuf:"varint,2,opt,name=userType,proto3" json:"userType,omitempty"`
	UserName string `protobuf:"bytes,3,opt,name=userName,proto3" json:"userName,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthData) Reset() {
	*x = AuthData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthData) ProtoMessage() {}

func (x *AuthData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthData.ProtoReflect.Descriptor instead.
func (*AuthData) Descriptor() ([]byte, []int) {
	return file_proto_auth_auth_proto_rawDescGZIP(), []int{4}
}

func (x *AuthData) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *AuthData) GetUserType() int32 {
	if x != nil {
		return x.UserType
	}
	return 0
}

func (x *AuthData) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *AuthData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CheckAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status CheckAuthResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=go.micro.service.auth.CheckAuthResponse_Status" json:"status,omitempty"`
	Msg    string                   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data   *AuthData                `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CheckAuthResponse) Reset() {
	*x = CheckAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAuthResponse) ProtoMessage() {}

func (x *CheckAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAuthResponse.ProtoReflect.Descriptor instead.
func (*CheckAuthResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_auth_proto_rawDescGZIP(), []int{5}
}

func (x *CheckAuthResponse) GetStatus() CheckAuthResponse_Status {
	if x != nil {
		return x.Status
	}
	return CheckAuthResponse_SUCCESS
}

func (x *CheckAuthResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CheckAuthResponse) GetData() *AuthData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_auth_auth_proto protoreflect.FileDescriptor

var file_proto_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x22, 0x44,
	0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x26, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x75, 0x74,
	0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xca, 0x01, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xdc, 0x01, 0x0a, 0x0d, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x29, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x22, 0x76, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0xd6, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x31, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x1a, 0x0a,
	0x0d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x32, 0xc1, 0x01, 0x0a, 0x0b, 0x41, 0x75,
	0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x24, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a,
	0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x12, 0x25, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x1a, 0x28, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a,
	0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_auth_auth_proto_rawDescOnce sync.Once
	file_proto_auth_auth_proto_rawDescData = file_proto_auth_auth_proto_rawDesc
)

func file_proto_auth_auth_proto_rawDescGZIP() []byte {
	file_proto_auth_auth_proto_rawDescOnce.Do(func() {
		file_proto_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_auth_auth_proto_rawDescData)
	})
	return file_proto_auth_auth_proto_rawDescData
}

var file_proto_auth_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_auth_auth_proto_goTypes = []interface{}{
	(LoginResponse_Status)(0),     // 0: go.micro.service.auth.LoginResponse.Status
	(CheckAuthResponse_Status)(0), // 1: go.micro.service.auth.CheckAuthResponse.Status
	(*LoginParam)(nil),            // 2: go.micro.service.auth.LoginParam
	(*CheckAuthParam)(nil),        // 3: go.micro.service.auth.CheckAuthParam
	(*UserData)(nil),              // 4: go.micro.service.auth.UserData
	(*LoginResponse)(nil),         // 5: go.micro.service.auth.LoginResponse
	(*AuthData)(nil),              // 6: go.micro.service.auth.AuthData
	(*CheckAuthResponse)(nil),     // 7: go.micro.service.auth.CheckAuthResponse
}
var file_proto_auth_auth_proto_depIdxs = []int32{
	0, // 0: go.micro.service.auth.LoginResponse.status:type_name -> go.micro.service.auth.LoginResponse.Status
	4, // 1: go.micro.service.auth.LoginResponse.data:type_name -> go.micro.service.auth.UserData
	1, // 2: go.micro.service.auth.CheckAuthResponse.status:type_name -> go.micro.service.auth.CheckAuthResponse.Status
	6, // 3: go.micro.service.auth.CheckAuthResponse.data:type_name -> go.micro.service.auth.AuthData
	2, // 4: go.micro.service.auth.AuthService.Login:input_type -> go.micro.service.auth.LoginParam
	3, // 5: go.micro.service.auth.AuthService.CheckAuth:input_type -> go.micro.service.auth.CheckAuthParam
	5, // 6: go.micro.service.auth.AuthService.Login:output_type -> go.micro.service.auth.LoginResponse
	7, // 7: go.micro.service.auth.AuthService.CheckAuth:output_type -> go.micro.service.auth.CheckAuthResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_auth_auth_proto_init() }
func file_proto_auth_auth_proto_init() {
	if File_proto_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginParam); i {
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
		file_proto_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAuthParam); i {
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
		file_proto_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_proto_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_proto_auth_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthData); i {
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
		file_proto_auth_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAuthResponse); i {
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
			RawDescriptor: file_proto_auth_auth_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_auth_auth_proto_goTypes,
		DependencyIndexes: file_proto_auth_auth_proto_depIdxs,
		EnumInfos:         file_proto_auth_auth_proto_enumTypes,
		MessageInfos:      file_proto_auth_auth_proto_msgTypes,
	}.Build()
	File_proto_auth_auth_proto = out.File
	file_proto_auth_auth_proto_rawDesc = nil
	file_proto_auth_auth_proto_goTypes = nil
	file_proto_auth_auth_proto_depIdxs = nil
}
