package crypto

import (
    "bytes"
    "crypto"
    "crypto/ed25519"
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "golang.org/x/crypto/sha3"
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

func NewKeyFromSeed(seed []byte) PrivateKey {
    p := ed25519.NewKeyFromSeed(seed)
    return PrivateKey(p)
}

func NewPrivateKeyFromHexString(key string) (PrivateKey, error) {
    k, err := hex.DecodeString(key)
    if err != nil {
        return nil, fmt.Errorf("account from private key: private key from b58: %w", err)
    }
    return k, nil
}

func NewPrivateKeyFromFile(file string) (PrivateKey, error) {
    content, err := ioutil.ReadFile(file)
    if err != nil {
        return nil, fmt.Errorf("read keygen file failed. err = %w", err)
    }
    seed, err := hex.DecodeString(string(content))
    if err != nil {
        return nil, fmt.Errorf("parse keygen file failed. err = %w", err)
    }
    return NewKeyFromSeed(seed), nil
}

func (p PrivateKey) Seed() []byte {
    seed := make([]byte, ed25519.SeedSize)
    copy(seed, p[:32])
    return seed
}

func (p PrivateKey) Save(file string) error {
    k := hex.EncodeToString(p.Seed())
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
    return hex.EncodeToString(p)
}

func (p PrivateKey) Equal(x PrivateKey) bool {
    return bytes.Equal(p, x)
}

func (pubK PublicKey) Address() string {
    hash := sha3.New256()
    hash.Write(pubK[:])
    hash.Write([]byte{0})
    return "0x" + hex.EncodeToString(hash.Sum(nil))
}

func (pubK PublicKey) String() string {
    return hex.EncodeToString(pubK)
}

func Verify(publicKey PublicKey, message, sig []byte) bool {
    return ed25519.Verify(ed25519.PublicKey(publicKey), message, sig)
}
