// repositories/todo_repository.go
package repositories

import (
	"fmt"
	"todo-list/config"
	"todo-list/models"
)

func GetAllTodos(todos *[]models.Todo) error {
	if err := config.DB.Find(todos).Error; err != nil {
		return err
	}
	return nil
}

func CreateTodo(todo *models.Todo) error {
	if err := config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetTodoByID(todo *models.Todo, id string) error {
	if err := config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodoByID(todo *models.Todo, id string) error {
	// Debug: Log values before updating
	fmt.Printf("Updating Todo ID %s: %#v\n", id, todo)

	// ตรวจสอบค่าของฟิลด์ก่อนการอัปเดต
	if todo.Completed != "true" && todo.Completed != "false" {
		return fmt.Errorf("invalid value for completed: %s", todo.Completed)
	}

	if err := config.DB.Model(todo).Where("id = ?", id).Updates(todo).Error; err != nil {
		return err
	}

	// Debug: Confirm update success
	fmt.Printf("Successfully updated Todo ID %s: %#v\n", id, todo)
	return nil
}

func DeleteTodoByID(id string) error {
	if err := config.DB.Delete(&models.Todo{}, id).Error; err != nil {
		return err
	}
	return nil
}
