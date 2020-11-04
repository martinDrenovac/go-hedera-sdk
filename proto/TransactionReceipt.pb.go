// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/TransactionReceipt.proto

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

//
//The summary of a transaction's result so far. If the transaction has not reached consensus, this result will be necessarily incomplete.
type TransactionReceipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The consensus status of the transaction; is UNKNOWN if consensus has not been reached, or if the
	// associated transaction did not have a valid payer signature
	Status ResponseCodeEnum `protobuf:"varint,1,opt,name=status,proto3,enum=proto.ResponseCodeEnum" json:"status,omitempty"`
	// In the receipt of a CryptoCreate, the id of the newly created account
	AccountID *AccountID `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	// In the receipt of a FileCreate, the id of the newly created file
	FileID *FileID `protobuf:"bytes,3,opt,name=fileID,proto3" json:"fileID,omitempty"`
	// In the receipt of a ContractCreate, the id of the newly created contract
	ContractID *ContractID `protobuf:"bytes,4,opt,name=contractID,proto3" json:"contractID,omitempty"`
	// The exchange rates in effect when the transaction reached consensus
	ExchangeRate *ExchangeRateSet `protobuf:"bytes,5,opt,name=exchangeRate,proto3" json:"exchangeRate,omitempty"`
	// In the receipt of a ConsensusCreateTopic, the id of the newly created topic.
	TopicID *TopicID `protobuf:"bytes,6,opt,name=topicID,proto3" json:"topicID,omitempty"`
	// In the receipt of a ConsensusSubmitMessage, the new sequence number of the topic that received the message
	TopicSequenceNumber uint64 `protobuf:"varint,7,opt,name=topicSequenceNumber,proto3" json:"topicSequenceNumber,omitempty"`
	// In the receipt of a ConsensusSubmitMessage, the new running hash of the topic that received the message.
	// This 48-byte field is the output of a particular SHA-384 digest whose input data are determined by the
	// value of the topicRunningHashVersion below. The bytes of each uint64 or uint32 are to be in Big-Endian
	// format.
	//
	// IF the topicRunningHashVersion is '0' or '1', then the input data to the SHA-384 digest are, in order:
	// ---
	// 1. The previous running hash of the topic (48 bytes)
	// 2. The topic's shard (8 bytes)
	// 3. The topic's realm (8 bytes)
	// 4. The topic's number (8 bytes)
	// 5. The number of seconds since the epoch before the ConsensusSubmitMessage reached consensus (8 bytes)
	// 6. The number of nanoseconds since 5. before the ConsensusSubmitMessage reached consensus (4 bytes)
	// 7. The topicSequenceNumber from above (8 bytes)
	// 8. The message bytes from the ConsensusSubmitMessage (variable).
	//
	// IF the topicRunningHashVersion is '2', then the input data to the SHA-384 digest are, in order:
	// ---
	// 1. The previous running hash of the topic (48 bytes)
	// 2. The topicRunningHashVersion below (8 bytes)
	// 3. The topic's shard (8 bytes)
	// 4. The topic's realm (8 bytes)
	// 5. The topic's number (8 bytes)
	// 6. The number of seconds since the epoch before the ConsensusSubmitMessage reached consensus (8 bytes)
	// 7. The number of nanoseconds since 6. before the ConsensusSubmitMessage reached consensus (4 bytes)
	// 8. The topicSequenceNumber from above (8 bytes)
	// 9. The output of the SHA-384 digest of the message bytes from the consensusSubmitMessage (48 bytes)
	//
	// Otherwise, IF the topicRunningHashVersion is '3', then the input data to the SHA-384 digest are, in order:
	// ---
	// 1. The previous running hash of the topic (48 bytes)
	// 2. The topicRunningHashVersion below (8 bytes)
	// 3. The payer account's shard (8 bytes)
	// 4. The payer account's realm (8 bytes)
	// 5. The payer account's number (8 bytes)
	// 6. The topic's shard (8 bytes)
	// 7. The topic's realm (8 bytes)
	// 8. The topic's number (8 bytes)
	// 9. The number of seconds since the epoch before the ConsensusSubmitMessage reached consensus (8 bytes)
	// 10. The number of nanoseconds since 9. before the ConsensusSubmitMessage reached consensus (4 bytes)
	// 11. The topicSequenceNumber from above (8 bytes)
	// 12. The output of the SHA-384 digest of the message bytes from the consensusSubmitMessage (48 bytes)
	TopicRunningHash []byte `protobuf:"bytes,8,opt,name=topicRunningHash,proto3" json:"topicRunningHash,omitempty"`
	// In the receipt of a ConsensusSubmitMessage, the version of the SHA-384 digest used to update the running hash.
	TopicRunningHashVersion uint64 `protobuf:"varint,9,opt,name=topicRunningHashVersion,proto3" json:"topicRunningHashVersion,omitempty"`
	// In the receipt of a CreateToken, the id of the newly created token
	TokenId *TokenID `protobuf:"bytes,10,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
}

func (x *TransactionReceipt) Reset() {
	*x = TransactionReceipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_TransactionReceipt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionReceipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionReceipt) ProtoMessage() {}

