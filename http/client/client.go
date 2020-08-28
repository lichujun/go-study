package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	httpClient *http.Client
)

const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 30
)

// init HTTPClient
func init() {
	httpClient = createHTTPClient()
}

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}
	return client
}

func GetHttpClient() *http.Client {
	return httpClient
}

func connPost() {
	path := "http://localhost:8080/hello"
	resp, err := httpClient.Post(path, "application/x-protobuf", nil)
	if err != nil {
		log.Println("err", err)
		return
	}
	buf, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(buf))
	if err := resp.Body.Close(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	connPost()
}
