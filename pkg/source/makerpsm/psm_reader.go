package makerpsm

import (
	"context"

	"github.com/Lyfebloc/ethrpc"
	"github.com/Lyfebloc/logger"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type PSMReader struct {
	abi          abi.ABI
	ethrpcClient *ethrpc.Client
}

func NewPSMReader(ethrpcClient *ethrpc.Client) *PSMReader {
	return &PSMReader{
		abi:          makerPSMPSM,
		ethrpcClient: ethrpcClient,
	}
}

func (r *PSMReader) Read(ctx context.Context, address string) (*PSM, error) {
	var psm PSM

	req := r.ethrpcClient.
		NewRequest().
		SetContext(ctx).
		AddCall(&ethrpc.Call{
			ABI:    r.abi,
			Target: address,
			Method: psmMethodTIn,
			Params: nil,
		}, []interface{}{&psm.TIn}).
		AddCall(&ethrpc.Call{
			ABI:    r.abi,
			Target: address,
			Method: psmMethodTOut,
			Params: nil,
		}, []interface{}{&psm.TOut}).
		AddCall(&ethrpc.Call{
			ABI:    r.abi,
			Target: address,
			Method: psmMethodVat,
			Params: nil,
		}, []interface{}{&psm.VatAddress}).
		AddCall(&ethrpc.Call{
			ABI:    r.abi,
			Target: address,
			Method: psmMethodIlk,
			Params: nil,
		}, []interface{}{&psm.ILK})

	_, err := req.Aggregate()
	if err != nil {
		logger.WithFields(logger.Fields{
			"dexID": DexTypeMakerPSM,
			"error": err,
		}).Error("eth rpc call error")
		return nil, err
	}

	return &psm, nil
}
