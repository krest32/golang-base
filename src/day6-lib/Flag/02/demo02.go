package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	//定义命令行参数方式1
	var name string
	var age int
	var married bool
	var delay time.Duration

	/*
		flag.Type(flag名, 默认值, 帮助信息)*Type
		例如我们要定义姓名、年龄、婚否三个命令行参数，我们可以按如下方式定义：
	 */

	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
	flag.Usage=printUsage
}

// 重写
func printUsage(){
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
	fmt.Printf("hihihi\n")
}