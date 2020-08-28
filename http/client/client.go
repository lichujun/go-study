package client

import (
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
