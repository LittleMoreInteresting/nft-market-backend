package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewStreamService,
	NewNftmarketService,
	NewERC20Event,
	NewNFTMarketEvent,
	NewERC721Event,
)
