// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.17.3
// source: nftmarket/v1/nftmarket.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type GetNFTMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NftAddress string `protobuf:"bytes,1,opt,name=nftAddress,proto3" json:"nftAddress,omitempty"`
	TokenId    string `protobuf:"bytes,2,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	ChainId    string `protobuf:"bytes,3,opt,name=chainId,proto3" json:"chainId,omitempty"`
}

func (x *GetNFTMetadataRequest) Reset() {
	*x = GetNFTMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nftmarket_v1_nftmarket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNFTMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNFTMetadataRequest) ProtoMessage() {}

func (x *GetNFTMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nftmarket_v1_nftmarket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNFTMetadataRequest.ProtoReflect.Descriptor instead.
func (*GetNFTMetadataRequest) Descriptor() ([]byte, []int) {
	return file_nftmarket_v1_nftmarket_proto_rawDescGZIP(), []int{0}
}

func (x *GetNFTMetadataRequest) GetNftAddress() string {
	if x != nil {
		return x.NftAddress
	}
	return ""
}

func (x *GetNFTMetadataRequest) GetTokenId() string {
	if x != nil {
		return x.TokenId
	}
	return ""
}

func (x *GetNFTMetadataRequest) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

type GetNFTMetadataReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data *structpb.Struct `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Msg  string           `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GetNFTMetadataReply) Reset() {
	*x = GetNFTMetadataReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nftmarket_v1_nftmarket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNFTMetadataReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNFTMetadataReply) ProtoMessage() {}

func (x *GetNFTMetadataReply) ProtoReflect() protoreflect.Message {
	mi := &file_nftmarket_v1_nftmarket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNFTMetadataReply.ProtoReflect.Descriptor instead.
func (*GetNFTMetadataReply) Descriptor() ([]byte, []int) {
	return file_nftmarket_v1_nftmarket_proto_rawDescGZIP(), []int{1}
}

func (x *GetNFTMetadataReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetNFTMetadataReply) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetNFTMetadataReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ListedPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page            int64  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize        int64  `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	MarketPlaceAddr string `protobuf:"bytes,3,opt,name=marketPlaceAddr,proto3" json:"marketPlaceAddr,omitempty"`
}

func (x *ListedPageRequest) Reset() {
	*x = ListedPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nftmarket_v1_nftmarket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListedPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListedPageRequest) ProtoMessage() {}

func (x *ListedPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nftmarket_v1_nftmarket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListedPageRequest.ProtoReflect.Descriptor instead.
func (*ListedPageRequest) Descriptor() ([]byte, []int) {
	return file_nftmarket_v1_nftmarket_proto_rawDescGZIP(), []int{2}
}

func (x *ListedPageRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListedPageRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListedPageRequest) GetMarketPlaceAddr() string {
	if x != nil {
		return x.MarketPlaceAddr
	}
	return ""
}

type ListedPageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32                      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data []*ListedPageReply_NftItem `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	Msg  string                     `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ListedPageReply) Reset() {
	*x = ListedPageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nftmarket_v1_nftmarket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListedPageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListedPageReply) ProtoMessage() {}

