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
type EventQuery struct {
	ChainId    string
	NftAddress string
	TokenId    string
}

type StreamRepo interface {
	FindOne(ctx context.Context, event *EventQuery) (*EventListed, bool, error)
	SaveListed(ctx context.Context, event *EventListed) error
}

type StreamUseCase struct {
	log    *log.Helper
	repo   StreamRepo
	abi    *abi.ABI
	market *contracts.HoraceNFTMarket // 仅用于解析日志，不做线上交互
}

func NewStreamUseCase(repo StreamRepo, logger log.Logger) *StreamUseCase {
	l := log.NewHelper(log.With(logger, "module", "StreamUseCase"))
	marketAbi, _ := contracts.HoraceNFTMarketMetaData.GetAbi()
	market, _ := contracts.NewHoraceNFTMarket(common.Address{}, nil)
	return &StreamUseCase{
		repo:   repo,
		log:    l,
		abi:    marketAbi,
		market: market,
	}
}

func (uc *StreamUseCase) Receive(ctx context.Context, req *v1.ReceiveRequest) error {
	if req.Confirmed {
		// 确认
		return nil
	}
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
		case uc.abi.Events["NftListed"].ID.Hex():
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
				OldPrice:           nftListed.OldPrice.String(),
				ListedTime:         time.Now().Unix(),
			}
			// list
			list, exist, err := uc.repo.FindOne(ctx, &EventQuery{
				ChainId:    eventList.ChainId,
				NftAddress: eventList.NftAddress,
				TokenId:    eventList.TokenId,
			})
			if err != nil {
				return err
			}
			if exist {
				// update
				uc.log.Info("update")
				eventList.Id = list.Id
				eventList.UpdateTime = time.Now().Unix()
			}
			return uc.repo.SaveListed(ctx, eventList)

		case uc.abi.Events["BuyNFT"].ID.Hex():
			// Buy
		case uc.abi.Events["CancelListing"].ID.Hex():
			// cancel

		}
	}

	return nil
}
