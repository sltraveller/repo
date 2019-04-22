package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Mkdir("h1", 0555)
	os.MkdirAll("h1/h2/h3", 0555)
	err := os.Remove("h1")
	if err != nil {
		fmt.Println("err:", err)
	}
	//os.RemoveAll("h1")
}
