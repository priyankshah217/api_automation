package http

import (
	"api_automation/header"
	"net/http"
)

type ICustomClient interface {
	Do(*http.Request) (*http.Response, error)
	SendHTTPRequest(string, string, header.Header, []byte) ([]byte, error)
}

type CustomHTTP struct {
	h *http.Client
}
