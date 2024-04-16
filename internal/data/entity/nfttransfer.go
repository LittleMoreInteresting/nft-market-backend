package entity

type NftTransfer struct {
	Id           string `bson:"_id,omitempty"`
	ChainId      string `bson:"chainId"`
	NftAddress   string `bson:"nftAddress"`
	TokenId      string `bson:"tokenId"`
	Owner        string `bson:"owner"`
	From         string `bson:"from"`
	TransferTime int64  `bson:"transferTime"`
}

func (l *NftTransfer) CollectionName() string {
	return "NftTransfer"
}
