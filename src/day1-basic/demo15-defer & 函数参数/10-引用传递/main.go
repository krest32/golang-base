package main

import (
	"fmt"
)

/*
	定义相互交换值的函数
	在默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
	注意
	1：无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，
	值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效。
	而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。
	2：map、slice、chan、指针、interface默认以引用的方式传递。

	不定参数传值 就是函数的参数不是固定的，后面的类型是固定的。（可变参数）
*/
func swap(x, y *int) {
	var temp int

	temp = *x /* 保存 x 的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y*/

}

func main() {
	var a, b int = 1, 2
	/*
	   调用 swap() 函数
	   &a 指向 a 指针，a 变量的地址
	   &b 指向 b 指针，b 变量的地址
	*/
	swap(&a, &b)

	fmt.Println(a, b)
}