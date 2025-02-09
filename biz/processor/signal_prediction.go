package processor

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/lyraaius/lyraAiX-exporters/biz/dal"
	"github.com/lyraaius/lyraAiX-exporters/contracts/lyraAiX"
	"github.com/lyraaius/lyraAiX-exporters/contracts/lyraAiX_mevm"
	"github.com/lyraaius/lyraAiX-exporters/contracts/lyraAiX_movement"
	"github.com/lyraaius/lyraAiX-exporters/pkg/cast"
)

type SignalPredictionProcessor struct {
}

func (p *SignalPredictionProcessor) Process(ctx context.Context, event *lyraAiX.lyraAiXSignalPredictionEvent, chainName string) {

	if event.Info.UserId == 0 || event.Info.SignalId == 0 {
		return
	}

	hlog.CtxInfof(ctx, "[SignalPredictionProcessor] userId: %v, signalId: %v, chainName: %v, choice: %v, txHash: %v, paymentAddr: %v",
		event.Info.UserId, event.Info.SignalId, chainName, event.Info.Choice, event.Raw.TxHash.Hex(), event.User.Hex())

	err := dal.AddUserPredictionLog(ctx, int64(event.Info.UserId),
		int64(event.Info.SignalId), chainName, int32(event.Info.Choice), event.Raw.TxHash.Hex(), event.User.Hex())

	if err != nil {
		hlog.CtxErrorf(ctx, "[SignalPredictionProcessor] chain name: %v,AddUserPredictionLog error: %v", chainName, err)
		return
	}
}

func (p *SignalPredictionProcessor) ProcessMove(ctx context.Context, event *lyraAiX_movement.PredictionEvent, chainName string) {

	if cast.ToInt64(event.Data.UserId) == 0 || event.Data.SignalId == 0 {
		return
	}

	hlog.CtxInfof(ctx, "[SignalPredictionProcessor] ProcessMove userId: %v, signalId: %v, chainName: %v, choice: %v, event type: %v, paymentAddr: %v",
		event.Data.UserId, event.Data.SignalId, chainName, event.Data.Choice, event.Type, event.Guid.AccountAddress)

	err := dal.AddUserPredictionLog(ctx, cast.ToInt64(event.Data.UserId),
		int64(event.Data.SignalId), chainName, int32(event.Data.Choice), event.Data.User, event.Data.User)

	if err != nil {
		hlog.CtxErrorf(ctx, "[SignalPredictionProcessor] ProcessMove chain name: %v, AddUserPredictionLog error: %v", chainName, err)
		return
	}
}

func (p *SignalPredictionProcessor) ProcessMEvm(ctx context.Context, event *lyraAiX_mevm.PredictionEvent, chainName string) {

	if event.UserId == 0 || event.SignalId == 0 {
		return
	}

	hlog.CtxInfof(ctx, "[SignalPredictionProcessor] ProcessMEvm userId: %v, signalId: %v, chainName: %v, choice: %v, event type: %v, paymentAddr: %v",
		event.UserId, event.SignalId, chainName, event.Choice, 0, event.Address)

	err := dal.AddUserPredictionLog(ctx, cast.ToInt64(event.UserId),
		event.SignalId, chainName, event.Choice, event.Address, event.Address)

	if err != nil {
		hlog.CtxErrorf(ctx, "[SignalPredictionProcessor] ProcessMove chain name: %v, AddUserPredictionLog error: %v", chainName, err)
		return
	}
}
