package setup

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// нужен пароль
const redisAddr = "redis://user:@localhost:6379/0"

func NewClientRedis() (*redis.Pool, error) {
	var redisPool *redis.Pool

	redisPool = &redis.Pool{
		MaxIdle:     4,
		IdleTimeout: 120 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(redisAddr)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}

	return redisPool, nil
}
