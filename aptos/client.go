package aptos

import (
	"encoding/json"
	"github.com/motoko9/aptos-go/rpc"
	"io/ioutil"
	"net/http"
)

type Client struct {
	*rpc.Client
	coinType map[string]*CoinInfo
}

func New(endpoint string, mainNet bool) *Client {
	coinType := make(map[string]*CoinInfo)
	if coins, err := readCoinFile(); err == nil {
		chainId := 1
		if !mainNet {
			chainId = 2
		}
		for _, coin := range coins {
			if coin.ChainId == chainId {
				coinType[coin.T] = coin
			}
		}
	}
	client := rpc.New(endpoint)
	return &Client{
		client,
		coinType,
	}
}

func readCoinFile() ([]*CoinInfo, error) {
	url := "https://raw.githubusercontent.com/pontem-network/coins-registry/main/src/coins.json"
	rsp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	//
	coins := make([]*CoinInfo, 0)
	err = json.Unmarshal(data, &coins)
	if err != nil {
		return nil, err
	}
	return coins, nil
}
