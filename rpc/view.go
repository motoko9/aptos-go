package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) View(ctx context.Context, view *rpcmodule.ViewRequest) (string, *rpcmodule.AptosError) {
	url := fmt.Sprintf("/view")
	var raw json.RawMessage
	var aptosError rpcmodule.AptosError
	cl.fetchClient.Post(url).SetJSONBody(view).Execute(&raw, &cl.rsp, &aptosError)
	if aptosError.IsError() {
		return "", &aptosError
	}
	return string(raw), nil
}
