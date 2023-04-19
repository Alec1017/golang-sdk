package client

import (
	"context"

	wasmdtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// This function takes custom instantiateMsg and instantiates cosmwasm contract
// it returns the instantiated contract address if succeeds
func (c *Client) InstantiateContract(
	code uint64,
	label string,
	instantiateMsg string,
	funds sdk.Coins,
	noAdmin bool,
	options ...TxOption,
) (*sdk.TxResponse, error) {

	// Create a new transaction builder
	txBuilder := c.encodingConfig.TxConfig.NewTxBuilder()

	// Get the sender address from the private key
	senderAddr := sdk.AccAddress(c.privKey.PubKey().Address()).String()

	// Construct the contract instantiation message
	msg := wasmdtypes.MsgInstantiateContract{
		Sender: senderAddr,
		CodeID: code,
		Label:  label,
		Msg:    asciiDecodeString(instantiateMsg),
		Funds:  funds,
	}

	// set the sender as an admin address if specified
	if !noAdmin {
		msg.Admin = senderAddr
	}

	// set the message on the transaction builder
	err := txBuilder.SetMsgs(&msg)
	if err != nil {
		panic(err)
	}

	// handle each transaction option passed in
	for _, option := range options {
		option(&txBuilder)
	}

	// Sign and send the transaction
	return c.signAndSendTx(&txBuilder)
}

// This function takes custom executeMsg and call designated cosmwasm contract execute endpoint
// it returns the instantiated contract address if succeeds
// Input fund example: "1000usei". Empty string can be passed if this execution doesn't intend to attach any fund.
func (c *Client) ExecuteContract(
	contractAddr string,
	code uint64,
	executeMsg string,
	funds sdk.Coins,
	options ...TxOption,
) (*sdk.TxResponse, error) {

	// Create a new transaction builder
	txBuilder := c.encodingConfig.TxConfig.NewTxBuilder()

	// Get the sender address from the private key
	senderAddr := sdk.AccAddress(c.privKey.PubKey().Address()).String()

	msg := wasmdtypes.MsgExecuteContract{
		Sender:   senderAddr,
		Contract: contractAddr,
		Msg:      asciiDecodeString(executeMsg),
		Funds:    funds,
	}

	// set the message on the transaction builder
	err := txBuilder.SetMsgs(&msg)
	if err != nil {
		panic(err)
	}

	// handle each transaction option passed in
	for _, option := range options {
		option(&txBuilder)
	}

	return c.signAndSendTx(&txBuilder)
}

// This function takes custom queryMsg and get the corresponding state from the contract
func (c *Client) QueryContract(queryMsg string, contractAddr string) (*wasmdtypes.QuerySmartContractStateResponse, error) {
	client := wasmdtypes.NewQueryClient(c.clientCtx)
	res, err := client.SmartContractState(
		context.Background(),
		&wasmdtypes.QuerySmartContractStateRequest{
			Address:   contractAddr,
			QueryData: asciiDecodeString(queryMsg),
		},
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
