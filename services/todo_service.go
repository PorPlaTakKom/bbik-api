package services

import (
	"todo-list/models"
	"todo-list/repositories"
)

func GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo
	err := repositories.GetAllTodos(&todos)
	return todos, err
}

func CreateTodo(todo *models.Todo) error {
	return repositories.CreateTodo(todo)
}

func GetTodoByID(id string) (models.Todo, error) {
	var todo models.Todo
	err := repositories.GetTodoByID(&todo, id)
	return todo, err
}

func UpdateTodoByID(id string, todo *models.Todo) error {
	return repositories.UpdateTodoByID(todo, id)
}

func DeleteTodoByID(id string) error {
	return repositories.DeleteTodoByID(id)
}
