package main

import (
	"encoding/json"
	"fmt"
)

func testInt() {
	var age = 100
	data, err := json.Marshal(age)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

func main() {
	testInt()
}
