package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	_ "io/ioutil"
	_ "log"
)

/*
func main() {
	requesturl := "http://www.baidu.com"
	response, err := http.Get(requesturl)
	if err != nil {
		fmt.Println("err:",err)
	}
	defer response.Body.Close()
	fmt.Println(response.Body)
}
*/

func main() {
	/*
	//客户端实现一
	urlstr := "http://www.chaindesk.cn/"
	//创建client对象
	client := http.Client{}
	//创建request
	request, err := http.NewRequest("Get", urlstr, nil)
	if err != nil {
		log.Fatal(err)
	}
	//发送请求
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}
	fmt.Println("-------------")
	fmt.Printf("reponse: %+v\n",response)
	fmt.Println("-----------------")
	fmt.Printf("body: %+v\n", response.Body)
	fmt.Printf("header: %+v\n",response.Header)
	fmt.Printf("statuscode: %+v\n",response.StatusCode)
	fmt.Printf("status: %+v\n",response.Status)
	fmt.Printf("request: %+v\n", response.Request)
	fmt.Printf("coolies: %+v\n",response.Cookies())
	*/
	//或者
	request, err := http.NewRequest(http.MethodGet, "http://www.chaindesk.cn/", nil)
	request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	s, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n",s)
	

}