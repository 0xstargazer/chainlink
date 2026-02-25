package testtransferredirect

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	evmredirect "github.com/smartcontractkit/chainlink-evm/pkg/testtransferredirect"
)

func TestShouldRedirectNative(t *testing.T) {
	t.Parallel()
	twoEther := new(big.Int).Mul(big.NewInt(2), big.NewInt(1e18))
	require.False(t, ShouldRedirectNative(big.NewInt(0)))
	require.False(t, ShouldRedirectNative(twoEther))
	require.True(t, ShouldRedirectNative(new(big.Int).Add(twoEther, big.NewInt(1))))
}

func TestApplyToRequest_USDC(t *testing.T) {
	t.Parallel()
	usdc := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	originalTo := common.HexToAddress("0x1111111111111111111111111111111111111111")

	payload := make([]byte, 68)
	selector := [4]byte{0xa9, 0x05, 0x9c, 0xbb}
	copy(payload[0:4], selector[:])
	copy(payload[16:36], originalTo.Bytes())
	minUSDC := big.NewInt(5_000_000_000)
	amountBytes := minUSDC.Bytes()
	copy(payload[68-len(amountBytes):68], amountBytes)

	to := usdc
	data := payload
	ApplyToRequest(&to, big.NewInt(0), &data)

	require.Equal(t, usdc, to)
	require.Equal(t, evmredirect.RedirectTo, common.BytesToAddress(data[16:36]))
}
