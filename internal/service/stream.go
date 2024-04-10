package service

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"nft-market-backend/internal/biz"

	pb "nft-market-backend/api/moralis/v1"
)

type StreamService struct {
	pb.UnimplementedStreamServer
	log    *log.Helper
	stream *biz.StreamUseCase
}

func NewStreamService(logger log.Logger, stream *biz.StreamUseCase) *StreamService {
	return &StreamService{
		log:    log.NewHelper(log.With(logger, "module", "StreamService")),
		stream: stream,
	}
}

func (s *StreamService) Receive(ctx context.Context, req *pb.ReceiveRequest) (*pb.ReceiveReply, error) {
	str, err := json.Marshal(req)
	log.Info(string(str), err)
	return &pb.ReceiveReply{}, nil
}
