// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.17.3
// source: moralis/v1/stream.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReceiveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReceiveReply) Reset() {
	*x = ReceiveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moralis_v1_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveReply) ProtoMessage() {}

func (x *ReceiveReply) ProtoReflect() protoreflect.Message {
	mi := &file_moralis_v1_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveReply.ProtoReflect.Descriptor instead.
func (*ReceiveReply) Descriptor() ([]byte, []int) {
	return file_moralis_v1_stream_proto_rawDescGZIP(), []int{0}
}

type ReceiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NftApprovals      *NftApprovals      `protobuf:"bytes,1,opt,name=nftApprovals,proto3" json:"nftApprovals,omitempty"`
	StreamId          string             `protobuf:"bytes,2,opt,name=streamId,proto3" json:"streamId,omitempty"`
	NftTransfers      []*structpb.Struct `protobuf:"bytes,3,rep,name=nftTransfers,proto3" json:"nftTransfers,omitempty"`
	Abi               []*Abi             `protobuf:"bytes,4,rep,name=abi,proto3" json:"abi,omitempty"`
	TxsInternal       []*anypb.Any       `protobuf:"bytes,5,rep,name=txsInternal,proto3" json:"txsInternal,omitempty"`
	Erc20Approvals    []*anypb.Any       `protobuf:"bytes,6,rep,name=erc20Approvals,proto3" json:"erc20Approvals,omitempty"`
	Confirmed         bool               `protobuf:"varint,7,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
	Txs               []*Txs             `protobuf:"bytes,8,rep,name=txs,proto3" json:"txs,omitempty"`
	Retries           int32              `protobuf:"varint,9,opt,name=retries,proto3" json:"retries,omitempty"`
	NftTokenApprovals []*anypb.Any       `protobuf:"bytes,10,rep,name=nftTokenApprovals,proto3" json:"nftTokenApprovals,omitempty"`
	ChainId           string             `protobuf:"bytes,11,opt,name=chainId,proto3" json:"chainId,omitempty"`
	NativeBalances    []*anypb.Any       `protobuf:"bytes,12,rep,name=nativeBalances,proto3" json:"nativeBalances,omitempty"`
	Erc20Transfers    []*anypb.Any       `protobuf:"bytes,13,rep,name=erc20Transfers,proto3" json:"erc20Transfers,omitempty"`
	Block             *Block             `protobuf:"bytes,14,opt,name=block,proto3" json:"block,omitempty"`
	Tag               string             `protobuf:"bytes,15,opt,name=tag,proto3" json:"tag,omitempty"`
	Logs              []*Logs            `protobuf:"bytes,16,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *ReceiveRequest) Reset() {
	*x = ReceiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moralis_v1_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveRequest) ProtoMessage() {}

func (x *ReceiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moralis_v1_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveRequest.ProtoReflect.Descriptor instead.
func (*ReceiveRequest) Descriptor() ([]byte, []int) {
	return file_moralis_v1_stream_proto_rawDescGZIP(), []int{1}
}

func (x *ReceiveRequest) GetNftApprovals() *NftApprovals {
	if x != nil {
		return x.NftApprovals
	}
	return nil
}

func (x *ReceiveRequest) GetStreamId() string {
	if x != nil {
		return x.StreamId
	}
	return ""
}

func (x *ReceiveRequest) GetNftTransfers() []*structpb.Struct {
	if x != nil {
		return x.NftTransfers
	}
	return nil
}

func (x *ReceiveRequest) GetAbi() []*Abi {
	if x != nil {
		return x.Abi
	}
	return nil
}

func (x *ReceiveRequest) GetTxsInternal() []*anypb.Any {
	if x != nil {
		return x.TxsInternal
	}
	return nil
}

func (x *ReceiveRequest) GetErc20Approvals() []*anypb.Any {
	if x != nil {
		return x.Erc20Approvals
	}
	return nil
}

func (x *ReceiveRequest) GetConfirmed() bool {
	if x != nil {
		return x.Confirmed
	}
	return false
}

func (x *ReceiveRequest) GetTxs() []*Txs {
	if x != nil {
		return x.Txs
	}
	return nil
}

func (x *ReceiveRequest) GetRetries() int32 {
	if x != nil {
		return x.Retries
	}
	return 0
}

func (x *ReceiveRequest) GetNftTokenApprovals() []*anypb.Any {
	if x != nil {
		return x.NftTokenApprovals
	}
	return nil
}

func (x *ReceiveRequest) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *ReceiveRequest) GetNativeBalances() []*anypb.Any {
	if x != nil {
		return x.NativeBalances
	}
	return nil
}

