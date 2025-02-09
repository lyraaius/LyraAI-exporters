package job

import (
	"context"
	"fmt"
	"github.com/bytedance/gopkg/util/gopool"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/lyraaius/lyraAiX-exporters/biz/processor"
	"github.com/lyraaius/lyraAiX-exporters/contracts"
	"github.com/lyraaius/lyraAiX-exporters/contracts/lyraAiX"
	"github.com/lyraaius/lyraAiX-exporters/pkg/cast"
	"github.com/lyraaius/lyraAiX-exporters/pkg/redis"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"sync"
	"time"
)

type KaiaEventListenerJob struct {
	startBlock        uint64
	endBlock          uint64
	checkInStartBlock uint64
	predictStartBlock uint64
	contractInstance  *contracts.Instance
	lyraAiXContract    *lyraAiX.lyraAiX
	workers           gopool.Pool
}

func NewKaiaEventListenerJob(ci *contracts.Instance) *KaiaEventListenerJob {
	ct, err := lyraAiX.NewlyraAiX(ci.Address, ci.Client)
	if err != nil {
		panic(err)
	}
	return &KaiaEventListenerJob{
		startBlock:       0,
		endBlock:         0,
		contractInstance: ci,
		lyraAiXContract:   ct,
		workers:          gopool.NewPool(fmt.Sprintf("%v-worker", ci.Name), 100, gopool.NewConfig()),
	}
}

func (e *KaiaEventListenerJob) Do(ctx context.Context) {
	hlog.CtxInfof(ctx, "[KaiaEventListenerJob] start")
	var err error
	lastListenOnBlockNumber := redis.ChainEventListenKey(e.contractInstance.Name)
	if e.startBlock == 0 {
		ret, err := redis.lyraAiXClient.Get(ctx, lastListenOnBlockNumber).Result()
		if err == nil {
			e.startBlock = cast.ToUint64(ret)
		}
	}
	chainName := e.contractInstance.Name
	e.endBlock, err = e.contractInstance.Client.BlockNumber(ctx)
	if err != nil {
		hlog.CtxErrorf(ctx, "[KaiaEventListenerJob] chain name: %v, get block number failed, err: %v", chainName, err)
	}

	if e.startBlock >= e.endBlock {
		return
	}
	hlog.CtxInfof(ctx, "[KaiaEventListenerJob] chain name: %v, start listen on block: %v, end block: %v", chainName, e.startBlock, e.endBlock)

	var end *uint64
	if e.endBlock > 0 {
		end = &e.endBlock
	}
	opts := bind.FilterOpts{
		Start:   e.startBlock,
		End:     end,
		Context: ctx,
	}

	e.checkInStartBlock = 0
	e.predictStartBlock = 0

	wg := &sync.WaitGroup{}
	wg.Add(2)
	e.workers.CtxGo(ctx, func() {
		defer wg.Done()
		e.processCheckinEvent(ctx, &opts)
	})
	e.workers.CtxGo(ctx, func() {
		defer wg.Done()
		e.processSignalPredictionEvent(ctx, &opts)
	})
	wg.Wait()

	if e.checkInStartBlock == 0 && e.predictStartBlock == 0 {
		return
	}

	e.startBlock = max(e.checkInStartBlock, e.predictStartBlock) + 1
	hlog.CtxInfof(ctx, "[KaiaEventListenerJob] chain name: %v, update last listen on block number: %d", chainName, e.startBlock)
	_, _ = redis.lyraAiXClient.Set(ctx, lastListenOnBlockNumber, e.startBlock, 0).Result()

}

func (e *KaiaEventListenerJob) processCheckinEvent(ctx context.Context, opts *bind.FilterOpts) {
	chainName := e.contractInstance.Name

	eventIterator, err := e.lyraAiXContract.FilterCheckinEvent(opts, nil)
	if err != nil {
		hlog.CtxErrorf(ctx, "[KaiaEventListenerJob] chain name: %v, get checkin event error: %v", chainName, err)
		return
	}

	wg := &sync.WaitGroup{}

	for {
		ret := eventIterator.Next()
		if !ret {
			break
		}
		wg.Add(1)

		event_ := eventIterator.Event
		if event_.Raw.BlockNumber > e.checkInStartBlock {
			e.checkInStartBlock = event_.Raw.BlockNumber
		}
		e.workers.CtxGo(ctx, func() {
			defer wg.Done()
			p := processor.CheckInProcessor{}
			p.Process(ctx, event_, chainName)
		})
	}
	wg.Wait()
}

func (e *KaiaEventListenerJob) processSignalPredictionEvent(ctx context.Context, opts *bind.FilterOpts) {
	chainName := e.contractInstance.Name

	eventIterator, err := e.lyraAiXContract.FilterSignalPredictionEvent(opts, nil)
	if err != nil {
		hlog.CtxErrorf(ctx, "[KaiaEventListenerJob] chain name: %v, get signal prediction event error: %v", chainName, err)
		return
	}

	wg := &sync.WaitGroup{}

	for {
		ret := eventIterator.Next()
		if !ret {
			break
		}
		wg.Add(1)

		event_ := eventIterator.Event
		if event_.Raw.BlockNumber > e.predictStartBlock {
			e.predictStartBlock = event_.Raw.BlockNumber
		}
		e.workers.CtxGo(ctx, func() {
			defer wg.Done()
			p := processor.SignalPredictionProcessor{}
			p.Process(ctx, event_, chainName)
		})
	}
	wg.Wait()

}

func (e *KaiaEventListenerJob) Interval() time.Duration {
	return 1 * time.Second
}

func (e *KaiaEventListenerJob) DisableSerializable() bool {
	return false
}
