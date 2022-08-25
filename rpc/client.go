package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func (cl *Client) Get(ctx context.Context, path string, params map[string]string, result interface{}) (int, error) {
	req, err := http.NewRequest("GET", cl.url+path, nil)
	if err != nil {
		return -1, err
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
		return -1, err
	}
	if resp.StatusCode != 200 {
		return resp.StatusCode, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return resp.StatusCode, err
	}
	err = json.Unmarshal(respBody, result)
	if err != nil {
		return -1, err
	}
	return resp.StatusCode, nil
}

func (cl *Client) Post(ctx context.Context, path string, params map[string]string, body interface{}, result interface{}) (int, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return -1, err
	}
	req, err := http.NewRequest("POST", cl.url+path, bytes.NewBuffer(reqBody))
	if err != nil {
		return -1, err
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
		return -1, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	// 202 - Transaction is accepted and submitted to mempool.
	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return resp.StatusCode, fmt.Errorf("%s", string(respBody))
	}
	err = json.Unmarshal(respBody, result)
	if err != nil {
		return -1, err
	}
	return resp.StatusCode, nil
}