func (x *ReceiveRequest) GetErc20Transfers() []*anypb.Any {
	if x != nil {
		return x.Erc20Transfers
	}
	return nil
}

func (x *ReceiveRequest) GetBlock() *Block {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *ReceiveRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *ReceiveRequest) GetLogs() []*Logs {
	if x != nil {
		return x.Logs
	}
	return nil
}

type NftApprovals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ERC721  []*anypb.Any `protobuf:"bytes,1,rep,name=ERC721,proto3" json:"ERC721,omitempty"`
	ERC1155 []*anypb.Any `protobuf:"bytes,2,rep,name=ERC1155,proto3" json:"ERC1155,omitempty"`
}

func (x *NftApprovals) Reset() {
	*x = NftApprovals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moralis_v1_stream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftApprovals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftApprovals) ProtoMessage() {}

func (x *NftApprovals) ProtoReflect() protoreflect.Message {
	mi := &file_moralis_v1_stream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftApprovals.ProtoReflect.Descriptor instead.
func (*NftApprovals) Descriptor() ([]byte, []int) {
	return file_moralis_v1_stream_proto_rawDescGZIP(), []int{2}
}

func (x *NftApprovals) GetERC721() []*anypb.Any {
	if x != nil {
		return x.ERC721
	}
	return nil
}

func (x *NftApprovals) GetERC1155() []*anypb.Any {
	if x != nil {
		return x.ERC1155
	}
	return nil
}

type Abi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inputs    []*Inputs `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Name      string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Anonymous bool      `protobuf:"varint,3,opt,name=anonymous,proto3" json:"anonymous,omitempty"`
	Type      string    `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Abi) Reset() {
	*x = Abi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moralis_v1_stream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Abi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Abi) ProtoMessage() {}

func (x *Abi) ProtoReflect() protoreflect.Message {
	mi := &file_moralis_v1_stream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Abi.ProtoReflect.Descriptor instead.
func (*Abi) Descriptor() ([]byte, []int) {
	return file_moralis_v1_stream_proto_rawDescGZIP(), []int{3}
}

func (x *Abi) GetInputs() []*Inputs {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *Abi) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Abi) GetAnonymous() bool {
	if x != nil {
		return x.Anonymous
	}
	return false
}

func (x *Abi) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Inputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Indexed      bool   `protobuf:"varint,1,opt,name=indexed,proto3" json:"indexed,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	InternalType string `protobuf:"bytes,3,opt,name=internalType,proto3" json:"internalType,omitempty"`
	Type         string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Inputs) Reset() {
	*x = Inputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moralis_v1_stream_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inputs) ProtoMessage() {}

func (x *Inputs) ProtoReflect() protoreflect.Message {
	mi := &file_moralis_v1_stream_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inputs.ProtoReflect.Descriptor instead.
func (*Inputs) Descriptor() ([]byte, []int) {
	return file_moralis_v1_stream_proto_rawDescGZIP(), []int{4}
}

func (x *Inputs) GetIndexed() bool {
	if x != nil {
		return x.Indexed
	}
	return false
}

func (x *Inputs) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Inputs) GetInternalType() string {
	if x != nil {
		return x.InternalType
	}
	return ""
}

func (x *Inputs) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Txs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReceiptGasUsed           string `protobuf:"bytes,1,opt,name=receiptGasUsed,proto3" json:"receiptGasUsed,omitempty"`
	TransactionIndex         string `protobuf:"bytes,2,opt,name=transactionIndex,proto3" json:"transactionIndex,omitempty"`
	Type                     string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Nonce                    string `protobuf:"bytes,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ToAddress                string `protobuf:"bytes,5,opt,name=toAddress,proto3" json:"toAddress,omitempty"`
	Input                    string `protobuf:"bytes,6,opt,name=input,proto3" json:"input,omitempty"`
	R                        string `protobuf:"bytes,7,opt,name=r,proto3" json:"r,omitempty"`
	S                        string `protobuf:"bytes,8,opt,name=s,proto3" json:"s,omitempty"`
	ReceiptCumulativeGasUsed string `protobuf:"bytes,9,opt,name=receiptCumulativeGasUsed,proto3" json:"receiptCumulativeGasUsed,omitempty"`
	V                        string `protobuf:"bytes,10,opt,name=v,proto3" json:"v,omitempty"`
	Gas                      string `protobuf:"bytes,11,opt,name=gas,proto3" json:"gas,omitempty"`
	FromAddress              string `protobuf:"bytes,12,opt,name=fromAddress,proto3" json:"fromAddress,omitempty"`
	Value                    string `protobuf:"bytes,13,opt,name=value,proto3" json:"value,omitempty"`
	Hash                     string `protobuf:"bytes,14,opt,name=hash,proto3" json:"hash,omitempty"`
	ReceiptStatus            string `protobuf:"bytes,15,opt,name=receiptStatus,proto3" json:"receiptStatus,omitempty"`
	GasPrice                 string `protobuf:"bytes,16,opt,name=gasPrice,proto3" json:"gasPrice,omitempty"`
}

