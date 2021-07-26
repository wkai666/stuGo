package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Addr string `json:"addr,omitempty"`
}

func main() {
	p1 := Person{
		Name: "Jack",
		Age: 23,
	}

	data1, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data1)

	p2 := Person{
		Name: "Tom",
		Age: 22,
		Addr: "China",
	}

	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data2)
}