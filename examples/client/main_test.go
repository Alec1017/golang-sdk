package clientexample

import (
	"testing"
)

func TestClient(t *testing.T) {
	// establish grpc connection

	// 	nodeURL := "tcp://rpc.atlantic-2.seinetwork.io"
	// 	rpcClient, err := rpchttp.New(nodeURL, "/websocket")
	// 	if err != nil {
	// 		// Handle error
	// 		panic(err)
	// 	}

	// 	txConfig := seiSdk.NewTxConfig2(
	// 		"tcp://rpc.atlantic-2.seinetwork.io",
	// 		"http://localhost:8088",
	// 		"atlantic-2",
	// 		2000000,
	// 		sdk.NewCoin("usei", sdk.NewInt(100000)),
	// 	)

	// 	client := seiSdk.NewClient(
	// 		secp256k1.GenPrivKey(),
	// 		txConfig,
	// 		seiSdk.NewDefaultEncodingConfig(),
	// 	)

	// 	fmt.Println(client)

	// 	response, err := client.QueryContract(
	// 		`{"config":{}}`,
	// 		"sei1t6k44ltqmr9alpenr8fu2g6rl8st0z8y9pl4vgg2t5p3dcqyhdyq3ny0qy",
	// 	)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	_ = response

	// fmt.Println(response)
}
