package crypto

import (
	"bytes"
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

type PrivateKey ed25519.PrivateKey
type PublicKey ed25519.PublicKey

func NewRandomPrivateKey() (PublicKey, PrivateKey, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return PublicKey(pub), PrivateKey(priv), nil
}

func NewPrivateKeyFromSeed(seed []byte) PrivateKey {
	p := ed25519.NewKeyFromSeed(seed)
	return PrivateKey(p)
}

func NewPrivateKeyFromHexString(key string) (PrivateKey, error) {
	k, err := hex.DecodeString(key)
	if err != nil {
		return nil, fmt.Errorf("account from private key: private key from b58: %w", err)
	}
	return NewPrivateKeyFromSeed(k), nil
}

func NewPrivateKeyFromFile(file string) (PrivateKey, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("read keygen file failed. err = %w", err)
	}
	return NewPrivateKeyFromHexString(string(content))
}

func (p PrivateKey) Seed() []byte {
	seed := make([]byte, ed25519.SeedSize)
	copy(seed, p[:32])
	return seed
}

func (p PrivateKey) HexString() string {
	return hex.EncodeToString(p.Seed())
}

func (p PrivateKey) SaveToFile(file string) error {
	k := p.HexString()
	return ioutil.WriteFile(file, []byte(k), 0666)
}

func (p PrivateKey) Sign(payload []byte) ([]byte, error) {
	return ed25519.PrivateKey(p).Sign(rand.Reader, payload, crypto.Hash(0))
}

func (p PrivateKey) PublicKey() PublicKey {
	publicKey := make([]byte, ed25519.PublicKeySize)
	copy(publicKey, p[32:])
	return publicKey
}

func (p PrivateKey) String() string {
	return p.HexString()
}

func (p PrivateKey) Equal(x PrivateKey) bool {
	return bytes.Equal(p, x)
}

func (pubK PublicKey) String() string {
	return hex.EncodeToString(pubK)
}

func Verify(publicKey PublicKey, message, sig []byte) bool {
	return ed25519.Verify(ed25519.PublicKey(publicKey), message, sig)
}
