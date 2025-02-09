package lyraAiX_mevm

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/client/retry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/lyraaius/lyraAiX-exporters/pkg/cast"
	"time"
)

type lyraAiXMEvm struct {
	moveAddress string
	moduleName  string
	rpcUrl      string
	client      *client.Client
}

func NewlyraAiXMEvm(rpcUrl string) (*lyraAiXMEvm, error) {
	clientCfg := &tls.Config{
		InsecureSkipVerify: true,
	}
	cli, err := client.NewClient(
		client.WithDialTimeout(5*time.Second),
		client.WithKeepAlive(true),
		client.WithRetryConfig(
			retry.WithMaxAttemptTimes(3),
			retry.WithInitDelay(1000),
			retry.WithMaxDelay(3000),
			retry.WithDelayPolicy(retry.DefaultDelayPolicy),
			retry.WithMaxJitter(1000),
		),
		client.WithWriteTimeout(10*time.Second),
		client.WithTLSConfig(clientCfg),
		client.WithDialer(standard.NewDialer()),
	)
	if err != nil {
		return nil, err
	}
	return &lyraAiXMEvm{
		rpcUrl: rpcUrl,
		client: cli,
	}, nil
}

type PredictionEvent struct {
	Address  string
	SignalId int64
	UserId   int64
	Choice   int32
}

type CheckInEvent struct {
	Address   string
	TaskId    int64
	UserId    int64
	Timestamp int64
}

func (m *lyraAiXMEvm) GetSignalPredictionEvent(ctx context.Context, lastIndex, length int64) ([]*PredictionEvent, error) {
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	defer func() {
		protocol.ReleaseRequest(req)
		protocol.ReleaseResponse(resp)
	}()

	eventUrl := fmt.Sprintf("%v/api/v1/get_signal_prediction_events?last_index=%v&length=%v",
		m.rpcUrl, lastIndex, length)

	req.SetRequestURI(eventUrl)
	req.SetMethod("GET")

	err := m.client.Do(ctx, req, resp)
	if err != nil {
		hlog.CtxErrorf(ctx, "get signal prediction event error, url: %v, error: %+v", eventUrl, err)
		return nil, err
	}
	type PredictionEventResp struct {
		Data  [][]interface{} `json:"data"`
		Error string          `json:"error"`
	}
	var eventResp PredictionEventResp
	err = json.Unmarshal(resp.Body(), &eventResp)
	if err != nil {
		hlog.CtxErrorf(ctx, "unmarshal signal prediction event error, body: %v, error: %+v", resp.Body(), err)
		return nil, err
	}
	if eventResp.Error != "" {
		return nil, errors.New(eventResp.Error)
	}
	var events []*PredictionEvent
	for _, ev := range eventResp.Data {
		if len(ev) != 5 {
			continue
		}
		events = append(events, &PredictionEvent{
			Address:  cast.ToString(ev[0]),
			SignalId: cast.ToInt64(ev[1]),
			UserId:   cast.ToInt64(ev[2]),
			Choice:   cast.ToInt32(ev[3]),
		})
	}
	return events, nil
}

func (m *lyraAiXMEvm) GetCheckInEvent(ctx context.Context, lastIndex, length int64) ([]*CheckInEvent, error) {
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	defer func() {
		protocol.ReleaseRequest(req)
		protocol.ReleaseResponse(resp)
	}()

	eventUrl := fmt.Sprintf("%v/api/v1/get_check_in_events?last_index=%v&length=%v",
		m.rpcUrl, lastIndex, length)

	req.SetRequestURI(eventUrl)
	req.SetMethod("GET")

	err := m.client.Do(ctx, req, resp)
	if err != nil {
		hlog.CtxErrorf(ctx, "get check in event error, url: %v, error: %+v", eventUrl, err)
		return nil, err
	}
	type CheckInEventResp struct {
		Data  [][]interface{} `json:"data"`
		Error string          `json:"error"`
	}
	var eventResp CheckInEventResp
	err = json.Unmarshal(resp.Body(), &eventResp)
	if err != nil {
		hlog.CtxErrorf(ctx, "unmarshal check in event error, body: %v, error: %+v", resp.Body(), err)
		return nil, err
	}

	if eventResp.Error != "" {
		return nil, errors.New(eventResp.Error)
	}

	var events []*CheckInEvent
	for _, ev := range eventResp.Data {
		if len(ev) != 4 {
			continue
		}
		events = append(events, &CheckInEvent{
			Address:   cast.ToString(ev[0]),
			TaskId:    cast.ToInt64(ev[1]),
			UserId:    cast.ToInt64(ev[2]),
			Timestamp: cast.ToInt64(ev[3]),
		})
	}
	return events, nil
}

