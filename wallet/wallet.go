package wallet

import (
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/crypto"
	"golang.org/x/crypto/sha3"
)

type Wallet struct {
	PrivateKey crypto.PrivateKey
}

func New() *Wallet {
	_, privateKey, err := crypto.NewRandomPrivateKey()
	if err != nil {
		panic(fmt.Sprintf("failed to generate private key: %s", err))
	}
	return &Wallet{
		PrivateKey: privateKey,
	}
}

func NewFromKey(key string) (*Wallet, error) {
	k, err := crypto.NewPrivateKeyFromHexString(key)
	if err != nil {
		return nil, fmt.Errorf("account from private key: private key from b58: %w", err)
	}
	return &Wallet{
		PrivateKey: k,
	}, nil
}

func NewFromKeygenFile(file string) (*Wallet, error) {
	k, err := crypto.NewPrivateKeyFromFile(file)
	if err != nil {
		return nil, fmt.Errorf("account from private key: private key from b58: %w", err)
	}
	return &Wallet{
		PrivateKey: k,
	}, nil
}

func (a *Wallet) SaveToKeygenFile(file string) error {
	return a.PrivateKey.SaveToFile(file)
}

func (a *Wallet) Key() string {
	return a.PrivateKey.HexString()
}

func (a *Wallet) PublicKey() crypto.PublicKey {
	return a.PrivateKey.PublicKey()
}

func (a *Wallet) Sign(data []byte) ([]byte, error) {
	return a.PrivateKey.Sign(data)
}

func (a *Wallet) Address() string {
	return PublicKey2Address(a.PublicKey())
}

func PublicKey2Address(pk crypto.PublicKey) string {
	hash := sha3.New256()
	hash.Write(pk[:])
	hash.Write([]byte{0})
	return "0x" + hex.EncodeToString(hash.Sum(nil))
}
