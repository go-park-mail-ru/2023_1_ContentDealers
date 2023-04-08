package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func NewClientRedis(cfg RedisConfig) (redis.Conn, error) {
	dsn := fmt.Sprintf("redis://%s:@%s:%s/%s", cfg.User, cfg.Host, cfg.Port, cfg.DBNum)
	redisConn, err := redis.DialURL(dsn)
	if err != nil {
		return nil, fmt.Errorf("cant connect to redis: %w", err)
	}
	return redisConn, nil
}
