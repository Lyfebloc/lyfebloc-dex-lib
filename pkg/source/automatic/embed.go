package automatic

import _ "embed"

//go:embed abis/AutomaticPool.json
var AutomaticPoolJson []byte

//go:embed abis/ERC20.json
var erc20Json []byte
