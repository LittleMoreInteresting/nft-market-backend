package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"nft-market-backend/internal/biz"
	"nft-market-backend/internal/data/entity"
	"time"
)

type StreamRepo struct {
	data *Data
	log  *log.Helper
}

func NewStreamRepo(data *Data, logger log.Logger) biz.StreamRepo {
	l := log.NewHelper(log.With(logger, "module", "StreamRepo"))
	return &StreamRepo{
		data: data,
		log:  l,
	}
}

func (s *StreamRepo) SaveListed(ctx context.Context, event *biz.EventListed) error {
	listed := &entity.NftListed{}
	db := s.data.mongo.Db.Collection(listed.CollectionName())
	err := copier.Copy(listed, event)
	if err != nil {
		return err
	}
	if event.Id == "" { // Insert
		listed.ListedTime = time.Now().Unix()
		_, err = db.InsertOne(ctx, listed)
		return err
	}
	// update
	hex, _ := primitive.ObjectIDFromHex(listed.Id)
	_, err = db.UpdateOne(ctx,
		bson.M{"_id": hex},
		bson.M{"$set": bson.M{
			"price":      listed.Price,
			"oldPrice":   listed.OldPrice,
			"updateTime": listed.UpdateTime,
		}},
	)
	return err
}

func (s *StreamRepo) FindOneListed(ctx context.Context, event *biz.EventQuery) (*biz.EventListed, bool, error) {
	listed := &entity.NftListed{}
	db := s.data.mongo.Db.Collection(listed.CollectionName())
	listResutl := db.FindOne(ctx, bson.D{
		{"chainId", event.ChainId},
		{"nftAddress", event.NftAddress},
		{"tokenId", event.TokenId},
	})
	err := listResutl.Decode(listed)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, false, nil
		}
		return nil, false, err
	}
	res := &biz.EventListed{}
	_ = copier.Copy(res, listed)
	return res, true, nil
}

func (s *StreamRepo) RemoveListed(ctx context.Context, event *biz.EventQuery) error {
	listed := &entity.NftListed{}
	_, err := s.data.mongo.Db.Collection(listed.CollectionName()).
		DeleteOne(ctx, bson.D{
			{"chainId", event.ChainId},
			{"nftAddress", event.NftAddress},
			{"tokenId", event.TokenId},
		})
	return err
}

func (s *StreamRepo) FindOneBought(ctx context.Context, event *biz.EventQuery) (*biz.EventBought, bool, error) {
	listed := &entity.NftBought{}
	db := s.data.mongo.Db.Collection(listed.CollectionName())
	listResutl := db.FindOne(ctx, bson.D{
		{"chainId", event.ChainId},
		{"nftAddress", event.NftAddress},
		{"tokenId", event.TokenId},
	})
	err := listResutl.Decode(listed)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, false, nil
		}
		return nil, false, err
	}
	res := &biz.EventBought{}
	_ = copier.Copy(res, listed)
	return res, true, nil
}

func (s *StreamRepo) SaveBought(ctx context.Context, event *biz.EventBought) error {
	entry := &entity.NftBought{}
	db := s.data.mongo.Db.Collection(entry.CollectionName())
	err := copier.Copy(entry, event)
	if err != nil {
		return err
	}
	entry.BoughtTime = time.Now().Unix()
	if event.Id == "" { // Insert
		_, err = db.InsertOne(ctx, entry)
		return err
	}
	// update
	hex, _ := primitive.ObjectIDFromHex(entry.Id)
	_, err = db.UpdateOne(ctx,
		bson.M{"_id": hex},
		bson.M{"$set": bson.M{
			"price":      entry.Price,
			"boughtTime": entry.BoughtTime,
		}},
	)
	return err
}

func (s *StreamRepo) FindOneCancel(ctx context.Context, event *biz.EventQuery) (*biz.EventCancel, bool, error) {
	listed := &entity.NftCanceled{}
	db := s.data.mongo.Db.Collection(listed.CollectionName())
	listResutl := db.FindOne(ctx, bson.D{
		{"chainId", event.ChainId},
		{"nftAddress", event.NftAddress},
		{"tokenId", event.TokenId},
	})
	err := listResutl.Decode(listed)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, false, nil
		}
		return nil, false, err
	}
	res := &biz.EventCancel{}
	_ = copier.Copy(res, listed)
	return res, true, nil
}

func (s *StreamRepo) SaveCancel(ctx context.Context, event *biz.EventCancel) error {
	entry := &entity.NftCanceled{}
	db := s.data.mongo.Db.Collection(entry.CollectionName())
	err := copier.Copy(entry, event)
	if err != nil {
		return err
	}
	entry.CanceledTime = time.Now().Unix()
	if event.Id == "" { // Insert
		_, err = db.InsertOne(ctx, entry)
		return err
	}
	// update
	hex, _ := primitive.ObjectIDFromHex(entry.Id)
	_, err = db.UpdateOne(ctx,
		bson.M{"_id": hex},
		bson.M{"$set": bson.M{
			"canceledTime": entry.CanceledTime,
		}},
	)
	return err
}
