package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "100"
	// 函数用于将字符串类型的整数转换为int类型
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}


	i2 := 200
	// 函数用于将int类型数据转换为对应的字符串表示
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2) //type:string value:"200"
}
