package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	response, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(body))
	src := string(body)
	re, _ := regexp.Compile("\\<[\\S\\s]+?>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
    re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
    src = re.ReplaceAllString(src, "")
    //去除SCRIPT
    re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
    src = re.ReplaceAllString(src, "")
    //去除所有尖括号内的HTML代码，并换成换行符
    re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
    src = re.ReplaceAllString(src, "\n")
    //去除连续的换行符
    re, _ = regexp.Compile("\\s{2,}")
    src = re.ReplaceAllString(src, "\n")
    fmt.Println(strings.TrimSpace(src))
}