package main

import "fmt"

/**
slice可变数组

*/
func main() {
	stack := []string{}
	stack = append(stack, "aaa", "bbb", "ccc")

	// 左开右闭的区间

	// 从1到2的元素，包含1，不包含2
	for _, str := range stack[1:2] {
		fmt.Println(str)
	}

	fmt.Println("---------")

	// 从1到结尾的元素,包含1
	for _, str := range stack[1:] {
		fmt.Println(str)
	}
	fmt.Println("---------")

	// 从0到2的元素，不包含2
	for _, str := range stack[:2] {
		fmt.Println(str)
	}

}
