package platypus

import (
	"math/big"
	"time"
)

const (
	oracleTypeNone       = "None"
	oracleTypeChainlink  = "Chainlink"
	oracleTypeStakedAvax = "StakedAvax"
)

const (
	poolTypePlatypusBase = "platypus-base"
	poolTypePlatypusAvax = "platypus-avax"
	poolTypePlatypusPure = "platypus-pure"
)

const (
	DexTypePlatypus = "platypus"

	addressZero       = "0x0000000000000000000000000000000000000000"
	addressStakedAvax = "0x2b2c81e08f1af8835a78bb2a90ae924ace0ea4be"

	graphQLRequestTimeout = 20 * time.Second
)

const (
	poolMethodAssetOf           = "assetOf"
	poolMethodGetC1             = "getC1"
	poolMethodGetHaircutRate    = "getHaircutRate"
	poolMethodGetPriceOracle    = "getPriceOracle"
	poolMethodGetRetentionRatio = "getRetentionRatio"
	poolMethodGetSlippageParamK = "getSlippageParamK"
	poolMethodGetSlippageParamN = "getSlippageParamN"
	poolMethodGetTokenAddresses = "getTokenAddresses"
	poolMethodGetXThreshold     = "getXThreshold"
	poolMethodPaused            = "paused"

	assetMethodCash             = "cash"
	assetMethodDecimals         = "decimals"
	assetMethodLiability        = "liability"
	assetMethodUnderlyingToken  = "underlyingToken"
	assetMethodAggregateAccount = "aggregateAccount"

	stakedAvaxMethodGetPooledAvaxByShares = "getPooledAvaxByShares"
)

var bOne = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
