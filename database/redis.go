package database

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

const (
	REDIS_HOST = "127.0.0.1"
	REDIS_PORT = 6379
)

var RedisDB redis.Conn

func init() {
	var err error
	RedisDB, err = redis.Dial("tcp", fmt.Sprintf("%s:%d", REDIS_HOST, REDIS_PORT))

	if err != nil {
		log.Println(err.Error())
	}
	ret, err := RedisDB.Do("ping")
	if ret != "PONG" || err != nil {
		log.Println(err)
	}
}
