package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)


var db *sqlx.DB

//数据库配置
const (
	userName = "root"
	password = "Bob.123456"
	ip = "127.0.0.1"
	port = "3306"
	dbName = "demo"
)

func initDB() (err error) {
	//构建连接："用户名:密码@tcp-01-basic(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp-01-basic(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", path)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	fmt.Println("connect success")
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

func main()  {
	initDB()
}