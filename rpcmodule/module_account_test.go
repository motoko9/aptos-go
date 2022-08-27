package rpcmodule

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAccountSignature(t *testing.T) {
	jsonText := `{
	  "type": "multi_ed25519_signature",
	  "public_keys": [
		"0x88fbd33f54e1126269769780feb24480428179f552e2313fbe571b72e62a1ca1 "
	  ],
	  "signatures": [
		"0x88fbd33f54e1126269769780feb24480428179f552e2313fbe571b72e62a1ca1 "
	  ],
	  "threshold": 0,
	  "bitmap": "0x88fbd33f54e1126269769780feb24480428179f552e2313fbe571b72e62a1ca1 "
	}`
	var accountSignature AccountSignature
	json.Unmarshal([]byte(jsonText), &accountSignature)
	fmt.Printf("%v\n", accountSignature)

	//
	jsonText1, err := json.Marshal(accountSignature)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(jsonText1))
}
