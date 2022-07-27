package http

import (
	"api_automation/header"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func (ch *CustomHttp) Do(req *http.Request) (*http.Response, error) {
	return ch.h.Do(req)
}

func NewCustomHttp() *CustomHttp {
	return &CustomHttp{h: http.DefaultClient}
}

func (ch *CustomHttp) SendHttpRequest(methodName string, url string, headers header.Header, bytes []byte) ([]byte, error) {
	req, err := http.NewRequest(methodName, url, strings.NewReader(string(bytes)))
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("error while closing body: %v", err)
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error while reading body: %v", err)
		return nil, err
	}
	return body, nil
}
