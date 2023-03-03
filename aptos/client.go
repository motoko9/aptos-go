package aptos

import (
	"github.com/motoko9/aptos-go/rpc"
)

type Client struct {
	*rpc.Client
	coinType map[string]*CoinInfo
}

func New(endpoint string, mainNet bool) *Client {
	coinType := make(map[string]*CoinInfo)
	pontemNetworkCoins, _ := readCoinFromPontemNetwork()
	pancakeCoins, _ := readCoinFromPancakeSwap()
	coins := make([]*CoinInfo, 0)
	coins = append(coins, pontemNetworkCoins...)
	coins = append(coins, pancakeCoins...)

	chainId := 1
	if !mainNet {
		chainId = 2
	}
	for _, coin := range coins {
		if coin.ChainId == chainId {
			coinType[coin.T] = coin
		}
	}

	client := rpc.New(endpoint)
	return &Client{
		client,
		coinType,
	}
}
