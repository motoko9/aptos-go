package faucet

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/crypto"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/wallet"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Faucet(t *testing.T) {
	ctx := context.Background()

	//wallet := wallet.New()
	address := "0x30764c46f4a43cda25680bc646f9c2e31bf420f8fe5f3570b389b48f9322ee07"
	//fmt.Printf("private is: %v\n", wallet.Key())
	//fmt.Printf("address is: %v\n", address)

	// fund (max: 20000)
	amount := uint64(20000)
	hashes, err := FundAccount(address, amount)
	if err != nil {
		panic(err)
	}
	fmt.Printf("faucet txs: %v\n", hashes)

	time.Sleep(time.Second * 10)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// latest ledger
	ledger, aptosErr := client.Ledger(ctx)
	assert.Equal(t, nil, aptosErr)

	// query balance of new account
	balance, aptosErr := client.AccountBalance(ctx, address, aptos.AptosCoin, ledger.LedgerVersion)
	assert.Equal(t, nil, aptosErr)
	assert.Equal(t, amount, balance)
	fmt.Printf("account balance is: %d\n", balance)
}

func Test_Faucet_Address(t *testing.T) {
	ctx := context.Background()

	privHexStr := "fc20bed4ec67f04b28f66faafc3e178c6c8936112c0e5f0a9c005fc056cf20fb729c5ad55087d8c9d2280c7d26e888a1ab4b463c56eb3901b5f9b150317cc3ae"
	priv, _ := crypto.NewPrivateKeyFromHexString(privHexStr)

	addr := wallet.PublicKey2Address(priv.PublicKey())

	// fund (max: 20000)
	amount := uint64(20000)
	hashes, err := FundAccount(addr, amount)
	if err != nil {
		panic(err)
	}
	fmt.Printf("faucet txs: %v\n", hashes)

	time.Sleep(time.Second * 10)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// latest ledger
	ledger, Equal := client.Ledger(ctx)
	assert.Equal(t, nil, Equal)

	// query balance of new account
	balance, Equal := client.AccountBalance(ctx, addr, aptos.AptosCoin, ledger.LedgerVersion)
	assert.Equal(t, nil, Equal)
	assert.Equal(t, amount, balance)
	fmt.Printf("account balance is: %d\n", balance)
}
