package client

import (
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func addGasFee(txBuilder *client.TxBuilder, gasLimit uint64, gasFee sdk.Coin) {
	(*txBuilder).SetGasLimit(gasLimit)
	(*txBuilder).SetFeeAmount([]sdk.Coin{gasFee})
}

func (c *Client) signAndSendTx(txBuilder *client.TxBuilder) (*sdk.TxResponse, error) {
	if err := c.signTx(txBuilder); err != nil {
		return nil, err
	}
	return c.sendTx(txBuilder)
}

func (c *Client) signTx(txBuilder *client.TxBuilder) error {
	var sigsV2 []signing.SignatureV2
	accountNum, seqNum, err := c.getAccountNumberSequenceNumber()
	if err != nil {
		return err
	}

	sigV2 := signing.SignatureV2{
		PubKey: c.privKey.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  c.clientCtx.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: seqNum,
	}
	sigsV2 = append(sigsV2, sigV2)
	if err := (*txBuilder).SetSignatures(sigsV2...); err != nil {
		return err
	}

	sigsV2 = []signing.SignatureV2{}
	signerData := xauthsigning.SignerData{
		ChainID:       c.clientCtx.ChainID,
		AccountNumber: accountNum,
		Sequence:      seqNum,
	}
	sigV2, _ = clienttx.SignWithPrivKey(
		c.clientCtx.TxConfig.SignModeHandler().DefaultMode(),
		signerData,
		*txBuilder,
		c.privKey,
		c.clientCtx.TxConfig,
		seqNum,
	)
	sigsV2 = append(sigsV2, sigV2)
	if err := (*txBuilder).SetSignatures(sigsV2...); err != nil {
		return err
	}

	return nil
}

func (c *Client) sendTx(txBuilder *client.TxBuilder) (*sdk.TxResponse, error) {

	txBytes, err := c.clientCtx.TxConfig.TxEncoder()((*txBuilder).GetTx())
	if err != nil {
		return nil, err
	}

	res, err := c.clientCtx.BroadcastTx(txBytes)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) getAccountNumberSequenceNumber() (uint64, uint64, error) {
	hexAccount := c.privKey.PubKey().Address()
	address, err := sdk.AccAddressFromHex(hexAccount.String())
	if err != nil {
		return 0, 0, err
	}
	accountRetriever := authtypes.AccountRetriever{}

	account, seq, err := accountRetriever.GetAccountNumberSequence(c.clientCtx, address)
	if err != nil {
		time.Sleep(5 * time.Second)
		// retry once after 5 seconds
		account, seq, err = accountRetriever.GetAccountNumberSequence(c.clientCtx, address)
		if err != nil {
			panic(err)
		}
	}
	return account, seq, nil
}
