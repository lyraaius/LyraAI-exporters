package contracts

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/lyraaius/lyraAiX-exporters/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	evmContract  = 0
	moveContract = 1
)

type Contract struct {
	contractMap map[string]*Instance
}

type Instance struct {
	Name         string
	Client       *ethclient.Client
	Address      common.Address
	RpcUrl       string
	BackupRpcUrl []string
	Type         int
	MoveAccount  string
	MoveModule   string
}

func Init(ctx context.Context, conf *conf.Config) (*Contract, error) {
	c := &Contract{contractMap: make(map[string]*Instance)}

	for _, ct := range conf.Contract {
		if ct.Type == evmContract {
			client, err := ethclient.Dial(ct.RpcUrl)
			if err != nil {
				hlog.CtxErrorf(ctx, "Failed to connect to the Ethereum client: %v", err)
				return nil, err
			}
			c.contractMap[ct.Name] = &Instance{
				Name:         ct.Name,
				Client:       client,
				Address:      common.HexToAddress(ct.Address),
				RpcUrl:       ct.RpcUrl,
				Type:         evmContract,
				BackupRpcUrl: ct.BackupRpcUrl,
			}
		}
		if ct.Type == moveContract {
			c.contractMap[ct.Name] = &Instance{
				Name:         ct.Name,
				RpcUrl:       ct.RpcUrl,
				Type:         evmContract,
				MoveAccount:  ct.MoveAccount,
				MoveModule:   ct.MoveModule,
				BackupRpcUrl: ct.BackupRpcUrl,
			}
		}
	}
	return c, nil
}

func (c *Contract) GetContractInstance(name string) *Instance {
	return c.contractMap[name]
}
