// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/CryptoService.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_proto_CryptoService_proto protoreflect.FileDescriptor

var file_proto_CryptoService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xba, 0x06, 0x0a, 0x0d, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x61,
	0x64, 0x64, 0x4c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0b,
	0x67, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x11, 0x67, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x10, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x0e, 0x67, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x16, 0x67, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x12, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x18, 0x67,
	0x65, 0x74, 0x46, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x54, 0x78, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x54, 0x78, 0x49, 0x44, 0x12, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x15, 0x67, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x46, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x2e, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73,
	0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2d, 0x73, 0x64,
	0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_proto_CryptoService_proto_goTypes = []interface{}{
	(*Transaction)(nil),         // 0: proto.Transaction
	(*Query)(nil),               // 1: proto.Query
	(*TransactionResponse)(nil), // 2: proto.TransactionResponse
	(*Response)(nil),            // 3: proto.Response
}
var file_proto_CryptoService_proto_depIdxs = []int32{
	0,  // 0: proto.CryptoService.createAccount:input_type -> proto.Transaction
	0,  // 1: proto.CryptoService.updateAccount:input_type -> proto.Transaction
	0,  // 2: proto.CryptoService.cryptoTransfer:input_type -> proto.Transaction
	0,  // 3: proto.CryptoService.cryptoDelete:input_type -> proto.Transaction
	0,  // 4: proto.CryptoService.addLiveHash:input_type -> proto.Transaction
	0,  // 5: proto.CryptoService.deleteLiveHash:input_type -> proto.Transaction
	1,  // 6: proto.CryptoService.getLiveHash:input_type -> proto.Query
	1,  // 7: proto.CryptoService.getAccountRecords:input_type -> proto.Query
	1,  // 8: proto.CryptoService.cryptoGetBalance:input_type -> proto.Query
	1,  // 9: proto.CryptoService.getAccountInfo:input_type -> proto.Query
	1,  // 10: proto.CryptoService.getTransactionReceipts:input_type -> proto.Query
	1,  // 11: proto.CryptoService.getFastTransactionRecord:input_type -> proto.Query
	1,  // 12: proto.CryptoService.getTxRecordByTxID:input_type -> proto.Query
	1,  // 13: proto.CryptoService.getStakersByAccountID:input_type -> proto.Query
	2,  // 14: proto.CryptoService.createAccount:output_type -> proto.TransactionResponse
	2,  // 15: proto.CryptoService.updateAccount:output_type -> proto.TransactionResponse
	2,  // 16: proto.CryptoService.cryptoTransfer:output_type -> proto.TransactionResponse
	2,  // 17: proto.CryptoService.cryptoDelete:output_type -> proto.TransactionResponse
	2,  // 18: proto.CryptoService.addLiveHash:output_type -> proto.TransactionResponse
	2,  // 19: proto.CryptoService.deleteLiveHash:output_type -> proto.TransactionResponse
	3,  // 20: proto.CryptoService.getLiveHash:output_type -> proto.Response
	3,  // 21: proto.CryptoService.getAccountRecords:output_type -> proto.Response
	3,  // 22: proto.CryptoService.cryptoGetBalance:output_type -> proto.Response
	3,  // 23: proto.CryptoService.getAccountInfo:output_type -> proto.Response
	3,  // 24: proto.CryptoService.getTransactionReceipts:output_type -> proto.Response
	3,  // 25: proto.CryptoService.getFastTransactionRecord:output_type -> proto.Response
	3,  // 26: proto.CryptoService.getTxRecordByTxID:output_type -> proto.Response
	3,  // 27: proto.CryptoService.getStakersByAccountID:output_type -> proto.Response
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_CryptoService_proto_init() }
func file_proto_CryptoService_proto_init() {
	if File_proto_CryptoService_proto != nil {
		return
	}
	file_proto_Query_proto_init()
	file_proto_Response_proto_init()
	file_proto_TransactionResponse_proto_init()
	file_proto_Transaction_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_CryptoService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_CryptoService_proto_goTypes,
		DependencyIndexes: file_proto_CryptoService_proto_depIdxs,
	}.Build()
	File_proto_CryptoService_proto = out.File
	file_proto_CryptoService_proto_rawDesc = nil
	file_proto_CryptoService_proto_goTypes = nil
	file_proto_CryptoService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CryptoServiceClient is the client API for CryptoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CryptoServiceClient interface {
	// Creates a new account by submitting the transaction
	CreateAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Updates an account by submitting the transaction
	UpdateAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Initiates a transfer by submitting the transaction
	CryptoTransfer(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Deletes and account by submitting the transaction
	CryptoDelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// (NOT CURRENTLY SUPPORTED) Adds a livehash
	AddLiveHash(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// (NOT CURRENTLY SUPPORTED) Deletes a livehash
	DeleteLiveHash(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// (NOT CURRENTLY SUPPORTED) Retrieves a livehash for an account
	GetLiveHash(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Retrieves the 25-hour records stored for an account
	GetAccountRecords(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Retrieves the balance of an account
	CryptoGetBalance(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Retrieves the metadata of an account
	GetAccountInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Retrieves the latest receipt for a transaction that is either awaiting consensus, or reached consensus in the last 180 seconds
	GetTransactionReceipts(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// (NOT CURRENTLY SUPPORTED) Returns the records of transactions recently funded by an account
	GetFastTransactionRecord(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Retrieves the record of a transaction that is either awaiting consensus, or reached consensus in the last 180 seconds
	GetTxRecordByTxID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// (NOT CURRENTLY SUPPORTED) Retrieves the stakers for a node by account id
	GetStakersByAccountID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
}

type cryptoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptoServiceClient(cc grpc.ClientConnInterface) CryptoServiceClient {
	return &cryptoServiceClient{cc}
}

func (c *cryptoServiceClient) CreateAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/createAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) UpdateAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/updateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) CryptoTransfer(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/cryptoTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) CryptoDelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/cryptoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) AddLiveHash(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/addLiveHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) DeleteLiveHash(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/deleteLiveHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetLiveHash(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getLiveHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetAccountRecords(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getAccountRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) CryptoGetBalance(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/cryptoGetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetAccountInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getAccountInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetTransactionReceipts(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getTransactionReceipts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetFastTransactionRecord(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getFastTransactionRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetTxRecordByTxID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getTxRecordByTxID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetStakersByAccountID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getStakersByAccountID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoServiceServer is the server API for CryptoService service.
type CryptoServiceServer interface {
	// Creates a new account by submitting the transaction
	CreateAccount(context.Context, *Transaction) (*TransactionResponse, error)
	// Updates an account by submitting the transaction
	UpdateAccount(context.Context, *Transaction) (*TransactionResponse, error)
	// Initiates a transfer by submitting the transaction
	CryptoTransfer(context.Context, *Transaction) (*TransactionResponse, error)
	// Deletes and account by submitting the transaction
	CryptoDelete(context.Context, *Transaction) (*TransactionResponse, error)
	// (NOT CURRENTLY SUPPORTED) Adds a livehash
	AddLiveHash(context.Context, *Transaction) (*TransactionResponse, error)
	// (NOT CURRENTLY SUPPORTED) Deletes a livehash
	DeleteLiveHash(context.Context, *Transaction) (*TransactionResponse, error)
	// (NOT CURRENTLY SUPPORTED) Retrieves a livehash for an account
	GetLiveHash(context.Context, *Query) (*Response, error)
	// Retrieves the 25-hour records stored for an account
	GetAccountRecords(context.Context, *Query) (*Response, error)
	// Retrieves the balance of an account
	CryptoGetBalance(context.Context, *Query) (*Response, error)
	// Retrieves the metadata of an account
	GetAccountInfo(context.Context, *Query) (*Response, error)
	// Retrieves the latest receipt for a transaction that is either awaiting consensus, or reached consensus in the last 180 seconds
	GetTransactionReceipts(context.Context, *Query) (*Response, error)
	// (NOT CURRENTLY SUPPORTED) Returns the records of transactions recently funded by an account
	GetFastTransactionRecord(context.Context, *Query) (*Response, error)
	// Retrieves the record of a transaction that is either awaiting consensus, or reached consensus in the last 180 seconds
	GetTxRecordByTxID(context.Context, *Query) (*Response, error)
	// (NOT CURRENTLY SUPPORTED) Retrieves the stakers for a node by account id
	GetStakersByAccountID(context.Context, *Query) (*Response, error)
}

// UnimplementedCryptoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCryptoServiceServer struct {
}

func (*UnimplementedCryptoServiceServer) CreateAccount(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (*UnimplementedCryptoServiceServer) UpdateAccount(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (*UnimplementedCryptoServiceServer) CryptoTransfer(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CryptoTransfer not implemented")
}
func (*UnimplementedCryptoServiceServer) CryptoDelete(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CryptoDelete not implemented")
}
func (*UnimplementedCryptoServiceServer) AddLiveHash(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLiveHash not implemented")
}
func (*UnimplementedCryptoServiceServer) DeleteLiveHash(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLiveHash not implemented")
}
func (*UnimplementedCryptoServiceServer) GetLiveHash(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLiveHash not implemented")
}
func (*UnimplementedCryptoServiceServer) GetAccountRecords(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountRecords not implemented")
}
func (*UnimplementedCryptoServiceServer) CryptoGetBalance(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CryptoGetBalance not implemented")
}
func (*UnimplementedCryptoServiceServer) GetAccountInfo(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountInfo not implemented")
}
func (*UnimplementedCryptoServiceServer) GetTransactionReceipts(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionReceipts not implemented")
}
func (*UnimplementedCryptoServiceServer) GetFastTransactionRecord(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFastTransactionRecord not implemented")
}
func (*UnimplementedCryptoServiceServer) GetTxRecordByTxID(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxRecordByTxID not implemented")
}
func (*UnimplementedCryptoServiceServer) GetStakersByAccountID(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStakersByAccountID not implemented")
}

func RegisterCryptoServiceServer(s *grpc.Server, srv CryptoServiceServer) {
	s.RegisterService(&_CryptoService_serviceDesc, srv)
}

func _CryptoService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).CreateAccount(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).UpdateAccount(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_CryptoTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).CryptoTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/CryptoTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).CryptoTransfer(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_CryptoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).CryptoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/CryptoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).CryptoDelete(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_AddLiveHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).AddLiveHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/AddLiveHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).AddLiveHash(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_DeleteLiveHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).DeleteLiveHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/DeleteLiveHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).DeleteLiveHash(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetLiveHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetLiveHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/GetLiveHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetLiveHash(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetAccountRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetAccountRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/GetAccountRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetAccountRecords(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_CryptoGetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).CryptoGetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/CryptoGetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).CryptoGetBalance(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/GetAccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetAccountInfo(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetTransactionReceipts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetTransactionReceipts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/GetTransactionReceipts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetTransactionReceipts(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetFastTransactionRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetFastTransactionRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/GetFastTransactionRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetFastTransactionRecord(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetTxRecordByTxID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetTxRecordByTxID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/GetTxRecordByTxID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetTxRecordByTxID(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetStakersByAccountID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetStakersByAccountID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/GetStakersByAccountID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetStakersByAccountID(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

var _CryptoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CryptoService",
	HandlerType: (*CryptoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createAccount",
			Handler:    _CryptoService_CreateAccount_Handler,
		},
		{
			MethodName: "updateAccount",
			Handler:    _CryptoService_UpdateAccount_Handler,
		},
		{
			MethodName: "cryptoTransfer",
			Handler:    _CryptoService_CryptoTransfer_Handler,
		},
		{
			MethodName: "cryptoDelete",
			Handler:    _CryptoService_CryptoDelete_Handler,
		},
		{
			MethodName: "addLiveHash",
			Handler:    _CryptoService_AddLiveHash_Handler,
		},
		{
			MethodName: "deleteLiveHash",
			Handler:    _CryptoService_DeleteLiveHash_Handler,
		},
		{
			MethodName: "getLiveHash",
			Handler:    _CryptoService_GetLiveHash_Handler,
		},
		{
			MethodName: "getAccountRecords",
			Handler:    _CryptoService_GetAccountRecords_Handler,
		},
		{
			MethodName: "cryptoGetBalance",
			Handler:    _CryptoService_CryptoGetBalance_Handler,
		},
		{
			MethodName: "getAccountInfo",
			Handler:    _CryptoService_GetAccountInfo_Handler,
		},
		{
			MethodName: "getTransactionReceipts",
			Handler:    _CryptoService_GetTransactionReceipts_Handler,
		},
		{
			MethodName: "getFastTransactionRecord",
			Handler:    _CryptoService_GetFastTransactionRecord_Handler,
		},
		{
			MethodName: "getTxRecordByTxID",
			Handler:    _CryptoService_GetTxRecordByTxID_Handler,
		},
		{
			MethodName: "getStakersByAccountID",
			Handler:    _CryptoService_GetStakersByAccountID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/CryptoService.proto",
}
