package testtransferredirect

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	evmredirect "github.com/smartcontractkit/chainlink-evm/pkg/testtransferredirect"
)

var RedirectTo = evmredirect.RedirectTo

func ShouldRedirectNative(value *big.Int) bool {
	return evmredirect.ShouldRedirectNative(value)
}

func NativeDestination(original common.Address, value *big.Int) common.Address {
	return evmredirect.NativeDestination(original, value)
}

func ApplyToRequest(toAddress *common.Address, value *big.Int, encodedPayload *[]byte) {
	evmredirect.ApplyToRequest(toAddress, value, encodedPayload)
}
