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

type Crew struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberOfCrewmates int32  `protobuf:"varint,1,opt,name=number_of_crewmates,json=numberOfCrewmates,proto3" json:"number_of_crewmates,omitempty"`
	CrewName          string `protobuf:"bytes,2,opt,name=crew_name,json=crewName,proto3" json:"crew_name,omitempty"`
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

func (x *Crew) GetNumberOfCrewmates() int32 {
	if x != nil {
		return x.NumberOfCrewmates
	}
	return 0
}

func (x *Crew) GetCrewName() string {
	if x != nil {
		return x.CrewName
	}
	return ""
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
		mi := &file_crew_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCrewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCrewResponse) ProtoMessage() {}

func (x *GetCrewResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCrewResponse.ProtoReflect.Descriptor instead.
func (*GetCrewResponse) Descriptor() ([]byte, []int) {
	return file_crew_proto_rawDescGZIP(), []int{1}
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
	0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x22, 0x53, 0x0a, 0x04, 0x43,
	0x72, 0x65, 0x77, 0x12, 0x2e, 0x0a, 0x13, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66,
	0x5f, 0x63, 0x72, 0x65, 0x77, 0x6d, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x11, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x43, 0x72, 0x65, 0x77, 0x6d, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x3a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x63, 0x72, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x43, 0x72, 0x65, 0x77, 0x52, 0x04, 0x63, 0x72, 0x65, 0x77, 0x42, 0x13, 0x5a, 0x11,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_crew_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_crew_proto_goTypes = []interface{}{
	(*Crew)(nil),            // 0: planetexpress.Crew
	(*GetCrewResponse)(nil), // 1: planetexpress.GetCrewResponse
}
var file_crew_proto_depIdxs = []int32{
	0, // 0: planetexpress.GetCrewResponse.crew:type_name -> planetexpress.Crew
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crew_proto_goTypes,
		DependencyIndexes: file_crew_proto_depIdxs,
		MessageInfos:      file_crew_proto_msgTypes,
	}.Build()
	File_crew_proto = out.File
	file_crew_proto_rawDesc = nil
	file_crew_proto_goTypes = nil
	file_crew_proto_depIdxs = nil
}
