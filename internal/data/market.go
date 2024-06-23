package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	pb "nft-market-backend/api/nftmarket/v1"
	"nft-market-backend/internal/biz"
	"nft-market-backend/internal/data/entity"
)

type MarketRepo struct {
	data *Data
	log  *log.Helper
}

func NewMarketRepo(data *Data, logger log.Logger) biz.MarketRepo {
	l := log.NewHelper(log.With(logger, "module", "StreamRepo"))
	return &MarketRepo{
		data: data,
		log:  l,
	}
}

func (m *MarketRepo) NftListed(ctx context.Context, req *pb.ListedPageRequest) ([]*pb.ListedPageReply_NftItem, error) {
	collection := m.data.mongo.Db.Collection((&entity.NftListed{}).CollectionName())
	skip := req.PageSize * (req.Page - 1)
	sort := bson.M{"_id": -1}
	findOptions := options.Find()
	findOptions.SetLimit(req.PageSize)
	findOptions.SetSkip(skip)
	findOptions.SetSort(sort)

	filter := bson.M{}
	if req.MarketPlaceAddr != "" {
		filter["marketPlaceAddress"] = req.MarketPlaceAddr
	}
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var results []*pb.ListedPageReply_NftItem
	for cursor.Next(ctx) {
		var result entity.NftListed
		err := cursor.Decode(&result)
		if err != nil {
			m.log.Error(err)
			continue
		}
		results = append(results, &pb.ListedPageReply_NftItem{
			Id:                 result.Id,
			ChainId:            result.ChainId,
			MarketPlaceAddress: result.MarketPlaceAddress,
			Seller:             result.Seller,
			NftAddress:         result.NftAddress,
			TokenId:            result.TokenId,
			Price:              result.Price,
			ListedTime:         result.ListedTime,
			UpdateTime:         result.UpdateTime,
		})
	}

	if err := cursor.Err(); err != nil {
		m.log.Error(err)
		return nil, err
	}
	return results, nil
}

func (m *MarketRepo) SelfNftList(ctx context.Context, req *pb.SelfPageRequest) ([]*biz.SelfNft, error) {
	collection := m.data.mongo.Db.Collection((&entity.NftTransfer{}).CollectionName())
	skip := req.PageSize * (req.Page - 1)
	sort := bson.M{"_id": -1}
	findOptions := options.Find()
	findOptions.SetLimit(req.PageSize)
	findOptions.SetSkip(skip)
	findOptions.SetSort(sort)

	filter := bson.M{
		"chainId": req.ChainId,
		"owner":   req.Owner,
	}
	if req.NftAddress != "" {
		filter["nftAddress"] = req.NftAddress
	}
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var results []*biz.SelfNft
	for cursor.Next(ctx) {
		var result entity.NftTransfer
		err := cursor.Decode(&result)
		if err != nil {
			m.log.Error(err)
			continue
		}
		results = append(results, &biz.SelfNft{
			Id:         result.Id,
			ChainId:    result.ChainId,
			NftAddress: result.NftAddress,
			TokenId:    result.TokenId,
			From:       result.From,
			Owner:      result.Owner,
		})
	}

	if err := cursor.Err(); err != nil {
		m.log.Error(err)
		return nil, err
	}
	return results, nil
}
