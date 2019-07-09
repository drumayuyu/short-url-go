package infrastructure

import (
	"github.com/garyburd/redigo/redis"
)

type RedisClient struct {
}

func NewRedisClient() RedisClient {
	return RedisClient{}
}

const constantURL = "https://yahoo.co.jp"

func (*RedisClient) RedisGet(key string) (string, error) {

	c, err := redisConnection()
	if err != nil {
		return string(""), err
	}

	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		return constantURL, err
	}
	return s, nil
}

func redisConnection() (redis.Conn, error) {
	const ipPort = "localhost:6379"

	//redisに接続
	c, err := redis.Dial("tcp", ipPort)
	if err != nil {
		return nil, err
	}
	return c, nil
}
