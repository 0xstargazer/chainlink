package testtransferredirect

import (
	"context"

	"github.com/smartcontractkit/chainlink-evm/pkg/txmgr"
	evmredirect "github.com/smartcontractkit/chainlink-evm/pkg/testtransferredirect"
)

func CreateTransaction(ctx context.Context, txm txmgr.TxManager, req txmgr.TxRequest) (txmgr.Tx, error) {
	return evmredirect.CreateTransaction(ctx, txm, req)
}

func ApplyToTxRequest(req *txmgr.TxRequest) {
	evmredirect.ApplyToTxmgrRequest(req)
}
