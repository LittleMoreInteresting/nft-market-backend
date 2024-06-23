package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/a93BJAOxsYWxlpZJAGq_nPYPUWDpJSsq")
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress("0x43CBe3CB4f095D1d7230e427aA4582638B2BF1A8")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log, 128)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		fmt.Println(err)
		log.Fatal(9, err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(0, err)
		case vLog := <-logs:
			eventId := vLog.Topics[0]
			fmt.Println(1, vLog, eventId) // pointer to event log
		}
	}
}
