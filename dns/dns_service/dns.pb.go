// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: dns/dns_service/dns.proto

package proto

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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainName string `protobuf:"bytes,1,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty"`
	Ip         string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_dns_service_dns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dns_dns_service_dns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_dns_dns_service_dns_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

func (x *CreateRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VectorClock []int32 `protobuf:"varint,1,rep,packed,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_dns_service_dns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dns_dns_service_dns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_dns_dns_service_dns_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetVectorClock() []int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainName string `protobuf:"bytes,1,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_dns_service_dns_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dns_dns_service_dns_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_dns_dns_service_dns_proto_rawDescGZIP(), []int{2}
}

func (x *ReadRequest) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip          string  `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	VectorClock []int32 `protobuf:"varint,2,rep,packed,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_dns_service_dns_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dns_dns_service_dns_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_dns_dns_service_dns_proto_rawDescGZIP(), []int{3}
}

func (x *ReadResponse) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ReadResponse) GetVectorClock() []int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainName string `protobuf:"bytes,1,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty"`
	Option     bool   `protobuf:"varint,2,opt,name=option,proto3" json:"option,omitempty"`
	Parameter  string `protobuf:"bytes,3,opt,name=parameter,proto3" json:"parameter,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_dns_service_dns_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dns_dns_service_dns_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_dns_dns_service_dns_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRequest) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

func (x *UpdateRequest) GetOption() bool {
	if x != nil {
		return x.Option
	}
	return false
}

func (x *UpdateRequest) GetParameter() string {
	if x != nil {
		return x.Parameter
	}
	return ""
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VectorClock []int32 `protobuf:"varint,1,rep,packed,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_dns_service_dns_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dns_dns_service_dns_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_dns_dns_service_dns_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateResponse) GetVectorClock() []int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainName string `protobuf:"bytes,1,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_dns_service_dns_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dns_dns_service_dns_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_dns_dns_service_dns_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRequest) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VectorClock []int32 `protobuf:"varint,1,rep,packed,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_dns_service_dns_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dns_dns_service_dns_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_dns_dns_service_dns_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteResponse) GetVectorClock() []int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

var File_dns_dns_service_dns_proto protoreflect.FileDescriptor

var file_dns_dns_service_dns_proto_rawDesc = []byte{
	0x0a, 0x19, 0x64, 0x6e, 0x73, 0x2f, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x64, 0x6e, 0x73,
	0x22, 0x40, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x22, 0x33, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x2e, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x76,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x66, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x22, 0x33, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x30, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x32, 0xda,
	0x01, 0x0a, 0x0a, 0x44, 0x4e, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x64, 0x6e,
	0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x2d, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x10, 0x2e, 0x64, 0x6e, 0x73,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x64,
	0x6e, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x33, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x64, 0x6e,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x12, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x64,
	0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dns_dns_service_dns_proto_rawDescOnce sync.Once
	file_dns_dns_service_dns_proto_rawDescData = file_dns_dns_service_dns_proto_rawDesc
)

func file_dns_dns_service_dns_proto_rawDescGZIP() []byte {
	file_dns_dns_service_dns_proto_rawDescOnce.Do(func() {
		file_dns_dns_service_dns_proto_rawDescData = protoimpl.X.CompressGZIP(file_dns_dns_service_dns_proto_rawDescData)
	})
	return file_dns_dns_service_dns_proto_rawDescData
}

var file_dns_dns_service_dns_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_dns_dns_service_dns_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),  // 0: dns.CreateRequest
	(*CreateResponse)(nil), // 1: dns.CreateResponse
	(*ReadRequest)(nil),    // 2: dns.ReadRequest
	(*ReadResponse)(nil),   // 3: dns.ReadResponse
	(*UpdateRequest)(nil),  // 4: dns.UpdateRequest
	(*UpdateResponse)(nil), // 5: dns.UpdateResponse
	(*DeleteRequest)(nil),  // 6: dns.DeleteRequest
	(*DeleteResponse)(nil), // 7: dns.DeleteResponse
}
var file_dns_dns_service_dns_proto_depIdxs = []int32{
	0, // 0: dns.DNSService.Create:input_type -> dns.CreateRequest
	2, // 1: dns.DNSService.Read:input_type -> dns.ReadRequest
	4, // 2: dns.DNSService.Update:input_type -> dns.UpdateRequest
	6, // 3: dns.DNSService.Delete:input_type -> dns.DeleteRequest
	1, // 4: dns.DNSService.Create:output_type -> dns.CreateResponse
	3, // 5: dns.DNSService.Read:output_type -> dns.ReadResponse
	5, // 6: dns.DNSService.Update:output_type -> dns.UpdateResponse
	7, // 7: dns.DNSService.Delete:output_type -> dns.DeleteResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dns_dns_service_dns_proto_init() }
func file_dns_dns_service_dns_proto_init() {
	if File_dns_dns_service_dns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dns_dns_service_dns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_dns_dns_service_dns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_dns_dns_service_dns_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_dns_dns_service_dns_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResponse); i {
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
		file_dns_dns_service_dns_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_dns_dns_service_dns_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_dns_dns_service_dns_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_dns_dns_service_dns_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
			RawDescriptor: file_dns_dns_service_dns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dns_dns_service_dns_proto_goTypes,
		DependencyIndexes: file_dns_dns_service_dns_proto_depIdxs,
		MessageInfos:      file_dns_dns_service_dns_proto_msgTypes,
	}.Build()
	File_dns_dns_service_dns_proto = out.File
	file_dns_dns_service_dns_proto_rawDesc = nil
	file_dns_dns_service_dns_proto_goTypes = nil
	file_dns_dns_service_dns_proto_depIdxs = nil
}
