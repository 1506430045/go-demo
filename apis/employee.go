package apis

import (
	"demo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetEmployeeApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Fatalln(err)
	}

	e := models.Employee{}
	e.Id = id

	e.GetEmployee()
	c.JSON(http.StatusOK, gin.H{
		"employee": e,
	})
}
