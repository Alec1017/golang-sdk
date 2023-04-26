package main

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	client "github.com/sei-protocol/golang-sdk/client"
)

func main() {

	// Instantiate sei client
	privKey := secp256k1.GenPrivKey()

	account := sdk.AccAddress(privKey.PubKey().Address())
	_ = account

	nodeURI := "https://rpc.atlantic-2.seinetwork.io/"
	chainID := "atlantic-2"
	broadcastMode := "block"

	// create sei SDK client
	seiClient := client.NewClient(
		nodeURI,
		client.ChainID(chainID),
		client.PrivateKey(privKey),
		client.BroadcastMode(broadcastMode),
	)

	// create new transaction builder
	// txBuilder := clientCtx.TxConfig.NewTxBuilder()

	// create transaction options
	// memo := fmt.Sprint("This is a test memo")
	// gasFee := sdk.NewCoin("usei", sdk.NewInt(100000))
	// gasLimit := uint64(3000000)

	// txBuilder.SetFeeAmount([]sdk.Coin{gasFee})
	// txBuilder.SetGasLimit(gasLimit)

	response, err := seiClient.QueryContract(
		`{"config":{}}`,
		"sei1t6k44ltqmr9alpenr8fu2g6rl8st0z8y9pl4vgg2t5p3dcqyhdyq3ny0qy",
	)
	if err != nil {
		panic(err)
	}

	type QueryResp struct {
		Dao string
	}
	var resp QueryResp
	json.Unmarshal(response.Data.Bytes(), &resp)

	fmt.Println(resp.Dao)
}
