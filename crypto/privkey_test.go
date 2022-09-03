package crypto

import (
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GenerateRandomKey(t *testing.T) {
	pub, priv, err := NewRandomPrivateKey()
	assert.NoError(t, err)
	privHexStr := hex.EncodeToString(priv)
	pubHexStr := hex.EncodeToString(pub)
	addr := pub.Address()
	fmt.Printf("private key: %v \n", privHexStr)
	fmt.Printf("public  key: %v \n", pubHexStr)
	fmt.Printf("address  is: %v \n", addr)

	k, err := NewPrivateKeyFromHexString(privHexStr)
	assert.NoError(t, err)
	assert.Equal(t, addr, k.PublicKey().Address())
}

func Test_SignAndVerify(t *testing.T) {
	payload, _ := hex.DecodeString("abcdefg")

	pub, priv, err := NewRandomPrivateKey()
	assert.NoError(t, err)

	sig, err := priv.Sign(payload)
	assert.NoError(t, err)

	v := Verify(pub, payload, sig)
	assert.True(t, v, "verify signature failed")
}

func Test_SaveKey(t *testing.T) {
	file := "account_test.key"

	_, priv, err := NewRandomPrivateKey()
	assert.NoError(t, err)
	privHexStr := hex.EncodeToString(priv)
	fmt.Printf("private key: %v \n", privHexStr)

	err = priv.Save(file)
	assert.NoError(t, err)

	np, err := NewPrivateKeyFromFile(file)
	assert.NoError(t, err)
	assert.True(t, priv.Equal(np))
}

func Test_NewFromHexString(t *testing.T) {
	priv, err := NewPrivateKeyFromHexString("fc20bed4ec67f04b28f66faafc3e178c6c8936112c0e5f0a9c005fc056cf20fb729c5ad55087d8c9d2280c7d26e888a1ab4b463c56eb3901b5f9b150317cc3ae")
	assert.NoError(t, err)

	fmt.Println(hex.EncodeToString(priv.Seed()))
	fmt.Println(priv.PublicKey().String())
	fmt.Println(priv.PublicKey().Address())
}
