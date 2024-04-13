package entity

type NftCanceled struct {
	Id                 string `bson:"_id"`
	ChainId            string `bson:"chainId"`
	MarketPlaceAddress string `bson:"marketPlaceAddress"`
	Seller             string `bson:"seller"`
	NftAddress         string `bson:"nftAddress"`
	TokenId            string `bson:"tokenId"`
	Price              string `bson:"price"`
	CanceledTime       int64  `bson:"canceledTime"`
}

func (l *NftCanceled) CollectionName() string {
	return "NftCanceled"
}
