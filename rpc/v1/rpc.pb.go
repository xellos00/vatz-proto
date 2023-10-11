// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/vatz/rpc/v1/rpc.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_OK   Status = 0
	Status_FAIL Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "OK",
		1: "FAIL",
	}
	Status_value = map[string]int32{
		"OK":   0,
		"FAIL": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_vatz_rpc_v1_rpc_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_proto_vatz_rpc_v1_rpc_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_vatz_rpc_v1_rpc_proto_rawDescGZIP(), []int{0}
}

type PluginStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       Status          `protobuf:"varint,1,opt,name=status,proto3,enum=vatz.rpc.Status" json:"status,omitempty"`
	PluginStatus []*PluginStatus `protobuf:"bytes,2,rep,name=plugin_status,json=pluginStatus,proto3" json:"plugin_status,omitempty"`
}

func (x *PluginStatusResponse) Reset() {
	*x = PluginStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vatz_rpc_v1_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginStatusResponse) ProtoMessage() {}

func (x *PluginStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vatz_rpc_v1_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginStatusResponse.ProtoReflect.Descriptor instead.
func (*PluginStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_vatz_rpc_v1_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *PluginStatusResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

func (x *PluginStatusResponse) GetPluginStatus() []*PluginStatus {
	if x != nil {
		return x.PluginStatus
	}
	return nil
}

type PluginStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     Status `protobuf:"varint,1,opt,name=status,proto3,enum=vatz.rpc.Status" json:"status,omitempty"`
	PluginName string `protobuf:"bytes,2,opt,name=plugin_name,json=pluginName,proto3" json:"plugin_name,omitempty"`
}

func (x *PluginStatus) Reset() {
	*x = PluginStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vatz_rpc_v1_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginStatus) ProtoMessage() {}

func (x *PluginStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vatz_rpc_v1_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginStatus.ProtoReflect.Descriptor instead.
func (*PluginStatus) Descriptor() ([]byte, []int) {
	return file_proto_vatz_rpc_v1_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *PluginStatus) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

func (x *PluginStatus) GetPluginName() string {
	if x != nil {
		return x.PluginName
	}
	return ""
}

var File_proto_vatz_rpc_v1_rpc_proto protoreflect.FileDescriptor

var file_proto_vatz_rpc_v1_rpc_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x74, 0x7a, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76,
	0x61, 0x74, 0x7a, 0x2e, 0x72, 0x70, 0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x14, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x76,
	0x61, 0x74, 0x7a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x0d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x76, 0x61, 0x74, 0x7a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x59, 0x0a, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x76, 0x61, 0x74, 0x7a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x1a,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x32, 0x6c, 0x0a, 0x07, 0x56, 0x61,
	0x74, 0x7a, 0x52, 0x50, 0x43, 0x12, 0x61, 0x0a, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e,
	0x76, 0x61, 0x74, 0x7a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0x5a, 0x06, 0x72, 0x70, 0x63, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_vatz_rpc_v1_rpc_proto_rawDescOnce sync.Once
	file_proto_vatz_rpc_v1_rpc_proto_rawDescData = file_proto_vatz_rpc_v1_rpc_proto_rawDesc
)

func file_proto_vatz_rpc_v1_rpc_proto_rawDescGZIP() []byte {
	file_proto_vatz_rpc_v1_rpc_proto_rawDescOnce.Do(func() {
		file_proto_vatz_rpc_v1_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_vatz_rpc_v1_rpc_proto_rawDescData)
	})
	return file_proto_vatz_rpc_v1_rpc_proto_rawDescData
}

var file_proto_vatz_rpc_v1_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_vatz_rpc_v1_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_vatz_rpc_v1_rpc_proto_goTypes = []interface{}{
	(Status)(0),                  // 0: vatz.rpc.Status
	(*PluginStatusResponse)(nil), // 1: vatz.rpc.PluginStatusResponse
	(*PluginStatus)(nil),         // 2: vatz.rpc.PluginStatus
	(*emptypb.Empty)(nil),        // 3: google.protobuf.Empty
}
var file_proto_vatz_rpc_v1_rpc_proto_depIdxs = []int32{
	0, // 0: vatz.rpc.PluginStatusResponse.status:type_name -> vatz.rpc.Status
	2, // 1: vatz.rpc.PluginStatusResponse.plugin_status:type_name -> vatz.rpc.PluginStatus
	0, // 2: vatz.rpc.PluginStatus.status:type_name -> vatz.rpc.Status
	3, // 3: vatz.rpc.VatzRPC.PluginStatus:input_type -> google.protobuf.Empty
	1, // 4: vatz.rpc.VatzRPC.PluginStatus:output_type -> vatz.rpc.PluginStatusResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_vatz_rpc_v1_rpc_proto_init() }
func file_proto_vatz_rpc_v1_rpc_proto_init() {
	if File_proto_vatz_rpc_v1_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_vatz_rpc_v1_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginStatusResponse); i {
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
		file_proto_vatz_rpc_v1_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginStatus); i {
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
			RawDescriptor: file_proto_vatz_rpc_v1_rpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_vatz_rpc_v1_rpc_proto_goTypes,
		DependencyIndexes: file_proto_vatz_rpc_v1_rpc_proto_depIdxs,
		EnumInfos:         file_proto_vatz_rpc_v1_rpc_proto_enumTypes,
		MessageInfos:      file_proto_vatz_rpc_v1_rpc_proto_msgTypes,
	}.Build()
	File_proto_vatz_rpc_v1_rpc_proto = out.File
	file_proto_vatz_rpc_v1_rpc_proto_rawDesc = nil
	file_proto_vatz_rpc_v1_rpc_proto_goTypes = nil
	file_proto_vatz_rpc_v1_rpc_proto_depIdxs = nil
}
