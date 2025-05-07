package routes

import (
	"todo-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Todo
	r.GET("/todos", controllers.GetTodos)
	r.GET("/todos/:id", controllers.GetTodo)
	r.POST("/todos", controllers.CreateTodo)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)

	// Projects
	r.GET("/projects", controllers.GetProjects)
	r.POST("/projects", controllers.CreateProject)

	// Categories
	r.GET("/categories", controllers.GetCategories)
	r.POST("/categories", controllers.CreateCategory)
}
