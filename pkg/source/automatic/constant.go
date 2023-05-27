package automatic

import (
	"math/big"
	"time"
)

const (
	DexTypeElastic        = "automatic"
	graphSkipLimit        = 5000
	graphFirstLimit       = 1000
	defaultTokenDecimals  = 18
	defaultTokenWeight    = 50
	reserveZero           = "0"
	graphQLRequestTimeout = 20 * time.Second
)

const (
	methodGetLiquidityState = "getLiquidityState"
	methodGetPoolState      = "getPoolState"
	erc20MethodBalanceOf    = "balanceOf"
)

var (
	ZeroBI = big.NewInt(0)
)
