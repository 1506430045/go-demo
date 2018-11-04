package main

import (
	. "demo/apis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", IndexApi)

	router.POST("/person", AddPersonApi)

	router.GET("/persons", GetPersonsApi)

	router.GET("/person/:id", GetPersonApi)

	router.PUT("/person/:id", ModPersonApi)

	router.DELETE("/person/:id", DelPersonApi)

	router.PUT("/user/:name", SetNameApi)

	router.GET("/user/:name", GetNameApi)

	router.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "hello.html", gin.H{
			"title": "Main website",
		})
	})

	router.NoRoute(go404)

	return router
}

func go404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "hello.html", gin.H{})
}
