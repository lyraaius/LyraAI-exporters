package lyraAiX_movement

import (
	"bytes"
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

	"math/rand"
	"time"
)

var movementErr = []byte("ZXJyb3IgY29kZ")

type MovelyraAiX struct {
	moveAddress string
	moduleName  string
	rpcUrl      string
	backupUrl   []string
	client      *client.Client
}

type CheckInEvent struct {
	Version string `json:"version"`
	Guid    struct {
		CreationNumber string `json:"creation_number"`
		AccountAddress string `json:"account_address"`
	} `json:"guid"`
	SequenceNumber string `json:"sequence_number"`
	Type           string `json:"type"`
	Data           struct {
		TaskId    int    `json:"taskId"`
		Timestamp string `json:"timestamp"`
		User      string `json:"user"`
		UserId    string `json:"userId"`
	} `json:"data"`
}

type PredictionEvent struct {
	Version string `json:"version"`
	Guid    struct {
		CreationNumber string `json:"creation_number"`
		AccountAddress string `json:"account_address"`
	} `json:"guid"`
	SequenceNumber string `json:"sequence_number"`
	Type           string `json:"type"`
	Data           struct {
		Choice      int    `json:"choice"`
		HasInvolved bool   `json:"hasInvolved"`
		SignalId    int    `json:"signalId"`
		User        string `json:"user"`
		UserId      string `json:"userId"`
	} `json:"data"`
}

func NewMovelyraAiX(moveAddr, rpcUrl, moduleName string, backupUrl []string) (*MovelyraAiX, error) {
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
		client.WithWriteTimeout(5*time.Second),
		client.WithTLSConfig(clientCfg),
		client.WithDialer(standard.NewDialer()),
	)
	if err != nil {
		return nil, err
	}
	return &MovelyraAiX{
		moveAddress: moveAddr,
		moduleName:  moduleName,
		rpcUrl:      rpcUrl,
		backupUrl:   backupUrl,
		client:      cli,
	}, nil
}

func (m *MovelyraAiX) GetSignalPredictionEvent(ctx context.Context, start, limit int64) ([]*PredictionEvent, error) {
	events, err := m.getSignalPredictionEvent(ctx, m.rpcUrl, start, limit)
	if err != nil {
		for _, url := range m.backupUrl {
			events, err = m.getSignalPredictionEvent(ctx, url, start, limit)
			if err == nil {
				return events, nil
			}
		}
	}
	return events, nil
}

func (m *MovelyraAiX) getSignalPredictionEvent(ctx context.Context, rpcUrl string, start, limit int64) ([]*PredictionEvent, error) {

	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	defer func() {
		protocol.ReleaseRequest(req)
		protocol.ReleaseResponse(resp)
	}()

	eventUrl := fmt.Sprintf("%v/accounts/%v/events/%v::%v::State/signalPredictionEventSet?start=%v&limit=%v",
		rpcUrl, m.moveAddress, m.moveAddress, m.moduleName, start, limit)

	req.SetRequestURI(eventUrl)
	req.SetMethod("GET")
	if rand.Intn(2) > 0 {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")
	} else {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36")
	}

	err := m.client.Do(ctx, req, resp)
	if err != nil {
		hlog.CtxErrorf(ctx, "get signal prediction event error, url: %v, error: %+v", eventUrl, err)
		return nil, err
	}

	if bytes.Contains(resp.Body(), movementErr) {
		hlog.CtxErrorf(ctx, "get signal prediction event error")
		return nil, errors.New("error code")
	}

	var events []*PredictionEvent
	err = json.Unmarshal(resp.Body(), &events)
	if err != nil {
		hlog.CtxErrorf(ctx, "unmarshal signal prediction event error, body: %v, error: %+v", resp.Body(), err)
		return nil, err
	}
	return events, nil
}

func (m *MovelyraAiX) GetCheckInEvent(ctx context.Context, start, limit int64) ([]*CheckInEvent, error) {
	events, err := m.getCheckInEvent(ctx, m.rpcUrl, start, limit)
	if err != nil {
		for _, url := range m.backupUrl {
			events, err = m.getCheckInEvent(ctx, url, start, limit)
			if err == nil {
				return events, nil
			}
		}
	}
	return events, nil
}
func (m *MovelyraAiX) getCheckInEvent(ctx context.Context, rpcUrl string, start, limit int64) ([]*CheckInEvent, error) {
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	defer func() {
		protocol.ReleaseRequest(req)
		protocol.ReleaseResponse(resp)
	}()

	eventUrl := fmt.Sprintf("%v/accounts/%v/events/%v::%v::State/checkInEventSet?start=%v&limit=%v",
		rpcUrl, m.moveAddress, m.moveAddress, m.moduleName, start, limit)

	req.SetRequestURI(eventUrl)
	req.SetMethod("GET")
	if rand.Intn(2) > 0 {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")
	} else {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36")
	}

	err := m.client.Do(ctx, req, resp)
	if err != nil {
		hlog.CtxErrorf(ctx, "get check in event error, url: %v, error: %+v", eventUrl, err)
		return nil, err
	}

	if bytes.Contains(resp.Body(), movementErr) {
		hlog.CtxErrorf(ctx, "get check in event error")
		return nil, nil
	}
	var events []*CheckInEvent
	err = json.Unmarshal(resp.Body(), &events)
	if err != nil {
		hlog.CtxErrorf(ctx, "unmarshal check in event error, body: %v, error: %+v", resp.Body(), err)
		return nil, err
	}
	return events, nil
}

