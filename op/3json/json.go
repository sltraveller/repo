package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type User struct {
    Name string
    Age int8
}

func testMarshal() []byte {
    user := User{
        Name: "Tab",
        Age: 18,
    }
    data, err := json.Marshal(user)
    if err != nil {
        log.Fatal(err)
    }
    return data
}
func testUnmarshal(data []byte) {
    var user User
    err := json.Unmarshal(data, &user)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(user)
}


func main() {
    var data []byte
    data = testMarshal()
    fmt.Println(string(data))
    testUnmarshal(data)
}