// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"nft-market-backend/internal/biz"
	"nft-market-backend/internal/conf"
	"nft-market-backend/internal/data"
	"nft-market-backend/internal/server"
	"nft-market-backend/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, moralis *conf.Moralis, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	streamRepo := data.NewStreamRepo(dataData, logger)
	streamUseCase := biz.NewStreamUseCase(streamRepo, logger)
	streamService := service.NewStreamService(logger, streamUseCase)
	marketRepo := data.NewMarketRepo(dataData, logger)
	marketUseCase := biz.NewMarketUseCase(logger, marketRepo, moralis)
	nftmarketService := service.NewNftmarketService(logger, marketUseCase)
	httpServer := server.NewHTTPServer(confServer, moralis, streamService, nftmarketService, logger)
	erc20Event := service.NewERC20Event()
	nftMarketUseCase := biz.NewNFTMarketUseCase(streamRepo, logger)
	nftMarketEvent := service.NewNFTMarketEvent(logger, nftMarketUseCase)
	erc721Event := service.NewERC721Event(logger, nftMarketUseCase)
	eventserverServer := server.NewEventServer(erc20Event, nftMarketEvent, erc721Event)
	app := newApp(logger, httpServer, eventserverServer)
	return app, func() {
		cleanup()
	}, nil
}