func (m *MovelyraAiX) CheckInResult(ctx context.Context, addr string, currentDay int64) (bool, error) {
	ret, err := m.checkInResult(ctx, m.rpcUrl, addr, currentDay)
	if err != nil {

		for _, url := range m.backupUrl {
			ret, err = m.checkInResult(ctx, url, addr, currentDay)
			if err == nil {
				return ret, nil
			}
		}
	}
	return ret, err
}
func (m *MovelyraAiX) checkInResult(ctx context.Context, rpcUrl string, addr string, currentDay int64) (bool, error) {
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	defer func() {
		protocol.ReleaseRequest(req)
		protocol.ReleaseResponse(resp)
	}()

	eventUrl := fmt.Sprintf("%v/view", rpcUrl)

	// {"function":"0x21737c31a43334567c18a56a4061579b36cc07f5a3a3bb95d8572aa210e64de2::lyraAiXv3::check_in_result",
	// "type_arguments":[],
	// "arguments":["0x21737c31a43334567c18a56a4061579b36cc07f5a3a3bb95d8572aa210e64de2","1234"]}
	// {address}::{module name}::{function name}
	payload := map[string]interface{}{
		"function":       fmt.Sprintf("%v::%v::check_in_result", m.moveAddress, m.moduleName),
		"type_arguments": []interface{}{},
		"arguments": []interface{}{
			addr,
			cast.ToString(currentDay),
		},
	}
	str, err := json.Marshal(payload)
	if err != nil {
		hlog.CtxErrorf(ctx, "marshal check in event error, payload: %v, error: %+v", payload, err)
		return false, err
	}

	req.SetRequestURI(eventUrl)
	req.SetMethod("POST")
	req.SetBody(str)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	if rand.Intn(2) > 0 {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")
	} else {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36")
	}

	err = m.client.Do(ctx, req, resp)
	if err != nil {
		hlog.CtxErrorf(ctx, "get CheckInResult error, body: %v, error: %+v", string(str), err)
		return false, err
	}
	var ret []interface{}
	err = json.Unmarshal(resp.Body(), &ret)
	if err != nil {
		hlog.CtxErrorf(ctx, "unmarshal CheckInResult error, body: %v, error: %+v", resp.Body(), err)
		return false, err
	}
	if len(ret) == 0 {
		return false, nil
	}
	return cast.ToBool(ret[0]), nil
}

func (m *MovelyraAiX) SignalPredictionResult(ctx context.Context, addr string, signalId uint32) (bool, uint32, uint8, error) {
	exist, sigId, choice, err := m.signalPredictionResult(ctx, m.rpcUrl, addr, signalId)
	if err != nil {

		for _, url := range m.backupUrl {
			exist, sigId, choice, err = m.signalPredictionResult(ctx, url, addr, signalId)
			if err == nil {
				return exist, 0, 0, nil
			}
		}
	}
	return exist, sigId, choice, nil
}
func (m *MovelyraAiX) signalPredictionResult(ctx context.Context, rpcUrl string, addr string, signalId uint32) (bool, uint32, uint8, error) {
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	defer func() {
		protocol.ReleaseRequest(req)
		protocol.ReleaseResponse(resp)
	}()

	eventUrl := fmt.Sprintf("%v/view", rpcUrl)

	// {"function":"0x21737c31a43334567c18a56a4061579b36cc07f5a3a3bb95d8572aa210e64de2::lyraAiXv3::signal_predict_result",
	// "type_arguments":[],
	// "arguments":["0x21737c31a43334567c18a56a4061579b36cc07f5a3a3bb95d8572aa210e64de2",2]}

	// {address}::{module name}::{function name}
	payload := map[string]interface{}{
		"function":       fmt.Sprintf("%v::%v::signal_predict_result", m.moveAddress, m.moduleName),
		"type_arguments": []interface{}{},
		"arguments": []interface{}{
			addr,
			signalId,
		},
	}
	str, err := json.Marshal(payload)
	if err != nil {
		hlog.CtxErrorf(ctx, "marshal check in event error, payload: %v, error: %+v", payload, err)
		return false, 0, 0, err
	}

	req.SetRequestURI(eventUrl)
	req.SetMethod("POST")
	req.SetBody(str)
	req.Header.SetContentTypeBytes([]byte("application/json"))

	if rand.Intn(2) > 0 {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")
	} else {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36")
	}

	err = m.client.Do(ctx, req, resp)
	if err != nil {
		hlog.CtxErrorf(ctx, "get CheckInResult error, body: %v, error: %+v", string(str), err)
		return false, 0, 0, err
	}
	var ret []interface{}
	err = json.Unmarshal(resp.Body(), &ret)
	if err != nil {
		hlog.CtxErrorf(ctx, "unmarshal CheckInResult error, body: %v, error: %+v", resp.Body(), err)
		return false, 0, 0, err
	}
	if len(ret) != 3 {
		return false, 0, 0, err
	}
	return cast.ToBool(ret[0]), cast.ToUint32(ret[1]), cast.ToUint8(ret[2]), nil
}
