package main

import (
	"fmt"
	"log"
	"os"
)

// 定义好全局变量
const (
	// 控制输出日志信息的细节，不能控制输出的顺序和格式。
	// 输出的日志在每一项后会有一个冒号分隔：例如2009/01-对象方法-for & defer/23 01-对象方法-for & defer:23:23.123123 /a/b/c/d.go:23: message
	Ldate         = 1 << iota     // 日期：2009/01-对象方法-for & defer/23
	Ltime                         // 时间：01-对象方法-for & defer:23:23
	Lmicroseconds                 // 微秒级别的时间：01-对象方法-for & defer:23:23.123123（用于增强Ltime位）
	Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
	Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
	LUTC                          // 使用UTC时间
	LstdFlags     = Ldate | Ltime // 标准logger的初始值
)

// 显示基本的日志
func  basic()  {
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")
}

// 自定义log基本信息
func basicLog()  {

	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
}

// 设置日志的前缀
func basicPrefix ()  {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[小王子]")
	log.Println("这是一条很普通的日志。")

}

func writeLogToFile()  {
	logFile, err := os.OpenFile("src/utils/demo17_log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open day6-lib file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[小王子]")
	log.Println("这是一条很普通的日志。")
}


func main() {
	// 显示日志的基本功能
	//basic()

	// 自定义日志
	//basicLog()

	//设置日志的前缀
	basicPrefix()

	// log写出文件
	//writeLogToFile()
}