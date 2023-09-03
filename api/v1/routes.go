package v1

import "github.com/gin-gonic/gin"

func InitializeRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.POST("/tasks", CreateTask)
		api.GET("/tasks", GetAllTasks)
		api.GET("/tasks/:id", GetTaskByID)
		api.PUT("/tasks/:id", UpdateTask)
		api.DELETE("/tasks/:id", DeleteTask)
	}
}