func (m *lyraAiXMEvm) CheckInResult(ctx context.Context, addr string, currentDay int64) (bool, error) {
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	defer func() {
		protocol.ReleaseRequest(req)
		protocol.ReleaseResponse(resp)
	}()

	eventUrl := fmt.Sprintf("%v/api/v1/check_in_result?address=%v&current_day=%v", m.rpcUrl, addr, currentDay)

	req.SetRequestURI(eventUrl)
	req.SetMethod("GET")
	req.Header.SetContentTypeBytes([]byte("application/json"))

	err := m.client.Do(ctx, req, resp)
	if err != nil {
		hlog.CtxErrorf(ctx, "get CheckInResult error, body: %v, error: %+v", eventUrl, err)
		return false, err
	}
	var ret map[string]interface{}
	err = json.Unmarshal(resp.Body(), &ret)
	if err != nil {
		hlog.CtxErrorf(ctx, "unmarshal CheckInResult error, body: %v, error: %+v", resp.Body(), err)
		return false, err
	}
	val := ret["data"]
	if val == nil {
		return false, errors.New(cast.ToString(ret["error"]))
	}
	return cast.ToBool(ret["data"]), nil
}

func (m *lyraAiXMEvm) SignalPredictionResult(ctx context.Context, addr string, signalId uint32) (bool, uint32, uint8, error) {
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	defer func() {
		protocol.ReleaseRequest(req)
		protocol.ReleaseResponse(resp)
	}()

	eventUrl := fmt.Sprintf("%v/api/v1/signal_prediction_result?address=%v&signal_id=%v", m.rpcUrl, addr, signalId)

	req.SetRequestURI(eventUrl)
	req.SetMethod("GET")
	req.Header.SetContentTypeBytes([]byte("application/json"))

	err := m.client.Do(ctx, req, resp)
	if err != nil {
		hlog.CtxErrorf(ctx, "get CheckInResult error, body: %v, error: %+v", string(eventUrl), err)
		return false, 0, 0, err
	}
	var ret map[string]interface{}
	err = json.Unmarshal(resp.Body(), &ret)
	if err != nil {
		hlog.CtxErrorf(ctx, "unmarshal CheckInResult error, body: %v, error: %+v", resp.Body(), err)
		return false, 0, 0, err
	}
	val := ret["data"]
	if val != nil {
		valList, ok := val.([]interface{})
		if ok && len(valList) == 3 {
			return cast.ToBool(valList[0]), cast.ToUint32(valList[1]), cast.ToUint8(valList[2]), nil
		}
	}
	return false, 0, 0, errors.New(cast.ToString(ret["error"]))
}

func (m *lyraAiXMEvm) GetTransactionReceiptStatus(ctx context.Context, txHashStr string) (int, error) {
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	defer func() {
		protocol.ReleaseRequest(req)
		protocol.ReleaseResponse(resp)
	}()

	eventUrl := fmt.Sprintf("%v/api/v1/get_transaction_receipt_status?tx_hash=%v", m.rpcUrl, txHashStr)

	req.SetRequestURI(eventUrl)
	req.SetMethod("GET")
	req.Header.SetContentTypeBytes([]byte("application/json"))

	err := m.client.Do(ctx, req, resp)
	if err != nil {
		hlog.CtxErrorf(ctx, "get CheckInResult error, body: %v, error: %+v", string(eventUrl), err)
		return -1, err
	}
	var ret map[string]interface{}
	err = json.Unmarshal(resp.Body(), &ret)
	if err != nil {
		hlog.CtxErrorf(ctx, "unmarshal CheckInResult error, body: %v, error: %+v", resp.Body(), err)
		return -1, err
	}
	val := ret["data"]
	if val != nil {
		return cast.ToInt(val), nil
	}
	return -1, errors.New(cast.ToString(ret["error"]))
}
