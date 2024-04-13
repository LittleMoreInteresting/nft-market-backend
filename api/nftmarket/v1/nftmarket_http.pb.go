// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v3.17.3
// source: nftmarket/v1/nftmarket.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationNftmarketGetNFTMetadata = "/api.nftmarket.v1.Nftmarket/GetNFTMetadata"
const OperationNftmarketListedPage = "/api.nftmarket.v1.Nftmarket/ListedPage"

type NftmarketHTTPServer interface {
	GetNFTMetadata(context.Context, *GetNFTMetadataRequest) (*GetNFTMetadataReply, error)
	ListedPage(context.Context, *ListedPageRequest) (*ListedPageReply, error)
}

func RegisterNftmarketHTTPServer(s *http.Server, srv NftmarketHTTPServer) {
	r := s.Route("/")
	r.GET("/nft/metaData/getNFTMetadata", _Nftmarket_GetNFTMetadata0_HTTP_Handler(srv))
	r.GET("/nft/listed/listedPage", _Nftmarket_ListedPage0_HTTP_Handler(srv))
}

func _Nftmarket_GetNFTMetadata0_HTTP_Handler(srv NftmarketHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetNFTMetadataRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNftmarketGetNFTMetadata)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNFTMetadata(ctx, req.(*GetNFTMetadataRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetNFTMetadataReply)
		return ctx.Result(200, reply)
	}
}

func _Nftmarket_ListedPage0_HTTP_Handler(srv NftmarketHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListedPageRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNftmarketListedPage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListedPage(ctx, req.(*ListedPageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListedPageReply)
		return ctx.Result(200, reply)
	}
}

type NftmarketHTTPClient interface {
	GetNFTMetadata(ctx context.Context, req *GetNFTMetadataRequest, opts ...http.CallOption) (rsp *GetNFTMetadataReply, err error)
	ListedPage(ctx context.Context, req *ListedPageRequest, opts ...http.CallOption) (rsp *ListedPageReply, err error)
}

type NftmarketHTTPClientImpl struct {
	cc *http.Client
}

func NewNftmarketHTTPClient(client *http.Client) NftmarketHTTPClient {
	return &NftmarketHTTPClientImpl{client}
}

func (c *NftmarketHTTPClientImpl) GetNFTMetadata(ctx context.Context, in *GetNFTMetadataRequest, opts ...http.CallOption) (*GetNFTMetadataReply, error) {
	var out GetNFTMetadataReply
	pattern := "/nft/metaData/getNFTMetadata"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNftmarketGetNFTMetadata))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NftmarketHTTPClientImpl) ListedPage(ctx context.Context, in *ListedPageRequest, opts ...http.CallOption) (*ListedPageReply, error) {
	var out ListedPageReply
	pattern := "/nft/listed/listedPage"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNftmarketListedPage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}