func (x *Txs) Reset() {
	*x = Txs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moralis_v1_stream_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Txs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Txs) ProtoMessage() {}

func (x *Txs) ProtoReflect() protoreflect.Message {
	mi := &file_moralis_v1_stream_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Txs.ProtoReflect.Descriptor instead.
func (*Txs) Descriptor() ([]byte, []int) {
	return file_moralis_v1_stream_proto_rawDescGZIP(), []int{5}
}

func (x *Txs) GetReceiptGasUsed() string {
	if x != nil {
		return x.ReceiptGasUsed
	}
	return ""
}

func (x *Txs) GetTransactionIndex() string {
	if x != nil {
		return x.TransactionIndex
	}
	return ""
}

func (x *Txs) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Txs) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *Txs) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *Txs) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *Txs) GetR() string {
	if x != nil {
		return x.R
	}
	return ""
}

func (x *Txs) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

func (x *Txs) GetReceiptCumulativeGasUsed() string {
	if x != nil {
		return x.ReceiptCumulativeGasUsed
	}
	return ""
}

func (x *Txs) GetV() string {
	if x != nil {
		return x.V
	}
	return ""
}

func (x *Txs) GetGas() string {
	if x != nil {
		return x.Gas
	}
	return ""
}

func (x *Txs) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *Txs) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Txs) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Txs) GetReceiptStatus() string {
	if x != nil {
		return x.ReceiptStatus
	}
	return ""
}

func (x *Txs) GetGasPrice() string {
	if x != nil {
		return x.GasPrice
	}
	return ""
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number    string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Hash      string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Timestamp string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moralis_v1_stream_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_moralis_v1_stream_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_moralis_v1_stream_proto_rawDescGZIP(), []int{6}
}

func (x *Block) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Block) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Block) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type Logs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic1          string `protobuf:"bytes,1,opt,name=topic1,proto3" json:"topic1,omitempty"`
	Topic2          string `protobuf:"bytes,2,opt,name=topic2,proto3" json:"topic2,omitempty"`
	LogIndex        string `protobuf:"bytes,3,opt,name=logIndex,proto3" json:"logIndex,omitempty"`
	Address         string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Topic0          string `protobuf:"bytes,5,opt,name=topic0,proto3" json:"topic0,omitempty"`
	Data            string `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	Topic3          string `protobuf:"bytes,7,opt,name=topic3,proto3" json:"topic3,omitempty"`
	TransactionHash string `protobuf:"bytes,8,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
}

func (x *Logs) Reset() {
	*x = Logs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moralis_v1_stream_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logs) ProtoMessage() {}

