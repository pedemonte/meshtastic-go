// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: admin.proto

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

//
// This message is handled by the Admin plugin and is responsible for all settings/channel read/write operations.  This message
// is used to do settings operations to both remote AND local nodes.
// (Prior to 1.2 these operations were done via special ToRadio operations)
type AdminMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Variant:
	//	*AdminMessage_SetRadio
	//	*AdminMessage_SetOwner
	//	*AdminMessage_SetChannel
	//	*AdminMessage_GetRadioRequest
	//	*AdminMessage_GetRadioResponse
	//	*AdminMessage_GetChannelRequest
	//	*AdminMessage_GetChannelResponse
	//	*AdminMessage_ConfirmSetChannel
	//	*AdminMessage_ConfirmSetRadio
	//	*AdminMessage_ExitSimulator
	//	*AdminMessage_RebootSeconds
	Variant isAdminMessage_Variant `protobuf_oneof:"variant"`
}

func (x *AdminMessage) Reset() {
	*x = AdminMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminMessage) ProtoMessage() {}

func (x *AdminMessage) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminMessage.ProtoReflect.Descriptor instead.
func (*AdminMessage) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{0}
}

func (m *AdminMessage) GetVariant() isAdminMessage_Variant {
	if m != nil {
		return m.Variant
	}
	return nil
}

func (x *AdminMessage) GetSetRadio() *RadioConfig {
	if x, ok := x.GetVariant().(*AdminMessage_SetRadio); ok {
		return x.SetRadio
	}
	return nil
}

func (x *AdminMessage) GetSetOwner() *User {
	if x, ok := x.GetVariant().(*AdminMessage_SetOwner); ok {
		return x.SetOwner
	}
	return nil
}

func (x *AdminMessage) GetSetChannel() *Channel {
	if x, ok := x.GetVariant().(*AdminMessage_SetChannel); ok {
		return x.SetChannel
	}
	return nil
}

func (x *AdminMessage) GetGetRadioRequest() bool {
	if x, ok := x.GetVariant().(*AdminMessage_GetRadioRequest); ok {
		return x.GetRadioRequest
	}
	return false
}

func (x *AdminMessage) GetGetRadioResponse() *RadioConfig {
	if x, ok := x.GetVariant().(*AdminMessage_GetRadioResponse); ok {
		return x.GetRadioResponse
	}
	return nil
}

func (x *AdminMessage) GetGetChannelRequest() uint32 {
	if x, ok := x.GetVariant().(*AdminMessage_GetChannelRequest); ok {
		return x.GetChannelRequest
	}
	return 0
}

func (x *AdminMessage) GetGetChannelResponse() *Channel {
	if x, ok := x.GetVariant().(*AdminMessage_GetChannelResponse); ok {
		return x.GetChannelResponse
	}
	return nil
}

func (x *AdminMessage) GetConfirmSetChannel() bool {
	if x, ok := x.GetVariant().(*AdminMessage_ConfirmSetChannel); ok {
		return x.ConfirmSetChannel
	}
	return false
}

func (x *AdminMessage) GetConfirmSetRadio() bool {
	if x, ok := x.GetVariant().(*AdminMessage_ConfirmSetRadio); ok {
		return x.ConfirmSetRadio
	}
	return false
}

func (x *AdminMessage) GetExitSimulator() bool {
	if x, ok := x.GetVariant().(*AdminMessage_ExitSimulator); ok {
		return x.ExitSimulator
	}
	return false
}

func (x *AdminMessage) GetRebootSeconds() int32 {
	if x, ok := x.GetVariant().(*AdminMessage_RebootSeconds); ok {
		return x.RebootSeconds
	}
	return 0
}

type isAdminMessage_Variant interface {
	isAdminMessage_Variant()
}

type AdminMessage_SetRadio struct {
	//
	// set the radio provisioning for this node
	SetRadio *RadioConfig `protobuf:"bytes,1,opt,name=set_radio,json=setRadio,proto3,oneof"`
}

