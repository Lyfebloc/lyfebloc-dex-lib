package dystopia

import (
	"context"

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
	log := logger.WithFields(logger.Fields{
		"poolAddress": p.Address,
	})
	log.Infof("[Dystopia] Start updating state ...")

	var reserve Reserves

	calls := d.ethrpcClient.NewRequest().SetContext(ctx)

	calls.AddCall(&ethrpc.Call{
		ABI:    pairABI,
		Target: p.Address,
		Method: poolMethodGetReserves,
		Params: nil,
	}, []interface{}{&reserve})

	if _, err := calls.TryAggregate(); err != nil {
		log.WithFields(logger.Fields{
			"error": err,
		}).Errorf("[Dystopia] failed to aggregate to get pool data")

		return entity.Pool{}, err
	}

	p.Reserves = entity.PoolReserves{reserve.Reserve0.String(), reserve.Reserve1.String()}
	p.Timestamp = int64(reserve.BlockTimestampLast)

	log.Infof("[Dystopia] Finish getting new state of pool")

	return p, nil
}
