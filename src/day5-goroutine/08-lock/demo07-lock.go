package main

import (
	"fmt"
	"sync"
)

var x7 int64
var wg7 sync.WaitGroup
var lock sync.Mutex

func add7() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x7 = x7 + 1
		lock.Unlock() // 解锁
	}
	wg7.Done()
}
func main() {
	wg7.Add(2)
	go add7()
	go add7()
	wg7.Wait()
	fmt.Println(x7)
}