func (x *Logs) ProtoReflect() protoreflect.Message {
	mi := &file_moralis_v1_stream_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logs.ProtoReflect.Descriptor instead.
func (*Logs) Descriptor() ([]byte, []int) {
	return file_moralis_v1_stream_proto_rawDescGZIP(), []int{7}
}

func (x *Logs) GetTopic1() string {
	if x != nil {
		return x.Topic1
	}
	return ""
}

func (x *Logs) GetTopic2() string {
	if x != nil {
		return x.Topic2
	}
	return ""
}

func (x *Logs) GetLogIndex() string {
	if x != nil {
		return x.LogIndex
	}
	return ""
}

func (x *Logs) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Logs) GetTopic0() string {
	if x != nil {
		return x.Topic0
	}
	return ""
}

func (x *Logs) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Logs) GetTopic3() string {
	if x != nil {
		return x.Topic3
	}
	return ""
}

func (x *Logs) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

var File_moralis_v1_stream_proto protoreflect.FileDescriptor

var file_moralis_v1_stream_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x6f, 0x72, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x0e, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0xea, 0x05, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0c, 0x6e, 0x66, 0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x66, 0x74, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x73, 0x52, 0x0c, 0x6e, 0x66, 0x74, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x12, 0x3b, 0x0a, 0x0c, 0x6e, 0x66, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x0c, 0x6e, 0x66, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x12, 0x25,
	0x0a, 0x03, 0x61, 0x62, 0x69, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x62, 0x69,
	0x52, 0x03, 0x61, 0x62, 0x69, 0x12, 0x36, 0x0a, 0x0b, 0x74, 0x78, 0x73, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x0b, 0x74, 0x78, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x3c, 0x0a,
	0x0e, 0x65, 0x72, 0x63, 0x32, 0x30, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0e, 0x65, 0x72, 0x63,
	0x32, 0x30, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x03, 0x74, 0x78, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x72,
	0x61, 0x6c, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x73, 0x52, 0x03, 0x74, 0x78, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x11, 0x6e, 0x66,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x11, 0x6e, 0x66, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0e, 0x6e, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0e, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x0e, 0x65, 0x72, 0x63, 0x32, 0x30, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0e, 0x65, 0x72, 0x63, 0x32, 0x30, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x72, 0x61, 0x6c, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x74, 0x61, 0x67, 0x12, 0x28, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x6c, 0x0a,
	0x0c, 0x4e, 0x66, 0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x73, 0x12, 0x2c, 0x0a,
	0x06, 0x45, 0x52, 0x43, 0x37, 0x32, 0x31, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x06, 0x45, 0x52, 0x43, 0x37, 0x32, 0x31, 0x12, 0x2e, 0x0a, 0x07, 0x45,
	0x52, 0x43, 0x31, 0x31, 0x35, 0x35, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x07, 0x45, 0x52, 0x43, 0x31, 0x31, 0x35, 0x35, 0x22, 0x7b, 0x0a, 0x03, 0x41,
	0x62, 0x69, 0x12, 0x2e, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d,
	0x6f, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x6e, 0x6f, 0x6e, 0x79,
	0x6d, 0x6f, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x6e, 0x0a, 0x06, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xbd, 0x03, 0x0a, 0x03, 0x54, 0x78, 0x73,
	0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x47, 0x61, 0x73, 0x55, 0x73,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x47, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x72,
	0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x73, 0x12, 0x3a,
	0x0a, 0x18, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x43, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x47, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x18, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x43, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x47, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x76, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x61, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x61, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x72,
	0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x51, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xda, 0x01, 0x0a, 0x04,
	0x4c, 0x6f, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x31, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x31, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x32, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x30, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x30, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x33,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x33, 0x12, 0x28,
	0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x32, 0x6d, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x63, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x12, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x42, 0x38, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x6f, 0x72, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x24, 0x6e, 0x66, 0x74,
	0x2d, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moralis_v1_stream_proto_rawDescOnce sync.Once
	file_moralis_v1_stream_proto_rawDescData = file_moralis_v1_stream_proto_rawDesc
)

func file_moralis_v1_stream_proto_rawDescGZIP() []byte {
	file_moralis_v1_stream_proto_rawDescOnce.Do(func() {
		file_moralis_v1_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_moralis_v1_stream_proto_rawDescData)
	})
	return file_moralis_v1_stream_proto_rawDescData
}

