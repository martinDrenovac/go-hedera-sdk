// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/ConsensusUpdateTopic.proto

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

// All fields left null will not be updated.
// See [ConsensusService.updateTopic()](#proto.ConsensusService)
type ConsensusUpdateTopicTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicID *TopicID `protobuf:"bytes,1,opt,name=topicID,proto3" json:"topicID,omitempty"`
	// Short publicly visible memo about the topic. No guarantee of uniqueness. Null for "do not update".
	Memo *StringValue `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty"`
	// Effective consensus timestamp at (and after) which all consensus transactions and queries will fail.
	// The expirationTime may be no longer than MAX_AUTORENEW_PERIOD (8000001 seconds) from the consensus timestamp of
	// this transaction.
	// On topics with no adminKey, extending the expirationTime is the only updateTopic option allowed on the topic.
	// If unspecified, no change.
	ExpirationTime *Timestamp `protobuf:"bytes,4,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	// Access control for update/delete of the topic.
	// If unspecified, no change.
	// If empty keyList - the adminKey is cleared.
	AdminKey *Key `protobuf:"bytes,6,opt,name=adminKey,proto3" json:"adminKey,omitempty"`
	// Access control for ConsensusService.submitMessage.
	// If unspecified, no change.
	// If empty keyList - the submitKey is cleared.
	SubmitKey *Key `protobuf:"bytes,7,opt,name=submitKey,proto3" json:"submitKey,omitempty"`
	// The amount of time to extend the topic's lifetime automatically at expirationTime if the autoRenewAccount is
	// configured and has funds (once autoRenew functionality is supported by HAPI).
	// Limited to between MIN_AUTORENEW_PERIOD (6999999 seconds) and MAX_AUTORENEW_PERIOD (8000001 seconds) by
	// servers-side configuration (which may change).
	// If unspecified, no change.
	AutoRenewPeriod *Duration `protobuf:"bytes,8,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`
	// Optional account to be used at the topic's expirationTime to extend the life of the topic.
	// Once autoRenew functionality is supported by HAPI, the topic lifetime will be extended up to a maximum of the
	// autoRenewPeriod or however long the topic can be extended using all funds on the account (whichever is the
	// smaller duration/amount).
	// If specified as the default value (0.0.0), the autoRenewAccount will be removed.
	// If unspecified, no change.
	AutoRenewAccount *AccountID `protobuf:"bytes,9,opt,name=autoRenewAccount,proto3" json:"autoRenewAccount,omitempty"`
}

func (x *ConsensusUpdateTopicTransactionBody) Reset() {
	*x = ConsensusUpdateTopicTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ConsensusUpdateTopic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusUpdateTopicTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusUpdateTopicTransactionBody) ProtoMessage() {}

func (x *ConsensusUpdateTopicTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ConsensusUpdateTopic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusUpdateTopicTransactionBody.ProtoReflect.Descriptor instead.
func (*ConsensusUpdateTopicTransactionBody) Descriptor() ([]byte, []int) {
	return file_proto_ConsensusUpdateTopic_proto_rawDescGZIP(), []int{0}
}

func (x *ConsensusUpdateTopicTransactionBody) GetTopicID() *TopicID {
	if x != nil {
		return x.TopicID
	}
	return nil
}

func (x *ConsensusUpdateTopicTransactionBody) GetMemo() *StringValue {
	if x != nil {
		return x.Memo
	}
	return nil
}

func (x *ConsensusUpdateTopicTransactionBody) GetExpirationTime() *Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

func (x *ConsensusUpdateTopicTransactionBody) GetAdminKey() *Key {
	if x != nil {
		return x.AdminKey
	}
	return nil
}

func (x *ConsensusUpdateTopicTransactionBody) GetSubmitKey() *Key {
	if x != nil {
		return x.SubmitKey
	}
	return nil
}

func (x *ConsensusUpdateTopicTransactionBody) GetAutoRenewPeriod() *Duration {
	if x != nil {
		return x.AutoRenewPeriod
	}
	return nil
}

func (x *ConsensusUpdateTopicTransactionBody) GetAutoRenewAccount() *AccountID {
	if x != nil {
		return x.AutoRenewAccount
	}
	return nil
}

var File_proto_ConsensusUpdateTopic_proto protoreflect.FileDescriptor

var file_proto_ConsensusUpdateTopic_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x02, 0x0a, 0x23, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73,
	0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x28, 0x0a, 0x07,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x44, 0x52, 0x07, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x38,
	0x0a, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x08, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x4b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79,
	0x12, 0x28, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52,
	0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x0f, 0x61, 0x75,
	0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x3c, 0x0a, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e,
	0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x44, 0x52, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x48, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ConsensusUpdateTopic_proto_rawDescOnce sync.Once
	file_proto_ConsensusUpdateTopic_proto_rawDescData = file_proto_ConsensusUpdateTopic_proto_rawDesc
)

