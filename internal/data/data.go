package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"nft-market-backend/internal/conf"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewStreamRepo)

// Data .
type Data struct {
	mongo *MongoManager
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	newMongo, err := NewMongo(c.Mongodb)
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		newMongo.Close()
	}
	if err != nil {
		return nil, cleanup, err
	}
	return &Data{
		mongo: newMongo,
	}, cleanup, nil
}

type MongoManager struct {
	Cli *mongo.Client
	Db  *mongo.Database
}

func NewMongo(conf *conf.Data_Mongodb) (*MongoManager, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(conf.Url)
	clientOptions.SetAuth(options.Credential{
		Username: conf.Username,
		Password: conf.Password,
	})
	clientOptions.SetMinPoolSize(uint64(conf.MinPoolSize))
	clientOptions.SetMaxPoolSize(uint64(conf.MaxPoolSize))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	m := &MongoManager{
		Cli: client,
	}
	m.Db = m.Cli.Database(conf.Db)
	return m, nil
}

func (m *MongoManager) Close() {
	_ = m.Cli.Disconnect(context.TODO())
}
