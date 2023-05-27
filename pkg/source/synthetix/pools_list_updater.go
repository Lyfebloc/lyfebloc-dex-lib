package synthetix

import (
	"context"
	"strings"

	"github.com/Lyfebloc/ethrpc"
	"github.com/Lyfebloc/logger"

	"github.com/Lyfebloc/lyfebloc-dex-lib/pkg/entity"
)

type PoolsListUpdater struct {
	cfg             *Config
	poolStateReader IPoolStateReader
	hasInitialized  bool
}

func NewPoolsListUpdater(cfg *Config, ethrpcClient *ethrpc.Client) *PoolsListUpdater {
	return &PoolsListUpdater{
		cfg:             cfg,
		poolStateReader: NewPoolStateReader(cfg, ethrpcClient),
	}
}

func (d *PoolsListUpdater) GetNewPools(ctx context.Context, metadataBytes []byte) ([]entity.Pool, []byte, error) {
	logger.WithFields(logger.Fields{"dexID": d.cfg.DexID}).Info("get new pools")

	if d.hasInitialized {
		return nil, nil, nil
	}

	poolAddress := strings.ToLower(d.cfg.Addresses.Synthetix)

	poolState, err := d.poolStateReader.Read(ctx, poolAddress)
	if err != nil {
		logger.WithFields(logger.Fields{
			"dexID": d.cfg.DexID,
			"error": err,
		}).Error("can not read pool state")
		return nil, nil, err
	}

	var (
		poolTokens = make([]*entity.PoolToken, 0, len(poolState.Synths))
		reserves   = make(entity.PoolReserves, 0, len(poolState.Synths))
	)

	for _, currencyKey := range poolState.CurrencyKeys {
		synthAddress := poolState.Synths[currencyKey]
		poolTokens = append(poolTokens, &entity.PoolToken{
			Address:   strings.ToLower(synthAddress.String()),
			Swappable: true,
		})
		reserves = append(reserves, poolState.SynthsTotalSupply[currencyKey].String())
	}

	pool := entity.Pool{
		Address:  poolAddress,
		Exchange: d.cfg.DexID,
		Type:     DexTypeSynthetix,
		Tokens:   poolTokens,
		Reserves: reserves,
	}

	d.hasInitialized = true

	return []entity.Pool{pool}, nil, nil
}
