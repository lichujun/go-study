package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{
		Name: "joseph",
		Age:  12,
	}
	res, _ := json.Marshal(p)
	fmt.Println(string(res))
}
