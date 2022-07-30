package wallet

import (
	"encoding/json"
	"fmt"
	"github.com/motoko9/aptos-go/utils"
	"io/ioutil"
)

type Wallet struct {
	PrivateKey utils.PrivateKey
}

func New() *Wallet {
	privateKey, err := utils.NewRandomPrivateKey()
	if err != nil {
		panic(fmt.Sprintf("failed to generate private key: %s", err))
	}
	return &Wallet{
		PrivateKey: privateKey,
	}
}

func NewFromKey(key string) (*Wallet, error) {
	k, err := utils.PrivateKeyFromHex(key)
	if err != nil {
		return nil, fmt.Errorf("account from private key: private key from b58: %w", err)
	}
	return &Wallet{
		PrivateKey: k,
	}, nil
}

func NewFromKeygenFile(file string) (*Wallet, error) {
	k, err := utils.PrivateKeyFromKeygenFile(file)
	if err != nil {
		return nil, fmt.Errorf("account from private key: private key from b58: %w", err)
	}
	return &Wallet{
		PrivateKey: k,
	}, nil
}

func (a *Wallet) PublicKey() utils.PublicKey {
	return a.PrivateKey.PublicKey()
}

func (a *Wallet) Save(file string) error {
	keyJson, _ := json.Marshal(a.PrivateKey)
	return ioutil.WriteFile(file, keyJson, 0666)
}

func (a *Wallet) Sign(data []byte) ([]byte, error) {
	return a.PrivateKey.Sign(data)
}

func (a *Wallet) Address() string {
	return utils.PublicKey2Address(a.PublicKey())
}
