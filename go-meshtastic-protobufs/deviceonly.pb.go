// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: deviceonly.proto

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

// This is a stub version of the old 1.1 representation of RadioConfig.  But only keeping the region info.  The device firmware uses
//this stub while migrating old nodes to the new preferences system.
type LegacyRadioConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Preferences *LegacyRadioConfig_LegacyPreferences `protobuf:"bytes,1,opt,name=preferences,proto3" json:"preferences,omitempty"`
}

func (x *LegacyRadioConfig) Reset() {
	*x = LegacyRadioConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deviceonly_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LegacyRadioConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LegacyRadioConfig) ProtoMessage() {}

func (x *LegacyRadioConfig) ProtoReflect() protoreflect.Message {
	mi := &file_deviceonly_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LegacyRadioConfig.ProtoReflect.Descriptor instead.
func (*LegacyRadioConfig) Descriptor() ([]byte, []int) {
	return file_deviceonly_proto_rawDescGZIP(), []int{0}
}

func (x *LegacyRadioConfig) GetPreferences() *LegacyRadioConfig_LegacyPreferences {
	if x != nil {
		return x.Preferences
	}
	return nil
}

//
// This message is never sent over the wire, but it is used for serializing DB
// state to flash in the device code
// FIXME, since we write this each time we enter deep sleep (and have infinite
// flash) it would be better to use some sort of append only data structure for
// the receive queue and use the preferences store for the other stuff
type DeviceState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Moved to its own file, but we keep this here so we can automatically migrate old radio.region settings
	LegacyRadio *LegacyRadioConfig `protobuf:"bytes,1,opt,name=legacyRadio,proto3" json:"legacyRadio,omitempty"`
	//
	// Read only settings/info about this node
	MyNode *MyNodeInfo `protobuf:"bytes,2,opt,name=my_node,json=myNode,proto3" json:"my_node,omitempty"`
	//
	// My owner info
	Owner  *User       `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	NodeDb []*NodeInfo `protobuf:"bytes,4,rep,name=node_db,json=nodeDb,proto3" json:"node_db,omitempty"`
	//
	// Received packets saved for delivery to the phone
	ReceiveQueue []*MeshPacket `protobuf:"bytes,5,rep,name=receive_queue,json=receiveQueue,proto3" json:"receive_queue,omitempty"`
	//
	// A version integer used to invalidate old save files when we make
	// incompatible changes This integer is set at build time and is private to
	// NodeDB.cpp in the device code.
	Version uint32 `protobuf:"varint,8,opt,name=version,proto3" json:"version,omitempty"`
	//
	// We keep the last received text message (only) stored in the device flash,
	// so we can show it on the screen.
	// Might be null
	RxTextMessage *MeshPacket `protobuf:"bytes,7,opt,name=rx_text_message,json=rxTextMessage,proto3" json:"rx_text_message,omitempty"`
	//
	// Used only during development.  Indicates developer is testing and changes
	// should never be saved to flash.
	NoSave bool `protobuf:"varint,9,opt,name=no_save,json=noSave,proto3" json:"no_save,omitempty"`
	//
	// Some GPSes seem to have bogus settings from the factory, so we always do one factory reset.
	DidGpsReset bool `protobuf:"varint,11,opt,name=did_gps_reset,json=didGpsReset,proto3" json:"did_gps_reset,omitempty"`
}

func (x *DeviceState) Reset() {
	*x = DeviceState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deviceonly_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceState) ProtoMessage() {}

func (x *DeviceState) ProtoReflect() protoreflect.Message {
	mi := &file_deviceonly_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceState.ProtoReflect.Descriptor instead.
func (*DeviceState) Descriptor() ([]byte, []int) {
	return file_deviceonly_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceState) GetLegacyRadio() *LegacyRadioConfig {
	if x != nil {
		return x.LegacyRadio
	}
	return nil
}

func (x *DeviceState) GetMyNode() *MyNodeInfo {
	if x != nil {
		return x.MyNode
	}
	return nil
}

func (x *DeviceState) GetOwner() *User {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *DeviceState) GetNodeDb() []*NodeInfo {
	if x != nil {
		return x.NodeDb
	}
	return nil
}

func (x *DeviceState) GetReceiveQueue() []*MeshPacket {
	if x != nil {
		return x.ReceiveQueue
	}
	return nil
}

func (x *DeviceState) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *DeviceState) GetRxTextMessage() *MeshPacket {
	if x != nil {
		return x.RxTextMessage
	}
	return nil
}

func (x *DeviceState) GetNoSave() bool {
	if x != nil {
		return x.NoSave
	}
	return false
}

func (x *DeviceState) GetDidGpsReset() bool {
	if x != nil {
		return x.DidGpsReset
	}
	return false
}

//* The on-disk saved channels
type ChannelFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	// The channels our node knows about
	Channels []*Channel `protobuf:"bytes,1,rep,name=channels,proto3" json:"channels,omitempty"`
}

func (x *ChannelFile) Reset() {
	*x = ChannelFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deviceonly_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelFile) ProtoMessage() {}

func (x *ChannelFile) ProtoReflect() protoreflect.Message {
	mi := &file_deviceonly_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelFile.ProtoReflect.Descriptor instead.
func (*ChannelFile) Descriptor() ([]byte, []int) {
	return file_deviceonly_proto_rawDescGZIP(), []int{2}
}

func (x *ChannelFile) GetChannels() []*Channel {
	if x != nil {
		return x.Channels
	}
	return nil
}

type LegacyRadioConfig_LegacyPreferences struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	// The region code for my radio (US, CN, EU433, etc...)
	Region RegionCode `protobuf:"varint,15,opt,name=region,proto3,enum=RegionCode" json:"region,omitempty"`
}

func (x *LegacyRadioConfig_LegacyPreferences) Reset() {
	*x = LegacyRadioConfig_LegacyPreferences{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deviceonly_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LegacyRadioConfig_LegacyPreferences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LegacyRadioConfig_LegacyPreferences) ProtoMessage() {}

func (x *LegacyRadioConfig_LegacyPreferences) ProtoReflect() protoreflect.Message {
	mi := &file_deviceonly_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LegacyRadioConfig_LegacyPreferences.ProtoReflect.Descriptor instead.
func (*LegacyRadioConfig_LegacyPreferences) Descriptor() ([]byte, []int) {
	return file_deviceonly_proto_rawDescGZIP(), []int{0, 0}
}

func (x *LegacyRadioConfig_LegacyPreferences) GetRegion() RegionCode {
	if x != nil {
		return x.Region
	}
	return RegionCode_Unset
}

var File_deviceonly_proto protoreflect.FileDescriptor

var file_deviceonly_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x6f, 0x6e, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0a, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x72,
	0x61, 0x64, 0x69, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x95, 0x01, 0x0a, 0x11, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x52, 0x61, 0x64, 0x69, 0x6f,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x46, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x4c, 0x65,
	0x67, 0x61, 0x63, 0x79, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x1a, 0x38,
	0x0a, 0x11, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0xee, 0x02, 0x0a, 0x0b, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x6c, 0x65, 0x67, 0x61,
	0x63, 0x79, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0b, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x12, 0x24,
	0x0a, 0x07, 0x6d, 0x79, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x4d, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6d, 0x79,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x12, 0x22, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x64, 0x62, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6e,
	0x6f, 0x64, 0x65, 0x44, 0x62, 0x12, 0x30, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d,
	0x65, 0x73, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x33, 0x0a, 0x0f, 0x72, 0x78, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x65, 0x73,
	0x68, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0d, 0x72, 0x78, 0x54, 0x65, 0x78, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x5f, 0x73, 0x61, 0x76,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6e, 0x6f, 0x53, 0x61, 0x76, 0x65, 0x12,
	0x22, 0x0a, 0x0d, 0x64, 0x69, 0x64, 0x5f, 0x67, 0x70, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x64, 0x69, 0x64, 0x47, 0x70, 0x73, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x4a, 0x04, 0x08, 0x0c, 0x10, 0x0d, 0x22, 0x33, 0x0a, 0x0b, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x42, 0x3e,
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x73, 0x76, 0x69, 0x6c, 0x6c, 0x65,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x42, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x6e, 0x6c,
	0x79, 0x48, 0x03, 0x5a, 0x19, 0x2e, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61,
	0x73, 0x74, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deviceonly_proto_rawDescOnce sync.Once
	file_deviceonly_proto_rawDescData = file_deviceonly_proto_rawDesc
)

func file_deviceonly_proto_rawDescGZIP() []byte {
	file_deviceonly_proto_rawDescOnce.Do(func() {
		file_deviceonly_proto_rawDescData = protoimpl.X.CompressGZIP(file_deviceonly_proto_rawDescData)
	})
	return file_deviceonly_proto_rawDescData
}

var file_deviceonly_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_deviceonly_proto_goTypes = []interface{}{
	(*LegacyRadioConfig)(nil),                   // 0: LegacyRadioConfig
	(*DeviceState)(nil),                         // 1: DeviceState
	(*ChannelFile)(nil),                         // 2: ChannelFile
	(*LegacyRadioConfig_LegacyPreferences)(nil), // 3: LegacyRadioConfig.LegacyPreferences
	(*MyNodeInfo)(nil),                          // 4: MyNodeInfo
	(*User)(nil),                                // 5: User
	(*NodeInfo)(nil),                            // 6: NodeInfo
	(*MeshPacket)(nil),                          // 7: MeshPacket
	(*Channel)(nil),                             // 8: Channel
	(RegionCode)(0),                             // 9: RegionCode
}
var file_deviceonly_proto_depIdxs = []int32{
	3, // 0: LegacyRadioConfig.preferences:type_name -> LegacyRadioConfig.LegacyPreferences
	0, // 1: DeviceState.legacyRadio:type_name -> LegacyRadioConfig
	4, // 2: DeviceState.my_node:type_name -> MyNodeInfo
	5, // 3: DeviceState.owner:type_name -> User
	6, // 4: DeviceState.node_db:type_name -> NodeInfo
	7, // 5: DeviceState.receive_queue:type_name -> MeshPacket
	7, // 6: DeviceState.rx_text_message:type_name -> MeshPacket
	8, // 7: ChannelFile.channels:type_name -> Channel
	9, // 8: LegacyRadioConfig.LegacyPreferences.region:type_name -> RegionCode
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_deviceonly_proto_init() }
func file_deviceonly_proto_init() {
	if File_deviceonly_proto != nil {
		return
	}
	file_mesh_proto_init()
	file_channel_proto_init()
	file_radioconfig_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_deviceonly_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LegacyRadioConfig); i {
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
		file_deviceonly_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceState); i {
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
		file_deviceonly_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelFile); i {
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
		file_deviceonly_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LegacyRadioConfig_LegacyPreferences); i {
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
			RawDescriptor: file_deviceonly_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_deviceonly_proto_goTypes,
		DependencyIndexes: file_deviceonly_proto_depIdxs,
		MessageInfos:      file_deviceonly_proto_msgTypes,
	}.Build()
	File_deviceonly_proto = out.File
	file_deviceonly_proto_rawDesc = nil
	file_deviceonly_proto_goTypes = nil
	file_deviceonly_proto_depIdxs = nil
}
