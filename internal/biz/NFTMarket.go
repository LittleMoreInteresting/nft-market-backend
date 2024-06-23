package biz

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kratos/kratos/v2/log"
	contracts "nft-market-backend/abi"
	"time"
)

type NFTMarketUseCase struct {
	log    *log.Helper
	repo   StreamRepo
	market *contracts.HoraceNFTMarket
}

func NewNFTMarketUseCase(repo StreamRepo, logger log.Logger) *NFTMarketUseCase {
	l := log.NewHelper(log.With(logger, "module", "NFTMarketUseCase"))
	market, _ := contracts.NewHoraceNFTMarket(common.Address{}, nil)
	return &NFTMarketUseCase{
		repo:   repo,
		log:    l,
		market: market,
	}
}

func (uc *NFTMarketUseCase) Listed(ctx context.Context, eventLog types.Log, chainId string, MarketPlaceAddress string) error {
	nftListed, err := uc.market.ParseNftListed(eventLog)
	if err != nil {
		return err
	}
	eventList := &EventListed{
		ChainId:            chainId,
		MarketPlaceAddress: MarketPlaceAddress,
		Seller:             nftListed.Seller.Hex(),
		NftAddress:         nftListed.NftAddress.Hex(),
		TokenId:            nftListed.TokenId.String(),
		Price:              nftListed.Price.String(),
	}
	// list
	list, exist, err := uc.repo.FindOneListed(ctx, &EventQuery{
		ChainId:    eventList.ChainId,
		NftAddress: eventList.NftAddress,
		TokenId:    eventList.TokenId,
	})
	if err != nil {
		return err
	}
	if exist {
		// update
		eventList.Id = list.Id
		eventList.UpdateTime = time.Now().Unix()
	}
	err = uc.repo.SaveListed(ctx, eventList)
	if err != nil {
		return err
	}
	return nil
}

func (uc *NFTMarketUseCase) Buy(ctx context.Context, eventLog types.Log, chainId string, MarketPlaceAddress string) error {
	nftBought, err := uc.market.ParseBuyNFT(eventLog)
	if err != nil {
		return err
	}
	event := &EventBought{
		ChainId:            chainId,
		MarketPlaceAddress: MarketPlaceAddress,
		Buyer:              nftBought.Buyer.Hex(),
		NftAddress:         nftBought.NftAddress.Hex(),
		TokenId:            nftBought.TokenId.String(),
		Price:              nftBought.Price.String(),
	}
	query := &EventQuery{
		ChainId:    event.ChainId,
		NftAddress: event.NftAddress,
		TokenId:    event.TokenId,
	}
	bought, exist, err := uc.repo.FindOneBought(ctx, query)
	if err != nil {
		return err
	}
	if exist {
		event.Id = bought.Id
	}
	err = uc.repo.SaveBought(ctx, event)
	if err != nil {
		return err
	}
	err = uc.repo.RemoveListed(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func (uc *NFTMarketUseCase) CancelListing(ctx context.Context, eventLog types.Log, chainId string, MarketPlaceAddress string) error {
	cancel, err := uc.market.ParseCancelListing(eventLog)
	if err != nil {
		return err
	}
	event := &EventCancel{
		ChainId:            chainId,
		MarketPlaceAddress: MarketPlaceAddress,
		Seller:             cancel.Seller.Hex(),
		NftAddress:         cancel.NftAddress.Hex(),
		TokenId:            cancel.TokenId.String(),
	}
	query := &EventQuery{
		ChainId:    event.ChainId,
		NftAddress: event.NftAddress,
		TokenId:    event.TokenId,
	}
	cancelList, exist, err := uc.repo.FindOneCancel(ctx, query)
	if err != nil {
		return err
	}
	if exist {
		event.Id = cancelList.Id
	}
	err = uc.repo.SaveCancel(ctx, event)
	if err != nil {
		return err
	}
	err = uc.repo.RemoveListed(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
