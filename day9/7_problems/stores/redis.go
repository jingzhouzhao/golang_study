package stores

import (
	"time"
	"github.com/gomodule/redigo/redis"
	"go_dev/day9/7_problems/config"
	
)

var redisPool *redis.Pool

func init(){
	getRedisPool(config.Redis_Max_Active, config.Redis_Max_Idle, config.Redis_Idle_Timeout, config.Redis_Address)
}

func getRedisPool(maxActive,maxIdle int,idleTimeout time.Duration,address string,options ...redis.DialOption){
	redisPool = &redis.Pool{
		MaxActive:maxActive,
		MaxIdle:maxIdle,
		IdleTimeout:idleTimeout,
		Dial: func() (redis.Conn, error){
			return redis.Dial("tcp", address, options...)
		},
	}
}

func GetConn()(redis.Conn){
	return redisPool.Get()
}

func HGet(conn redis.Conn,table,key string)(interface{},error){
	return conn.Do("HGet", table,key)
}

func HSet(conn redis.Conn,table,key,value string)(interface{},error){
	return conn.Do("HSet", table,key,value)
}