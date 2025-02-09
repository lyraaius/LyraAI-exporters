package lyraAiX_movement

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"testing"
	"time"
)

func TestMovelyraAiX(t *testing.T) {
	m, _ := NewMovelyraAiX("0x21737c31a43334567c18a56a4061579b36cc07f5a3a3bb95d8572aa210e64de2",
		"https://aptos.testnet.porto.movementlabs.xyz/v1",
		"lyraAiXv4", []string{"https://testnet.porto.movementnetwork.xyz/v1"},
	)

	ctx := context.Background()

	addr := "0x21737c31a43334567c18a56a4061579b36cc07f5a3a3bb95d8572aa210e64de2"
	nowDay := time.Now().Unix() / 86400

	ret, err := m.CheckInResult(ctx, addr, nowDay)
	fmt.Println(ret, err)

	signalId := 2

	ret, sId, choice, err := m.SignalPredictionResult(ctx, addr, uint32(signalId))
	fmt.Println(ret, sId, choice, err)

	events, err := m.GetCheckInEvent(ctx, 1279094, 10)
	fmt.Println(sonic.MarshalString(events))

	sevents, err := m.GetSignalPredictionEvent(ctx, 7415158, 100)
	fmt.Println(sonic.MarshalString(sevents))
}
