package main

import (
	"fmt"
	"sync"
	"time"
)
// 同步锁
var lock sync.Mutex



/**

 */
func main() {

	// 相当于Java的匿名内部类
	func() {
		// 设定现在的时间
		t1 := time.Now()
		// 循环执行一个方法
		for i := 0; i < 1000000; i++ {
			test15_2()
		}
		// 计算花费的时间
		elapsed := time.Since(t1)
		fmt.Println("test elapsed: ", elapsed)
	}()

	// 相当于Java的匿名内部类
	// 测试使用defer上锁、解锁，性能感觉会低了一些
	func() {
		t1 := time.Now()
		for i := 0; i < 1000000; i++ {
			testDefer15_2()
		}
		elapsed := time.Since(t1)
		fmt.Println("test defer elapsed: ", elapsed)
	}()

}


func test15_2() {
	lock.Lock()
	lock.Unlock()
}

func testDefer15_2() {
	lock.Lock()
	defer lock.Unlock()
}