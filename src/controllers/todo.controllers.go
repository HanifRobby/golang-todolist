package controllers

import (
	"net/http"

	"golang-todolist/src/config"
	"golang-todolist/src/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Define database client
var db *gorm.DB = config.ConnectDB()

// Todo struct for request body
type todoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Defining struct for response
type todoResponse struct {
	todoRequest
	ID uint `json:"id"`
}

// Create todo data to database by run this function
func CreateTodo(context *gin.Context) {
	var data todoRequest

	// Binding request body json to request body struct
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Marching todo models struct with todo request struct
	todo := models.Todo{}
	todo.Name = data.Name
	todo.Description = data.Description

	// Querying to database
	result := db.Create(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	// Matching result to create response
	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	// Creating http response
	context.JSON(http.StatusCreated, response)

}

// Getting all todo datas
func GetAllTodos(context *gin.Context) {
	var todos []models.Todo

	//Querying to find todo datas.
	err := db.Find(&todos)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	// Creating http response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    todos,
	})
}
