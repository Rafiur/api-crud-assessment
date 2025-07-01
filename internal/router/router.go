package router

import (
	"github.com/Rafiur/api-crud-assessment/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handler.TaskHandler) {
	tasks := r.Group("/tasks")
	{
		tasks.POST("/", h.CreateTask)
		tasks.GET("/:id", h.GetByIDTask)
		tasks.GET("/list", h.ListTask)
		tasks.PUT("/:id", h.UpdateTask)
		tasks.DELETE("/:id", h.DeleteTask)
	}
}
