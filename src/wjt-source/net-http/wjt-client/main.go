package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)
//建立的client客户端，发起的get请求，在请求体中添加header，然后使用client执行这个请求
func main() {
    client:=&http.Client{}
    res,err:=http.NewRequest("GET","https://www.baidu.com",nil)
    res.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
    resp,err:=client.Do(res)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}