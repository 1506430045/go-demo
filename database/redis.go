package database

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

var RedisDB redis.Conn

func init() {
	var err error
	RedisDB, err = redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		log.Println(err.Error())
	}
	ret, err := RedisDB.Do("ping")
	if ret != "PONG" || err != nil {
		log.Println(err)
	}
}
