package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/motoko9/aptos-go/rpcmodule"
	"io/ioutil"
	"net/http"
)

type Client struct {
	url     string
	headers map[string]string
	client  *http.Client
}

func New(endpoint string) *Client {
	return &Client{
		url:    endpoint,
		client: &http.Client{},
	}
}

func (cl *Client) SetHeaders(headers map[string]string) {
	cl.headers = headers
}

func (cl *Client) Get(ctx context.Context, path string, params map[string]string, result interface{}) (error, *rpcmodule.AptosError) {
	req, err := http.NewRequest("GET", cl.url+path, nil)
	if err != nil {
		return err, nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accepts", "application/json")

	if cl.headers != nil {
		for k, v := range cl.headers {
			req.Header.Set(k, v)
		}
	}

	if params != nil {
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	req = req.WithContext(ctx)
	resp, err := cl.client.Do(req)
	if err != nil {
		return err, nil
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err, nil
	}
	if resp.StatusCode != 200 {
		// fetch aptos error
		var aptosError rpcmodule.AptosError
		if err = json.Unmarshal(respBody, &aptosError); err != nil {
			return err, nil
		}
		return nil, &aptosError
	}
	// try to
	// todo
	/*
		var aptosError rpcmodule.AptosError
		if err = json.Unmarshal(respBody, &aptosError); err != nil {
			return err, nil
		}
		if aptosError.ErrorCode != "" {
			return nil, &aptosError
		}
	*/
	err = json.Unmarshal(respBody, result)
	if err != nil {
		return err, nil
	}
	return nil, nil
}

func (cl *Client) Post(ctx context.Context, path string, params map[string]string, body interface{}, result interface{}) (error, *rpcmodule.AptosError) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return err, nil
	}
	req, err := http.NewRequest("POST", cl.url+path, bytes.NewBuffer(reqBody))
	if err != nil {
		return err, nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accepts", "application/json")

	if cl.headers != nil {
		for k, v := range cl.headers {
			req.Header.Set(k, v)
		}
	}

	if params != nil {
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	req = req.WithContext(ctx)
	resp, err := cl.client.Do(req)
	if err != nil {
		return err, nil
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err, nil
	}
	// 202 - Transaction is accepted and submitted to mempool.
	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		// fetch aptos error
		var aptosError rpcmodule.AptosError
		if err = json.Unmarshal(respBody, &aptosError); err != nil {
			return err, nil
		}
		return nil, &aptosError
	}
	// try to
	// todo
	/*
		var aptosError rpcmodule.AptosError
		if err = json.Unmarshal(respBody, &aptosError); err != nil {
			return err, nil
		}
		if aptosError.ErrorCode != "" {
			return nil, &aptosError
		}
	*/
	err = json.Unmarshal(respBody, result)
	if err != nil {
		return err, nil
	}
	return nil, nil
}
