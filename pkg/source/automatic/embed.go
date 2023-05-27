package automatic

import _ "embed"

//go:embed abis/AutomaticPool.json
var automaticPool.Json []byte

//go:embed abis/ERC20.json
var erc20Json []byte
