package polydex

import (
	"context"

	"github.com/Lyfebloc/ethrpc"
	"github.com/Lyfebloc/logger"

	"github.com/Lyfebloc/lyfebloc-dex-lib/pkg/entity"
)

type PoolTracker struct {
	ethrpcClient *ethrpc.Client
}

func NewPoolTracker(
	ethrpcClient *ethrpc.Client,
) (*PoolTracker, error) {
	return &PoolTracker{
		ethrpcClient: ethrpcClient,
	}, nil
}

func (d *PoolTracker) GetNewPoolState(ctx context.Context, p entity.Pool) (entity.Pool, error) {
	log := logger.WithFields(logger.Fields{
		"liquiditySource": DexTypePolydex,
		"poolAddress":     p.Address,
	})
	log.Infof("Start getting new state of pool: %v", p.Address)

	rpcRequest := d.ethrpcClient.NewRequest().SetContext(ctx)

	var reserves Reserves
	var swapFee uint32

	rpcRequest.AddCall(&ethrpc.Call{
		ABI:    pairABI,
		Target: p.Address,
		Method: pairMethodGetReserves,
		Params: nil,
	}, []interface{}{&reserves})

	rpcRequest.AddCall(&ethrpc.Call{
		ABI:    pairABI,
		Target: p.Address,
		Method: pairMethodGetSwapFee,
		Params: nil,
	}, []interface{}{&swapFee})

	_, err := rpcRequest.TryAggregate()
	if err != nil {
		log.Errorf("failed to process tryAggregate for pool: %v, err: %v", p.Address, err)
		return entity.Pool{}, err
	}

	p.Timestamp = int64(reserves.BlockTimestampLast)
	p.Reserves = entity.PoolReserves{
		reserves.Reserve0.String(),
		reserves.Reserve1.String(),
	}
	p.SwapFee = float64(swapFee) / bps

	log.Infof("Finish getting new state of pool: %v", p.Address)

	return p, nil
}
