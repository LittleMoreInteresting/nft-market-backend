package eventserver

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/go-kratos/kratos/v2/log"
)

type ServerOption func(o *Server)

func WithContext(ctx context.Context) ServerOption {
	return func(s *Server) {
		s.baseCtx = ctx
	}
}

func Logger(logger log.Logger) ServerOption {
	return func(s *Server) {
		s.log = log.NewHelper(logger)
	}
}

func WithClient(cli *ethclient.Client) ServerOption {
	return func(s *Server) {
		s.client = cli
	}
}
