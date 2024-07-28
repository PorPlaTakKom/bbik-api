package controllers

import (
	"log"
	"net/http"
	"todo-list/models"
	"todo-list/services"
	"todo-list/utils"

	"github.com/gin-gonic/gin"
)

func GetAllTodos(c *gin.Context) {
	todos, err := services.GetAllTodos()
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(c, http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := services.CreateTodo(&todo); err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(c, http.StatusCreated, todo)
}

func GetTodoByID(c *gin.Context) {
	id := c.Param("id")
	todo, err := services.GetTodoByID(id)
	if err != nil {
		utils.RespondJSON(c, http.StatusNotFound, err.Error())
		return
	}
	utils.RespondJSON(c, http.StatusOK, todo)
}

func UpdateTodoByID(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	// Debug: Log the received todo for update
	log.Printf("Received update request for ID %s: %#v", id, todo)

	if err := services.UpdateTodoByID(id, &todo); err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Debug: Log the successfully updated todo
	log.Printf("Successfully updated todo with ID %s: %#v", id, todo)

	utils.RespondJSON(c, http.StatusOK, todo)
}

func DeleteTodoByID(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteTodoByID(id); err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(c, http.StatusOK, "Todo deleted successfully")
}
