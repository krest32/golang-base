package main

import (
	"fmt"
)

//定义结构体
type User struct {
	Name  string
	Email string
}

func main() {
	// 通过值类型调用方法
	u1 := User{"golang", "golang@golang.com"}
	// 通过对象调用方法
	u1.Notify()

	// 通过指针类型调用方法
	u2 := User{"go", "go@go.com"}
	// 指针传递
	u3 := &u2
	u3.Notify()
}

//定义方法
func (u *User) Notify() {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}
