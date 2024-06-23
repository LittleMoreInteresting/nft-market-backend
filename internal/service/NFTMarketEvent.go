package service

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kratos/kratos/v2/log"
	contracts "nft-market-backend/abi"
	"nft-market-backend/internal/biz"
	"nft-market-backend/pkg/eventserver"
)

type NFTMarketEvent struct {
	log      *log.Helper
	marketUC *biz.NFTMarketUseCase
}

var _ eventserver.EventSever = (*NFTMarketEvent)(nil)

var (
	NFTMarketAddr = common.HexToAddress("0x35D7A0A429f9fe06D79Ac351E1DAcC8d8B228914")
	ChainId       = "11155111"
)

func NewNFTMarketEvent(logger log.Logger, uc *biz.NFTMarketUseCase) *NFTMarketEvent {
	return &NFTMarketEvent{
		log:      log.NewHelper(log.With(logger, "module", "NFTMarketEvent")),
		marketUC: uc,
	}
}

func (m *NFTMarketEvent) Addr() common.Address {
	return NFTMarketAddr
}

func (m *NFTMarketEvent) Metadata() *bind.MetaData {
	return contracts.HoraceNFTMarketMetaData
}

func (m *NFTMarketEvent) Handlers() map[string]eventserver.EventHandler {

	return map[string]eventserver.EventHandler{
		"NftListed": func(ctx context.Context, log types.Log) {
			err := m.marketUC.Listed(ctx, log, ChainId, NFTMarketAddr.String())
			if err != nil {
				m.log.Error(err)
			}
		},
		"BuyNFT": func(ctx context.Context, log types.Log) {
			err := m.marketUC.Buy(ctx, log, ChainId, NFTMarketAddr.String())
			if err != nil {
				m.log.Error(err)
			}
		},

		"CancelListing": func(ctx context.Context, log types.Log) {
			err := m.marketUC.CancelListing(ctx, log, ChainId, NFTMarketAddr.String())
			if err != nil {
				m.log.Error(err)
			}
		},
	}
}
