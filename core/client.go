package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

type Client struct {
	privKey   cryptotypes.PrivKey
	clientCtx client.Context
}

type ClientOption func(c *Client)

func NewClient(
	node string,
	options ...ClientOption,
) *Client {

	// create a new client context
	clientCtx := client.Context{
		TxConfig: NewTxConfig(),
	}

	// Create a new client
	client := Client{
		clientCtx: clientCtx,
	}

	// set up the client with the node URI and some defaults
	client.WithNode(node)
	client.WithBroadcastMode("block")
	client.WithChainID(DEFAULT_CHAIN_ID)

	// handle each option passed in
	for _, option := range options {
		option(&client)
	}

	return &client
}

func (c *Client) WithNode(nodeURI string) {

	nodeClient, err := rpchttp.New(nodeURI)
	if err != nil {
		panic(err)
	}

	c.clientCtx = c.clientCtx.
		WithNodeURI(nodeURI).
		WithClient(nodeClient)
}

func (c *Client) WithChainID(chainID string) {
	c.clientCtx = c.clientCtx.WithChainID(chainID)
}

func (c *Client) WithPrivateKey(key cryptotypes.PrivKey) {
	c.privKey = key
}

func (c *Client) WithBroadcastMode(mode string) {
	c.clientCtx = c.clientCtx.WithBroadcastMode(mode)
}

// CLIENT INSTANTIATION OPTIONS

func ChainID(chainID string) ClientOption {
	return func(c *Client) {
		c.WithChainID(chainID)
	}
}

func PrivateKey(key cryptotypes.PrivKey) ClientOption {
	return func(c *Client) {
		c.WithPrivateKey(key)
	}
}

func BroadcastMode(mode string) ClientOption {
	return func(c *Client) {
		c.WithBroadcastMode(mode)
	}
}
