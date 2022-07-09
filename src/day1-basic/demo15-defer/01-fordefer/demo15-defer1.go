package main

import "fmt"

// defer 延缓打印
func main() {
	var whatever [5]struct{}
	//因为最后一次调用 defer 时传入了 fmt.Println(4)，所以这段代码会优先打印 4。
	for i := range whatever {
		defer fmt.Println(i)
	}
}