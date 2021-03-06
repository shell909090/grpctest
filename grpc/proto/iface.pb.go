// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: iface.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_iface_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *AddRequest) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N int32 `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_iface_proto_rawDescGZIP(), []int{1}
}

func (x *AddResponse) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N int32 `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iface_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iface_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_iface_proto_rawDescGZIP(), []int{2}
}

func (x *SumRequest) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N int32 `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iface_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iface_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_iface_proto_rawDescGZIP(), []int{3}
}

func (x *SumResponse) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

type RangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N   int32 `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
	Len int32 `protobuf:"varint,2,opt,name=len,proto3" json:"len,omitempty"`
}

func (x *RangeRequest) Reset() {
	*x = RangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iface_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RangeRequest) ProtoMessage() {}

func (x *RangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iface_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RangeRequest.ProtoReflect.Descriptor instead.
func (*RangeRequest) Descriptor() ([]byte, []int) {
	return file_iface_proto_rawDescGZIP(), []int{4}
}

func (x *RangeRequest) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *RangeRequest) GetLen() int32 {
	if x != nil {
		return x.Len
	}
	return 0
}

type RangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N int32 `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
}

func (x *RangeResponse) Reset() {
	*x = RangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iface_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RangeResponse) ProtoMessage() {}

func (x *RangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iface_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RangeResponse.ProtoReflect.Descriptor instead.
func (*RangeResponse) Descriptor() ([]byte, []int) {
	return file_iface_proto_rawDescGZIP(), []int{5}
}

func (x *RangeResponse) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N int32  `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
	S string `protobuf:"bytes,2,opt,name=s,proto3" json:"s,omitempty"`
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iface_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iface_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_iface_proto_rawDescGZIP(), []int{6}
}

func (x *EchoRequest) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *EchoRequest) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

type EchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N int32  `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
	S string `protobuf:"bytes,2,opt,name=s,proto3" json:"s,omitempty"`
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iface_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iface_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResponse.ProtoReflect.Descriptor instead.
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return file_iface_proto_rawDescGZIP(), []int{7}
}

func (x *EchoResponse) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *EchoResponse) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

var File_iface_proto protoreflect.FileDescriptor

var file_iface_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61,
	0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x62, 0x22, 0x1b,
	0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a,
	0x01, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x6e, 0x22, 0x1a, 0x0a, 0x0a, 0x53,
	0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x6e, 0x22, 0x1b, 0x0a, 0x0b, 0x53, 0x75, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x6e, 0x22, 0x2e, 0x0a, 0x0c, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x01, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x6c, 0x65, 0x6e, 0x22, 0x1d, 0x0a, 0x0d, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x6e, 0x22, 0x29, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x6e,
	0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x73, 0x22, 0x2a,
	0x0a, 0x0c, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c,
	0x0a, 0x01, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x6e, 0x12, 0x0c, 0x0a, 0x01,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x73, 0x32, 0x67, 0x0a, 0x03, 0x41, 0x64,
	0x64, 0x12, 0x2e, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x30, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x32, 0x3f, 0x0a, 0x05, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x05,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x32, 0x3d, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x35, 0x0a, 0x04,
	0x45, 0x63, 0x68, 0x6f, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x63, 0x68,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_iface_proto_rawDescOnce sync.Once
	file_iface_proto_rawDescData = file_iface_proto_rawDesc
)

func file_iface_proto_rawDescGZIP() []byte {
	file_iface_proto_rawDescOnce.Do(func() {
		file_iface_proto_rawDescData = protoimpl.X.CompressGZIP(file_iface_proto_rawDescData)
	})
	return file_iface_proto_rawDescData
}

var file_iface_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_iface_proto_goTypes = []interface{}{
	(*AddRequest)(nil),    // 0: proto.AddRequest
	(*AddResponse)(nil),   // 1: proto.AddResponse
	(*SumRequest)(nil),    // 2: proto.SumRequest
	(*SumResponse)(nil),   // 3: proto.SumResponse
	(*RangeRequest)(nil),  // 4: proto.RangeRequest
	(*RangeResponse)(nil), // 5: proto.RangeResponse
	(*EchoRequest)(nil),   // 6: proto.EchoRequest
	(*EchoResponse)(nil),  // 7: proto.EchoResponse
}
var file_iface_proto_depIdxs = []int32{
	0, // 0: proto.Add.Add:input_type -> proto.AddRequest
	2, // 1: proto.Add.Sum:input_type -> proto.SumRequest
	4, // 2: proto.Range.Range:input_type -> proto.RangeRequest
	6, // 3: proto.Echo.Echo:input_type -> proto.EchoRequest
	1, // 4: proto.Add.Add:output_type -> proto.AddResponse
	3, // 5: proto.Add.Sum:output_type -> proto.SumResponse
	5, // 6: proto.Range.Range:output_type -> proto.RangeResponse
	7, // 7: proto.Echo.Echo:output_type -> proto.EchoResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_iface_proto_init() }
func file_iface_proto_init() {
	if File_iface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_iface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
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
		file_iface_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_iface_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
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
		file_iface_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RangeRequest); i {
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
		file_iface_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RangeResponse); i {
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
		file_iface_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoRequest); i {
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
		file_iface_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResponse); i {
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
			RawDescriptor: file_iface_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_iface_proto_goTypes,
		DependencyIndexes: file_iface_proto_depIdxs,
		MessageInfos:      file_iface_proto_msgTypes,
	}.Build()
	File_iface_proto = out.File
	file_iface_proto_rawDesc = nil
	file_iface_proto_goTypes = nil
	file_iface_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AddClient is the client API for Add service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	Sum(ctx context.Context, opts ...grpc.CallOption) (Add_SumClient, error)
}

