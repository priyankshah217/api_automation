package api

import (
	"api_automation/header"
	customHttp "api_automation/http"
	"encoding/json"
	"net/http"
)

func (c *Client) Get(url string, headers header.Header) ([]byte, error) {
	body, err := c.ch.SendHttpRequest(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) Post(url string, headers header.Header, payload interface{}) ([]byte, error) {
	bytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	body, err := c.ch.SendHttpRequest(http.MethodPost, url, headers, bytes)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) Put(url string, headers header.Header, payload interface{}) ([]byte, error) {
	bytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	body, err := c.ch.SendHttpRequest(http.MethodPut, url, headers, bytes)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) Delete(url string, headers header.Header) ([]byte, error) {
	body, err := c.ch.SendHttpRequest(http.MethodDelete, url, headers, nil)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func NewClient() IClient {
	return &Client{ch: customHttp.NewCustomHttp()}
}
