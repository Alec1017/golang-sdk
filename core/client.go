package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

type Client struct {
	Account        sdk.AccAddress
	privKey        cryptotypes.PrivKey
	clientCtx      client.Context
	encodingConfig EncodingConfig
}

type ClientOption func(c *Client)

func NewClient(
	node string,
	options ...ClientOption,
) *Client {

	// create a default encoding config
	encodingConfig := NewDefaultEncodingConfig()

	// Create a new client
	client := Client{
		clientCtx:      client.Context{},
		encodingConfig: encodingConfig,
	}

	// set up the client with the node URI and some defaults
	client.WithNode(node)
	client.WithBroadcastMode("block")
	client.WithChainID(DEFAULT_CHAIN_ID)
	client.WithInterfaceRegistry(encodingConfig.InterfaceRegistry)

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

func (c *Client) WithPrivateKey(privKey cryptotypes.PrivKey) {
	c.Account = sdk.AccAddress(privKey.PubKey().Address())
	c.privKey = privKey
}

func (c *Client) WithBroadcastMode(mode string) {
	c.clientCtx = c.clientCtx.WithBroadcastMode(mode)
}

func (c *Client) WithInterfaceRegistry(interfaceRegistry codectypes.InterfaceRegistry) {
	c.clientCtx = c.clientCtx.WithInterfaceRegistry(interfaceRegistry)
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
