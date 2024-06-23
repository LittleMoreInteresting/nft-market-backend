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

type ERC721Event struct {
	marketUC *biz.NFTMarketUseCase
	log      *log.Helper
}

var _ eventserver.EventSever = (*ERC721Event)(nil)
var (
	NFTAddr = common.HexToAddress("0x43CBe3CB4f095D1d7230e427aA4582638B2BF1A8")
)

func NewERC721Event(logger log.Logger, uc *biz.NFTMarketUseCase) *ERC721Event {
	return &ERC721Event{
		log:      log.NewHelper(log.With(logger, "module", "NFTMarketEvent")),
		marketUC: uc,
	}
}

func (e *ERC721Event) Addr() common.Address {
	return NFTAddr
}

func (e *ERC721Event) Metadata() *bind.MetaData {
	return contracts.HoraceNFTMetaData
}

func (e *ERC721Event) Handlers() map[string]eventserver.EventHandler {
	return map[string]eventserver.EventHandler{
		"Transfer": func(ctx context.Context, log types.Log) {
			err := e.marketUC.NftTransfer(ctx, log, ChainId, NFTAddr.String())
			if err != nil {
				e.log.Error(err)
			}
		},
	}
}
