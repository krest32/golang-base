package main

import (
	"database/sql"
	"fmt"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

//数据库配置
const (
	userName = "root"
	password = "123456"
	ip = "127.0.0.1"
	port = "3306"
	dbName = "demo"
)
//Db数据库连接池
var DB *sql.DB

//注意方法名大写，就是public
func InitDB()  {
	//构建连接："用户名:密码@tcp-01-basic(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp-01-basic(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/day7-mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}


// 插入数据
func InsertUser(user User) (bool){
	//开启事务
	tx, err := DB.Begin()
	if err != nil{
		fmt.Println("tx fail")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO demo_user (`name`, `password`) VALUES (?, ?)")
	if err != nil{
		fmt.Println("Prepare fail")
		return false
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(user.UserName, user.Password)
	if err != nil{
		fmt.Println("Exec fail")
		return false
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}

func DeleteUser(user User) (bool) {
	//开启事务
	tx, err := DB.Begin()
	if err != nil{
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM demo_user WHERE name = ?")
	if err != nil{
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(user.UserName)
	if err != nil{
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	//获得上一个insert的id
	fmt.Println(res.LastInsertId())
	return true
}


// 更新数据
func UpdateUser(user User) (bool) {
	//开启事务
	tx, err := DB.Begin()
	if err != nil{
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("UPDATE demo_user SET name = ?, password = ? WHERE name = ?")
	if err != nil{
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(user.UserName, user.Password, user.UserName)
	if err != nil{
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}


func SelectUser(name string) (User) {
	var user User
	err := DB.QueryRow("SELECT * FROM demo_user WHERE name = ?", name).Scan(&user.UserName, &user.Password)
	if err != nil{
		fmt.Println("查询出错了")
	}
	return user
}

func SelectAllUser() ([]User) {
	//执行查询语句
	rows, err := DB.Query("SELECT * from user")
	if err != nil{
		fmt.Println("查询出错了")
	}
	var users []User
	//循环读取结果
	for rows.Next(){
		var user User
		//将每一行的结果都赋值到一个user对象中
		err := rows.Scan(&user.UserName, &user.Password)
		if err != nil {
			fmt.Println("rows fail")
		}
		//将user追加到users的这个数组中
		users = append(users, user)
	}
	return users
}

// 定义一个 DivideError 结构
type User struct {
	UserName string
	Password string
}

func main() {
	InitDB()
	//user1 := User{
	//	UserName : "duxin",
	//	Password : "123456",
	//}

	//user2 := User{
	//	UserName : "duxin12",
	//	Password : "123456789",
	//}
	//Fmt.Println(InsertUser(user2))
	//Fmt.Println(DeleteUser(user))
	//Fmt.Println(UpdateUser(user2))
	//Fmt.Println(SelectUser("duxin"))
	// 返回数组数组
	userList := SelectAllUser()
	fmt.Println(userList[1])
}