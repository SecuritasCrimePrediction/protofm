// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: testproto/testproto.proto

package testproto

import (
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

type Single struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"` // a
}

func (x *Single) Reset() {
	*x = Single{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testproto_testproto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Single) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Single) ProtoMessage() {}

func (x *Single) ProtoReflect() protoreflect.Message {
	mi := &file_testproto_testproto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Single.ProtoReflect.Descriptor instead.
func (*Single) Descriptor() ([]byte, []int) {
	return file_testproto_testproto_proto_rawDescGZIP(), []int{0}
}

func (x *Single) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type RepeatedAndSingle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SingleValue   *Single  `protobuf:"bytes,1,opt,name=single_value,json=singleValue,proto3" json:"single_value,omitempty"`       // a
	RepeatedValue []string `protobuf:"bytes,2,rep,name=repeated_value,json=repeatedValue,proto3" json:"repeated_value,omitempty"` // b
}

func (x *RepeatedAndSingle) Reset() {
	*x = RepeatedAndSingle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testproto_testproto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepeatedAndSingle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepeatedAndSingle) ProtoMessage() {}

func (x *RepeatedAndSingle) ProtoReflect() protoreflect.Message {
	mi := &file_testproto_testproto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepeatedAndSingle.ProtoReflect.Descriptor instead.
func (*RepeatedAndSingle) Descriptor() ([]byte, []int) {
	return file_testproto_testproto_proto_rawDescGZIP(), []int{1}
}

func (x *RepeatedAndSingle) GetSingleValue() *Single {
	if x != nil {
		return x.SingleValue
	}
	return nil
}

func (x *RepeatedAndSingle) GetRepeatedValue() []string {
	if x != nil {
		return x.RepeatedValue
	}
	return nil
}

type ComplexObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepeatedAndSingleValue *RepeatedAndSingle `protobuf:"bytes,1,opt,name=repeated_and_single_value,json=repeatedAndSingleValue,proto3" json:"repeated_and_single_value,omitempty"`
	SingleValue            string             `protobuf:"bytes,2,opt,name=single_value,json=singleValue,proto3" json:"single_value,omitempty"` // c
}

func (x *ComplexObject) Reset() {
	*x = ComplexObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testproto_testproto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexObject) ProtoMessage() {}

func (x *ComplexObject) ProtoReflect() protoreflect.Message {
	mi := &file_testproto_testproto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexObject.ProtoReflect.Descriptor instead.
func (*ComplexObject) Descriptor() ([]byte, []int) {
	return file_testproto_testproto_proto_rawDescGZIP(), []int{2}
}

func (x *ComplexObject) GetRepeatedAndSingleValue() *RepeatedAndSingle {
	if x != nil {
		return x.RepeatedAndSingleValue
	}
	return nil
}

func (x *ComplexObject) GetSingleValue() string {
	if x != nil {
		return x.SingleValue
	}
	return ""
}

type SimpleObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pow string `protobuf:"bytes,1,opt,name=pow,proto3" json:"pow,omitempty"`
	Wow string `protobuf:"bytes,2,opt,name=wow,proto3" json:"wow,omitempty"`
	Foo int32  `protobuf:"varint,3,opt,name=foo,proto3" json:"foo,omitempty"`
	Baz int32  `protobuf:"varint,4,opt,name=baz,proto3" json:"baz,omitempty"`
}

func (x *SimpleObject) Reset() {
	*x = SimpleObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testproto_testproto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleObject) ProtoMessage() {}

