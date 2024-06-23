package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"sync"
)

type Server struct {
	client    *ethclient.Client
	contracts []*ContractEvent

	running bool
	lock    sync.RWMutex
}

type ContractEvent struct {
	Addr   common.Address
	Events []string
}

func (s *Server) Bind(c []*ContractEvent) {
	s.contracts = append(s.contracts, c...)
}
