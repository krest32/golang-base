package main

import "fmt"

// 定义结构体
type Data struct {
	x int
}


func main() {
	d := Data{
		x: 20,
	}

	fmt.Printf("%d\n",d)
	fmt.Printf("=========================\n")
	// 得到在内存中的地址
	p :=&d
	// 打印内存中的地址
	fmt.Printf("Data: %p\n", p)
	// 打印内存中的数据
	fmt.Printf("Data: %d\n", *p)
	fmt.Printf("=========================\n")
	d.VAlueTest()
	d.PointerTest()
	fmt.Printf("=========================\n")
	p.VAlueTest()
	p.PointerTest()
}

// 定义方法
func (self Data) VAlueTest()  {
	// 打印内存中的地址
	fmt.Printf("Value : %p\n", &self)
	// 打印值
	fmt.Printf("Value : %d\n", self)
}

// 定义指针方法,获取得到的是Data的指针
func (self *Data) PointerTest() {
	fmt.Printf("PointerTest: %p\n", self)
	// 从指针中获取数值
	fmt.Printf("PointerTest: %d\n", *self)
}
