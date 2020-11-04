// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/FileGetContents.proto

package proto

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

// Get the contents of a file. The content field is empty (no bytes) if the file is empty.
type FileGetContentsQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"` // Standard info sent from client to node, including the signed payment, and what kind of response is requested (cost, state proof, both, or neither).
	FileID *FileID      `protobuf:"bytes,2,opt,name=fileID,proto3" json:"fileID,omitempty"` // The file ID of the file whose contents are requested
}

func (x *FileGetContentsQuery) Reset() {
	*x = FileGetContentsQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_FileGetContents_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileGetContentsQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileGetContentsQuery) ProtoMessage() {}

func (x *FileGetContentsQuery) ProtoReflect() protoreflect.Message {
	mi := &file_proto_FileGetContents_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileGetContentsQuery.ProtoReflect.Descriptor instead.
func (*FileGetContentsQuery) Descriptor() ([]byte, []int) {
	return file_proto_FileGetContents_proto_rawDescGZIP(), []int{0}
}

func (x *FileGetContentsQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *FileGetContentsQuery) GetFileID() *FileID {
	if x != nil {
		return x.FileID
	}
	return nil
}

// Response when the client sends the node FileGetContentsQuery
type FileGetContentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header       *ResponseHeader                       `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`             //Standard response from node to client, including the requested fields: cost, or state proof, or both, or neither
	FileContents *FileGetContentsResponse_FileContents `protobuf:"bytes,2,opt,name=fileContents,proto3" json:"fileContents,omitempty"` // the file ID and contents (a state proof can be generated for this)
}

func (x *FileGetContentsResponse) Reset() {
	*x = FileGetContentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_FileGetContents_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileGetContentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileGetContentsResponse) ProtoMessage() {}

func (x *FileGetContentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_FileGetContents_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileGetContentsResponse.ProtoReflect.Descriptor instead.
func (*FileGetContentsResponse) Descriptor() ([]byte, []int) {
	return file_proto_FileGetContents_proto_rawDescGZIP(), []int{1}
}

func (x *FileGetContentsResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *FileGetContentsResponse) GetFileContents() *FileGetContentsResponse_FileContents {
	if x != nil {
		return x.FileContents
	}
	return nil
}

type FileGetContentsResponse_FileContents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileID   *FileID `protobuf:"bytes,1,opt,name=fileID,proto3" json:"fileID,omitempty"`     // The file ID of the file whose contents are being returned
	Contents []byte  `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"` // The bytes contained in the file
}

func (x *FileGetContentsResponse_FileContents) Reset() {
	*x = FileGetContentsResponse_FileContents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_FileGetContents_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileGetContentsResponse_FileContents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileGetContentsResponse_FileContents) ProtoMessage() {}

func (x *FileGetContentsResponse_FileContents) ProtoReflect() protoreflect.Message {
	mi := &file_proto_FileGetContents_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileGetContentsResponse_FileContents.ProtoReflect.Descriptor instead.
func (*FileGetContentsResponse_FileContents) Descriptor() ([]byte, []int) {
	return file_proto_FileGetContents_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FileGetContentsResponse_FileContents) GetFileID() *FileID {
	if x != nil {
		return x.FileID
	}
	return nil
}

func (x *FileGetContentsResponse_FileContents) GetContents() []byte {
	if x != nil {
		return x.Contents
	}
	return nil
}

var File_proto_FileGetContents_proto protoreflect.FileDescriptor

var file_proto_FileGetContents_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x69, 0x0a, 0x14, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x44, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x22, 0xec, 0x01, 0x0a,
	0x17, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x4f, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x51, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x48, 0x0a, 0x1a, 0x63,
	0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2f, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_FileGetContents_proto_rawDescOnce sync.Once
	file_proto_FileGetContents_proto_rawDescData = file_proto_FileGetContents_proto_rawDesc
)

func file_proto_FileGetContents_proto_rawDescGZIP() []byte {
	file_proto_FileGetContents_proto_rawDescOnce.Do(func() {
		file_proto_FileGetContents_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_FileGetContents_proto_rawDescData)
	})
	return file_proto_FileGetContents_proto_rawDescData
}

var file_proto_FileGetContents_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_FileGetContents_proto_goTypes = []interface{}{
	(*FileGetContentsQuery)(nil),                 // 0: proto.FileGetContentsQuery
	(*FileGetContentsResponse)(nil),              // 1: proto.FileGetContentsResponse
	(*FileGetContentsResponse_FileContents)(nil), // 2: proto.FileGetContentsResponse.FileContents
	(*QueryHeader)(nil),                          // 3: proto.QueryHeader
	(*FileID)(nil),                               // 4: proto.FileID
	(*ResponseHeader)(nil),                       // 5: proto.ResponseHeader
}
var file_proto_FileGetContents_proto_depIdxs = []int32{
	3, // 0: proto.FileGetContentsQuery.header:type_name -> proto.QueryHeader
	4, // 1: proto.FileGetContentsQuery.fileID:type_name -> proto.FileID
	5, // 2: proto.FileGetContentsResponse.header:type_name -> proto.ResponseHeader
	2, // 3: proto.FileGetContentsResponse.fileContents:type_name -> proto.FileGetContentsResponse.FileContents
	4, // 4: proto.FileGetContentsResponse.FileContents.fileID:type_name -> proto.FileID
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_FileGetContents_proto_init() }
func file_proto_FileGetContents_proto_init() {
	if File_proto_FileGetContents_proto != nil {
		return
	}
	file_proto_BasicTypes_proto_init()
	file_proto_QueryHeader_proto_init()
	file_proto_ResponseHeader_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_FileGetContents_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileGetContentsQuery); i {
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
		file_proto_FileGetContents_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileGetContentsResponse); i {
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
		file_proto_FileGetContents_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileGetContentsResponse_FileContents); i {
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
			RawDescriptor: file_proto_FileGetContents_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_FileGetContents_proto_goTypes,
		DependencyIndexes: file_proto_FileGetContents_proto_depIdxs,
		MessageInfos:      file_proto_FileGetContents_proto_msgTypes,
	}.Build()
	File_proto_FileGetContents_proto = out.File
	file_proto_FileGetContents_proto_rawDesc = nil
	file_proto_FileGetContents_proto_goTypes = nil
	file_proto_FileGetContents_proto_depIdxs = nil
}
