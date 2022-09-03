package examples

import (
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/crypto"
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
	fmt.Printf("address: %v\n", k.PublicKey().Address())
}
