package main

import "fmt"

/**
	指针操作
 */
func main() {
	 a := 20   /* 声明实际变量 */
	 var ip *int        /* 声明指针变量 */

	ip = &a  /* 指针变量的存储地址，通过这种方式，内存的使用会得到优化 */

	fmt.Printf("a 变量的地址是: %x \n", &a  )

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x \n", ip )

	// 使用指针访问值，将变量为指针的数据，转化为刻度数据
	fmt.Printf("*ip 变量的值: %d\n", *ip )


}