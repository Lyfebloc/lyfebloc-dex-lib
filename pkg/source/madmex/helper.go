package madmex

import (
	"github.com/Lyfebloc/ethrpc"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

func CallParamsFactory(abi abi.ABI, address string) func(callMethod string, params []interface{}) *ethrpc.Call {
	return func(callMethod string, params []interface{}) *ethrpc.Call {
		return &ethrpc.Call{
			ABI:    abi,
			Target: address,
			Method: callMethod,
			Params: params,
		}
	}
}
