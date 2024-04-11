package entity

type NftListed struct {
	Id                 string `bson:"_id,omitempty"`
	ChainId            string `bson:"chainId"`
	MarketPlaceAddress string `bson:"marketPlaceAddress"`
	Seller             string `bson:"seller"`
	NftAddress         string `bson:"nftAddress"`
	TokenId            string `bson:"tokenId"`
	Price              string `bson:"price"`
	OldPrice           string `bson:"oldPrice"`
	ListedTime         int64  `bson:"listedTime"`
	UpdateTime         int64  `bson:"updateTime"`
}

func (l *NftListed) CollectionName() string {
	return "NftListed"
}
