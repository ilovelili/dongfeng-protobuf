// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/ilovelili/dongfeng-protobuf/dongfeng.proto

package dongfeng_protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ******************************************************* message definition *******************************************************
// /////////////////////////////////////////////////
type DashboardRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DashboardRequest) Reset()         { *m = DashboardRequest{} }
func (m *DashboardRequest) String() string { return proto.CompactTextString(m) }
func (*DashboardRequest) ProtoMessage()    {}
func (*DashboardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dongfeng_a8a5cdfedef06db9, []int{0}
}
func (m *DashboardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DashboardRequest.Unmarshal(m, b)
}
func (m *DashboardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DashboardRequest.Marshal(b, m, deterministic)
}
func (dst *DashboardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DashboardRequest.Merge(dst, src)
}
func (m *DashboardRequest) XXX_Size() int {
	return xxx_messageInfo_DashboardRequest.Size(m)
}
func (m *DashboardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DashboardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DashboardRequest proto.InternalMessageInfo

func (m *DashboardRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DashboardResponse struct {
	UserId               string          `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Notifications        []*Notification `protobuf:"bytes,2,rep,name=notifications,proto3" json:"notifications"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DashboardResponse) Reset()         { *m = DashboardResponse{} }
func (m *DashboardResponse) String() string { return proto.CompactTextString(m) }
func (*DashboardResponse) ProtoMessage()    {}
func (*DashboardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dongfeng_a8a5cdfedef06db9, []int{1}
}
func (m *DashboardResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DashboardResponse.Unmarshal(m, b)
}
func (m *DashboardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DashboardResponse.Marshal(b, m, deterministic)
}
func (dst *DashboardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DashboardResponse.Merge(dst, src)
}
func (m *DashboardResponse) XXX_Size() int {
	return xxx_messageInfo_DashboardResponse.Size(m)
}
func (m *DashboardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DashboardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DashboardResponse proto.InternalMessageInfo

func (m *DashboardResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *DashboardResponse) GetNotifications() []*Notification {
	if m != nil {
		return m.Notifications
	}
	return nil
}

// /////////////////////////////////////////////////
type LoginRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dongfeng_a8a5cdfedef06db9, []int{2}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (dst *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(dst, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LoginResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dongfeng_a8a5cdfedef06db9, []int{3}
}
func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (dst *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(dst, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// /////////////////////////////////////////////////
type UploadAvatarResponse struct {
	Uri                  string   `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadAvatarResponse) Reset()         { *m = UploadAvatarResponse{} }
func (m *UploadAvatarResponse) String() string { return proto.CompactTextString(m) }
func (*UploadAvatarResponse) ProtoMessage()    {}
func (*UploadAvatarResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dongfeng_a8a5cdfedef06db9, []int{4}
}
func (m *UploadAvatarResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadAvatarResponse.Unmarshal(m, b)
}
func (m *UploadAvatarResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadAvatarResponse.Marshal(b, m, deterministic)
}
func (dst *UploadAvatarResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadAvatarResponse.Merge(dst, src)
}
func (m *UploadAvatarResponse) XXX_Size() int {
	return xxx_messageInfo_UploadAvatarResponse.Size(m)
}
func (m *UploadAvatarResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadAvatarResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadAvatarResponse proto.InternalMessageInfo

func (m *UploadAvatarResponse) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

// /////////////////////////////////////////////////
type User struct {
	Newuser              bool       `protobuf:"varint,1,opt,name=newuser,proto3" json:"newuser"`
	Role                 string     `protobuf:"bytes,2,opt,name=role,proto3" json:"role"`
	Id                   string     `protobuf:"bytes,3,opt,name=id,proto3" json:"id"`
	Email                string     `protobuf:"bytes,4,opt,name=email,proto3" json:"email"`
	Name                 string     `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	Avatar               string     `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar"`
	Settings             []*Setting `protobuf:"bytes,7,rep,name=settings,proto3" json:"settings"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_dongfeng_a8a5cdfedef06db9, []int{5}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetNewuser() bool {
	if m != nil {
		return m.Newuser
	}
	return false
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *User) GetSettings() []*Setting {
	if m != nil {
		return m.Settings
	}
	return nil
}

// /////////////////////////////////////////////////
type Setting struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Enabled              bool     `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Setting) Reset()         { *m = Setting{} }
func (m *Setting) String() string { return proto.CompactTextString(m) }
func (*Setting) ProtoMessage()    {}
func (*Setting) Descriptor() ([]byte, []int) {
	return fileDescriptor_dongfeng_a8a5cdfedef06db9, []int{6}
}
func (m *Setting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Setting.Unmarshal(m, b)
}
func (m *Setting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Setting.Marshal(b, m, deterministic)
}
func (dst *Setting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Setting.Merge(dst, src)
}
func (m *Setting) XXX_Size() int {
	return xxx_messageInfo_Setting.Size(m)
}
func (m *Setting) XXX_DiscardUnknown() {
	xxx_messageInfo_Setting.DiscardUnknown(m)
}

var xxx_messageInfo_Setting proto.InternalMessageInfo

func (m *Setting) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Setting) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Setting) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

// /////////////////////////////////////////////////
type Notification struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id"`
	CustomCode           string   `protobuf:"bytes,3,opt,name=custom_code,json=customCode,proto3" json:"custom_code"`
	Category             string   `protobuf:"bytes,4,opt,name=category,proto3" json:"category"`
	Details              string   `protobuf:"bytes,5,opt,name=details,proto3" json:"details"`
	Link                 string   `protobuf:"bytes,6,opt,name=link,proto3" json:"link"`
	Time                 string   `protobuf:"bytes,7,opt,name=time,proto3" json:"time"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_dongfeng_a8a5cdfedef06db9, []int{7}
}
func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (dst *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(dst, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Notification) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Notification) GetCustomCode() string {
	if m != nil {
		return m.CustomCode
	}
	return ""
}

func (m *Notification) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Notification) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *Notification) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Notification) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func init() {
	proto.RegisterType((*DashboardRequest)(nil), "dongfeng.protobuf.DashboardRequest")
	proto.RegisterType((*DashboardResponse)(nil), "dongfeng.protobuf.DashboardResponse")
	proto.RegisterType((*LoginRequest)(nil), "dongfeng.protobuf.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "dongfeng.protobuf.LoginResponse")
	proto.RegisterType((*UploadAvatarResponse)(nil), "dongfeng.protobuf.UploadAvatarResponse")
	proto.RegisterType((*User)(nil), "dongfeng.protobuf.User")
	proto.RegisterType((*Setting)(nil), "dongfeng.protobuf.Setting")
	proto.RegisterType((*Notification)(nil), "dongfeng.protobuf.Notification")
}

func init() {
	proto.RegisterFile("github.com/ilovelili/dongfeng-protobuf/dongfeng.proto", fileDescriptor_dongfeng_a8a5cdfedef06db9)
}

var fileDescriptor_dongfeng_a8a5cdfedef06db9 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0x65, 0xd3, 0xdd, 0xcd, 0x7a, 0xb7, 0x95, 0x76, 0x28, 0x76, 0xe8, 0x4b, 0x4b, 0xf0, 0x61,
	0x41, 0xdc, 0x42, 0x45, 0x9f, 0x7c, 0x11, 0x15, 0x11, 0xc4, 0x87, 0x48, 0x9f, 0xcb, 0x24, 0x73,
	0x37, 0x1d, 0x3a, 0x99, 0xbb, 0x66, 0x26, 0x15, 0xff, 0x98, 0x8f, 0xfe, 0x36, 0x99, 0x8f, 0x0d,
	0x59, 0x5a, 0x7c, 0xbb, 0xe7, 0x70, 0x6e, 0xce, 0xb9, 0x27, 0x03, 0x6f, 0x1b, 0xe5, 0xee, 0xfa,
	0x6a, 0x5d, 0x53, 0x7b, 0xa5, 0x34, 0x3d, 0xa0, 0x56, 0x5a, 0x5d, 0x49, 0x32, 0xcd, 0x06, 0x4d,
	0xf3, 0x7a, 0xdb, 0x91, 0xa3, 0xaa, 0xdf, 0x0c, 0xcc, 0x3a, 0x30, 0xec, 0x64, 0x1f, 0x57, 0xfd,
	0xa6, 0x58, 0xc1, 0xf1, 0x27, 0x61, 0xef, 0x2a, 0x12, 0x9d, 0x2c, 0xf1, 0x67, 0x8f, 0xd6, 0xb1,
	0x53, 0x98, 0x39, 0xba, 0x47, 0xc3, 0x27, 0x97, 0x93, 0xd5, 0xb3, 0x32, 0x82, 0xc2, 0xc2, 0xc9,
	0x48, 0x69, 0xb7, 0x64, 0x2c, 0xb2, 0x33, 0xc8, 0x7b, 0x8b, 0xdd, 0xad, 0x92, 0x49, 0x3c, 0xf7,
	0xf0, 0xab, 0x64, 0x9f, 0xe1, 0xc8, 0x90, 0x53, 0x1b, 0x55, 0x0b, 0xa7, 0xc8, 0x58, 0x9e, 0x5d,
	0x1e, 0xac, 0x96, 0xd7, 0x17, 0xeb, 0x47, 0x11, 0xd6, 0xdf, 0x47, 0xba, 0x72, 0x7f, 0xab, 0x78,
	0x09, 0x87, 0xdf, 0xa8, 0x51, 0xe6, 0xff, 0xd1, 0xde, 0xc3, 0x51, 0x52, 0xa5, 0x58, 0xaf, 0x60,
	0xea, 0x73, 0x04, 0xd5, 0xf2, 0xfa, 0xec, 0x09, 0xd3, 0x1b, 0x8b, 0x5d, 0x19, 0x44, 0xc5, 0x0a,
	0x4e, 0x6f, 0xb6, 0x9a, 0x84, 0xfc, 0xf0, 0x20, 0x9c, 0xe8, 0x86, 0x8f, 0x1c, 0xc3, 0x41, 0xdf,
	0xa9, 0xe4, 0xe4, 0xc7, 0xe2, 0xef, 0x04, 0xa6, 0x7e, 0x91, 0x71, 0xc8, 0x0d, 0xfe, 0x1a, 0x2c,
	0x16, 0xe5, 0x0e, 0x32, 0x06, 0xd3, 0x8e, 0x34, 0xf2, 0x2c, 0x6c, 0x85, 0x99, 0x3d, 0x87, 0x4c,
	0x49, 0x7e, 0x10, 0x98, 0x4c, 0x49, 0x7f, 0x04, 0xb6, 0x42, 0x69, 0x3e, 0x8d, 0x47, 0x04, 0xe0,
	0x37, 0x8d, 0x68, 0x91, 0xcf, 0xe2, 0xa6, 0x9f, 0xd9, 0x0b, 0x98, 0x8b, 0x10, 0x8a, 0xcf, 0x63,
	0xbb, 0x11, 0xb1, 0x77, 0xb0, 0xb0, 0xe8, 0x9c, 0x32, 0x8d, 0xe5, 0x79, 0x28, 0xf6, 0xfc, 0x89,
	0x1b, 0x7f, 0x44, 0x49, 0x39, 0x68, 0x8b, 0x2f, 0x90, 0x27, 0x32, 0x85, 0xf2, 0xe9, 0x67, 0x21,
	0xd4, 0xce, 0x3e, 0x1b, 0xd9, 0x73, 0xc8, 0xd1, 0x88, 0x4a, 0x63, 0x4c, 0xbf, 0x28, 0x77, 0xb0,
	0xf8, 0x33, 0x81, 0xc3, 0xf1, 0x7f, 0x7b, 0xf4, 0xb9, 0xd1, 0xc3, 0xc8, 0xf6, 0x1e, 0xc6, 0x05,
	0x2c, 0xeb, 0xde, 0x3a, 0x6a, 0x6f, 0x6b, 0x92, 0x98, 0x5a, 0x81, 0x48, 0x7d, 0x24, 0x89, 0xec,
	0x1c, 0x16, 0xb5, 0x70, 0xd8, 0x50, 0xf7, 0x3b, 0x15, 0x34, 0x60, 0x1f, 0x48, 0xa2, 0x13, 0x4a,
	0xdb, 0x54, 0xd3, 0x0e, 0xfa, 0xf8, 0x5a, 0x99, 0xfb, 0xd4, 0x53, 0x98, 0x3d, 0xe7, 0x54, 0x8b,
	0x3c, 0x8f, 0x9c, 0x9f, 0xab, 0x79, 0x68, 0xe7, 0xcd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32,
	0xa6, 0x84, 0x04, 0x42, 0x03, 0x00, 0x00,
}
