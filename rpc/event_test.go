package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_EventsByKey(t *testing.T) {
	client := New(DevNet_RPC)
	events, err := client.EventsByKey(context.Background(), "0x0200000000000000697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f")
	if err != nil {
		panic(err)
	}
	eventsJson, _ := json.MarshalIndent(events, "", "    ")
	fmt.Printf("events: %s\n", string(eventsJson))
}

func TestClient_EventsByHandle(t *testing.T) {
	client := New(DevNet_RPC)
	events, err := client.EventsByHandle(context.Background(),
		"0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f",
		"0x1::coin::CoinStore<0x1::aptos_coin::AptosCoin>",
		"deposit_events")
	if err != nil {
		panic(err)
	}
	eventsJson, _ := json.MarshalIndent(events, "", "    ")
	fmt.Printf("events: %s\n", string(eventsJson))
}
