package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("connection to redis failed:", err)
		return
	}
	defer conn.Close()

	//Set
	_, err = conn.Do("Set", "user_id", 123)
	if err != nil {
		fmt.Println("Redis Execute Set Command Failed:", err)
	}
	r, err := redis.Int(conn.Do("Get", "user_id"))
	if err != nil {
		fmt.Println("Redis Execute Get Command Failed:", err)
	}
	fmt.Println(r)

	//MSet(批量Set)
	_, err = conn.Do("MSet", "username", "lishi", "nickname", "ls")
	ms, err := redis.Strings(conn.Do("MGet", "username", "nickname"))
	for _, v := range ms {
		fmt.Println("mget:", v)
	}

	//HSet
	_, err = conn.Do("HSet", "user", "username", "zhangsan")
	if err != nil {
		fmt.Println("Redis Execute Hset Command Failed:", err)
	}
	s, err := redis.String(conn.Do("HGet", "user", "username"))
	fmt.Println(s)
	//获取全部
	m, err := redis.StringMap(conn.Do("HGetAll", "user"))
	fmt.Println(m)
	//lpush/lpop
	for i := 0; i < 100; i++ {
		_, err = conn.Do("lPush", "u_queue", i,fmt.Sprintf("i:%d", i))
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	for {
		uq, err := redis.String(conn.Do("lPop", "u_queue"))
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("lPop:", uq)
		time.Sleep(time.Second)
	}
	//expire (unit:second)
	_, err = conn.Do("Expire", "user_id", 3)
	for {
		uid, err := redis.Int(conn.Do("Get", "user_id"))
		if err != nil {
			if err == redis.ErrNil {
				fmt.Println("user_id expired")
				break
			}
			fmt.Println("Redis Execute Get Command Failed:", err)
		}
		fmt.Println("uid:", uid)
		time.Sleep(time.Second)
	}
}
