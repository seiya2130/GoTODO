package v1

import "github.com/gin-gonic/gin"

func InitializeRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.POST("/tasks", CreateTask)
	}
}
