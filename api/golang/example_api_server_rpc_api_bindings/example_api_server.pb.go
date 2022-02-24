// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: example_api_server.proto

package example_api_server_rpc_api_bindings

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// ==============================================================================================
//                                         AddPerson
// ==============================================================================================
type AddPersonArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//The person ID
	PersonId string `protobuf:"bytes,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
}

func (x *AddPersonArgs) Reset() {
	*x = AddPersonArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_api_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPersonArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPersonArgs) ProtoMessage() {}

func (x *AddPersonArgs) ProtoReflect() protoreflect.Message {
	mi := &file_example_api_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPersonArgs.ProtoReflect.Descriptor instead.
func (*AddPersonArgs) Descriptor() ([]byte, []int) {
	return file_example_api_server_proto_rawDescGZIP(), []int{0}
}

func (x *AddPersonArgs) GetPersonId() string {
	if x != nil {
		return x.PersonId
	}
	return ""
}

// ==============================================================================================
//                                         GetPerson
// ==============================================================================================
type GetPersonArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//The person ID
	PersonId string `protobuf:"bytes,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
}

func (x *GetPersonArgs) Reset() {
	*x = GetPersonArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_api_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonArgs) ProtoMessage() {}

func (x *GetPersonArgs) ProtoReflect() protoreflect.Message {
	mi := &file_example_api_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonArgs.ProtoReflect.Descriptor instead.
func (*GetPersonArgs) Descriptor() ([]byte, []int) {
	return file_example_api_server_proto_rawDescGZIP(), []int{1}
}

func (x *GetPersonArgs) GetPersonId() string {
	if x != nil {
		return x.PersonId
	}
	return ""
}

type GetPersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//The amount of book read by the person
	BooksRead uint32 `protobuf:"varint,1,opt,name=books_read,json=booksRead,proto3" json:"books_read,omitempty"`
}

func (x *GetPersonResponse) Reset() {
	*x = GetPersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_api_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonResponse) ProtoMessage() {}

func (x *GetPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_api_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonResponse.ProtoReflect.Descriptor instead.
func (*GetPersonResponse) Descriptor() ([]byte, []int) {
	return file_example_api_server_proto_rawDescGZIP(), []int{2}
}

func (x *GetPersonResponse) GetBooksRead() uint32 {
	if x != nil {
		return x.BooksRead
	}
	return 0
}

// ==============================================================================================
//                                      IncrementBooksRead
// ==============================================================================================
type IncrementBooksReadArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//The person ID
	PersonId string `protobuf:"bytes,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
}

func (x *IncrementBooksReadArgs) Reset() {
	*x = IncrementBooksReadArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_api_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrementBooksReadArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrementBooksReadArgs) ProtoMessage() {}

func (x *IncrementBooksReadArgs) ProtoReflect() protoreflect.Message {
	mi := &file_example_api_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrementBooksReadArgs.ProtoReflect.Descriptor instead.
func (*IncrementBooksReadArgs) Descriptor() ([]byte, []int) {
	return file_example_api_server_proto_rawDescGZIP(), []int{3}
}

func (x *IncrementBooksReadArgs) GetPersonId() string {
	if x != nil {
		return x.PersonId
	}
	return ""
}

var File_example_api_server_proto protoreflect.FileDescriptor

var file_example_api_server_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61,
	0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2c, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x41, 0x72, 0x67, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x2c, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x41, 0x72, 0x67, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x61, 0x64, 0x22,
	0x35, 0x0a, 0x16, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x73, 0x52, 0x65, 0x61, 0x64, 0x41, 0x72, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x32, 0xe9, 0x02, 0x0a, 0x17, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x49, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x12, 0x25, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x5f, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x25,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x29, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x5e, 0x0a, 0x12, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x2e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x52, 0x65, 0x61, 0x64, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x42, 0x5c, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x75, 0x72, 0x74, 0x6f, 0x73, 0x69, 0x73, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_api_server_proto_rawDescOnce sync.Once
	file_example_api_server_proto_rawDescData = file_example_api_server_proto_rawDesc
)

func file_example_api_server_proto_rawDescGZIP() []byte {
	file_example_api_server_proto_rawDescOnce.Do(func() {
		file_example_api_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_api_server_proto_rawDescData)
	})
	return file_example_api_server_proto_rawDescData
}

var file_example_api_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_example_api_server_proto_goTypes = []interface{}{
	(*AddPersonArgs)(nil),          // 0: example_api_server_api.AddPersonArgs
	(*GetPersonArgs)(nil),          // 1: example_api_server_api.GetPersonArgs
	(*GetPersonResponse)(nil),      // 2: example_api_server_api.GetPersonResponse
	(*IncrementBooksReadArgs)(nil), // 3: example_api_server_api.IncrementBooksReadArgs
	(*emptypb.Empty)(nil),          // 4: google.protobuf.Empty
}
var file_example_api_server_proto_depIdxs = []int32{
	4, // 0: example_api_server_api.ExampleAPIServerService.IsAvailable:input_type -> google.protobuf.Empty
	0, // 1: example_api_server_api.ExampleAPIServerService.AddPerson:input_type -> example_api_server_api.AddPersonArgs
	1, // 2: example_api_server_api.ExampleAPIServerService.GetPerson:input_type -> example_api_server_api.GetPersonArgs
	3, // 3: example_api_server_api.ExampleAPIServerService.IncrementBooksRead:input_type -> example_api_server_api.IncrementBooksReadArgs
	4, // 4: example_api_server_api.ExampleAPIServerService.IsAvailable:output_type -> google.protobuf.Empty
	4, // 5: example_api_server_api.ExampleAPIServerService.AddPerson:output_type -> google.protobuf.Empty
	2, // 6: example_api_server_api.ExampleAPIServerService.GetPerson:output_type -> example_api_server_api.GetPersonResponse
	4, // 7: example_api_server_api.ExampleAPIServerService.IncrementBooksRead:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_example_api_server_proto_init() }
func file_example_api_server_proto_init() {
	if File_example_api_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_api_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPersonArgs); i {
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
		file_example_api_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonArgs); i {
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
		file_example_api_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonResponse); i {
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
		file_example_api_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrementBooksReadArgs); i {
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
			RawDescriptor: file_example_api_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_api_server_proto_goTypes,
		DependencyIndexes: file_example_api_server_proto_depIdxs,
		MessageInfos:      file_example_api_server_proto_msgTypes,
	}.Build()
	File_example_api_server_proto = out.File
	file_example_api_server_proto_rawDesc = nil
	file_example_api_server_proto_goTypes = nil
	file_example_api_server_proto_depIdxs = nil
}
