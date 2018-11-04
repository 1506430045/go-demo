package apis

import (
	"demo/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func IndexApi(c *gin.Context) () {
	c.String(http.StatusOK, "Hello World")
}

func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	p := models.Person{
		FirstName: firstName,
		LastName:  lastName,
	}

	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetPersonsApi(c *gin.Context) {
	p := models.Person{}
	persons := p.GetPersons()

	c.JSON(http.StatusOK, gin.H{
		"person": persons,
	})
}

func GetPersonApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Fatalln(err)
	}

	p := models.Person{}
	p.Id = id

	p.GetPerson()
	c.JSON(http.StatusOK, gin.H{
		"person": p,
	})
}

func ModPersonApi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	p := models.Person{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
	}

	re := p.ModPerson()

	if re {
		c.JSON(http.StatusOK, gin.H{
			"msg": "mod success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "mod failed",
		})
	}
}

func DelPersonApi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	p := models.Person{
		Id: id,
	}

	re := p.DelPerson()
	if re {
		c.JSON(http.StatusOK, gin.H{
			"msg": "del success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "del failed",
		})
	}
}
