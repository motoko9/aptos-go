package examples

import (
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/crypto"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
)

func Test_PrintKeys(t *testing.T) {
	printKey(PrivateKey)
	printKey(AlicePrivateKey)
	printKey(BobPrivateKey)
}

func printKey(k crypto.PrivateKey) {
	fmt.Println("#####################")
	fmt.Printf("seed   : %v\n", hex.EncodeToString(k.Seed()))
	fmt.Printf("key    : %v\n", k.PublicKey().String())
	fmt.Printf("address: %v\n", wallet.PublicKey2Address(k.PublicKey()))
}
