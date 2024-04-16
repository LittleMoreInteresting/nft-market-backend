package biz

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kratos/kratos/v2/log"
	contracts "nft-market-backend/abi"
	v1 "nft-market-backend/api/moralis/v1"
	"time"
)

type EventListed struct {
	Id                 string
	ChainId            string
	MarketPlaceAddress string
	Seller             string
	NftAddress         string
	TokenId            string
	Price              string
	OldPrice           string
	ListedTime         int64
	UpdateTime         int64
}
type EventBought struct {
	Id                 string
	ChainId            string
	MarketPlaceAddress string
	Buyer              string
	NftAddress         string
	TokenId            string
	Price              string
}
type EventCancel struct {
	Id                 string
	ChainId            string
	MarketPlaceAddress string
	Seller             string
	NftAddress         string
	TokenId            string
	Price              string
}
type EventQuery struct {
	ChainId    string
	NftAddress string
	TokenId    string
}

type EventTransfer struct {
	Id         string
	ChainId    string
	NftAddress string
	TokenId    string
	Owner      string
	From       string
}

type StreamRepo interface {
	FindOneListed(ctx context.Context, event *EventQuery) (*EventListed, bool, error)
	SaveListed(ctx context.Context, event *EventListed) error
	RemoveListed(ctx context.Context, event *EventQuery) error
	FindOneBought(ctx context.Context, event *EventQuery) (*EventBought, bool, error)
	SaveBought(ctx context.Context, event *EventBought) error
	FindOneCancel(ctx context.Context, event *EventQuery) (*EventCancel, bool, error)
	SaveCancel(ctx context.Context, event *EventCancel) error
	SaveNftTransfer(ctx context.Context, event *EventTransfer) error
	FindOneTransfer(ctx context.Context, event *EventQuery) (*EventTransfer, bool, error)
}

type StreamUseCase struct {
	log       *log.Helper
	repo      StreamRepo
	marketAbi *abi.ABI
	market    *contracts.HoraceNFTMarket // 仅用于解析日志，不做线上交互
	nftAbi    *abi.ABI
	nft       *contracts.HoraceNFT
}

func NewStreamUseCase(repo StreamRepo, logger log.Logger) *StreamUseCase {
	l := log.NewHelper(log.With(logger, "module", "StreamUseCase"))
	marketAbi, _ := contracts.HoraceNFTMarketMetaData.GetAbi()
	market, _ := contracts.NewHoraceNFTMarket(common.Address{}, nil)
	nftAbi, _ := contracts.HoraceNFTMetaData.GetAbi()
	nft, _ := contracts.NewHoraceNFT(common.Address{}, nil)
	return &StreamUseCase{
		repo:      repo,
		log:       l,
		marketAbi: marketAbi,
		nftAbi:    nftAbi,
		market:    market,
		nft:       nft,
	}
}

func (uc *StreamUseCase) Receive(ctx context.Context, req *v1.ReceiveRequest) error {
	chainId, _ := hexutil.DecodeBig(req.ChainId)
	logs := req.Logs
	for _, l := range logs {
		decode, _ := hexutil.Decode(l.Data)
		eventLog := types.Log{
			Address: common.HexToAddress(l.Address),
			Topics: []common.Hash{
				common.HexToHash(l.Topic0),
				common.HexToHash(l.Topic1),
				common.HexToHash(l.Topic2),
				common.HexToHash(l.Topic3),
			},
			Data: decode,
		}

		switch l.Topic0 {
		case uc.marketAbi.Events["NftListed"].ID.Hex():
			nftListed, err := uc.market.ParseNftListed(eventLog)
			if err != nil {
				return err
			}
			eventList := &EventListed{
				ChainId:            chainId.String(),
				MarketPlaceAddress: l.Address,
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
		case uc.marketAbi.Events["BuyNFT"].ID.Hex():
			// Buy
			nftBought, err := uc.market.ParseBuyNFT(eventLog)
			if err != nil {
				log.Error("ParseBuyNFT error:", err)
				continue
			}
			event := &EventBought{
				ChainId:            chainId.String(),
				MarketPlaceAddress: l.Address,
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
				log.Error("FindOneBought error:", err)
				continue
			}
			if exist {
				event.Id = bought.Id
			}
			err = uc.repo.SaveBought(ctx, event)
			if err != nil {
				log.Error("SaveBought error:", err)
				continue
			}
			err = uc.repo.RemoveListed(ctx, query)
			if err != nil {
				log.Error("SaveBought error:", err)
				continue
			}
		case uc.marketAbi.Events["CancelListing"].ID.Hex():
			// cancel
			cancel, err := uc.market.ParseCancelListing(eventLog)
			if err != nil {
				log.Error("ParseCancelListing error:", err)
				continue
			}
			event := &EventCancel{
				ChainId:            chainId.String(),
				MarketPlaceAddress: l.Address,
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
				log.Error("FindOneCancel error:", err)
				continue
			}
			if exist {
				event.Id = cancelList.Id
			}
			err = uc.repo.SaveCancel(ctx, event)
			if err != nil {
				log.Error("SaveCancel error:", err)
				continue
			}
			err = uc.repo.RemoveListed(ctx, query)
			if err != nil {
				log.Error("SaveCancel error:", err)
				continue
			}
		case uc.nftAbi.Events["Transfer"].ID.Hex():
			// Transfer
			transfer, err := uc.nft.ParseTransfer(eventLog)
			if err != nil {
				log.Error("ParseTransfer error:", err)
				continue
			}
			event := &EventTransfer{
				ChainId:    chainId.String(),
				From:       transfer.From.Hex(),
				Owner:      transfer.To.Hex(),
				NftAddress: l.Address,
				TokenId:    transfer.TokenId.String(),
			}
			query := &EventQuery{
				ChainId:    event.ChainId,
				NftAddress: event.NftAddress,
				TokenId:    event.TokenId,
			}

			tr, exist, err := uc.repo.FindOneTransfer(ctx, query)
			if err != nil {
				log.Error("FindOneCancel error:", err)
				continue
			}
			if exist {
				event.Id = tr.Id
			}
			err = uc.repo.SaveNftTransfer(ctx, event)
			if err != nil {
				log.Error("SaveCancel error:", err)
				continue
			}
		}
	}
	return nil
}
