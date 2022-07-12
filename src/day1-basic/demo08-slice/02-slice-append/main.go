package main

import "fmt"

/**
slice可变数组

*/
func main() {
	stack := []string{}
	stack = append(stack, "aaa", "bbb", "ccc")
	// _ 代表索引，str 代表元素
	for _, str := range stack {
		fmt.Println(str)
	}
}
