package main

import "fmt"

// 变量必须使用，否则编译无法通过
func main() {
	// 编译错误，a没有使用
	//var a string = "abc"
	fmt.Println("hello, world")
}