package rpctmp

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) AccountResourceByAddressAndType(ctx context.Context,
	address string, resourceType string, version uint64) (*rpcmodule.MoveResource, error) {
	params := make(map[string]string)
	if version != 0 {
		params["version"] = fmt.Sprintf("%d", version)
	}
	url := fmt.Sprintf("/accounts/%v/resource/%v", address, resourceType)
	var moveResource rpcmodule.MoveResource
	err := cl.fetchClient.Get(url).SetQueryParams(params).Execute(&moveResource)
	if err != nil {
		return nil, err
	}
	return &moveResource, nil
}