func (x *ListedPageReply) ProtoReflect() protoreflect.Message {
	mi := &file_nftmarket_v1_nftmarket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListedPageReply.ProtoReflect.Descriptor instead.
func (*ListedPageReply) Descriptor() ([]byte, []int) {
	return file_nftmarket_v1_nftmarket_proto_rawDescGZIP(), []int{3}
}

func (x *ListedPageReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListedPageReply) GetData() []*ListedPageReply_NftItem {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListedPageReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SelfPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       int64  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize   int64  `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	NftAddress string `protobuf:"bytes,3,opt,name=nftAddress,proto3" json:"nftAddress,omitempty"`
	Owner      string `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	ChainId    string `protobuf:"bytes,5,opt,name=chainId,proto3" json:"chainId,omitempty"`
}

func (x *SelfPageRequest) Reset() {
	*x = SelfPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nftmarket_v1_nftmarket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelfPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelfPageRequest) ProtoMessage() {}

func (x *SelfPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nftmarket_v1_nftmarket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelfPageRequest.ProtoReflect.Descriptor instead.
func (*SelfPageRequest) Descriptor() ([]byte, []int) {
	return file_nftmarket_v1_nftmarket_proto_rawDescGZIP(), []int{4}
}

func (x *SelfPageRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SelfPageRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SelfPageRequest) GetNftAddress() string {
	if x != nil {
		return x.NftAddress
	}
	return ""
}

func (x *SelfPageRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *SelfPageRequest) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

type SelfPageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32              `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data []*structpb.Struct `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	Msg  string             `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SelfPageReply) Reset() {
	*x = SelfPageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nftmarket_v1_nftmarket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelfPageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelfPageReply) ProtoMessage() {}

func (x *SelfPageReply) ProtoReflect() protoreflect.Message {
	mi := &file_nftmarket_v1_nftmarket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelfPageReply.ProtoReflect.Descriptor instead.
func (*SelfPageReply) Descriptor() ([]byte, []int) {
	return file_nftmarket_v1_nftmarket_proto_rawDescGZIP(), []int{5}
}

func (x *SelfPageReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SelfPageReply) GetData() []*structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SelfPageReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ListedPageReply_NftItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ChainId            string `protobuf:"bytes,2,opt,name=chainId,proto3" json:"chainId,omitempty"`
	MarketPlaceAddress string `protobuf:"bytes,3,opt,name=marketPlaceAddress,proto3" json:"marketPlaceAddress,omitempty"`
	Seller             string `protobuf:"bytes,4,opt,name=seller,proto3" json:"seller,omitempty"`
	NftAddress         string `protobuf:"bytes,5,opt,name=nftAddress,proto3" json:"nftAddress,omitempty"`
	TokenId            string `protobuf:"bytes,6,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	Price              string `protobuf:"bytes,7,opt,name=price,proto3" json:"price,omitempty"`
	ListedTime         int64  `protobuf:"varint,8,opt,name=listedTime,proto3" json:"listedTime,omitempty"`
	UpdateTime         int64  `protobuf:"varint,9,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *ListedPageReply_NftItem) Reset() {
	*x = ListedPageReply_NftItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nftmarket_v1_nftmarket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListedPageReply_NftItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListedPageReply_NftItem) ProtoMessage() {}

func (x *ListedPageReply_NftItem) ProtoReflect() protoreflect.Message {
	mi := &file_nftmarket_v1_nftmarket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListedPageReply_NftItem.ProtoReflect.Descriptor instead.
func (*ListedPageReply_NftItem) Descriptor() ([]byte, []int) {
	return file_nftmarket_v1_nftmarket_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ListedPageReply_NftItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListedPageReply_NftItem) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *ListedPageReply_NftItem) GetMarketPlaceAddress() string {
	if x != nil {
		return x.MarketPlaceAddress
	}
	return ""
}

func (x *ListedPageReply_NftItem) GetSeller() string {
	if x != nil {
		return x.Seller
	}
	return ""
}

func (x *ListedPageReply_NftItem) GetNftAddress() string {
	if x != nil {
		return x.NftAddress
	}
	return ""
}

func (x *ListedPageReply_NftItem) GetTokenId() string {
	if x != nil {
		return x.TokenId
	}
	return ""
}

func (x *ListedPageReply_NftItem) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *ListedPageReply_NftItem) GetListedTime() int64 {
	if x != nil {
		return x.ListedTime
	}
	return 0
}

func (x *ListedPageReply_NftItem) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

var File_nftmarket_v1_nftmarket_proto protoreflect.FileDescriptor

var file_nftmarket_v1_nftmarket_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6e, 0x66, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6e,
	0x66, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4e, 0x46, 0x54,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x27, 0x0a, 0x0a, 0x6e, 0x66, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x6e, 0x66,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x07, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x68,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x4e, 0x46, 0x54, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x7f, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x64, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x22, 0x02, 0x20, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x41, 0x64,
	0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x22, 0x84, 0x03, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x64, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x3d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x2e, 0x4e, 0x66, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x1a, 0x8b, 0x02, 0x0a, 0x07, 0x4e, 0x66, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72,
	0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x66, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x66, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0xb5, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x6c, 0x66, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x23, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x66, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x66, 0x74, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x62, 0x0a, 0x0d, 0x53, 0x65, 0x6c, 0x66,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2b, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xf7, 0x02, 0x0a,
	0x09, 0x4e, 0x66, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x86, 0x01, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x4e, 0x46, 0x54, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x27, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x46, 0x54, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x46, 0x54,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x6e, 0x66, 0x74, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x44, 0x61, 0x74, 0x61, 0x2f, 0x67, 0x65, 0x74, 0x4e, 0x46, 0x54, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x74, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x18, 0x12, 0x16, 0x2f, 0x6e, 0x66, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x64, 0x50, 0x61, 0x67, 0x65, 0x12, 0x6b, 0x0a, 0x08, 0x53, 0x65, 0x6c,
	0x66, 0x50, 0x61, 0x67, 0x65, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x66, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e,
	0x66, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x66,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x15, 0x12, 0x13, 0x2f, 0x6e, 0x66, 0x74, 0x2f, 0x73, 0x65, 0x6c, 0x66, 0x2d, 0x70, 0x61, 0x67,
	0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x3c, 0x0a, 0x10, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66,
	0x74, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x26, 0x6e, 0x66,
	0x74, 0x2d, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x66, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nftmarket_v1_nftmarket_proto_rawDescOnce sync.Once
	file_nftmarket_v1_nftmarket_proto_rawDescData = file_nftmarket_v1_nftmarket_proto_rawDesc
)

func file_nftmarket_v1_nftmarket_proto_rawDescGZIP() []byte {
	file_nftmarket_v1_nftmarket_proto_rawDescOnce.Do(func() {
		file_nftmarket_v1_nftmarket_proto_rawDescData = protoimpl.X.CompressGZIP(file_nftmarket_v1_nftmarket_proto_rawDescData)
	})
	return file_nftmarket_v1_nftmarket_proto_rawDescData
}

var file_nftmarket_v1_nftmarket_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_nftmarket_v1_nftmarket_proto_goTypes = []interface{}{
	(*GetNFTMetadataRequest)(nil),   // 0: api.nftmarket.v1.GetNFTMetadataRequest
	(*GetNFTMetadataReply)(nil),     // 1: api.nftmarket.v1.GetNFTMetadataReply
	(*ListedPageRequest)(nil),       // 2: api.nftmarket.v1.ListedPageRequest
	(*ListedPageReply)(nil),         // 3: api.nftmarket.v1.ListedPageReply
	(*SelfPageRequest)(nil),         // 4: api.nftmarket.v1.SelfPageRequest
	(*SelfPageReply)(nil),           // 5: api.nftmarket.v1.SelfPageReply
	(*ListedPageReply_NftItem)(nil), // 6: api.nftmarket.v1.ListedPageReply.NftItem
	(*structpb.Struct)(nil),         // 7: google.protobuf.Struct
}
var file_nftmarket_v1_nftmarket_proto_depIdxs = []int32{
	7, // 0: api.nftmarket.v1.GetNFTMetadataReply.data:type_name -> google.protobuf.Struct
	6, // 1: api.nftmarket.v1.ListedPageReply.data:type_name -> api.nftmarket.v1.ListedPageReply.NftItem
	7, // 2: api.nftmarket.v1.SelfPageReply.data:type_name -> google.protobuf.Struct
	0, // 3: api.nftmarket.v1.Nftmarket.GetNFTMetadata:input_type -> api.nftmarket.v1.GetNFTMetadataRequest
	2, // 4: api.nftmarket.v1.Nftmarket.ListedPage:input_type -> api.nftmarket.v1.ListedPageRequest
	4, // 5: api.nftmarket.v1.Nftmarket.SelfPage:input_type -> api.nftmarket.v1.SelfPageRequest
	1, // 6: api.nftmarket.v1.Nftmarket.GetNFTMetadata:output_type -> api.nftmarket.v1.GetNFTMetadataReply
	3, // 7: api.nftmarket.v1.Nftmarket.ListedPage:output_type -> api.nftmarket.v1.ListedPageReply
	5, // 8: api.nftmarket.v1.Nftmarket.SelfPage:output_type -> api.nftmarket.v1.SelfPageReply
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_nftmarket_v1_nftmarket_proto_init() }
func file_nftmarket_v1_nftmarket_proto_init() {
	if File_nftmarket_v1_nftmarket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nftmarket_v1_nftmarket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNFTMetadataRequest); i {
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
		file_nftmarket_v1_nftmarket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNFTMetadataReply); i {
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
		file_nftmarket_v1_nftmarket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListedPageRequest); i {
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
		file_nftmarket_v1_nftmarket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListedPageReply); i {
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
		file_nftmarket_v1_nftmarket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelfPageRequest); i {
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
		file_nftmarket_v1_nftmarket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelfPageReply); i {
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
		file_nftmarket_v1_nftmarket_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListedPageReply_NftItem); i {
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
			RawDescriptor: file_nftmarket_v1_nftmarket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nftmarket_v1_nftmarket_proto_goTypes,
		DependencyIndexes: file_nftmarket_v1_nftmarket_proto_depIdxs,
		MessageInfos:      file_nftmarket_v1_nftmarket_proto_msgTypes,
	}.Build()
	File_nftmarket_v1_nftmarket_proto = out.File
	file_nftmarket_v1_nftmarket_proto_rawDesc = nil
	file_nftmarket_v1_nftmarket_proto_goTypes = nil
	file_nftmarket_v1_nftmarket_proto_depIdxs = nil
}
