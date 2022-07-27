package http

import (
	"api_automation/header"
	"net/http"
)

type ICustomClient interface {
	Do(*http.Request) (*http.Response, error)
	SendHttpRequest(string, string, header.Header, []byte) ([]byte, error)
}

type CustomHttp struct {
	h *http.Client
}
