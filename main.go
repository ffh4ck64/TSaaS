package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ffh4ck64/TSaaS/docs"
)

// @template
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
	router.GET("/templates/:id", getTemplateByID)
	router.POST("/templates", postTemplate)
	router.GET("/document", getDocument)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}

// getTemplates responds with the list of all books as JSON.
// getTemplates     godoc
// @Summary         Get an array of template
// @Description     Responds with the list of all templates as JSON.
// @Tags            templates
// @Produce         json
// @Success         200  {array}  template
// @Router          /templates [get]
func getTemplates(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, templates_test_data)
}

// postTemplate		responds with the list of all books as JSON.
// postTemplate		godoc
// @Summary         Get an array of template
// @Description     Responds with the list of all templates as JSON.
// @Tags            templates
// @Produce         json
// @Success         200  {array}  template
// @Router          /template [post]
func postTemplate(c *gin.Context) {
	var newTemplate template

	if err := c.BindJSON(&newTemplate); err != nil {
		return
	}

	templates_test_data = append(templates_test_data, newTemplate)
	c.IndentedJSON(http.StatusCreated, newTemplate)
}

// getTemplateByID responds a template by id
// getTemplateByID	godoc
// @Summary         Get an name of template by id
// @Tags            template
// @Produce         json
// @Success         200  {object}  template
// @Failure			400 {string} string "We need ID!!"
// @Failure			404 {string} string "Can not find ID"
// @Router          /template/{id} [get]
func getTemplateByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "We need ID!!")
		return
	}

	for _, a := range templates_test_data {
		if a.Id == int(id) {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, "Can not find ID")
}

// getDocument get generated document by params
// getDocument		godoc
// @Summary       	Get generated document by params
// @Success         200
// @Param			any_query_param query string false "Произвольный query-параметр" style=form explode=true
// @Produce			application/pdf
// @Router          /document [get]
func getDocument(c *gin.Context) {
	// Получаем все query параметры в виде map[string][]string
	params := c.Request.URL.Query()

	// Преобразуем в map[string]string (берем только первое значение каждого параметра)
	simpleParams := make(map[string]string)
	for key, values := range params {
		if len(values) > 0 {
			simpleParams[key] = values[0]
		}
	}

	// Передаём параметры в другую функцию
	result := put(simpleParams)

	// Возвращаем результат
	c.JSON(http.StatusOK, gin.H{
		"received": simpleParams,
		"result":   result,
	})
}
