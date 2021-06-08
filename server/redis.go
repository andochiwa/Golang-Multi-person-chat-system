package main

import "github.com/gomodule/redigo/redis"

var RedisPool *redis.Pool

func initPool(address string) {
	RedisPool = &redis.Pool{
		MaxActive:   0,
		MaxIdle:     8,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
