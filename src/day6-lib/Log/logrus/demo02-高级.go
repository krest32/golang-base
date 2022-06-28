package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

// 创建一个新的logger实例。可以创建任意多个。
var log2 = logrus.New()

func main() {
	// 设置日志输出为os.Stdout
	log2.Out = os.Stdout

	// 可以设置像文件等任意`io.Writer`类型作为日志输出
	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	//  log.Out = file
	// } else {
	//  log.Info("Failed to log to file, using default stderr")
	// }

	log2.WithFields(logrus.Fields{
		"animal": "dog",
		"size":   10,
	}).Info("一群舔狗出现了。")

	log2.Trace("Something very low level.")
	log2.Debug("Useful debugging information.")
	log2.Info("Something noteworthy happened!")
	log2.Warn("You should probably take a look at this.")
	log2.Error("Something failed but I'm not quitting.")
	// 记完日志后会调用os.Exit(1)
	log2.Fatal("Bye.")
	// 记完日志后会调用 panic()
	log2.Panic("I'm bailing.")
}