// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/test.proto

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

type TestMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (x *TestMsg) Reset() {
	*x = TestMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMsg) ProtoMessage() {}

func (x *TestMsg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMsg.ProtoReflect.Descriptor instead.
func (*TestMsg) Descriptor() ([]byte, []int) {
	return file_proto_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestMsg) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *TestMsg) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type TeestTestMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   *string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Msg  *TestMsg `protobuf:"bytes,3,opt,name=msg" json:"msg,omitempty"`
}

func (x *TeestTestMsg) Reset() {
	*x = TeestTestMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeestTestMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeestTestMsg) ProtoMessage() {}

func (x *TeestTestMsg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeestTestMsg.ProtoReflect.Descriptor instead.
func (*TeestTestMsg) Descriptor() ([]byte, []int) {
	return file_proto_test_proto_rawDescGZIP(), []int{1}
}

func (x *TeestTestMsg) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *TeestTestMsg) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *TeestTestMsg) GetMsg() *TestMsg {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_proto_test_proto protoreflect.FileDescriptor

var file_proto_test_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x07, 0x54, 0x65, 0x73,
	0x74, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x54, 0x0a, 0x0c, 0x54, 0x65, 0x65, 0x73,
	0x74, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x03, 0x6d, 0x73, 0x67,
}

var (
	file_proto_test_proto_rawDescOnce sync.Once
	file_proto_test_proto_rawDescData = file_proto_test_proto_rawDesc
)

func file_proto_test_proto_rawDescGZIP() []byte {
	file_proto_test_proto_rawDescOnce.Do(func() {
		file_proto_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_test_proto_rawDescData)
	})
	return file_proto_test_proto_rawDescData
}

var file_proto_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_test_proto_goTypes = []interface{}{
	(*TestMsg)(nil),      // 0: proto.TestMsg
	(*TeestTestMsg)(nil), // 1: proto.TeestTestMsg
}
var file_proto_test_proto_depIdxs = []int32{
	0, // 0: proto.TeestTestMsg.msg:type_name -> proto.TestMsg
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_test_proto_init() }
func file_proto_test_proto_init() {
	if File_proto_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMsg); i {
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
		file_proto_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeestTestMsg); i {
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
			RawDescriptor: file_proto_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_test_proto_goTypes,
		DependencyIndexes: file_proto_test_proto_depIdxs,
		MessageInfos:      file_proto_test_proto_msgTypes,
	}.Build()
	File_proto_test_proto = out.File
	file_proto_test_proto_rawDesc = nil
	file_proto_test_proto_goTypes = nil
	file_proto_test_proto_depIdxs = nil
}
