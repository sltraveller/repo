package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/xml"
)

// Recurlyservers is
type Recurlyservers struct{
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs []server `xml:"server"`
	Description string `xml:"serverIP"`
}

type server struct {
	XMLName xml.Name `xml:"server"`
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}

func main() {
	file, err := os.Open("xml.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil{
		fmt.Printf("error:%v",err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil{
		fmt.Printf("error:%v", err)
		return
	}
	fmt.Println(v)

}