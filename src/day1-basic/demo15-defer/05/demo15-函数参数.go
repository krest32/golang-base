package main

import "fmt"

// 函数多参数
func main() {
	println(test("sum: %d",1,2,3))
}

// ... 代表可以输入多个数组参数
func test(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}

