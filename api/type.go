package api

import (
	"api_automation/header"
	customHttp "api_automation/http"
)

type IClient interface {
	Get(string, header.Header) ([]byte, error)
	Post(string, header.Header, interface{}) ([]byte, error)
	Put(string, header.Header, interface{}) ([]byte, error)
	Delete(string, header.Header) ([]byte, error)
}

type Client struct {
	ch customHttp.ICustomClient
}
