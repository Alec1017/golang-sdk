package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TxOption func(t *client.TxBuilder)

func GasFee(gasFee sdk.Coin) TxOption {
	return func(t *client.TxBuilder) {
		(*t).SetFeeAmount([]sdk.Coin{gasFee})
	}
}

func GasLimit(gasLimit uint64) TxOption {
	return func(t *client.TxBuilder) {
		(*t).SetGasLimit(gasLimit)
	}
}
