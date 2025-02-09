package job

import (
	"context"
	"fmt"
	"github.com/bytedance/gopkg/util/gopool"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/lyraaius/lyraAiX-exporters/biz/processor"
	"github.com/lyraaius/lyraAiX-exporters/contracts"
	"github.com/lyraaius/lyraAiX-exporters/contracts/lyraAiX_mevm"
	"github.com/lyraaius/lyraAiX-exporters/pkg/cast"
	"github.com/lyraaius/lyraAiX-exporters/pkg/redis"
	"sync"
	"time"
)

type MEvmEventListenerJob struct {
	checkInStartSequence          int64
	signalPredictionStartSequence int64
	chainName                     string
	lyraAiXContract                *lyraAiX_mevm.lyraAiXMEvm
	workers                       gopool.Pool
}

func NewMEvmEventListenerJob(ci *contracts.Instance) *MEvmEventListenerJob {
	ct, err := lyraAiX_mevm.NewlyraAiXMEvm(ci.RpcUrl)
	if err != nil {
		panic(err)
	}
	return &MEvmEventListenerJob{
		checkInStartSequence:          0,
		signalPredictionStartSequence: 0,
		chainName:                     ci.Name,
		lyraAiXContract:                ct,
		workers:                       gopool.NewPool(fmt.Sprintf("%v-worker", ci.Name), 500, gopool.NewConfig()),
	}
}

func (e *MEvmEventListenerJob) Do(ctx context.Context) {
	hlog.CtxInfof(ctx, "[EventListenerJob] start")
	checkInSeqKey := redis.MoveChainCheckInEventSeqKey(e.chainName)
	if e.checkInStartSequence == 0 {

		ret, err := redis.lyraAiXClient.Get(ctx, checkInSeqKey).Result()
		if err == nil {
			e.checkInStartSequence = cast.ToInt64(ret)
		}
	}

	signalPredSeqKey := redis.MoveChainSignalPredictionEventSeqKey(e.chainName)
	if e.signalPredictionStartSequence == 0 {
		ret, err := redis.lyraAiXClient.Get(ctx, signalPredSeqKey).Result()
		if err == nil {
			e.signalPredictionStartSequence = cast.ToInt64(ret)
		}
	}

	newCheckInStartSequence := int64(0)
	newSignalPredictionStartSequence := int64(0)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	e.workers.CtxGo(ctx, func() {
		defer wg.Done()
		newCheckInStartSequence = e.processCheckinEvent(ctx)
	})
	e.workers.CtxGo(ctx, func() {
		defer wg.Done()
		newSignalPredictionStartSequence = e.processSignalPredictionEvent(ctx)
	})
	wg.Wait()

	e.checkInStartSequence = newCheckInStartSequence
	e.signalPredictionStartSequence = newSignalPredictionStartSequence

	hlog.CtxInfof(ctx, "[MEvmEventListenerJob] chain name: %v, update newCheckInStartSequence: %d", e.chainName, newCheckInStartSequence)
	hlog.CtxInfof(ctx, "[MEvmEventListenerJob] chain name: %v, update newSignalPredictionStartSequence: %d", e.chainName, newSignalPredictionStartSequence)

	_, _ = redis.lyraAiXClient.Set(ctx, checkInSeqKey, newCheckInStartSequence, 0).Result()
	_, _ = redis.lyraAiXClient.Set(ctx, signalPredSeqKey, newSignalPredictionStartSequence, 0).Result()

}

func (e *MEvmEventListenerJob) processCheckinEvent(ctx context.Context) int64 {

	events, err := e.lyraAiXContract.GetCheckInEvent(ctx, e.checkInStartSequence, 100)
	if err != nil {
		hlog.CtxErrorf(ctx, "[MEvmEventListenerJob] chain name: %v, get checkin event error: %v", e.chainName, err)
		return e.checkInStartSequence
	}

	wg := &sync.WaitGroup{}

	size := int64(len(events))

	for _, event := range events {
		wg.Add(1)
		event_ := event
		e.workers.CtxGo(ctx, func() {
			defer wg.Done()
			p := processor.CheckInProcessor{}
			p.ProcessMEvmEvent(ctx, event_, e.chainName)
		})
	}

	wg.Wait()

	if len(events) == 0 {
		return e.checkInStartSequence
	}

	return e.checkInStartSequence + size
}

func (e *MEvmEventListenerJob) processSignalPredictionEvent(ctx context.Context) int64 {

	events, err := e.lyraAiXContract.GetSignalPredictionEvent(ctx, e.signalPredictionStartSequence, 100)
	if err != nil {
		hlog.CtxErrorf(ctx, "[MEvmEventListenerJob] chain name: %v, get signal prediction event error: %v", e.chainName, err)
		return e.signalPredictionStartSequence
	}

	wg := &sync.WaitGroup{}

	for _, event := range events {
		wg.Add(1)
		event_ := event

		e.workers.CtxGo(ctx, func() {
			defer wg.Done()
			p := processor.SignalPredictionProcessor{}
			p.ProcessMEvm(ctx, event_, e.chainName)
		})
	}

	wg.Wait()

	if len(events) == 0 {
		return e.signalPredictionStartSequence
	}

	return e.signalPredictionStartSequence + int64(len(events))
}

func (e *MEvmEventListenerJob) Interval() time.Duration {
	return 200 * time.Millisecond
}

func (e *MEvmEventListenerJob) DisableSerializable() bool {
	return false
}
