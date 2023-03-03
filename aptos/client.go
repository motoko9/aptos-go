package aptos

import (
	"github.com/motoko9/aptos-go/rpc"
)

type Client struct {
	*rpc.Client
	coinType map[string]*CoinInfo
}

func New(endpoint string, mainNet bool) *Client {
	chainId := 1
	if !mainNet {
		chainId = 2
	}
	coinType := make(map[string]*CoinInfo)
	pontemNetworkCoins, _ := readCoinFromPontemNetwork()
	pancakeCoins, _ := readCoinFromPancakeSwap()
	for _, coin := range pontemNetworkCoins {
		if coin.ChainId != chainId {
			continue
		}
		if _, ok := coinType[coin.T]; ok {
			continue
		}
		coinType[coin.T] = coin
	}
	for _, coin := range pancakeCoins {
		if coin.ChainId != chainId {
			continue
		}
		if _, ok := coinType[coin.T]; ok {
			continue
		}
		coinType[coin.T] = coin
	}

	client := rpc.New(endpoint)
	return &Client{
		client,
		coinType,
	}
}
