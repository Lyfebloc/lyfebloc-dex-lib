package oneswap

import (
	"context"
	"encoding/json"
	"math/big"
	"time"

	"github.com/Lyfebloc/ethrpc"
	"github.com/Lyfebloc/logger"

	"github.com/Lyfebloc/lyfebloc-dex-lib/pkg/entity"
)

type PoolTracker struct {
	ethrpcClient *ethrpc.Client
}

func NewPoolTracker(ethrpcClient *ethrpc.Client) *PoolTracker {
	return &PoolTracker{
		ethrpcClient: ethrpcClient,
	}
}

func (d *PoolTracker) GetNewPoolState(ctx context.Context, p entity.Pool) (entity.Pool, error) {
	logger.Infof("[Oneswap] Start getting new state of pool %v", p.Address)

	var (
		swapStorage SwapStorage
		balances    = make([]*big.Int, len(p.Tokens))
	)

	calls := d.ethrpcClient.NewRequest().SetContext(ctx)

	calls.AddCall(&ethrpc.Call{
		ABI:    oneSwapABI,
		Target: p.Address,
		Method: poolMethodSwapStorage,
		Params: nil,
	}, []interface{}{&swapStorage})

	calls.AddCall(&ethrpc.Call{
		ABI:    oneSwapABI,
		Target: p.Address,
		Method: poolMethodGetBalances,
		Params: nil,
	}, []interface{}{&balances})

	if _, err := calls.TryAggregate(); err != nil {
		logger.WithFields(logger.Fields{
			"poolAddress": p.Address,
			"error":       err,
		}).Errorf("failed to aggregate pool data")
		return entity.Pool{}, err
	}

	var extra = Extra{
		InitialA:           swapStorage.InitialA.String(),
		FutureA:            swapStorage.FutureA.String(),
		InitialATime:       swapStorage.InitialATime.Int64(),
		FutureATime:        swapStorage.FutureATime.Int64(),
		SwapFee:            swapStorage.SwapFee.String(),
		AdminFee:           swapStorage.AdminFee.String(),
		DefaultWithdrawFee: swapStorage.DefaultWithdrawFee.String(),
	}
	extraBytes, err := json.Marshal(extra)
	if err != nil {
		logger.WithFields(logger.Fields{
			"poolAddress": p.Address,
			"error":       err,
		}).Errorf("failed to marshal extra data")
		return entity.Pool{}, err
	}

	var reserves = make([]string, len(balances))
	for i := range balances {
		reserves[i] = balances[i].String()
	}

	p.Extra = string(extraBytes)
	p.Reserves = reserves
	p.Timestamp = time.Now().Unix()

	logger.Infof("[Oneswap] Finish getting new state of pool %v", p.Address)

	return p, nil
}
