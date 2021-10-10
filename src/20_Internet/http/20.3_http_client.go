package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Go 的 Web 框架 ==> beego、gin
func main() {
	//http 包
	client := http.Client{}

	//func (c *Client) Get(url string) (resp *Response, err error)
	res, err := client.Get("https://www.baidu.com")

	if err != nil {
		fmt.Println("client.Get err:", err)
		return
	}

	ct := res.Header.Get("Content-Type")
	date := res.Header.Get("Date")
	ser := res.Header.Get("Server")
	//fmt.Println("response:", res)
	fmt.Println("Content-Type:", ct)
	fmt.Println("Date:", date)
	fmt.Println("Server:", ser)

	fmt.Println("Status", res.Status)
	fmt.Println("StatusCode:", res.StatusCode)
	fmt.Println("URL:", res.Request.URL)

	readBodyStr, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll err:", err)
	}
	fmt.Println("body:", string(readBodyStr))
}
