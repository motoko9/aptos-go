package rpc

import (
	"context"
	"github.com/klauspost/compress/gzhttp"
	"github.com/motoko9/aptos-go/rpc/jsonrpc"
	"io"
	"net"
	"net/http"
	"time"
)

type JSONRPCClient interface {
	CallForInto(ctx context.Context, out interface{}, method string, params []interface{}) error
	CallWithCallback(ctx context.Context, method string, params []interface{}, callback func(*http.Request, *http.Response) error) error
}

type Client struct {
	url    string
	client JSONRPCClient
}

func New(endpoint string) *Client {
	opts := &jsonrpc.RPCClientOpts{
		HTTPClient: newHTTP(),
	}
	rpcClient := jsonrpc.NewClientWithOpts(endpoint, opts)
	return &Client{
		client: rpcClient,
	}
}

func NewWithHeaders(endpoint string, headers map[string]string) *Client {
	opts := &jsonrpc.RPCClientOpts{
		HTTPClient:    newHTTP(),
		CustomHeaders: headers,
	}
	rpcClient := jsonrpc.NewClientWithOpts(endpoint, opts)
	return &Client{
		client: rpcClient,
	}
}

func (cl *Client) Close() error {
	if cl.client == nil {
		return nil
	}
	if c, ok := cl.client.(io.Closer); ok {
		return c.Close()
	}
	return nil
}

func (cl *Client) RPCCallForInto(ctx context.Context, out interface{}, method string, params []interface{}) error {
	return cl.client.CallForInto(ctx, out, method, params)
}

func (cl *Client) RPCCallWithCallback(
	ctx context.Context,
	method string,
	params []interface{},
	callback func(*http.Request, *http.Response) error) error {
	return cl.client.CallWithCallback(ctx, method, params, callback)
}

var (
	defaultMaxIdleConnsPerHost = 9
	defaultTimeout             = 5 * time.Minute
	defaultKeepAlive           = 180 * time.Second
)

func newHTTPTransport() *http.Transport {
	return &http.Transport{
		IdleConnTimeout:     defaultTimeout,
		MaxConnsPerHost:     defaultMaxIdleConnsPerHost,
		MaxIdleConnsPerHost: defaultMaxIdleConnsPerHost,
		Proxy:               http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   defaultTimeout,
			KeepAlive: defaultKeepAlive,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2: true,
		// MaxIdleConns:          100,
		TLSHandshakeTimeout: 10 * time.Second,
		// ExpectContinueTimeout: 1 * time.Second,
	}
}

func newHTTP() *http.Client {
	tr := newHTTPTransport()

	return &http.Client{
		Timeout:   defaultTimeout,
		Transport: gzhttp.Transport(tr),
	}
}
