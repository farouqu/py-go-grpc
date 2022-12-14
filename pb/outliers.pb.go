// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: outliers.proto

package pb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time  *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Name  string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value float64              `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outliers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_outliers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_outliers_proto_rawDescGZIP(), []int{0}
}

func (x *Metric) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Metric) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metric) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type OutliersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *OutliersRequest) Reset() {
	*x = OutliersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outliers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutliersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutliersRequest) ProtoMessage() {}

func (x *OutliersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_outliers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutliersRequest.ProtoReflect.Descriptor instead.
func (*OutliersRequest) Descriptor() ([]byte, []int) {
	return file_outliers_proto_rawDescGZIP(), []int{1}
}

func (x *OutliersRequest) GetMetrics() []*Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type OutliersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Indices []int32 `protobuf:"varint,1,rep,packed,name=indices,proto3" json:"indices,omitempty"`
}

func (x *OutliersResponse) Reset() {
	*x = OutliersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outliers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutliersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutliersResponse) ProtoMessage() {}

func (x *OutliersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_outliers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutliersResponse.ProtoReflect.Descriptor instead.
func (*OutliersResponse) Descriptor() ([]byte, []int) {
	return file_outliers_proto_rawDescGZIP(), []int{2}
}

func (x *OutliersResponse) GetIndices() []int32 {
	if x != nil {
		return x.Indices
	}
	return nil
}

var File_outliers_proto protoreflect.FileDescriptor

var file_outliers_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12,
	0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x37, 0x0a, 0x0f, 0x4f, 0x75, 0x74,
	0x6c, 0x69, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x07,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x22, 0x2c, 0x0a, 0x10, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73,
	0x32, 0x41, 0x0a, 0x08, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x12, 0x35, 0x0a, 0x06,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x75, 0x74, 0x6c,
	0x69, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62,
	0x2e, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x66, 0x61, 0x72, 0x6f, 0x75, 0x71, 0x75, 0x2f, 0x70, 0x79, 0x2d, 0x67, 0x6f, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_outliers_proto_rawDescOnce sync.Once
	file_outliers_proto_rawDescData = file_outliers_proto_rawDesc
)

func file_outliers_proto_rawDescGZIP() []byte {
	file_outliers_proto_rawDescOnce.Do(func() {
		file_outliers_proto_rawDescData = protoimpl.X.CompressGZIP(file_outliers_proto_rawDescData)
	})
	return file_outliers_proto_rawDescData
}

var file_outliers_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_outliers_proto_goTypes = []interface{}{
	(*Metric)(nil),              // 0: pb.Metric
	(*OutliersRequest)(nil),     // 1: pb.OutliersRequest
	(*OutliersResponse)(nil),    // 2: pb.OutliersResponse
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_outliers_proto_depIdxs = []int32{
	3, // 0: pb.Metric.time:type_name -> google.protobuf.Timestamp
	0, // 1: pb.OutliersRequest.metrics:type_name -> pb.Metric
	1, // 2: pb.Outliers.Detect:input_type -> pb.OutliersRequest
	2, // 3: pb.Outliers.Detect:output_type -> pb.OutliersResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_outliers_proto_init() }
func file_outliers_proto_init() {
	if File_outliers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_outliers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
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
		file_outliers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutliersRequest); i {
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
		file_outliers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutliersResponse); i {
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
			RawDescriptor: file_outliers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_outliers_proto_goTypes,
		DependencyIndexes: file_outliers_proto_depIdxs,
		MessageInfos:      file_outliers_proto_msgTypes,
	}.Build()
	File_outliers_proto = out.File
	file_outliers_proto_rawDesc = nil
	file_outliers_proto_goTypes = nil
	file_outliers_proto_depIdxs = nil
}
