package main

import (
	"fmt"
	"sync"
)

var x int64
var wg6 sync.WaitGroup

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg6.Done()
}

/**
	上面的代码中我们开启了两个goroutine去累加变量x的值，
	这两个goroutine在访问和修改x变量的时候就会存在数据
	竞争，导致最后的结果与期待的不符。
	每次的运行结果不一致
 */
func main() {
	wg6.Add(2)
	go add()
	go add()
	wg6.Wait()
	fmt.Println(x)
}
