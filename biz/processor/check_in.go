package processor

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/lyraaius/lyraAiX-exporters/biz/dal"
	"github.com/lyraaius/lyraAiX-exporters/contracts/lyraAiX"
	"github.com/lyraaius/lyraAiX-exporters/contracts/lyraAiX_mevm"
	"github.com/lyraaius/lyraAiX-exporters/contracts/lyraAiX_movement"
	"github.com/lyraaius/lyraAiX-exporters/pkg/cast"
)

type CheckInProcessor struct {
}

func (p *CheckInProcessor) Process(ctx context.Context, event *lyraAiX.lyraAiXCheckinEvent, chainName string) {

	hlog.CtxInfof(ctx, "[CheckInProcessor] user: %v, userId: %v, taskId: %v, timestamp: %v",
		event.User.Hex(), event.Info.UserId, event.Info.TaskId, event.Info.Timestamp.Int64())

	ret, err := dal.DailySignIn(ctx, int64(event.Info.UserId), int32(event.Info.TaskId), chainName)
	if err != nil {
		hlog.CtxErrorf(ctx, "[CheckInProcessor] DailySignIn error: %v", err)
		return
	}
	fmt.Println(ret, err)
}

func (p *CheckInProcessor) ProcessMoveEvent(ctx context.Context, event *lyraAiX_movement.CheckInEvent, chainName string) {

	hlog.CtxInfof(ctx, "[CheckInProcessor] ProcessMoveEvent user: %v, userId: %v, taskId: %v, timestamp: %v",
		event.Data.User, event.Data.UserId, event.Data.TaskId, event.Data.Timestamp)

	ret, err := dal.DailySignIn(ctx, cast.ToInt64(event.Data.UserId), int32(event.Data.TaskId), chainName)
	if err != nil {
		hlog.CtxErrorf(ctx, "[CheckInProcessor] move DailySignIn error: %v", err)
		return
	}
	fmt.Println(ret, err)
}

func (p *CheckInProcessor) ProcessMEvmEvent(ctx context.Context, event *lyraAiX_mevm.CheckInEvent, chainName string) {

	hlog.CtxInfof(ctx, "[CheckInProcessor] ProcessMEvmEvent user: %v, userId: %v, taskId: %v, timestamp: %v",
		event.Address, event.UserId, event.TaskId, event.Timestamp)

	ret, err := dal.DailySignIn(ctx, cast.ToInt64(event.UserId), int32(event.TaskId), chainName)
	if err != nil {
		hlog.CtxErrorf(ctx, "[CheckInProcessor] move DailySignIn error: %v", err)
		return
	}
	fmt.Println(ret, err)
}
