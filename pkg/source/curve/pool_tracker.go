package curve

import (
	"context"
	"errors"

	"github.com/Lyfebloc/ethrpc"
	"github.com/Lyfebloc/logger"

	"github.com/Lyfebloc/lyfebloc-dex-lib/pkg/entity"
)

type PoolTracker struct {
	config       *Config
	ethrpcClient *ethrpc.Client
}

func NewPoolTracker(
	cfg *Config,
	ethrpcClient *ethrpc.Client,
) (*PoolTracker, error) {
	if err := initConfig(cfg, ethrpcClient); err != nil {
		return nil, err
	}

	return &PoolTracker{
		config:       cfg,
		ethrpcClient: ethrpcClient,
	}, nil
}

func (d *PoolTracker) GetNewPoolState(
	ctx context.Context,
	p entity.Pool,
) (entity.Pool, error) {
	switch p.Type {
	case poolTypeBase:
		return d.getNewPoolStateTypeBase(ctx, p)
	case poolTypePlainOracle:
		return d.getNewPoolStateTypePlainOracle(ctx, p)
	case poolTypeMeta:
		return d.getNewPoolStateTypeMeta(ctx, p)
	case poolTypeAave:
		return d.getNewPoolStateTypeAave(ctx, p)
	case poolTypeCompound:
		return d.getNewPoolStateTypeCompound(ctx, p)
	case poolTypeTwo:
		return d.getNewPoolStateTypeTwo(ctx, p)
	case poolTypeTricrypto:
		return d.getNewPoolStateTypeTricrypto(ctx, p)
	default:
		logger.WithFields(logger.Fields{
			"poolAddress": p.Address,
			"poolType":    p.Type,
		}).Errorf("pool type is not implemented")

		return entity.Pool{}, errors.New("pool type is not implemented")
	}
}
