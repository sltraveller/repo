package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"web/3w/token"
)

type Instance struct {
	Idc             string `json:"idc"`
	Env             string `json:"env"`
	EnvDesc         string `json:"env_desc"`
	Product         string `json:"product"`
	ProductDesc     string `json:"product_desc"`
	Application     string `json:"application"`
	ApplicationDesc string `json:"application_desc"`
	Domain          string `json:"domain"`
	Services        string `json:"services"`
	Hostname        string `json:"hostname"`
	InternalIP      string `json:"internal_ip"`
	ServiceIP       string `json:"service_ip"`
	UseState        string `json:"use_state"`
	Type            string `json:"type"`
	Department      string `json:"department"`
	Path            string `json:"path"`
}

type Datas struct {
	Total   int         `json:"total"`
	Results int         `json:"results"`
	Page    int         `json:"page"`
	Data    []*Instance `json:"data"`
}

type Appinstance struct {
	token.Statu
	Data Datas `json:"data"`
}

func main() {
	url := "http://beta-tree.op.kingnet.com/api/v1/instances/detail?page=1&results=10"
	request, err := http.NewRequest("GET", url, nil)
	token.Ferr(err)
	tokens, _ := token.Gettoken()
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Authorization-Token", tokens)
	response, err := http.DefaultClient.Do(request)
	token.Ferr(err)
	
	resp, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	instance := new(Appinstance)
	json.Unmarshal(resp, instance)
	fmt.Println(instance.Data.Data[0].Idc)
}
