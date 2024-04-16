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
	b, err := s.findOneEntity(ctx, listed, event)
	if err != nil {
		return nil, b, err
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
	b, err := s.findOneEntity(ctx, listed, event)
	if err != nil {
		return nil, b, err
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
	b, err := s.findOneEntity(ctx, listed, event)
	if err != nil {
		return nil, b, err
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

func (s *StreamRepo) SaveNftTransfer(ctx context.Context, event *biz.EventTransfer) error {
	entry := &entity.NftTransfer{}
	db := s.data.mongo.Db.Collection(entry.CollectionName())
	err := copier.Copy(entry, event)
	if err != nil {
		return err
	}
	entry.TransferTime = time.Now().Unix()
	if entry.Id == "" {
		_, err = db.InsertOne(ctx, entry)
		return err
	}
	// update
	hex, _ := primitive.ObjectIDFromHex(entry.Id)
	_, err = db.UpdateOne(ctx,
		bson.M{"_id": hex},
		bson.M{"$set": bson.M{
			"from":         entry.From,
			"owner":        entry.Owner,
			"transferTime": entry.TransferTime,
		}},
	)
	return err
}

func (s *StreamRepo) FindOneTransfer(ctx context.Context, event *biz.EventQuery) (*biz.EventTransfer, bool, error) {
	listed := &entity.NftTransfer{}
	b, err := s.findOneEntity(ctx, listed, event)
	if err != nil {
		return nil, b, err
	}
	res := &biz.EventTransfer{}
	_ = copier.Copy(res, listed)
	return res, b, nil
}

func (s *StreamRepo) findOneEntity(ctx context.Context, e entity.Entity, query *biz.EventQuery) (bool, error) {
	result := s.data.mongo.Db.Collection(e.CollectionName()).FindOne(ctx, bson.D{
		{"chainId", query.ChainId},
		{"nftAddress", query.NftAddress},
		{"tokenId", query.TokenId},
	})

	err := result.Decode(e)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
