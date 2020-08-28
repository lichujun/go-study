package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

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
		log.Println(err)
	}
}

func Test(t *testing.T) {
	connPost()
}
