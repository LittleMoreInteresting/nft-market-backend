package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "nft-market-backend/api/moralis/v1"
)

type StreamRepo interface {
	Receive(ctx context.Context, req *v1.ReceiveRequest) error
}

type StreamUseCase struct {
	log  *log.Helper
	repo StreamRepo
}

func NewStreamUseCase(repo StreamRepo, logger log.Logger) *StreamUseCase {
	l := log.NewHelper(log.With(logger, "module", "StreamUseCase"))
	return &StreamUseCase{
		repo: repo,
		log:  l,
	}
}

func (uc *StreamUseCase) Receive(ctx context.Context, req *v1.ReceiveRequest) error {
	return uc.repo.Receive(ctx, req)
}
