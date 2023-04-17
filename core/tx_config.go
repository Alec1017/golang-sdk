package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
)

type TxOption func(t *client.TxBuilder)

func NewTxConfig() client.TxConfig {
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	return tx.NewTxConfig(marshaler, tx.DefaultSignModes)
}

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
