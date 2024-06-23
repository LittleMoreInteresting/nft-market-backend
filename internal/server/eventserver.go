package server

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"nft-market-backend/internal/service"
	"nft-market-backend/pkg/eventserver"
)

func NewEventServer(erc20 *service.ERC20Event, nftMarket *service.NFTMarketEvent) *eventserver.Server {
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/a93BJAOxsYWxlpZJAGq_nPYPUWDpJSsq")
	if err != nil {
		panic(err)
	}
	opts := []eventserver.ServerOption{
		eventserver.WithClient(client),
	}
	esrv := eventserver.NewServer(opts...)
	esrv.Bind(erc20)
	esrv.Bind(nftMarket)
	return esrv
}
