package rpc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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

func (cl *Client) SetHeader(headers map[string]string) {
	cl.headers = headers
}

func (cl *Client) Get(path string, params map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", cl.url+path, nil)
	if err != nil {
		return nil, err
	}

	if cl.headers != nil {
		for k, v := range cl.headers {
			req.Header.Set(k, v)
		}
	}

	if params != nil {
		q := url.Values{}
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := cl.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func (cl *Client) Post(path string, params map[string]string, body []byte) ([]byte, error) {
	reqReader := bytes.NewReader(body)
	req, err := http.NewRequest("POST", cl.url+path, reqReader)
	if err != nil {
		return nil, err
	}

	if cl.headers != nil {
		for k, v := range cl.headers {
			req.Header.Set(k, v)
		}
	}

	if params != nil {
		q := url.Values{}
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := cl.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
