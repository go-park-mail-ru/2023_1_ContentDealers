package redis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func NewClientRedis(cfg RedisConfig) (*redis.Pool, error) {
	dsn := fmt.Sprintf("redis://%s:@%s:%s/%s", cfg.User, cfg.Host, cfg.Port, cfg.DBNum)

	var redisPool *redis.Pool

	redisPool = &redis.Pool{
		MaxIdle:     4,
		IdleTimeout: 120 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(dsn)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
	return redisPool, nil
}