func (x *SimpleObject) ProtoReflect() protoreflect.Message {
	mi := &file_testproto_testproto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleObject.ProtoReflect.Descriptor instead.
func (*SimpleObject) Descriptor() ([]byte, []int) {
	return file_testproto_testproto_proto_rawDescGZIP(), []int{3}
}

func (x *SimpleObject) GetPow() string {
	if x != nil {
		return x.Pow
	}
	return ""
}

func (x *SimpleObject) GetWow() string {
	if x != nil {
		return x.Wow
	}
	return ""
}

func (x *SimpleObject) GetFoo() int32 {
	if x != nil {
		return x.Foo
	}
	return 0
}

func (x *SimpleObject) GetBaz() int32 {
	if x != nil {
		return x.Baz
	}
	return 0
}

type NestedObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pow           int32                  `protobuf:"varint,1,opt,name=pow,proto3" json:"pow,omitempty"`
	Wow           string                 `protobuf:"bytes,2,opt,name=wow,proto3" json:"wow,omitempty"`
	FooBaz        []*NestedObject_FooBaz `protobuf:"bytes,3,rep,name=foo_baz,json=fooBaz,proto3" json:"foo_baz,omitempty"`
	ComplexObject *ComplexObject         `protobuf:"bytes,4,opt,name=complex_object,json=complexObject,proto3" json:"complex_object,omitempty"` // c
}

func (x *NestedObject) Reset() {
	*x = NestedObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testproto_testproto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedObject) ProtoMessage() {}