func file_proto_ConsensusUpdateTopic_proto_rawDescGZIP() []byte {
	file_proto_ConsensusUpdateTopic_proto_rawDescOnce.Do(func() {
		file_proto_ConsensusUpdateTopic_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ConsensusUpdateTopic_proto_rawDescData)
	})
	return file_proto_ConsensusUpdateTopic_proto_rawDescData
}

var file_proto_ConsensusUpdateTopic_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_ConsensusUpdateTopic_proto_goTypes = []interface{}{
	(*ConsensusUpdateTopicTransactionBody)(nil), // 0: proto.ConsensusUpdateTopicTransactionBody
	(*TopicID)(nil),     // 1: proto.TopicID
	(*StringValue)(nil), // 2: proto.StringValue
	(*Timestamp)(nil),   // 3: proto.Timestamp
	(*Key)(nil),         // 4: proto.Key
	(*Duration)(nil),    // 5: proto.Duration
	(*AccountID)(nil),   // 6: proto.AccountID
}
var file_proto_ConsensusUpdateTopic_proto_depIdxs = []int32{
	1, // 0: proto.ConsensusUpdateTopicTransactionBody.topicID:type_name -> proto.TopicID
	2, // 1: proto.ConsensusUpdateTopicTransactionBody.memo:type_name -> proto.StringValue
	3, // 2: proto.ConsensusUpdateTopicTransactionBody.expirationTime:type_name -> proto.Timestamp
	4, // 3: proto.ConsensusUpdateTopicTransactionBody.adminKey:type_name -> proto.Key
	4, // 4: proto.ConsensusUpdateTopicTransactionBody.submitKey:type_name -> proto.Key
	5, // 5: proto.ConsensusUpdateTopicTransactionBody.autoRenewPeriod:type_name -> proto.Duration
	6, // 6: proto.ConsensusUpdateTopicTransactionBody.autoRenewAccount:type_name -> proto.AccountID
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_ConsensusUpdateTopic_proto_init() }
func file_proto_ConsensusUpdateTopic_proto_init() {
	if File_proto_ConsensusUpdateTopic_proto != nil {
		return
	}
	file_proto_wrappers_proto_init()
	file_proto_BasicTypes_proto_init()
	file_proto_Duration_proto_init()
	file_proto_Timestamp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_ConsensusUpdateTopic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusUpdateTopicTransactionBody); i {
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
			RawDescriptor: file_proto_ConsensusUpdateTopic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ConsensusUpdateTopic_proto_goTypes,
		DependencyIndexes: file_proto_ConsensusUpdateTopic_proto_depIdxs,
		MessageInfos:      file_proto_ConsensusUpdateTopic_proto_msgTypes,
	}.Build()
	File_proto_ConsensusUpdateTopic_proto = out.File
	file_proto_ConsensusUpdateTopic_proto_rawDesc = nil
	file_proto_ConsensusUpdateTopic_proto_goTypes = nil
	file_proto_ConsensusUpdateTopic_proto_depIdxs = nil
}
