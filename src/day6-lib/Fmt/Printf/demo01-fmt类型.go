package main

import "fmt"

func main()  {
	var n int=100
	fmt.Printf("%T\n", n) // 查看类型
	fmt.Printf("%v\n",n ) //查看变量的值
	fmt.Printf("%d\n",n) // 查看十进制
	fmt.Printf("%b\n",n) // 查看二进制
	fmt.Printf("%b\n",n) // 查看二进制
	fmt.Printf("%o\n",n) // 查看八进制
	fmt.Printf("%x\n",n) // 查看十六进制

	var s string="hello"
	fmt.Printf("%s\n",s) // 查看字符串
}
