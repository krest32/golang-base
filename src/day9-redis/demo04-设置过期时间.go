package main


import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp-01-basic", "121.196.111.229:6379")
	if err != nil {
		fmt.Println("conn day9-redis failed,", err)
		return
	}

	defer c.Close()
	_, err = c.Do("expire", "abc", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
}