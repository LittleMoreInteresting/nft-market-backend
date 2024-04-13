package entity

type NftBought struct {
	Id                 string `bson:"_id"`
	ChainId            string `bson:"chainId"`
	MarketPlaceAddress string `bson:"marketPlaceAddress"`
	Buyer              string `bson:"buyer"`
	NftAddress         string `bson:"nftAddress"`
	TokenId            string `bson:"tokenId"`
	Price              string `bson:"price"`
	BoughtTime         int64  `bson:"boughtTime"`
}

func (l *NftBought) CollectionName() string {
	return "NftBought"
}
