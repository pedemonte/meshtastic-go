// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: environmental_measurement.proto

package go_meshtastic_protobufs

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

type EnvironmentalMeasurement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperature        float32 `protobuf:"fixed32,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	RelativeHumidity   float32 `protobuf:"fixed32,2,opt,name=relative_humidity,json=relativeHumidity,proto3" json:"relative_humidity,omitempty"`
	BarometricPressure float32 `protobuf:"fixed32,3,opt,name=barometric_pressure,json=barometricPressure,proto3" json:"barometric_pressure,omitempty"`
}

func (x *EnvironmentalMeasurement) Reset() {
	*x = EnvironmentalMeasurement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_environmental_measurement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentalMeasurement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentalMeasurement) ProtoMessage() {}

func (x *EnvironmentalMeasurement) ProtoReflect() protoreflect.Message {
	mi := &file_environmental_measurement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentalMeasurement.ProtoReflect.Descriptor instead.
func (*EnvironmentalMeasurement) Descriptor() ([]byte, []int) {
	return file_environmental_measurement_proto_rawDescGZIP(), []int{0}
}

func (x *EnvironmentalMeasurement) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *EnvironmentalMeasurement) GetRelativeHumidity() float32 {
	if x != nil {
		return x.RelativeHumidity
	}
	return 0
}

func (x *EnvironmentalMeasurement) GetBarometricPressure() float32 {
	if x != nil {
		return x.BarometricPressure
	}
	return 0
}

var File_environmental_measurement_proto protoreflect.FileDescriptor

var file_environmental_measurement_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f,
	0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x18, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x68, 0x75, 0x6d,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x48, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x2f, 0x0a,
	0x13, 0x62, 0x61, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x62, 0x61, 0x72, 0x6f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x42, 0x1b,
	0x5a, 0x19, 0x2e, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_environmental_measurement_proto_rawDescOnce sync.Once
	file_environmental_measurement_proto_rawDescData = file_environmental_measurement_proto_rawDesc
)

func file_environmental_measurement_proto_rawDescGZIP() []byte {
	file_environmental_measurement_proto_rawDescOnce.Do(func() {
		file_environmental_measurement_proto_rawDescData = protoimpl.X.CompressGZIP(file_environmental_measurement_proto_rawDescData)
	})
	return file_environmental_measurement_proto_rawDescData
}

var file_environmental_measurement_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_environmental_measurement_proto_goTypes = []interface{}{
	(*EnvironmentalMeasurement)(nil), // 0: EnvironmentalMeasurement
}
var file_environmental_measurement_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_environmental_measurement_proto_init() }
func file_environmental_measurement_proto_init() {
	if File_environmental_measurement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_environmental_measurement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentalMeasurement); i {
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
			RawDescriptor: file_environmental_measurement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_environmental_measurement_proto_goTypes,
		DependencyIndexes: file_environmental_measurement_proto_depIdxs,
		MessageInfos:      file_environmental_measurement_proto_msgTypes,
	}.Build()
	File_environmental_measurement_proto = out.File
	file_environmental_measurement_proto_rawDesc = nil
	file_environmental_measurement_proto_goTypes = nil
	file_environmental_measurement_proto_depIdxs = nil
}
