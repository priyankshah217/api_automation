package http

import (
	"api_automation/header"
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func (ch *CustomHTTP) Do(req *http.Request) (*http.Response, error) {
	return ch.h.Do(req)
}

func NewCustomHTTP() *CustomHTTP {
	return &CustomHTTP{h: http.DefaultClient}
}

func (ch *CustomHTTP) SendHTTPRequest(name string, url string, headers header.Header, bytes []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(context.Background(), name, url, strings.NewReader(string(bytes)))
	if err != nil {
		log.Fatalf("error while creating request: %v", err)
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, err := ch.Do(req)
	if err != nil {
		log.Fatalf("error while making request: %v", err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
