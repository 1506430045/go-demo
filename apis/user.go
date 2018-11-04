package apis

import (
	red "demo/database"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SetNameApi(c *gin.Context) {

	username := c.Param("name")
	_, err := red.RedisDB.Do("SET", "username", username)

	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
	})
}

func GetNameApi(c *gin.Context) {
	username, err := redis.String(red.RedisDB.Do("GET", "username"))
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    username,
	})
}
