package move_example

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
)

func TestMoveRead(t *testing.T) {
	ctx := context.Background()

	// move Module account
	moveModule, err := wallet.NewFromKeygenFile("account_move_publish")
	if err != nil {
		panic(err)
	}
	moduleAddress := moveModule.Address()
	fmt.Printf("move module address: %s\n", moduleAddress)

	// user account
	wallet, err := wallet.NewFromKeygenFile("account_user")
	if err != nil {
		panic(err)
	}
	address := wallet.Address()
	fmt.Printf("user address: %s\n", address)

	// new rpc
	client := rpc.New(rpc.DevNet_RPC)

	// todo,
	// can not ready resource type Message::MessageHolder
	// only support CoinStore type
	// need update AccountResourceByAddressAndType
	//
	resourceType := fmt.Sprintf("%s::Message::MessageHolder", moduleAddress)
	accountResource, err := client.AccountResourceByAddressAndType(ctx, address, resourceType, 0)
	if err != nil {
		panic(err)
	}
	accountResourceJson, _ := json.MarshalIndent(accountResource, "", "    ")
	fmt.Printf("account resource: %s\n", string(accountResourceJson))
}