func (x *NestedObject) ProtoReflect() protoreflect.Message {
	mi := &file_testproto_testproto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedObject.ProtoReflect.Descriptor instead.
func (*NestedObject) Descriptor() ([]byte, []int) {
	return file_testproto_testproto_proto_rawDescGZIP(), []int{4}
}

func (x *NestedObject) GetPow() int32 {
	if x != nil {
		return x.Pow
	}
	return 0
}

func (x *NestedObject) GetWow() string {
	if x != nil {
		return x.Wow
	}
	return ""
}

func (x *NestedObject) GetFooBaz() []*NestedObject_FooBaz {
	if x != nil {
		return x.FooBaz
	}
	return nil
}

func (x *NestedObject) GetComplexObject() *ComplexObject {
	if x != nil {
		return x.ComplexObject
	}
	return nil
}

type FakeFieldMaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldMask   []string `protobuf:"bytes,1,rep,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	NrResponses int32    `protobuf:"varint,2,opt,name=nr_responses,json=nrResponses,proto3" json:"nr_responses,omitempty"`
	RetError    string   `protobuf:"bytes,3,opt,name=ret_error,json=retError,proto3" json:"ret_error,omitempty"`
}

func (x *FakeFieldMaskRequest) Reset() {
	*x = FakeFieldMaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testproto_testproto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FakeFieldMaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FakeFieldMaskRequest) ProtoMessage() {}

func (x *FakeFieldMaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testproto_testproto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FakeFieldMaskRequest.ProtoReflect.Descriptor instead.
func (*FakeFieldMaskRequest) Descriptor() ([]byte, []int) {
	return file_testproto_testproto_proto_rawDescGZIP(), []int{5}
}

func (x *FakeFieldMaskRequest) GetFieldMask() []string {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

func (x *FakeFieldMaskRequest) GetNrResponses() int32 {
	if x != nil {
		return x.NrResponses
	}
	return 0
}

func (x *FakeFieldMaskRequest) GetRetError() string {
	if x != nil {
		return x.RetError
	}
	return ""
}

type NestedObject_FooBaz struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo int32  `protobuf:"varint,1,opt,name=foo,proto3" json:"foo,omitempty"`
	Baz string `protobuf:"bytes,2,opt,name=baz,proto3" json:"baz,omitempty"`
}

func (x *NestedObject_FooBaz) Reset() {
	*x = NestedObject_FooBaz{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testproto_testproto_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedObject_FooBaz) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedObject_FooBaz) ProtoMessage() {}

func (x *NestedObject_FooBaz) ProtoReflect() protoreflect.Message {
	mi := &file_testproto_testproto_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedObject_FooBaz.ProtoReflect.Descriptor instead.
func (*NestedObject_FooBaz) Descriptor() ([]byte, []int) {
	return file_testproto_testproto_proto_rawDescGZIP(), []int{4, 0}
}

func (x *NestedObject_FooBaz) GetFoo() int32 {
	if x != nil {
		return x.Foo
	}
	return 0
}

func (x *NestedObject_FooBaz) GetBaz() string {
	if x != nil {
		return x.Baz
	}
	return ""
}

var File_testproto_testproto_proto protoreflect.FileDescriptor

var file_testproto_testproto_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x66, 0x6d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x22, 0x1e, 0x0a, 0x06, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x73, 0x0a, 0x11, 0x52, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x6e, 0x64, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x37,
	0x0a, 0x0c, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x6d, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x0b, 0x73, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x8e,
	0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x5a, 0x0a, 0x19, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x6e, 0x64,
	0x5f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x6d, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x6e, 0x64, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x52, 0x16, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x6e,
	0x64, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x56, 0x0a, 0x0c, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x70, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x6f,
	0x77, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x77, 0x6f, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x61, 0x7a, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x62, 0x61, 0x7a, 0x22, 0xe0, 0x01, 0x0a, 0x0c, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x77, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x6f, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x6f,
	0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x6f, 0x77, 0x12, 0x3a, 0x0a, 0x07,
	0x66, 0x6f, 0x6f, 0x5f, 0x62, 0x61, 0x7a, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x6d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x46, 0x6f, 0x6f, 0x42, 0x61, 0x7a,
	0x52, 0x06, 0x66, 0x6f, 0x6f, 0x42, 0x61, 0x7a, 0x12, 0x42, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x78, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x6d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0d, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x2c, 0x0a, 0x06,
	0x46, 0x6f, 0x6f, 0x42, 0x61, 0x7a, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x61, 0x7a, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x61, 0x7a, 0x22, 0x75, 0x0a, 0x14, 0x46, 0x61,
	0x6b, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73,
	0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6e, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testproto_testproto_proto_rawDescOnce sync.Once
	file_testproto_testproto_proto_rawDescData = file_testproto_testproto_proto_rawDesc
)

func file_testproto_testproto_proto_rawDescGZIP() []byte {
	file_testproto_testproto_proto_rawDescOnce.Do(func() {
		file_testproto_testproto_proto_rawDescData = protoimpl.X.CompressGZIP(file_testproto_testproto_proto_rawDescData)
	})
	return file_testproto_testproto_proto_rawDescData
}

var file_testproto_testproto_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_testproto_testproto_proto_goTypes = []interface{}{
	(*Single)(nil),               // 0: protofm.test.Single
	(*RepeatedAndSingle)(nil),    // 1: protofm.test.RepeatedAndSingle
	(*ComplexObject)(nil),        // 2: protofm.test.ComplexObject
	(*SimpleObject)(nil),         // 3: protofm.test.SimpleObject
	(*NestedObject)(nil),         // 4: protofm.test.NestedObject
	(*FakeFieldMaskRequest)(nil), // 5: protofm.test.FakeFieldMaskRequest
	(*NestedObject_FooBaz)(nil),  // 6: protofm.test.NestedObject.FooBaz
}
var file_testproto_testproto_proto_depIdxs = []int32{
	0, // 0: protofm.test.RepeatedAndSingle.single_value:type_name -> protofm.test.Single
	1, // 1: protofm.test.ComplexObject.repeated_and_single_value:type_name -> protofm.test.RepeatedAndSingle
	6, // 2: protofm.test.NestedObject.foo_baz:type_name -> protofm.test.NestedObject.FooBaz
	2, // 3: protofm.test.NestedObject.complex_object:type_name -> protofm.test.ComplexObject
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_testproto_testproto_proto_init() }
func file_testproto_testproto_proto_init() {
	if File_testproto_testproto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testproto_testproto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Single); i {
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
		file_testproto_testproto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepeatedAndSingle); i {
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
		file_testproto_testproto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexObject); i {
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
		file_testproto_testproto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleObject); i {
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
		file_testproto_testproto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedObject); i {
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
		file_testproto_testproto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FakeFieldMaskRequest); i {
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
		file_testproto_testproto_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedObject_FooBaz); i {
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
			RawDescriptor: file_testproto_testproto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testproto_testproto_proto_goTypes,
		DependencyIndexes: file_testproto_testproto_proto_depIdxs,
		MessageInfos:      file_testproto_testproto_proto_msgTypes,
	}.Build()
	File_testproto_testproto_proto = out.File
	file_testproto_testproto_proto_rawDesc = nil
	file_testproto_testproto_proto_goTypes = nil
	file_testproto_testproto_proto_depIdxs = nil
}
