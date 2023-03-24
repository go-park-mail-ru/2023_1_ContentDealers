package setup

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

const redisAddr = "redis://user:@localhost:6379/0"

func NewClientRedis() (redis.Conn, error) {
	redisConn, err := redis.DialURL(redisAddr)
	if err != nil {
		return nil, fmt.Errorf("cant connect to redis: %w", err)
	}
	return redisConn, nil
}
