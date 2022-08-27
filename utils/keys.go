package utils

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type PrivateKey []byte

func PrivateKeyFromKeygenFile(file string) (PrivateKey, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("read keygen file: %w", err)
	}

	var values []byte
	err = json.Unmarshal(content, &values)
	if err != nil {
		return nil, fmt.Errorf("decode keygen file: %w", err)
	}

	return values, nil
}

func PrivateKeyFromHex(key string) (PrivateKey, error) {
	k, err := hex.DecodeString(key)
	if err != nil {
		panic(err)
	}
	return k, nil
}

func NewRandomPrivateKey() (PrivateKey, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	var publicKey PublicKey
	copy(publicKey[:], pub)
	return PrivateKey(priv), nil
}

func (k PrivateKey) String() string {
	return hex.EncodeToString(k)
}

func (k PrivateKey) Sign(payload []byte) ([]byte, error) {
	p := ed25519.PrivateKey(k)
	signData, err := p.Sign(rand.Reader, payload, crypto.Hash(0))
	if err != nil {
		return nil, err
	}
	return signData, err
}

func (k PrivateKey) PublicKey() PublicKey {
	p := ed25519.PrivateKey(k)
	pub := p.Public().(ed25519.PublicKey)

	var publicKey PublicKey
	copy(publicKey[:], pub)

	return publicKey
}

const (
	/// Number of bytes in a pubkey.
	PublicKeyLength = 32
)

type PublicKey [PublicKeyLength]byte

func PublicKeyFromBytes(in []byte) (out PublicKey) {
	byteCount := len(in)
	if byteCount == 0 {
		return
	}

	max := PublicKeyLength
	if byteCount < max {
		max = byteCount
	}

	copy(out[:], in[0:max])
	return
}

func (p PublicKey) Equals(pb PublicKey) bool {
	return p == pb
}

// ToPointer returns a pointer to the pubkey.
func (p PublicKey) ToPointer() *PublicKey {
	return &p
}

func (p PublicKey) Bytes() []byte {
	return p[:]
}

var zeroPublicKey = PublicKey{}

// IsZero returns whether the public key is zero.
// NOTE: the System Program public key is also zero.
func (p PublicKey) IsZero() bool {
	return p == zeroPublicKey
}

func (p PublicKey) String() string {
	return hex.EncodeToString(p[:])
}
