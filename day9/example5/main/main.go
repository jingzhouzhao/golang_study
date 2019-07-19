package main

import (
	"time"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init(){
	pool = &redis.Pool{
		MaxActive:0,
		MaxIdle:16,
		IdleTimeout:300,
		Dial:func() (redis.Conn, error){
			return redis.Dial("tcp", "localhost:6379",redis.DialConnectTimeout(10*time.Millisecond))
		},
	}
}

func main(){
	conn:= pool.Get()
	defer conn.Close()
	// _, err := conn.Do("Set", "user_id", 123)
	// if err != nil {
	// 	fmt.Println("Redis Execute Set Command Failed:", err)
	// }
	r, err := redis.Int(conn.Do("Get", "user_id"))
	if err != nil {
		fmt.Println("Redis Execute Get Command Failed:", err)
	}
	fmt.Println(r)
	defer pool.Close()
}