// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: ship.proto

package planetexpress

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

type Ship_FuelLevel int32

const (
	Ship_EMPTY Ship_FuelLevel = 0
	Ship_LOW   Ship_FuelLevel = 1
	Ship_FULL  Ship_FuelLevel = 2
)

// Enum value maps for Ship_FuelLevel.
var (
	Ship_FuelLevel_name = map[int32]string{
		0: "EMPTY",
		1: "LOW",
		2: "FULL",
	}
	Ship_FuelLevel_value = map[string]int32{
		"EMPTY": 0,
		"LOW":   1,
		"FULL":  2,
	}
)

func (x Ship_FuelLevel) Enum() *Ship_FuelLevel {
	p := new(Ship_FuelLevel)
	*p = x
	return p
}

func (x Ship_FuelLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ship_FuelLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_ship_proto_enumTypes[0].Descriptor()
}

func (Ship_FuelLevel) Type() protoreflect.EnumType {
	return &file_ship_proto_enumTypes[0]
}

func (x Ship_FuelLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ship_FuelLevel.Descriptor instead.
func (Ship_FuelLevel) EnumDescriptor() ([]byte, []int) {
	return file_ship_proto_rawDescGZIP(), []int{0, 0}
}

type Ship struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location  string         `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	FuelLevel Ship_FuelLevel `protobuf:"varint,3,opt,name=fuel_level,json=fuelLevel,proto3,enum=planetexpress.Ship_FuelLevel" json:"fuel_level,omitempty"`
	Crew      *Crew          `protobuf:"bytes,4,opt,name=crew,proto3" json:"crew,omitempty"` // Delivery delivery = 5;
}

func (x *Ship) Reset() {
	*x = Ship{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ship_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ship) ProtoMessage() {}

func (x *Ship) ProtoReflect() protoreflect.Message {
	mi := &file_ship_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ship.ProtoReflect.Descriptor instead.
func (*Ship) Descriptor() ([]byte, []int) {
	return file_ship_proto_rawDescGZIP(), []int{0}
}

func (x *Ship) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Ship) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Ship) GetFuelLevel() Ship_FuelLevel {
	if x != nil {
		return x.FuelLevel
	}
	return Ship_EMPTY
}

func (x *Ship) GetCrew() *Crew {
	if x != nil {
		return x.Crew
	}
	return nil
}

type GetShipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ship *Ship `protobuf:"bytes,1,opt,name=ship,proto3" json:"ship,omitempty"`
}

func (x *GetShipResponse) Reset() {
	*x = GetShipResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ship_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShipResponse) ProtoMessage() {}

func (x *GetShipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ship_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShipResponse.ProtoReflect.Descriptor instead.
func (*GetShipResponse) Descriptor() ([]byte, []int) {
	return file_ship_proto_rawDescGZIP(), []int{1}
}

func (x *GetShipResponse) GetShip() *Ship {
	if x != nil {
		return x.Ship
	}
	return nil
}

var File_ship_proto protoreflect.FileDescriptor

var file_ship_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x0a, 0x63, 0x72, 0x65,
	0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x04, 0x53, 0x68, 0x69, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3c, 0x0a, 0x0a, 0x66, 0x75, 0x65, 0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x2e, 0x46, 0x75, 0x65, 0x6c, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x52, 0x09, 0x66, 0x75, 0x65, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x27,
	0x0a, 0x04, 0x63, 0x72, 0x65, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x43, 0x72, 0x65,
	0x77, 0x52, 0x04, 0x63, 0x72, 0x65, 0x77, 0x22, 0x29, 0x0a, 0x09, 0x46, 0x75, 0x65, 0x6c, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x55, 0x4c, 0x4c,
	0x10, 0x02, 0x22, 0x3a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x68, 0x69, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x73, 0x68, 0x69, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x52, 0x04, 0x73, 0x68, 0x69, 0x70, 0x42, 0x13,
	0x5a, 0x11, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ship_proto_rawDescOnce sync.Once
	file_ship_proto_rawDescData = file_ship_proto_rawDesc
)

func file_ship_proto_rawDescGZIP() []byte {
	file_ship_proto_rawDescOnce.Do(func() {
		file_ship_proto_rawDescData = protoimpl.X.CompressGZIP(file_ship_proto_rawDescData)
	})
	return file_ship_proto_rawDescData
}

var file_ship_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ship_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ship_proto_goTypes = []interface{}{
	(Ship_FuelLevel)(0),     // 0: planetexpress.Ship.FuelLevel
	(*Ship)(nil),            // 1: planetexpress.Ship
	(*GetShipResponse)(nil), // 2: planetexpress.GetShipResponse
	(*Crew)(nil),            // 3: planetexpress.Crew
}
var file_ship_proto_depIdxs = []int32{
	0, // 0: planetexpress.Ship.fuel_level:type_name -> planetexpress.Ship.FuelLevel
	3, // 1: planetexpress.Ship.crew:type_name -> planetexpress.Crew
	1, // 2: planetexpress.GetShipResponse.ship:type_name -> planetexpress.Ship
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ship_proto_init() }
func file_ship_proto_init() {
	if File_ship_proto != nil {
		return
	}
	file_crew_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ship_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ship); i {
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
		file_ship_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShipResponse); i {
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
			RawDescriptor: file_ship_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ship_proto_goTypes,
		DependencyIndexes: file_ship_proto_depIdxs,
		EnumInfos:         file_ship_proto_enumTypes,
		MessageInfos:      file_ship_proto_msgTypes,
	}.Build()
	File_ship_proto = out.File
	file_ship_proto_rawDesc = nil
	file_ship_proto_goTypes = nil
	file_ship_proto_depIdxs = nil
}
