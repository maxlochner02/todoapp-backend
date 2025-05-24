package routes

import (
	"todo-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/todos", controllers.GetTodos)
	r.GET("/todos/:id", controllers.GetTodo)
	r.POST("/todos", controllers.CreateTodo)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)

	r.GET("/projects", controllers.GetProjects)
	r.POST("/projects", controllers.CreateProject)
	r.PUT("/projects/:id", controllers.UpdateProject)
	r.DELETE("/projects/:id", controllers.DeleteProject)

	r.GET("/categories", controllers.GetCategories)
	r.POST("/categories", controllers.CreateCategory)
	r.PUT("/categories/:id", controllers.UpdateCategory)
	r.DELETE("/categories/:id", controllers.DeleteCategory)
}
