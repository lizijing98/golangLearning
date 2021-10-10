package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   int
	Name string
	Age  int
	//gender 是小写,小写字母开头在 JSON 编码时会忽略掉
	gender string
}

func main() {
	xiaoming := Student{
		Id:     1,
		Name:   "xiaoming",
		Age:    20,
		gender: "man",
	}

	//1.编码（序列化）
	//func Marshal(v interface{}) ([]byte, error)
	encodeInfo, err := json.Marshal(&xiaoming) //函数参数是接口，传地址
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	fmt.Println("encodeInfo:", string(encodeInfo))

	//2.解码
	var xiaoming2 Student
	err = json.Unmarshal(encodeInfo, &xiaoming2)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	fmt.Println("xiaoming name:", xiaoming2.Name)
	fmt.Println("xiaoming gender:", xiaoming2.gender)
}
