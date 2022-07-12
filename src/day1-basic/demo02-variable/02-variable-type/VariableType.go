package main

import "fmt"

// 变量类型
func main() {
	// 字符串类型
	var a string = "菜鸟教材"
	fmt.Println(a)

	// 整数类型
	var b, c int = 1, 2
	fmt.Println(b, c)
	// int16 int32 int64
	var f int8 = 3
	fmt.Println(f)

	// 浮点类型
	var d float32 = 3.14 * 2
	fmt.Println(d)

	// 浮点类型 默认类型 float64
	var e float64 = 3.14 * 2
	fmt.Println(e)

	var h bool = true
	fmt.Println(h)
}