func (x *TransactionReceipt) ProtoReflect() protoreflect.Message {
	mi := &file_proto_TransactionReceipt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionReceipt.ProtoReflect.Descriptor instead.
func (*TransactionReceipt) Descriptor() ([]byte, []int) {
	return file_proto_TransactionReceipt_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionReceipt) GetStatus() ResponseCodeEnum {
	if x != nil {
		return x.Status
	}
	return ResponseCodeEnum_OK
}

func (x *TransactionReceipt) GetAccountID() *AccountID {
	if x != nil {
		return x.AccountID
	}
	return nil
}

func (x *TransactionReceipt) GetFileID() *FileID {
	if x != nil {
		return x.FileID
	}
	return nil
}

func (x *TransactionReceipt) GetContractID() *ContractID {
	if x != nil {
		return x.ContractID
	}
	return nil
}

func (x *TransactionReceipt) GetExchangeRate() *ExchangeRateSet {
	if x != nil {
		return x.ExchangeRate
	}
	return nil
}

func (x *TransactionReceipt) GetTopicID() *TopicID {
	if x != nil {
		return x.TopicID
	}
	return nil
}

func (x *TransactionReceipt) GetTopicSequenceNumber() uint64 {
	if x != nil {
		return x.TopicSequenceNumber
	}
	return 0
}

func (x *TransactionReceipt) GetTopicRunningHash() []byte {
	if x != nil {
		return x.TopicRunningHash
	}
	return nil
}

func (x *TransactionReceipt) GetTopicRunningHashVersion() uint64 {
	if x != nil {
		return x.TopicRunningHashVersion
	}
	return 0
}

func (x *TransactionReceipt) GetTokenId() *TokenID {
	if x != nil {
		return x.TokenId
	}
	return nil
}

var File_proto_TransactionReceipt_proto protoreflect.FileDescriptor

var file_proto_TransactionReceipt_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x03, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x25, 0x0a, 0x06, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x44, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x49, 0x44, 0x12, 0x3a, 0x0a, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x74, 0x52, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x28, 0x0a, 0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x49, 0x44, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x13, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x53,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a,
	0x10, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x75,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73, 0x68, 0x12, 0x38, 0x0a, 0x17, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73, 0x68, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x17, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73, 0x68, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x49, 0x44, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x42, 0x48, 0x0a,
	0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2f, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_TransactionReceipt_proto_rawDescOnce sync.Once
	file_proto_TransactionReceipt_proto_rawDescData = file_proto_TransactionReceipt_proto_rawDesc
)

func file_proto_TransactionReceipt_proto_rawDescGZIP() []byte {
	file_proto_TransactionReceipt_proto_rawDescOnce.Do(func() {
		file_proto_TransactionReceipt_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_TransactionReceipt_proto_rawDescData)
	})
	return file_proto_TransactionReceipt_proto_rawDescData
}

var file_proto_TransactionReceipt_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_TransactionReceipt_proto_goTypes = []interface{}{
	(*TransactionReceipt)(nil), // 0: proto.TransactionReceipt
	(ResponseCodeEnum)(0),      // 1: proto.ResponseCodeEnum
	(*AccountID)(nil),          // 2: proto.AccountID
	(*FileID)(nil),             // 3: proto.FileID
	(*ContractID)(nil),         // 4: proto.ContractID
	(*ExchangeRateSet)(nil),    // 5: proto.ExchangeRateSet
	(*TopicID)(nil),            // 6: proto.TopicID
	(*TokenID)(nil),            // 7: proto.TokenID
}
var file_proto_TransactionReceipt_proto_depIdxs = []int32{
	1, // 0: proto.TransactionReceipt.status:type_name -> proto.ResponseCodeEnum
	2, // 1: proto.TransactionReceipt.accountID:type_name -> proto.AccountID
	3, // 2: proto.TransactionReceipt.fileID:type_name -> proto.FileID
	4, // 3: proto.TransactionReceipt.contractID:type_name -> proto.ContractID
	5, // 4: proto.TransactionReceipt.exchangeRate:type_name -> proto.ExchangeRateSet
	6, // 5: proto.TransactionReceipt.topicID:type_name -> proto.TopicID
	7, // 6: proto.TransactionReceipt.tokenId:type_name -> proto.TokenID
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_TransactionReceipt_proto_init() }
func file_proto_TransactionReceipt_proto_init() {
	if File_proto_TransactionReceipt_proto != nil {
		return
	}
	file_proto_BasicTypes_proto_init()
	file_proto_ResponseCode_proto_init()
	file_proto_ExchangeRate_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_TransactionReceipt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionReceipt); i {
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
			RawDescriptor: file_proto_TransactionReceipt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_TransactionReceipt_proto_goTypes,
		DependencyIndexes: file_proto_TransactionReceipt_proto_depIdxs,
		MessageInfos:      file_proto_TransactionReceipt_proto_msgTypes,
	}.Build()
	File_proto_TransactionReceipt_proto = out.File
	file_proto_TransactionReceipt_proto_rawDesc = nil
	file_proto_TransactionReceipt_proto_goTypes = nil
	file_proto_TransactionReceipt_proto_depIdxs = nil
}
