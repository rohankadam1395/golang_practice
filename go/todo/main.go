package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var todos = []todo{
	{ID: "1", Title: "First todo"},
	{ID: "2", Title: "Second todo"},
}

func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.StaticFile("/", "./static/index.html")
	router.GET("/todos", getTodos)
	router.POST("/todos", postTodos)

	router.Run("localhost:8080")
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func postTodos(c *gin.Context) {
	var newTodo todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}
