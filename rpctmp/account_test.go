package rpctmp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/motoko9/aptos-go/crypto"
	"github.com/motoko9/aptos-go/rpcmodule"
	"github.com/motoko9/aptos-go/wallet"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_AccountResourceByAddressAndType_Success(t *testing.T) {
	client := New("https://fullnode.devnet.aptoslabs.com/v1")

	// 正常的
	accountResource, err := client.AccountResourceByAddressAndType(
		context.Background(),
		"0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f",
		"0x1::coin::CoinStore<0x1::aptos_coin::AptosCoin>",
		0)
	assert.NoError(t, err)
	fmt.Printf("account resource: \n")
	accountResourceJson, _ := json.MarshalIndent(accountResource, "", "    ")
	fmt.Printf(string(accountResourceJson))
}


func TestClient_AccountResourceByAddressAndType_ResourceNotFound(t *testing.T) {
	client := New("https://fullnode.devnet.aptoslabs.com/v1")

	// new nil address
	pk, _, _ := crypto.NewRandomPrivateKey()
	addr := wallet.PublicKey2Address(pk)

	accountResource, err := client.AccountResourceByAddressAndType(
		context.Background(),
		addr,
		"0x1::coin::CoinStore<0x1::aptos_coin::AptosCoin>",
		0)
	if aptErr, ok := err.(rpcmodule.AptosError); ok {
		// 是不是想要这个
		assert.Equal(t, rpcmodule.ResourceNotFound, aptErr.ErrorCode)
		fmt.Println(aptErr.ErrorCode)
		fmt.Println(aptErr.VmErrorCode)
		fmt.Println(aptErr.Error())
	}
	fmt.Printf("account resource: \n")
	accountResourceJson, _ := json.MarshalIndent(accountResource, "", "    ")
	fmt.Println(string(accountResourceJson))
}