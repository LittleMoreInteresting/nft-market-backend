package service

import (
	"context"
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
	/*if req.Confirmed {
		// 确认
		return &pb.ReceiveReply{}, nil
	}*/
	err := s.stream.Receive(ctx, req)
	return &pb.ReceiveReply{}, err
}