type addClient struct {
	cc grpc.ClientConnInterface
}

func NewAddClient(cc grpc.ClientConnInterface) AddClient {
	return &addClient{cc}
}

func (c *addClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/proto.Add/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addClient) Sum(ctx context.Context, opts ...grpc.CallOption) (Add_SumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Add_serviceDesc.Streams[0], "/proto.Add/Sum", opts...)
	if err != nil {
		return nil, err
	}
	x := &addSumClient{stream}
	return x, nil
}

type Add_SumClient interface {
	Send(*SumRequest) error
	CloseAndRecv() (*SumResponse, error)
	grpc.ClientStream
}

type addSumClient struct {
	grpc.ClientStream
}

func (x *addSumClient) Send(m *SumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *addSumClient) CloseAndRecv() (*SumResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AddServer is the server API for Add service.
type AddServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	Sum(Add_SumServer) error
}

// UnimplementedAddServer can be embedded to have forward compatible implementations.
type UnimplementedAddServer struct {
}

func (*UnimplementedAddServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedAddServer) Sum(Add_SumServer) error {
	return status.Errorf(codes.Unimplemented, "method Sum not implemented")
}

func RegisterAddServer(s *grpc.Server, srv AddServer) {
	s.RegisterService(&_Add_serviceDesc, srv)
}

func _Add_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Add/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Add_Sum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AddServer).Sum(&addSumServer{stream})
}

type Add_SumServer interface {
	SendAndClose(*SumResponse) error
	Recv() (*SumRequest, error)
	grpc.ServerStream
}

type addSumServer struct {
	grpc.ServerStream
}

func (x *addSumServer) SendAndClose(m *SumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *addSumServer) Recv() (*SumRequest, error) {
	m := new(SumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Add_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Add",
	HandlerType: (*AddServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Add_Add_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sum",
			Handler:       _Add_Sum_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "iface.proto",
}

// RangeClient is the client API for Range service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RangeClient interface {
	Range(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (Range_RangeClient, error)
}

type rangeClient struct {
	cc grpc.ClientConnInterface
}

func NewRangeClient(cc grpc.ClientConnInterface) RangeClient {
	return &rangeClient{cc}
}

func (c *rangeClient) Range(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (Range_RangeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Range_serviceDesc.Streams[0], "/proto.Range/Range", opts...)
	if err != nil {
		return nil, err
	}
	x := &rangeRangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Range_RangeClient interface {
	Recv() (*RangeResponse, error)
	grpc.ClientStream
}

type rangeRangeClient struct {
	grpc.ClientStream
}

func (x *rangeRangeClient) Recv() (*RangeResponse, error) {
	m := new(RangeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RangeServer is the server API for Range service.
type RangeServer interface {
	Range(*RangeRequest, Range_RangeServer) error
}

// UnimplementedRangeServer can be embedded to have forward compatible implementations.
type UnimplementedRangeServer struct {
}

func (*UnimplementedRangeServer) Range(*RangeRequest, Range_RangeServer) error {
	return status.Errorf(codes.Unimplemented, "method Range not implemented")
}

func RegisterRangeServer(s *grpc.Server, srv RangeServer) {
	s.RegisterService(&_Range_serviceDesc, srv)
}

func _Range_Range_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RangeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RangeServer).Range(m, &rangeRangeServer{stream})
}

type Range_RangeServer interface {
	Send(*RangeResponse) error
	grpc.ServerStream
}

type rangeRangeServer struct {
	grpc.ServerStream
}

func (x *rangeRangeServer) Send(m *RangeResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Range_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Range",
	HandlerType: (*RangeServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Range",
			Handler:       _Range_Range_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "iface.proto",
}

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoClient interface {
	Echo(ctx context.Context, opts ...grpc.CallOption) (Echo_EchoClient, error)
}

type echoClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoClient(cc grpc.ClientConnInterface) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Echo(ctx context.Context, opts ...grpc.CallOption) (Echo_EchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Echo_serviceDesc.Streams[0], "/proto.Echo/Echo", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoEchoClient{stream}
	return x, nil
}

type Echo_EchoClient interface {
	Send(*EchoRequest) error
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoEchoClient struct {
	grpc.ClientStream
}

func (x *echoEchoClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoEchoClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	Echo(Echo_EchoServer) error
}

// UnimplementedEchoServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServer struct {
}

func (*UnimplementedEchoServer) Echo(Echo_EchoServer) error {
	return status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Echo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServer).Echo(&echoEchoServer{stream})
}

type Echo_EchoServer interface {
	Send(*EchoResponse) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type echoEchoServer struct {
	grpc.ServerStream
}

func (x *echoEchoServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoEchoServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Echo",
			Handler:       _Echo_Echo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "iface.proto",
}
