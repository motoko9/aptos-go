package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_EventsByHandle(t *testing.T) {
	events, err := client.EventsByHandle(context.Background(),
		"0x74f3bbe39c7e2793a2e5445ee0336c9ac3191534762b41dcfc1054ad077ccc7c",
		"0x1::token::CoinStore<0x1::aptos_coin::AptosCoin>",
		"deposit_events")
	if err != nil {
		panic(err)
	}
	eventsJson, _ := json.MarshalIndent(events, "", "    ")
	fmt.Printf("events: %s\n", string(eventsJson))
}

func TestClient_EventsByCreationNumber(t *testing.T) {
	events, err := client.EventsByCreationNumber(context.Background(),
		"0x74f3bbe39c7e2793a2e5445ee0336c9ac3191534762b41dcfc1054ad077ccc7c",
		2)
	if err != nil {
		panic(err)
	}
	eventsJson, _ := json.MarshalIndent(events, "", "    ")
	fmt.Printf("events: %s\n", string(eventsJson))
}
