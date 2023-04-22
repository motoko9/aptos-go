package rpc

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/rpcmodule"
	"golang.org/x/crypto/sha3"
	"strings"
)

func (cl *Client) Account(ctx context.Context, address string, version uint64) (*rpcmodule.AccountData, *rpcmodule.AptosError) {
	params := make(map[string]string)
	if version != 0 {
		params["version"] = fmt.Sprintf("%d", version)
	}
	url := fmt.Sprintf("/accounts/%s", address)
	var account rpcmodule.AccountData
	var aptosError rpcmodule.AptosError
	cl.fetchClient.Get(url).SetQueryParams(params).Execute(&account, &cl.rsp, &aptosError)
	if aptosError.IsError() {
		return nil, &aptosError
	}
	return &account, nil
}

func (cl *Client) AccountResources(ctx context.Context, address string, cursor string, limit uint64, version uint64) (*rpcmodule.MoveResources, *rpcmodule.AptosError) {
	params := make(map[string]string)
	if version != 0 {
		params["version"] = fmt.Sprintf("%d", version)
	}
	if limit != 0 {
		params["limit"] = fmt.Sprintf("%d", limit)
	}
	if cursor != "" {
		params["start"] = cursor
	}
	url := fmt.Sprintf("/accounts/%s/resources", address)
	var moveResources rpcmodule.MoveResources
	var aptosError rpcmodule.AptosError
	cl.fetchClient.Get(url).SetQueryParams(params).Execute(&moveResources, &cl.rsp, &aptosError)
	if aptosError.IsError() {
		return nil, &aptosError
	}
	return &moveResources, nil
}

func (cl *Client) AccountResourceByAddressAndType(ctx context.Context,
	address string, resourceType string, version uint64) (*rpcmodule.MoveResource, *rpcmodule.AptosError) {
	params := make(map[string]string)
	if version != 0 {
		params["version"] = fmt.Sprintf("%d", version)
	}
	url := fmt.Sprintf("/accounts/%v/resource/%v", address, resourceType)
	var moveResource rpcmodule.MoveResource
	var aptosError rpcmodule.AptosError
	cl.fetchClient.Get(url).SetQueryParams(params).Execute(&moveResource, &cl.rsp, &aptosError)
	if aptosError.IsError() {
		return nil, &aptosError
	}
	return &moveResource, nil
}

func (cl *Client) AccountModules(ctx context.Context, address string, version uint64) (*rpcmodule.MoveModules, *rpcmodule.AptosError) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	url := fmt.Sprintf("/accounts/%s/modules", address)
	var moveModules rpcmodule.MoveModules
	var aptosError rpcmodule.AptosError
	cl.fetchClient.Get(url).SetQueryParams(params).Execute(&moveModules, &cl.rsp, &aptosError)
	if aptosError.IsError() {
		return nil, &aptosError
	}
	return &moveModules, nil
}

func (cl *Client) AccountModuleByAddressAndName(ctx context.Context, address string, name string, version uint64) (*rpcmodule.MoveModule, *rpcmodule.AptosError) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	url := fmt.Sprintf("/accounts/%s/module/%s", address, name)
	var moveModule rpcmodule.MoveModule
	var aptosError rpcmodule.AptosError
	cl.fetchClient.Get(url).SetQueryParams(params).Execute(&moveModule, &cl.rsp, &aptosError)
	if aptosError.IsError() {
		return nil, &aptosError
	}
	return &moveModule, nil
}

func (cl *Client) GetResourceAccountAddress(address string, seed []byte) (string, *rpcmodule.AptosError) {
	if strings.Contains(address, "0x") {
		address = address[2:]
	}
	addr, err := hex.DecodeString(address)
	if err != nil {
		return "", rpcmodule.AptosErrorFromError(err)
	}
	addrReverse := make([]byte, len(addr))
	for i := 0; i < len(addr); i++ {
		addrReverse[i] = addr[len(addr)-1-i]
	}
	data := make([]byte, 0)
	data = append(data, addr...)
	data = append(data, seed...)
	data = append(data, []byte{255}...)
	hash := sha3.New256()
	hash.Write(data)
	return "0x" + hex.EncodeToString(hash.Sum(nil)), nil
}
