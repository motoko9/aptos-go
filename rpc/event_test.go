package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_EventsByKey(t *testing.T) {
	client := New(DevNet_RPC)
	events, err := client.EventsByKey(context.Background(), "0x0200000000000000b4dd6392c96cee32802cce841a99c4fd381ff9818d086b80d801657a240ba588")
	if err != nil {
		panic(err)
	}
	eventsJson, _ := json.MarshalIndent(events, "", "    ")
	fmt.Printf("events: %s\n", string(eventsJson))
}
