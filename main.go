package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/lyraaius/lyraAiX-exporters/biz/job"
	"github.com/lyraaius/lyraAiX-exporters/conf"
	"github.com/lyraaius/lyraAiX-exporters/contracts"
	"github.com/lyraaius/lyraAiX-exporters/pkg/cron_job"
	"github.com/lyraaius/lyraAiX-exporters/pkg/kafka"
	"github.com/lyraaius/lyraAiX-exporters/pkg/redis"
	hertzZerolog "github.com/hertz-contrib/logger/zerolog"
	"github.com/hertz-contrib/pprof"
	"os"
	"time"
)

func init() {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println("Failed to set UTC+0 timezone:", err)
		panic(err)
	}
	time.Local = loc
}

func main() {
	// init cancel context
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	srvConf := conf.GetConf("config")

	// init log
	initLogger(srvConf)

	h := server.New(server.WithHostPorts(srvConf.Hertz.Address))
	pprof.Register(h)

	// init redis
	redis.Init(srvConf)

	// init kafka
	kafka.Init(srvConf)

	// init contract client
	contract, err := contracts.Init(ctx, srvConf)
	if err != nil {
		panic(err)
	}

	if os.Getenv("MEVM_ONLY") == "1" {
		// Movement EVM Testnet
		moveEVMChainEventListenerJob := cron_job.NewCronJob(job.NewMEvmEventListenerJob(contract.GetContractInstance("Movement EVM Testnet")))
		go moveEVMChainEventListenerJob.Run(ctx)

	} else if os.Getenv("MOVEMENT_ONLY") == "1" {
		// Movement Aptos Testnet
		moveChainEventListenerJob := cron_job.NewCronJob(job.NewMovementEventListenerJob(contract.GetContractInstance("Movement Aptos Testnet")))
		go moveChainEventListenerJob.Run(ctx)

	} else if conf.GetEnv() == "prod" {
		// Bitlayer Mainnet
		chainEventListenerJob := cron_job.NewCronJob(job.NewEventListenerJob(contract.GetContractInstance("Bitlayer Mainnet")))
		go chainEventListenerJob.Run(ctx)

		// BSC main net
		BSCChainEventListenerJob := cron_job.NewCronJob(job.NewBSCEventListenerJob(contract.GetContractInstance("Binance Smart Chain")))
		go BSCChainEventListenerJob.Run(ctx)

		// kaia main net
		KaiaMainNetChainEventListenerJob := cron_job.NewCronJob(job.NewKaiaEventListenerJob(contract.GetContractInstance("Kaia Mainnet")))
		go KaiaMainNetChainEventListenerJob.Run(ctx)

	} else {
		// bitlayer Testnet
		chainEventListenerJob := cron_job.NewCronJob(job.NewEventListenerJob(contract.GetContractInstance("Bitlayer Testnet")))
		go chainEventListenerJob.Run(ctx)

		// BSC Testnet
		BSCTestNetChainEventListenerJob := cron_job.NewCronJob(job.NewBSCEventListenerJob(contract.GetContractInstance("Binance Smart Chain Testnet")))
		go BSCTestNetChainEventListenerJob.Run(ctx)

		// kaia testnet
		KaiaTestNetChainEventListenerJob := cron_job.NewCronJob(job.NewKaiaEventListenerJob(contract.GetContractInstance("Kaia Kairos Testnet")))
		go KaiaTestNetChainEventListenerJob.Run(ctx)
	}

	h.Spin()
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func initLogger(config *conf.Config) {
	hlog.SetLogger(hertzZerolog.New(
		hertzZerolog.WithOutput(os.Stdout),     // allows to specify output
		hertzZerolog.WithLevel(hlog.LevelInfo), // option with log level
		hertzZerolog.WithTimestamp(),           // option with timestamp
		hertzZerolog.WithCaller()))             // option with caller
}
