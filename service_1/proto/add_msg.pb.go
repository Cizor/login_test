// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: service_1/proto/add_msg.proto

package pb

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

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=A,proto3" json:"A,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=B,proto3" json:"B,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_1_proto_add_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_1_proto_add_msg_proto_msgTypes[0]
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
	return file_service_1_proto_add_msg_proto_rawDescGZIP(), []int{0}
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

	Sum int32 `protobuf:"varint,1,opt,name=Sum,proto3" json:"Sum,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_1_proto_add_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_1_proto_add_msg_proto_msgTypes[1]
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
	return file_service_1_proto_add_msg_proto_rawDescGZIP(), []int{1}
}

func (x *AddResponse) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type PrintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *PrintRequest) Reset() {
	*x = PrintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_1_proto_add_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintRequest) ProtoMessage() {}

func (x *PrintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_1_proto_add_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintRequest.ProtoReflect.Descriptor instead.
func (*PrintRequest) Descriptor() ([]byte, []int) {
	return file_service_1_proto_add_msg_proto_rawDescGZIP(), []int{2}
}

func (x *PrintRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PrintResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PrintResponse) Reset() {
	*x = PrintResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_1_proto_add_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrintResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintResponse) ProtoMessage() {}

func (x *PrintResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_1_proto_add_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintResponse.ProtoReflect.Descriptor instead.
func (*PrintResponse) Descriptor() ([]byte, []int) {
	return file_service_1_proto_add_msg_proto_rawDescGZIP(), []int{3}
}

var File_service_1_proto_add_msg_proto protoreflect.FileDescriptor

var file_service_1_proto_add_msg_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x61, 0x64, 0x64, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x28, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0c, 0x0a, 0x01, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x41, 0x12,
	0x0c, 0x0a, 0x01, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x42, 0x22, 0x1f, 0x0a,
	0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x53, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x53, 0x75, 0x6d, 0x22, 0x28,
	0x0a, 0x0c, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x31, 0x0a, 0x05, 0x41, 0x64, 0x64,
	0x65, 0x72, 0x12, 0x28, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x39, 0x0a, 0x07,
	0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x6e, 0x74,
	0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_1_proto_add_msg_proto_rawDescOnce sync.Once
	file_service_1_proto_add_msg_proto_rawDescData = file_service_1_proto_add_msg_proto_rawDesc
)

func file_service_1_proto_add_msg_proto_rawDescGZIP() []byte {
	file_service_1_proto_add_msg_proto_rawDescOnce.Do(func() {
		file_service_1_proto_add_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_1_proto_add_msg_proto_rawDescData)
	})
	return file_service_1_proto_add_msg_proto_rawDescData
}

var file_service_1_proto_add_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_1_proto_add_msg_proto_goTypes = []interface{}{
	(*AddRequest)(nil),    // 0: pb.AddRequest
	(*AddResponse)(nil),   // 1: pb.AddResponse
	(*PrintRequest)(nil),  // 2: pb.PrintRequest
	(*PrintResponse)(nil), // 3: pb.PrintResponse
}
var file_service_1_proto_add_msg_proto_depIdxs = []int32{
	0, // 0: pb.Adder.Add:input_type -> pb.AddRequest
	2, // 1: pb.Printer.Print:input_type -> pb.PrintRequest
	1, // 2: pb.Adder.Add:output_type -> pb.AddResponse
	3, // 3: pb.Printer.Print:output_type -> pb.PrintResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_1_proto_add_msg_proto_init() }
func file_service_1_proto_add_msg_proto_init() {
	if File_service_1_proto_add_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_1_proto_add_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_service_1_proto_add_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_service_1_proto_add_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrintRequest); i {
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
		file_service_1_proto_add_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrintResponse); i {
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
			RawDescriptor: file_service_1_proto_add_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_service_1_proto_add_msg_proto_goTypes,
		DependencyIndexes: file_service_1_proto_add_msg_proto_depIdxs,
		MessageInfos:      file_service_1_proto_add_msg_proto_msgTypes,
	}.Build()
	File_service_1_proto_add_msg_proto = out.File
	file_service_1_proto_add_msg_proto_rawDesc = nil
	file_service_1_proto_add_msg_proto_goTypes = nil
	file_service_1_proto_add_msg_proto_depIdxs = nil
}