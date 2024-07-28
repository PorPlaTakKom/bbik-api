package routes

import (
	"todo-list/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// Group v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("/todos", controllers.GetAllTodos)
		v1.POST("/todos", controllers.CreateTodo)
		v1.PUT("/todos/:id", controllers.UpdateTodoByID)
		v1.DELETE("/todos/:id", controllers.DeleteTodoByID)
	}

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello")
	})
}
