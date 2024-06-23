package service

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	contracts "nft-market-backend/abi"
	"nft-market-backend/pkg/eventserver"
)

type ERC20Event struct {
}

func NewERC20Event() *ERC20Event {
	return &ERC20Event{}
}

var _ eventserver.EventSever = (*ERC20Event)(nil)

func (E *ERC20Event) Addr() common.Address {
	return common.HexToAddress("0xf91898Eb6c134E421cdEE59Eb41031FBeD5a5d84")
}

func (E *ERC20Event) Metadata() *bind.MetaData {
	return contracts.OBTokenMetaData
}

func (E *ERC20Event) Handlers() map[string]eventserver.EventHandler {
	return map[string]eventserver.EventHandler{
		"Transfer": E.Print,
	}
}

func (E *ERC20Event) Print(ctx context.Context, log types.Log) {
	fmt.Println(log)
	fmt.Println(log.BlockNumber)
}
