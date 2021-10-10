package main

import (
	"encoding/json"
	"fmt"
)

type Teacher struct {
	Name    string `json:"-"`                 //在使用 JSON 编码时不编码该字段
	Subject string `json:"Sub_name"`          //在 JSON 编码时会用标签替换原字段名
	Age     int    `json:"age,string"`        //在 JSON 编码时更改字段数据类型,注意中间逗号左右不能加空格:字段名,字段类型
	Address string `json:"address,omitempty"` //omitempty表示如果该字段为空则不参与编码
	gender  string
}

func main() {
	T1 := Teacher{
		Name:    "Teacher1",
		Subject: "Golang",
		Age:     40,
		gender:  "man",
		//Address: "南京",
	}
	fmt.Println("Teacher1", T1)
	encodeInfo, err := json.Marshal(&T1)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	fmt.Println("T1 after encoding:", string(encodeInfo))
}
