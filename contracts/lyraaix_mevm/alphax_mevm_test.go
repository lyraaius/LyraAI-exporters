package lyraAiX_mevm

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"testing"
	"time"
)

func TestMovementEVM(t *testing.T) {
	m, _ := NewlyraAiXMEvm("")

	ctx := context.Background()

	addr := "0xe33d843df52bc45d9dc1941fcd1c78f134f114de"
	nowDay := time.Now().Unix() / 86400

	ret, err := m.CheckInResult(ctx, addr, nowDay)
	fmt.Println(ret, err)

	signalId := 2
	ret, sId, choice, err := m.SignalPredictionResult(ctx, addr, uint32(signalId))
	fmt.Println(ret, sId, choice, err)

	ret, err = m.CheckInResult(ctx, addr+"0", nowDay)
	fmt.Println(ret, err)

	ret, sId, choice, err = m.SignalPredictionResult(ctx, addr+"0", uint32(signalId))
	fmt.Println(ret, sId, choice, err)

	events, err := m.GetCheckInEvent(ctx, 0, 10)
	fmt.Println(sonic.MarshalString(events))

	sevents, err := m.GetSignalPredictionEvent(ctx, 0, 10)
	fmt.Println(sonic.MarshalString(sevents))

	events, err = m.GetCheckInEvent(ctx, 10, 10)
	fmt.Println(sonic.MarshalString(events))

	sevents, err = m.GetSignalPredictionEvent(ctx, 10, 10)
	fmt.Println(sonic.MarshalString(sevents))
}
