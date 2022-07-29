package wallet

import "fmt"

type Wallet struct {
	PrivateKey PrivateKey
}

func New() *Wallet {
	privateKey, err := NewRandomPrivateKey()
	if err != nil {
		panic(fmt.Sprintf("failed to generate private key: %s", err))
	}
	return &Wallet{
		PrivateKey: privateKey,
	}
}

func NewFromKey(key string) (*Wallet, error) {
	k, err := PrivateKeyFromHex(key)
	if err != nil {
		return nil, fmt.Errorf("account from private key: private key from b58: %w", err)
	}
	return &Wallet{
		PrivateKey: k,
	}, nil
}

func NewFromKeygenFile(file string) (*Wallet, error) {
	k, err := PrivateKeyFromKeygenFile(file)
	if err != nil {
		return nil, fmt.Errorf("account from private key: private key from b58: %w", err)
	}
	return &Wallet{
		PrivateKey: k,
	}, nil
}

func (a *Wallet) PublicKey() PublicKey {
	return a.PrivateKey.PublicKey()
}
