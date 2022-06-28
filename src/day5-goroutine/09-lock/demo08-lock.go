package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x8      int64
	wg8     sync.WaitGroup
	lock8   sync.Mutex
	rwlock8 sync.RWMutex
)

func write() {
	// lock.Lock()   // 加互斥锁
	rwlock8.Lock() // 加写锁
	x8 = x8 + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock8.Unlock()                   // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg8.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock8.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock8.RUnlock()             // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg8.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg8.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg8.Add(1)
		go read()
	}

	wg8.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}