package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "nft-market-backend/api/moralis/v1"
	"nft-market-backend/internal/biz"
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

func (s StreamRepo) Receive(ctx context.Context, req *v1.ReceiveRequest) error {
	//TODO implement me
	panic("implement me")
}
