package biz

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/structpb"
	"io"
	"net/http"
	pb "nft-market-backend/api/nftmarket/v1"
	"nft-market-backend/internal/conf"
	"nft-market-backend/pkg/chainId"
	"sync"
)

type SelfNft struct {
	Id         string
	ChainId    string
	NftAddress string
	TokenId    string
	Owner      string
	From       string
}
type MarketRepo interface {
	NftListed(ctx context.Context, req *pb.ListedPageRequest) ([]*pb.ListedPageReply_NftItem, error)
	SelfNftList(ctx context.Context, req *pb.SelfPageRequest) ([]*SelfNft, error)
}

type MarketUseCase struct {
	log     *log.Helper
	repo    MarketRepo
	moralis *conf.Moralis
}

func NewMarketUseCase(logger log.Logger, repo MarketRepo, moralis *conf.Moralis) *MarketUseCase {
	return &MarketUseCase{
		log:     log.NewHelper(log.With(logger, "module", "MarketUseCase")),
		repo:    repo,
		moralis: moralis,
	}
}

func (m *MarketUseCase) GetMetadata(ctx context.Context, req *pb.GetNFTMetadataRequest) (*structpb.Struct, error) {
	if m.moralis == nil {
		return nil, errors.New(500, "moralis config error", "moralis config error")
	}
	url := m.buildUrl(req)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("X-API-Key", m.moralis.Key)

	res, _ := http.DefaultClient.Do(request)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	ret := &structpb.Struct{}
	err = json.Unmarshal(body, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (m *MarketUseCase) buildUrl(req *pb.GetNFTMetadataRequest) string {
	return m.moralis.Url + req.NftAddress + "/" + req.TokenId + "?chain=" + req.ChainId +
		"&format=decimal&normalizeMetadata=true&media_items=false"
}

func (m *MarketUseCase) NftListed(ctx context.Context, req *pb.ListedPageRequest) ([]*pb.ListedPageReply_NftItem, error) {
	return m.repo.NftListed(ctx, req)
}

func (m *MarketUseCase) SelfPage(ctx context.Context, req *pb.SelfPageRequest) ([]*structpb.Struct, error) {
	list, err := m.repo.SelfNftList(ctx, req)
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	var metadataCh = make(chan *structpb.Struct, len(list))
	for _, item := range list {
		wg.Add(1)
		go func(i *SelfNft) {
			defer wg.Done()
			metadata, errGetMetadata := m.GetMetadata(ctx, &pb.GetNFTMetadataRequest{
				NftAddress: i.NftAddress,
				TokenId:    i.TokenId,
				ChainId:    chainId.ParseChainId(i.ChainId),
			})
			if errGetMetadata != nil {
				return
			}
			metadataCh <- metadata
		}(item)
	}
	wg.Wait()
	close(metadataCh)
	result := make([]*structpb.Struct, 0, len(list))
	for ch := range metadataCh {
		result = append(result, ch)
	}
	return result, nil
}
