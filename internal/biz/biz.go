package biz

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewStreamUseCase, NewMarketUseCase)

var ErrNFTNotFound = errors.New(404, "NFT not found", "NFT not found")