var file_moralis_v1_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_moralis_v1_stream_proto_goTypes = []interface{}{
	(*ReceiveReply)(nil),    // 0: api.moralis.v1.ReceiveReply
	(*ReceiveRequest)(nil),  // 1: api.moralis.v1.ReceiveRequest
	(*NftApprovals)(nil),    // 2: api.moralis.v1.NftApprovals
	(*Abi)(nil),             // 3: api.moralis.v1.Abi
	(*Inputs)(nil),          // 4: api.moralis.v1.Inputs
	(*Txs)(nil),             // 5: api.moralis.v1.Txs
	(*Block)(nil),           // 6: api.moralis.v1.Block
	(*Logs)(nil),            // 7: api.moralis.v1.Logs
	(*structpb.Struct)(nil), // 8: google.protobuf.Struct
	(*anypb.Any)(nil),       // 9: google.protobuf.Any
}
var file_moralis_v1_stream_proto_depIdxs = []int32{
	2,  // 0: api.moralis.v1.ReceiveRequest.nftApprovals:type_name -> api.moralis.v1.NftApprovals
	8,  // 1: api.moralis.v1.ReceiveRequest.nftTransfers:type_name -> google.protobuf.Struct
	3,  // 2: api.moralis.v1.ReceiveRequest.abi:type_name -> api.moralis.v1.Abi
	9,  // 3: api.moralis.v1.ReceiveRequest.txsInternal:type_name -> google.protobuf.Any
	9,  // 4: api.moralis.v1.ReceiveRequest.erc20Approvals:type_name -> google.protobuf.Any
	5,  // 5: api.moralis.v1.ReceiveRequest.txs:type_name -> api.moralis.v1.Txs
	9,  // 6: api.moralis.v1.ReceiveRequest.nftTokenApprovals:type_name -> google.protobuf.Any
	9,  // 7: api.moralis.v1.ReceiveRequest.nativeBalances:type_name -> google.protobuf.Any
	9,  // 8: api.moralis.v1.ReceiveRequest.erc20Transfers:type_name -> google.protobuf.Any
	6,  // 9: api.moralis.v1.ReceiveRequest.block:type_name -> api.moralis.v1.Block
	7,  // 10: api.moralis.v1.ReceiveRequest.logs:type_name -> api.moralis.v1.Logs
	9,  // 11: api.moralis.v1.NftApprovals.ERC721:type_name -> google.protobuf.Any
	9,  // 12: api.moralis.v1.NftApprovals.ERC1155:type_name -> google.protobuf.Any
	4,  // 13: api.moralis.v1.Abi.inputs:type_name -> api.moralis.v1.Inputs
	1,  // 14: api.moralis.v1.Stream.Receive:input_type -> api.moralis.v1.ReceiveRequest
	0,  // 15: api.moralis.v1.Stream.Receive:output_type -> api.moralis.v1.ReceiveReply
	15, // [15:16] is the sub-list for method output_type
	14, // [14:15] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_moralis_v1_stream_proto_init() }
func file_moralis_v1_stream_proto_init() {
	if File_moralis_v1_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_moralis_v1_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveReply); i {
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
		file_moralis_v1_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveRequest); i {
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
		file_moralis_v1_stream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftApprovals); i {
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
		file_moralis_v1_stream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Abi); i {
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
		file_moralis_v1_stream_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inputs); i {
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
		file_moralis_v1_stream_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Txs); i {
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
		file_moralis_v1_stream_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_moralis_v1_stream_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logs); i {
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
			RawDescriptor: file_moralis_v1_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_moralis_v1_stream_proto_goTypes,
		DependencyIndexes: file_moralis_v1_stream_proto_depIdxs,
		MessageInfos:      file_moralis_v1_stream_proto_msgTypes,
	}.Build()
	File_moralis_v1_stream_proto = out.File
	file_moralis_v1_stream_proto_rawDesc = nil
	file_moralis_v1_stream_proto_goTypes = nil
	file_moralis_v1_stream_proto_depIdxs = nil
}
