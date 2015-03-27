package shortener

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisApi struct {
	pool *redis.Pool
}

func getPool(server string, password string) (pool *redis.Pool) {
	pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, _ time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return pool
}

func NewRedisApi(server, password string) *RedisApi {
	return &RedisApi{
		getPool(server, password),
	}
}

func (api *RedisApi) urlHash() string {
	return RandomString(7)
}

func (api *RedisApi) CreateRecord(url string) (string, error) {
	conn := api.pool.Get()
	defer conn.Close()
	hash := api.urlHash()
	if err := conn.Send("SETEX", hash, 86400, url); err != nil {
		return "", err
	}
	return hash, nil
}

func (api *RedisApi) GetUrl(hash string) (string, error) {
	conn := api.pool.Get()
	defer conn.Close()
	if url, err := redis.String(conn.Do("GET", hash)); err != nil {
		if err == redis.ErrNil {
			err = errors.New("URL was not found")
		}
		return "", err
	} else {
		return fmt.Sprintf("%s", url), nil
	}
}
