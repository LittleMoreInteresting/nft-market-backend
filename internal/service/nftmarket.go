package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
	pb "nft-market-backend/api/nftmarket/v1"
	"nft-market-backend/internal/biz"
)

type NftmarketService struct {
	pb.UnimplementedNftmarketServer
	log    *log.Helper
	market *biz.MarketUseCase
}

var _ pb.NftmarketHTTPServer = (*NftmarketService)(nil)

func NewNftmarketService(logger log.Logger, market *biz.MarketUseCase) *NftmarketService {
	return &NftmarketService{
		log:    log.NewHelper(log.With(logger, "module", "NftmarketService")),
		market: market,
	}
}

func (s *NftmarketService) GetNFTMetadata(ctx context.Context, req *pb.GetNFTMetadataRequest) (*pb.GetNFTMetadataReply, error) {
	metadata, err := s.market.GetMetadata(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.GetNFTMetadataReply{
		Code: http.StatusOK,
		Data: metadata,
		Msg:  "Success",
	}, nil
}

func (s *NftmarketService) ListedPage(ctx context.Context, req *pb.ListedPageRequest) (*pb.ListedPageReply, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 50
	}
	listed, err := s.market.NftListed(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ListedPageReply{
		Code: http.StatusOK,
		Data: listed,
		Msg:  "Success",
	}, nil
}

func (s *NftmarketService) SelfPage(ctx context.Context, request *pb.SelfPageRequest) (*pb.SelfPageReply, error) {
	page, err := s.market.SelfPage(ctx, request)
	if err != nil {
		return nil, err
	}
	return &pb.SelfPageReply{
		Code: http.StatusOK,
		Data: page,
		Msg:  "Success",
	}, nil
}
