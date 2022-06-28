package main

import (
	"fmt"
)

// 接口
type Phone interface {
	call()
}

/* 定义结构体 */
type NokiaPhone struct {
}

/* 定义结构体 */
type IPhone struct {
}

/* 实现接口方法 */
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

/* 实现接口方法 */
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}