package main

import (
	"net/http"
	"strconv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"

	_ "app/docs"
)

type template struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var templates_test_data = []template{
	{Id: 1, Name: "Blue Train"},
	{Id: 2, Name: "Green Train"},
	{Id: 3, Name: "Red Train"},
}

// @title           TSaaS API
// @version         1.0
// @description     The TSaaS is Typesetting as a Service.

// @contact.name   ffh4ck64
// @contact.email  ctipax@gmail.com

// @host      localhost:8080
// @BasePath  /
func main() {
	router := gin.Default()

	router.GET("/templates", getTemplates)
	router.GET("/templates/:id", getAlbumByID)
	router.POST("/templates", postAlbums)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}

func getTemplates(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, templates_test_data)
}

func postAlbums(c *gin.Context) {
	var newTemplate template

	if err := c.BindJSON(&newTemplate); err != nil {
		return
	}

	templates_test_data = append(templates_test_data, newTemplate)
	c.IndentedJSON(http.StatusCreated, newTemplate)
}

func getAlbumByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	for _, a := range templates_test_data {
		if a.Id == int(id) {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "template not found"})
}
