package main

import (
	"fmt"
	"strings"
)

func main()  {
	// 转义字符串
	path := "\"d:\\go\\src\\code\""
	fmt.Println(path)

	//字符串
	s := "I'm ok"
	fmt.Println(s)

	// 多行字符串，原样输出
	s2 := `
		123
		234
	`
	fmt.Println(s2)

	// 求取字符串长度
	fmt.Println(len(s2))

	// 字符串拼接
	name1 := "你"
	name2 := "hao"
	name := name1+name2
	fmt.Println(name)

	// 分割字符串,成为一个数组
	ret := strings.Split(path,"\\")
	fmt.Println(ret)

	// 包含某个字符串
	fmt.Println(strings.Contains(path,"d:"))
}
