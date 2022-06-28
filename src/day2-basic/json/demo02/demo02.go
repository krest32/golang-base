package main

import (
	"encoding/json"
	"fmt"
)

type Person2 struct {
	// 指定json序列化/反序列化时使用小写name
	Name   string	`json:"name"`
	Age    int64
	Weight float64
}

func main() {
	p1 := Person2{
		Name:   "七米",
		Age:    18,
		Weight: 71.5,
	}
	// struct -> json string
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	// json string -> struct
	var p2 Person2
	err = json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}
