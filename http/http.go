package http

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func PostRequest(url string, headers map[string]string, payload interface{}) ([]byte, error) {
	bytes, err := json.Marshal(payload)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(bytes)))
	if err != nil {
		log.Fatalf("error while creating request: %v", err)
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, err := client.Do(req)
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

func PutRequest(url string, headers map[string]string, payload interface{}) ([]byte, error) {
	bytes, err := json.Marshal(payload)
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, strings.NewReader(string(bytes)))
	if err != nil {
		log.Fatalf("error while creating request: %v", err)
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, err := client.Do(req)
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

func DeleteRequest(url string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatalf("error while creating request: %v", err)
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, err := client.Do(req)
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

func GetRequest(url string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("error while creating request: %v", err)
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, err := client.Do(req)
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
