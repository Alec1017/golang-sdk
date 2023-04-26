package client

import (
	"encoding/hex"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/x/dex/types"
	dextypes "github.com/sei-protocol/sei-chain/x/dex/types"
)

func (c *Client) SendRegisterContract(
	contractAddr string,
	codeId uint64,
	needHook bool,
	options ...TxOption,
) (*sdk.TxResponse, error) {
	txBuilder := c.encodingConfig.TxConfig.NewTxBuilder()
	msg := dextypes.MsgRegisterContract{
		Creator: sdk.AccAddress(c.privKey.PubKey().Address()).String(),
		Contract: &dextypes.ContractInfoV2{
			CodeId:            codeId,
			ContractAddr:      contractAddr,
			NeedOrderMatching: true,
			NeedHook:          needHook,
		},
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

func (c *Client) RegisterPairAndWaitForApproval(
	title string,
	contractAddr string,
	pairs []*dextypes.Pair,
) error {
	proposalResp, err := c.RegisterPair(title, contractAddr, pairs)
	if err != nil {
		return err
	}

	proposalId := GetEventAttributeValue(*proposalResp, "submit_proposal", "proposal_id")
	for {
		if c.IsProposalHandled(proposalId) {
			return nil
		}
		time.Sleep(time.Second * VOTE_WAIT_SECONDS)
	}
}

func (c *Client) RegisterPair(
	title string,
	contractAddr string,
	pairs []*dextypes.Pair,
	options ...TxOption,
) (*sdk.TxResponse, error) {
	txBuilder := c.encodingConfig.TxConfig.NewTxBuilder()
	from := sdk.AccAddress(c.privKey.PubKey().Address())

	msg := types.NewMsgRegisterPairs(
		from.String(),
		[]dextypes.BatchContractPair{
			{
				ContractAddr: contractAddr,
				Pairs:        pairs,
			},
		},
	)

	// set the message on the transaction builder
	err := txBuilder.SetMsgs(msg)
	if err != nil {
		panic(err)
	}

	// handle each transaction option passed in
	for _, option := range options {
		option(&txBuilder)
	}

	return c.signAndSendTx(&txBuilder)
}

func (c *Client) SendOrder(
	order FundedOrder,
	contractAddr string,
	options ...TxOption,
) (dextypes.MsgPlaceOrdersResponse, error) {
	seiOrder := ToSeiOrderPlacement(order)
	orderPlacements := []*dextypes.Order{&seiOrder}
	amount, _ := sdk.ParseCoinsNormalized(order.Fund)
	txBuilder := c.encodingConfig.TxConfig.NewTxBuilder()
	msg := dextypes.MsgPlaceOrders{
		Creator:      sdk.AccAddress(c.privKey.PubKey().Address()).String(),
		Orders:       orderPlacements,
		ContractAddr: contractAddr,
		Funds:        amount,
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

	resp, err := c.signAndSendTx(&txBuilder)
	if err != nil {
		return dextypes.MsgPlaceOrdersResponse{}, err
	}

	msgResp := sdk.TxMsgData{}
	respDataBytes, err := hex.DecodeString(resp.Data)
	if err != nil {
		return dextypes.MsgPlaceOrdersResponse{}, err
	}

	if err := msgResp.Unmarshal(respDataBytes); err != nil {
		return dextypes.MsgPlaceOrdersResponse{}, err
	}

	orderPlacementResponse := dextypes.MsgPlaceOrdersResponse{}
	orderMsgData := msgResp.Data[0].Data
	if err := orderPlacementResponse.Unmarshal([]byte(orderMsgData)); err != nil {
		return orderPlacementResponse, err
	}

	return orderPlacementResponse, nil
}

func (c *Client) SendCancel(
	order CancelOrder,
	contractAddr string,
	options ...TxOption,
) (dextypes.MsgCancelOrdersResponse, error) {
	seiCancellation := ToSeiCancelOrderPlacement(order)
	orderCancellations := []*dextypes.Cancellation{&seiCancellation}
	txBuilder := c.encodingConfig.TxConfig.NewTxBuilder()
	msg := dextypes.MsgCancelOrders{
		Creator:       sdk.AccAddress(c.privKey.PubKey().Address()).String(),
		Cancellations: orderCancellations,
		ContractAddr:  contractAddr,
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

	resp, err := c.signAndSendTx(&txBuilder)
	if err != nil {
		return dextypes.MsgCancelOrdersResponse{}, err
	}

	msgResp := sdk.TxMsgData{}
	respDataBytes, err := hex.DecodeString(resp.Data)
	if err != nil {
		return dextypes.MsgCancelOrdersResponse{}, err
	}

	if err := msgResp.Unmarshal(respDataBytes); err != nil {
		return dextypes.MsgCancelOrdersResponse{}, err
	}

	cancelOrderResponse := dextypes.MsgCancelOrdersResponse{}
	cancelOrderMsgData := msgResp.Data[0].Data
	if err := cancelOrderResponse.Unmarshal([]byte(cancelOrderMsgData)); err != nil {
		return cancelOrderResponse, err
	}

	return cancelOrderResponse, nil
}