type AdminMessage_SetOwner struct {
	//
	// Set the owner for this node
	SetOwner *User `protobuf:"bytes,2,opt,name=set_owner,json=setOwner,proto3,oneof"`
}

type AdminMessage_SetChannel struct {
	//
	// Set channels (using the new API).  A special channel is the "primary channel".  The other records are secondary channels.
	// Note: only one channel can be marked as primary.  If the client sets a particular channel to be primary, the previous channel will be set to SECONDARY automatically
	SetChannel *Channel `protobuf:"bytes,3,opt,name=set_channel,json=setChannel,proto3,oneof"`
}

type AdminMessage_GetRadioRequest struct {
	//
	// Send the current RadioConfig in the response for this message
	GetRadioRequest bool `protobuf:"varint,4,opt,name=get_radio_request,json=getRadioRequest,proto3,oneof"`
}

type AdminMessage_GetRadioResponse struct {
	GetRadioResponse *RadioConfig `protobuf:"bytes,5,opt,name=get_radio_response,json=getRadioResponse,proto3,oneof"`
}

type AdminMessage_GetChannelRequest struct {
	//
	// Send the specified channel in the response for this message
	// NOTE: This field is sent with the channel index + 1 (to ensure we never try to send 'zero' - which protobufs treats as not present)
	GetChannelRequest uint32 `protobuf:"varint,6,opt,name=get_channel_request,json=getChannelRequest,proto3,oneof"`
}

type AdminMessage_GetChannelResponse struct {
	GetChannelResponse *Channel `protobuf:"bytes,7,opt,name=get_channel_response,json=getChannelResponse,proto3,oneof"`
}

type AdminMessage_ConfirmSetChannel struct {
	// Setting channels/radio config remotely carries the risk that you might
	// send an invalid config and the radio never talks to your mesh again.
	// Therefore if setting either of these properties remotely, you must send
	// a confirm_xxx message within 10 minutes.  If you fail to do so, the radio
	// will assume loss of comms and revert your changes.
	// These messages are optional when changing the local node.
	ConfirmSetChannel bool `protobuf:"varint,32,opt,name=confirm_set_channel,json=confirmSetChannel,proto3,oneof"`
}

type AdminMessage_ConfirmSetRadio struct {
	ConfirmSetRadio bool `protobuf:"varint,33,opt,name=confirm_set_radio,json=confirmSetRadio,proto3,oneof"`
}

type AdminMessage_ExitSimulator struct {
	// This message is only supported for the simulator porduino build.  If received the
	// simulator will exit successfully/
	ExitSimulator bool `protobuf:"varint,34,opt,name=exit_simulator,json=exitSimulator,proto3,oneof"`
}

type AdminMessage_RebootSeconds struct {
	//* Tell the node to reboot in this many seconds (or <0 to cancel reboot)
	RebootSeconds int32 `protobuf:"varint,35,opt,name=reboot_seconds,json=rebootSeconds,proto3,oneof"`
}

func (*AdminMessage_SetRadio) isAdminMessage_Variant() {}

func (*AdminMessage_SetOwner) isAdminMessage_Variant() {}

func (*AdminMessage_SetChannel) isAdminMessage_Variant() {}

func (*AdminMessage_GetRadioRequest) isAdminMessage_Variant() {}

func (*AdminMessage_GetRadioResponse) isAdminMessage_Variant() {}

func (*AdminMessage_GetChannelRequest) isAdminMessage_Variant() {}

func (*AdminMessage_GetChannelResponse) isAdminMessage_Variant() {}

func (*AdminMessage_ConfirmSetChannel) isAdminMessage_Variant() {}

func (*AdminMessage_ConfirmSetRadio) isAdminMessage_Variant() {}

func (*AdminMessage_ExitSimulator) isAdminMessage_Variant() {}

func (*AdminMessage_RebootSeconds) isAdminMessage_Variant() {}

