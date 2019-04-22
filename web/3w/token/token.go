package token

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"bytes"
)

type Statu struct{
	Code int `json:"code"`
	Message string `json:"message"`
	Status int `json:"status"`
}

type Data struct{
	Token string `json:"token"`
	User_name string `json:"user_name"`
	Expired int `json:"expired"`
}

type Plat struct{
	Statu
	Data *Data `json:"data"`
}

// Gettoken get
func Gettoken() (token string, err error){
	urlstr := "http://beta-tree.op.kingnet.com/api/v1/token"
	//创建客户端
	client := http.Client{}
	//创建request
	key := map[string]string{"plat_name": "cicd", "plat_key": "52695bc33a6b48ccbb3be1960e4e006a"}
	body, _ := json.Marshal(key)
	request, err := http.NewRequest("POST", urlstr, bytes.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/json")
	//发送请求
	response, err := client.Do(request)
	Ferr(err)
	//处理返回信息
	defer response.Body.Close()
	plat := new(Plat)
	resp, err := ioutil.ReadAll(response.Body)
	if err := json.Unmarshal(resp, plat); err != nil{
		fmt.Println(err)
	}
	//fmt.Println(plat.Status)
	return plat.Data.Token, err
}

func Ferr(err error){
	if err != nil {
		fmt.Println(err)
	}
}