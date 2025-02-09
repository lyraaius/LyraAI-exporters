package lyraAiX

import (
	"context"
	"fmt"
	conf2 "github.com/lyraaius/lyraAiX-exporters/conf"
	"github.com/lyraaius/lyraAiX-exporters/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
	"time"
)

func TestCheckInResult(t *testing.T) {

	conf := &conf2.Config{
		Env:     "",
		Hertz:   nil,
		MySQL:   nil,
		Redis:   nil,
		Metrics: nil,
		Sentry:  nil,
		Contract: []*conf2.Contract{{
			Name:        "bitlayer1",
			Address:     "0xcc146691726aae13b0a35c565a7f00b5ce0f1f16",
			RpcUrl:      "https://rpc.bitlayer.org",
			Type:        0,
			MoveAccount: "",
			MoveModule:  "",
		}, {
			Name:        "bitlayer2",
			Address:     "0x5BBbF0475605D32C1c40c1f08dA4b79e76Ac038f",
			RpcUrl:      "https://mevm.devnet.imola.movementnetwork.xyz",
			Type:        0,
			MoveAccount: "",
			MoveModule:  "",
		}},
	}

	dalContract, err := contracts.Init(context.Background(), conf)

	instance := dalContract.GetContractInstance("bitlayer1")
	alphaContract, _ := NewlyraAiX(instance.Address, instance.Client)

	opts := bind.CallOpts{
		Context: context.Background(),
	}
	currentDay := time.Now().Unix() / 86400
	result, err := alphaContract.CheckInResult(&opts, common.HexToAddress("0x2fd7596c8545443cfafdff33f0b864ee59d180b2"), big.NewInt(currentDay))

	fmt.Println(result, err)

}

func TestBNBTestnetCheckInResult(t *testing.T) {

	conf := &conf2.Config{
		Env:     "",
		Hertz:   nil,
		MySQL:   nil,
		Redis:   nil,
		Metrics: nil,
		Sentry:  nil,
		Contract: []*conf2.Contract{{
			Name:        "BNB Testnet",
			Address:     "0xf0C23A2B852E4845B92d27b872CbC5E84d7380BB",
			RpcUrl:      "https://bsc-testnet-rpc.publicnode.com",
			Type:        0,
			MoveAccount: "",
			MoveModule:  "",
		}},
	}

	dalContract, err := contracts.Init(context.Background(), conf)

	instance := dalContract.GetContractInstance("BNB Testnet")
	alphaContract, _ := NewlyraAiX(instance.Address, instance.Client)

	opts := bind.CallOpts{
		Context: context.Background(),
	}
	currentDay := time.Now().Unix() / 86400
	result, err := alphaContract.CheckInResult(&opts, common.HexToAddress("0x2fd7596c8545443cfafdff33f0b864ee59d180b2"), big.NewInt(currentDay))
	fmt.Println(result, err)

	r1, r2, r3, err := alphaContract.SignalPredictionResult(&opts, common.HexToAddress("0x2fd7596c8545443cfafdff33f0b864ee59d180b2"), 1)
	fmt.Println(r1, r2, r3, err)
}

func TestBNBMainnetCheckInResult(t *testing.T) {

	conf := &conf2.Config{
		Env:     "",
		Hertz:   nil,
		MySQL:   nil,
		Redis:   nil,
		Metrics: nil,
		Sentry:  nil,
		Contract: []*conf2.Contract{{
			Name:        "BNB Testnet",
			Address:     "0xA07F71451eD702669E9e08d97BAd2124777eD612",
			RpcUrl:      "https://bsc-dataseed.bnbchain.org",
			Type:        0,
			MoveAccount: "",
			MoveModule:  "",
		}},
	}

	dalContract, err := contracts.Init(context.Background(), conf)

	instance := dalContract.GetContractInstance("BNB Testnet")
	alphaContract, _ := NewlyraAiX(instance.Address, instance.Client)

	opts := bind.CallOpts{
		Context: context.Background(),
	}
	currentDay := time.Now().Unix() / 86400
	result, err := alphaContract.CheckInResult(&opts, common.HexToAddress("0x8d37fbf1145c0c3898a0078d22c1c34708e02b79"), big.NewInt(currentDay))
	fmt.Println(result, err)

	r1, r2, r3, err := alphaContract.SignalPredictionResult(&opts, common.HexToAddress("0x8d37fbf1145c0c3898a0078d22c1c34708e02b79"), 1)
	fmt.Println(r1, r2, r3, err)

	optsNew := bind.FilterOpts{
		Start: 45148747,
		End:   nil,
	}
	eventIterator, err := alphaContract.FilterCheckinEvent(&optsNew, nil)
	if err != nil {
		panic(err)
	}
	for {
		ret := eventIterator.Next()
		if !ret {
			break
		}
		fmt.Println(eventIterator.Event)
	}

	eventIterator2, err := alphaContract.FilterSignalPredictionEvent(&optsNew, nil)
	if err != nil {
		panic(err)
	}
	for {
		ret := eventIterator2.Next()
		if !ret {
			break
		}
		fmt.Println(eventIterator.Event)

	}

}

