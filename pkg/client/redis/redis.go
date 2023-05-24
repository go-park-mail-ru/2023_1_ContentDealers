package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	maxIdle     = 4
	idleTimeout = 120 * time.Second
)

func NewClientRedis(cfg RedisConfig) (*redis.Pool, error) {
	dsn := fmt.Sprintf("redis://%s:@%s:%s/%s", cfg.User, cfg.Host, cfg.Port, cfg.DBNum)

	redisPool := &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: idleTimeout,
		DialContext: func(ctx context.Context) (redis.Conn, error) {
			c, err := redis.DialURLContext(ctx, dsn)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
	return redisPool, nil
}
