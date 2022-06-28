package main

import (
	"fmt"
	"os"
)

//os.Args demo
func main() {

	//os.Args是一个[]string，获取命令的输入
	if len(os.Args) > 1 {
		// 遍历命令行输入的信息
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}


