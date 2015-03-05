package shortener

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type RedisApi struct {
	redisConn redis.Conn
}

func NewRedisApi(ip string, port string) *RedisApi {
	if conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", ip, port)); err != nil {
		panic(err)
	} else {
		return &RedisApi{conn}
	}
}

func (api *RedisApi) urlHash() string {
	if key, err := api.redisConn.Do("INCR", "url.pointer"); err != nil {
		panic(err)
	} else {
		return fmt.Sprintf("%x", key)
	}
}

func (api *RedisApi) CreateRecord(url string) string {
	hash := api.urlHash()
	if _, err := api.redisConn.Do("SET", hash, url); err != nil {
		panic(err)
	}
	return hash
}

func (api *RedisApi) GetUrl(hash string) string {
	if url, err := api.redisConn.Do("GET", hash); err != nil {
		panic(err)
	} else {
		return fmt.Sprintf("%s", url)
	}
}
