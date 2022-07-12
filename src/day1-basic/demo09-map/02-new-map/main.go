package main

import "fmt"

func main() {
	// 初始化
	data := map[byte]int{'a': 1, 'b': 2, 'c': 3}
	fmt.Println(data['a'])
	fmt.Println(data['b'])
	fmt.Println(data['c'])
}
