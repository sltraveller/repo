package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age      int
	Birthday string
	Sex      string
	Email    string
	Phone    string
}

/*map转json*/

func testMap() {
	var mmp map[string]interface{}
	mmp = make(map[string]interface{})

	mmp["username"] = "Murphy"
	mmp["age"] = 19
	mmp["sex"] = "man"

	data, err := json.Marshal(mmp)
	if err != nil {
		fmt.Println("json marshal failed,err:", err)
		return
	}
	fmt.Printf("%s\n", string(data))

}

func main() {
	testMap()
}