func TestKaiaTestnetCheckInResult(t *testing.T) {

	conf := &conf2.Config{
		Env:     "",
		Hertz:   nil,
		MySQL:   nil,
		Redis:   nil,
		Metrics: nil,
		Sentry:  nil,
		Contract: []*conf2.Contract{{
			Name:        "Kaia Kairos Testnet",
			Address:     "0x725da8E34204e4D2f98d9f68aF1Fbf5eA816c303",
			RpcUrl:      "https://rpc.ankr.com/kaia_testnet",
			Type:        0,
			MoveAccount: "",
			MoveModule:  "",
		}},
	}

	dalContract, err := contracts.Init(context.Background(), conf)

	instance := dalContract.GetContractInstance("Kaia Kairos Testnet")
	alphaContract, _ := NewlyraAiX(instance.Address, instance.Client)

	opts := bind.CallOpts{
		Context: context.Background(),
	}
	currentDay := time.Now().Unix() / 86400
	result, err := alphaContract.CheckInResult(&opts, common.HexToAddress("0x690344aa950f60adeaf5e0a5fa3da79cba98443b"), big.NewInt(currentDay))
	fmt.Println(result, err)

	r1, r2, r3, err := alphaContract.SignalPredictionResult(&opts, common.HexToAddress("0x690344aa950f60adeaf5e0a5fa3da79cba98443b"), 1)
	fmt.Println(r1, r2, r3, err)

	optsNew := bind.FilterOpts{
		Start: 0,
		End:   nil,
	}
	eventIterator, err := alphaContract.FilterCheckinEvent(&optsNew, nil)
	if err != nil {
		panic(err)
	}
	for {
		ret := eventIterator.Next()
		if !ret {
			break
		}
		fmt.Println(eventIterator.Event)
	}

	eventIterator2, err := alphaContract.FilterSignalPredictionEvent(&optsNew, nil)
	if err != nil {
		panic(err)
	}
	for {
		ret := eventIterator2.Next()
		if !ret {
			break
		}
		fmt.Println(eventIterator2.Event)

	}
}

func TestKaiaMainNet(t *testing.T) {

	conf := &conf2.Config{
		Env:     "",
		Hertz:   nil,
		MySQL:   nil,
		Redis:   nil,
		Metrics: nil,
		Sentry:  nil,
		Contract: []*conf2.Contract{{
			Name:        "Kaia Mainnet",
			Address:     "0x5C6A28569bB8588E25a67CAEce7914D585C897F9",
			RpcUrl:      "https://rpc.ankr.com/kaia",
			Type:        0,
			MoveAccount: "",
			MoveModule:  "",
		}},
	}

	dalContract, err := contracts.Init(context.Background(), conf)

	instance := dalContract.GetContractInstance("Kaia Mainnet")
	alphaContract, _ := NewlyraAiX(instance.Address, instance.Client)

	opts := bind.CallOpts{
		Context: context.Background(),
	}
	currentDay := time.Now().Unix() / 86400
	result, err := alphaContract.CheckInResult(&opts, common.HexToAddress("0x60DF23ee5312D0A981392bd7E15558Ed26d55Ea6"), big.NewInt(currentDay))
	fmt.Println(result, err)

	r1, r2, r3, err := alphaContract.SignalPredictionResult(&opts, common.HexToAddress("0x60DF23ee5312D0A981392bd7E15558Ed26d55Ea6"), 1)
	fmt.Println(r1, r2, r3, err)

	optsNew := bind.FilterOpts{
		Start: 174996750,
		End:   nil,
	}
	eventIterator, err := alphaContract.FilterCheckinEvent(&optsNew, nil)
	if err != nil {
		panic(err)
	}
	for {
		ret := eventIterator.Next()
		if !ret {
			break
		}
		fmt.Println(eventIterator.Event)
	}

	eventIterator2, err := alphaContract.FilterSignalPredictionEvent(&optsNew, nil)
	if err != nil {
		panic(err)
	}
	for {
		ret := eventIterator2.Next()
		if !ret {
			break
		}
		fmt.Println(eventIterator2.Event)

	}
}
