package redis

import (
	"fmt"
	"time"

	// "github.com/garyburd/redigo/redis"
	"github.com/gomodule/redigo/redis"
)

var (
	pool      *redis.Pool
	redisHost = "10.10.30.244:6379"
)

// newRedisPool : 创建redis连接池
func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,
		MaxActive:   30,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			// 1. 打开连接
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
}

func init() {
	pool = newRedisPool()
	data, err := pool.Get().Do("KEYS", "*")
	fmt.Println(err)
	fmt.Println(data)
}

func RedisPool() *redis.Pool {
	return pool
}
