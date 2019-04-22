package nodes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"web/3w/token"
)


type Node struct {
	Name     string                 `json:"name"`
	Alias    string                 `json:"alias"`
	Uuid     string                 `json:"uuid"`
	Path     string                 `json:"path"`
	Type     int                    `json:"type"`
	Tags     map[string]interface{} `json:"tags"`
	Children map[string]*Node       `json:"children"`
}
type AppNode struct {
	token.Statu
	Data map[string]*Node `json:"data"`
}

func GetNodes() *Node {
	urlstl := "http://beta-tree.op.kingnet.com/api/v1/tree/app"
	client := http.Client{}
	request, err := http.NewRequest("GET", urlstl, nil)
	token.Ferr(err)
	tokens, err := token.Gettoken()
	token.Ferr(err)
	//fmt.Println(token)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Authorization-Token", tokens)
	//fmt.Println(request)
	response, err := client.Do(request)
	token.Ferr(err)
	defer response.Body.Close()
	resp, err := ioutil.ReadAll(response.Body)
	token.Ferr(err)
	data := new(AppNode)
	json.Unmarshal(resp, data)
	//fmt.Println(data.Data["kingnet"])
	return data.Data["kingnet"]
}
