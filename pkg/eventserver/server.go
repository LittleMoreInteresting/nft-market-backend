package eventserver

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"sync"
)

var (
	_ transport.Server = (*Server)(nil)
)

type Server struct {
	client  *ethclient.Client
	watches []*WatchEvent
	log     *log.Helper
	baseCtx context.Context
	err     error
	running bool
	lock    sync.Mutex
	close   chan struct{}
}

type EventSever interface {
	Addr() common.Address
	Metadata() *bind.MetaData
	Handlers() map[string]EventHandler
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		baseCtx: context.Background(),
		close:   make(chan struct{}),
		log:     log.NewHelper(log.DefaultLogger),
	}
	srv.init(opts...)
	return srv
}

func (s *Server) init(opts ...ServerOption) {
	for _, o := range opts {
		o(s)
	}
}
func (s *Server) Bind(es EventSever) {

	hashHandlers := map[common.Hash]EventHandler{}
	md := es.Metadata()
	handlers := es.Handlers()
	abi, _ := md.GetAbi()
	for eventName, handle := range handlers {
		if e, ok := abi.Events[eventName]; ok {
			hashHandlers[e.ID] = handle
		}
	}
	watch := NewWatcher(s.baseCtx, s.client, es.Addr(), md, hashHandlers, s.close)
	s.watches = append(s.watches, watch)
}

func (s *Server) Start(ctx context.Context) error {
	if s.err != nil {
		return s.err
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.running {
		return nil
	}
	for _, watch := range s.watches {
		go watch.run()
	}
	s.running = true
	s.log.WithContext(s.baseCtx).Info("server starting")
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.log.WithContext(s.baseCtx).Info("server stopping")
	s.running = false
	s.close <- struct{}{}
	return nil
}
