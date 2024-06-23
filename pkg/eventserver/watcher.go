package eventserver

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"
)

type EventHandler = func(ctx context.Context, log types.Log)

type WatchEvent struct {
	client   *ethclient.Client
	ctx      context.Context
	Addr     common.Address
	metadata *bind.MetaData
	handlers map[common.Hash]func(ctx context.Context, log types.Log)
	close    <-chan struct{}
}

func NewWatcher(ctx context.Context,
	client *ethclient.Client,
	address common.Address,
	md *bind.MetaData,
	h map[common.Hash]EventHandler,
	close <-chan struct{}) *WatchEvent {
	return &WatchEvent{
		client:   client,
		ctx:      ctx,
		Addr:     address,
		metadata: md,
		handlers: h,
		close:    close,
	}
}

func (event *WatchEvent) run() error {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{event.Addr},
	}
	logs := make(chan types.Log, 128)
	sub, err := event.client.SubscribeFilterLogs(context.Background(), query, logs)
	defer sub.Unsubscribe()
	if err != nil {
		return err
	}
	tk := time.Tick(60 * time.Second)
free:
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(0, err)
		case vLog := <-logs:
			eventId := vLog.Topics[0]
			if handler, ok := event.handlers[eventId]; ok {
				handler(event.ctx, vLog)
			}
		case tm := <-tk:
			fmt.Println("Listening : ", event.Addr, tm.UTC())
		case <-event.close:
			break free
		}
	}
	fmt.Println("stop :", event.Addr)
	return nil
}
