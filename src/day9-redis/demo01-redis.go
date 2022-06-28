package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp-01", "121.196.111.229:6379")
	if err != nil {
		fmt.Println("Connect to day9-redis error", err)
		return
	}
	fmt.Println("Connect to day9-redis success")
	defer c.Close()
}
