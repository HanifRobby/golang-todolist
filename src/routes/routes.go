package routes

import (
	"golang-todolist/src/controllers"

	"github.com/gin-gonic/gin"
)

// Routes function to serve endpoints
func Routes() {
	route := gin.Default()

	route.POST("/todo", controllers.CreateTodo)
	route.GET("/todo", controllers.GetAllTodos)
	route.PUT("/todo", controllers.UpdateTodo)
	route.DELETE("/todo", controllers.DeleteTodo)

	// Run route whenever triggered
	route.Run()
}
