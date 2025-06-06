// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: loggerpb/message.proto

package loggerpb

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

type LogStatus int32

const (
	LogStatus_INFO    LogStatus = 0
	LogStatus_WARNING LogStatus = 1
	LogStatus_ERROR   LogStatus = 2
	LogStatus_DEBUG   LogStatus = 3
	LogStatus_FATAL   LogStatus = 4
	LogStatus_UNKNOWN LogStatus = 5
)

// Enum value maps for LogStatus.
var (
	LogStatus_name = map[int32]string{
		0: "INFO",
		1: "WARNING",
		2: "ERROR",
		3: "DEBUG",
		4: "FATAL",
		5: "UNKNOWN",
	}
	LogStatus_value = map[string]int32{
		"INFO":    0,
		"WARNING": 1,
		"ERROR":   2,
		"DEBUG":   3,
		"FATAL":   4,
		"UNKNOWN": 5,
	}
)

func (x LogStatus) Enum() *LogStatus {
	p := new(LogStatus)
	*p = x
	return p
}

func (x LogStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_loggerpb_message_proto_enumTypes[0].Descriptor()
}

func (LogStatus) Type() protoreflect.EnumType {
	return &file_loggerpb_message_proto_enumTypes[0]
}

func (x LogStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogStatus.Descriptor instead.
func (LogStatus) EnumDescriptor() ([]byte, []int) {
	return file_loggerpb_message_proto_rawDescGZIP(), []int{0}
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message         string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Status          LogStatus `protobuf:"varint,3,opt,name=status,proto3,enum=loggerpb.LogStatus" json:"status,omitempty"`
	FromApplication string    `protobuf:"bytes,4,opt,name=from_application,json=fromApplication,proto3" json:"from_application,omitempty"`
	DateTime        string    `protobuf:"bytes,5,opt,name=date_time,json=dateTime,proto3" json:"date_time,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loggerpb_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_loggerpb_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_loggerpb_message_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Log) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Log) GetStatus() LogStatus {
	if x != nil {
		return x.Status
	}
	return LogStatus_INFO
}

func (x *Log) GetFromApplication() string {
	if x != nil {
		return x.FromApplication
	}
	return ""
}

func (x *Log) GetDateTime() string {
	if x != nil {
		return x.DateTime
	}
	return ""
}

var File_loggerpb_message_proto protoreflect.FileDescriptor

var file_loggerpb_message_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x70, 0x62, 0x22, 0xa4, 0x01, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x70, 0x62, 0x2e,
	0x4c, 0x6f, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x29, 0x0a, 0x10, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x72, 0x6f,
	0x6d, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x2a, 0x50, 0x0a, 0x09, 0x4c, 0x6f, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x55,
	0x47, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x04, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x05, 0x42, 0xa5, 0x01, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x70, 0x62, 0x42, 0x0c, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x69, 0x66, 0x69, 0x79, 0x61, 0x2d,
	0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c, 0x2d, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x6f,
	0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x4c, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x4c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x70, 0x62, 0xca, 0x02, 0x08, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x70,
	0x62, 0xe2, 0x02, 0x14, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x70, 0x62, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x4c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loggerpb_message_proto_rawDescOnce sync.Once
	file_loggerpb_message_proto_rawDescData = file_loggerpb_message_proto_rawDesc
)

func file_loggerpb_message_proto_rawDescGZIP() []byte {
	file_loggerpb_message_proto_rawDescOnce.Do(func() {
		file_loggerpb_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_loggerpb_message_proto_rawDescData)
	})
	return file_loggerpb_message_proto_rawDescData
}

var file_loggerpb_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_loggerpb_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_loggerpb_message_proto_goTypes = []interface{}{
	(LogStatus)(0), // 0: loggerpb.LogStatus
	(*Log)(nil),    // 1: loggerpb.Log
}
var file_loggerpb_message_proto_depIdxs = []int32{
	0, // 0: loggerpb.Log.status:type_name -> loggerpb.LogStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_loggerpb_message_proto_init() }
func file_loggerpb_message_proto_init() {
	if File_loggerpb_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loggerpb_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
			RawDescriptor: file_loggerpb_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_loggerpb_message_proto_goTypes,
		DependencyIndexes: file_loggerpb_message_proto_depIdxs,
		EnumInfos:         file_loggerpb_message_proto_enumTypes,
		MessageInfos:      file_loggerpb_message_proto_msgTypes,
	}.Build()
	File_loggerpb_message_proto = out.File
	file_loggerpb_message_proto_rawDesc = nil
	file_loggerpb_message_proto_goTypes = nil
	file_loggerpb_message_proto_depIdxs = nil
}
