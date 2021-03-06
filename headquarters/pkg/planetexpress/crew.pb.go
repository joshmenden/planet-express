// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: crew.proto

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

type CrewMember_Role int32

const (
	CrewMember_PILOT     CrewMember_Role = 0
	CrewMember_GUNMAN    CrewMember_Role = 1
	CrewMember_NAVIGATOR CrewMember_Role = 2
)

// Enum value maps for CrewMember_Role.
var (
	CrewMember_Role_name = map[int32]string{
		0: "PILOT",
		1: "GUNMAN",
		2: "NAVIGATOR",
	}
	CrewMember_Role_value = map[string]int32{
		"PILOT":     0,
		"GUNMAN":    1,
		"NAVIGATOR": 2,
	}
)

func (x CrewMember_Role) Enum() *CrewMember_Role {
	p := new(CrewMember_Role)
	*p = x
	return p
}

func (x CrewMember_Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CrewMember_Role) Descriptor() protoreflect.EnumDescriptor {
	return file_crew_proto_enumTypes[0].Descriptor()
}

func (CrewMember_Role) Type() protoreflect.EnumType {
	return &file_crew_proto_enumTypes[0]
}

func (x CrewMember_Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CrewMember_Role.Descriptor instead.
func (CrewMember_Role) EnumDescriptor() ([]byte, []int) {
	return file_crew_proto_rawDescGZIP(), []int{1, 0}
}

type Crew struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CrewMembers []*CrewMember `protobuf:"bytes,1,rep,name=crew_members,json=crewMembers,proto3" json:"crew_members,omitempty"`
	CrewName    string        `protobuf:"bytes,2,opt,name=crew_name,json=crewName,proto3" json:"crew_name,omitempty"`
}

func (x *Crew) Reset() {
	*x = Crew{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crew_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Crew) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Crew) ProtoMessage() {}

func (x *Crew) ProtoReflect() protoreflect.Message {
	mi := &file_crew_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Crew.ProtoReflect.Descriptor instead.
func (*Crew) Descriptor() ([]byte, []int) {
	return file_crew_proto_rawDescGZIP(), []int{0}
}

func (x *Crew) GetCrewMembers() []*CrewMember {
	if x != nil {
		return x.CrewMembers
	}
	return nil
}

func (x *Crew) GetCrewName() string {
	if x != nil {
		return x.CrewName
	}
	return ""
}

type CrewMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Role CrewMember_Role `protobuf:"varint,2,opt,name=role,proto3,enum=planetexpress.CrewMember_Role" json:"role,omitempty"`
}

func (x *CrewMember) Reset() {
	*x = CrewMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crew_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrewMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrewMember) ProtoMessage() {}

func (x *CrewMember) ProtoReflect() protoreflect.Message {
	mi := &file_crew_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrewMember.ProtoReflect.Descriptor instead.
func (*CrewMember) Descriptor() ([]byte, []int) {
	return file_crew_proto_rawDescGZIP(), []int{1}
}

func (x *CrewMember) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CrewMember) GetRole() CrewMember_Role {
	if x != nil {
		return x.Role
	}
	return CrewMember_PILOT
}

type GetCrewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crew *Crew `protobuf:"bytes,1,opt,name=crew,proto3" json:"crew,omitempty"`
}

func (x *GetCrewResponse) Reset() {
	*x = GetCrewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crew_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCrewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCrewResponse) ProtoMessage() {}

func (x *GetCrewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crew_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCrewResponse.ProtoReflect.Descriptor instead.
func (*GetCrewResponse) Descriptor() ([]byte, []int) {
	return file_crew_proto_rawDescGZIP(), []int{2}
}

func (x *GetCrewResponse) GetCrew() *Crew {
	if x != nil {
		return x.Crew
	}
	return nil
}

var File_crew_proto protoreflect.FileDescriptor

var file_crew_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x22, 0x61, 0x0a, 0x04, 0x43,
	0x72, 0x65, 0x77, 0x12, 0x3c, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x77, 0x5f, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x77, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x77, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x82,
	0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x77, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x32, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1e, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x77, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x2c, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x09, 0x0a,
	0x05, 0x50, 0x49, 0x4c, 0x4f, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x55, 0x4e, 0x4d,
	0x41, 0x4e, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x41, 0x56, 0x49, 0x47, 0x41, 0x54, 0x4f,
	0x52, 0x10, 0x02, 0x22, 0x3a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x77, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x63, 0x72, 0x65, 0x77, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x77, 0x52, 0x04, 0x63, 0x72, 0x65, 0x77, 0x42,
	0x13, 0x5a, 0x11, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crew_proto_rawDescOnce sync.Once
	file_crew_proto_rawDescData = file_crew_proto_rawDesc
)

func file_crew_proto_rawDescGZIP() []byte {
	file_crew_proto_rawDescOnce.Do(func() {
		file_crew_proto_rawDescData = protoimpl.X.CompressGZIP(file_crew_proto_rawDescData)
	})
	return file_crew_proto_rawDescData
}

var file_crew_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_crew_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_crew_proto_goTypes = []interface{}{
	(CrewMember_Role)(0),    // 0: planetexpress.CrewMember.Role
	(*Crew)(nil),            // 1: planetexpress.Crew
	(*CrewMember)(nil),      // 2: planetexpress.CrewMember
	(*GetCrewResponse)(nil), // 3: planetexpress.GetCrewResponse
}
var file_crew_proto_depIdxs = []int32{
	2, // 0: planetexpress.Crew.crew_members:type_name -> planetexpress.CrewMember
	0, // 1: planetexpress.CrewMember.role:type_name -> planetexpress.CrewMember.Role
	1, // 2: planetexpress.GetCrewResponse.crew:type_name -> planetexpress.Crew
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_crew_proto_init() }
func file_crew_proto_init() {
	if File_crew_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crew_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Crew); i {
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
		file_crew_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrewMember); i {
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
		file_crew_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCrewResponse); i {
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
			RawDescriptor: file_crew_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crew_proto_goTypes,
		DependencyIndexes: file_crew_proto_depIdxs,
		EnumInfos:         file_crew_proto_enumTypes,
		MessageInfos:      file_crew_proto_msgTypes,
	}.Build()
	File_crew_proto = out.File
	file_crew_proto_rawDesc = nil
	file_crew_proto_goTypes = nil
	file_crew_proto_depIdxs = nil
}