var File_admin_proto protoreflect.FileDescriptor

var file_admin_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x72, 0x61, 0x64, 0x69, 0x6f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x04, 0x0a, 0x0c,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x09,
	0x73, 0x65, 0x74, 0x5f, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52,
	0x08, 0x73, 0x65, 0x74, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x12, 0x24, 0x0a, 0x09, 0x73, 0x65, 0x74,
	0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x08, 0x73, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12,
	0x2b, 0x0a, 0x0b, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x48, 0x00,
	0x52, 0x0a, 0x73, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x2c, 0x0a, 0x11,
	0x67, 0x65, 0x74, 0x5f, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0f, 0x67, 0x65, 0x74, 0x52, 0x61,
	0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x12, 0x67, 0x65,
	0x74, 0x5f, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x10, 0x67, 0x65, 0x74, 0x52, 0x61, 0x64, 0x69, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x11, 0x67, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x14, 0x67, 0x65,
	0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x48, 0x00, 0x52, 0x12, 0x67, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18,
	0x20, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x53, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x18,
	0x21, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x53, 0x65, 0x74, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x12, 0x27, 0x0a, 0x0e, 0x65, 0x78, 0x69, 0x74,
	0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x22, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x0d, 0x65, 0x78, 0x69, 0x74, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x27, 0x0a, 0x0e, 0x72, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x23, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0d, 0x72, 0x65, 0x62,
	0x6f, 0x6f, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x42, 0x3f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x65, 0x65,
	0x6b, 0x73, 0x76, 0x69, 0x6c, 0x6c, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x42, 0x0b, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48, 0x03, 0x5a, 0x19, 0x2e, 0x2f, 0x67,
	0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_proto_rawDescOnce sync.Once
	file_admin_proto_rawDescData = file_admin_proto_rawDesc
)

func file_admin_proto_rawDescGZIP() []byte {
	file_admin_proto_rawDescOnce.Do(func() {
		file_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_proto_rawDescData)
	})
	return file_admin_proto_rawDescData
}

var file_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_admin_proto_goTypes = []interface{}{
	(*AdminMessage)(nil), // 0: AdminMessage
	(*RadioConfig)(nil),  // 1: RadioConfig
	(*User)(nil),         // 2: User
	(*Channel)(nil),      // 3: Channel
}
var file_admin_proto_depIdxs = []int32{
	1, // 0: AdminMessage.set_radio:type_name -> RadioConfig
	2, // 1: AdminMessage.set_owner:type_name -> User
	3, // 2: AdminMessage.set_channel:type_name -> Channel
	1, // 3: AdminMessage.get_radio_response:type_name -> RadioConfig
	3, // 4: AdminMessage.get_channel_response:type_name -> Channel
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_admin_proto_init() }
func file_admin_proto_init() {
	if File_admin_proto != nil {
		return
	}
	file_mesh_proto_init()
	file_radioconfig_proto_init()
	file_channel_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminMessage); i {
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
	file_admin_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AdminMessage_SetRadio)(nil),
		(*AdminMessage_SetOwner)(nil),
		(*AdminMessage_SetChannel)(nil),
		(*AdminMessage_GetRadioRequest)(nil),
		(*AdminMessage_GetRadioResponse)(nil),
		(*AdminMessage_GetChannelRequest)(nil),
		(*AdminMessage_GetChannelResponse)(nil),
		(*AdminMessage_ConfirmSetChannel)(nil),
		(*AdminMessage_ConfirmSetRadio)(nil),
		(*AdminMessage_ExitSimulator)(nil),
		(*AdminMessage_RebootSeconds)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_admin_proto_goTypes,
		DependencyIndexes: file_admin_proto_depIdxs,
		MessageInfos:      file_admin_proto_msgTypes,
	}.Build()
	File_admin_proto = out.File
	file_admin_proto_rawDesc = nil
	file_admin_proto_goTypes = nil
	file_admin_proto_depIdxs = nil
}
