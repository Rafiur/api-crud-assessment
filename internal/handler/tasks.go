package handler

import (
	"github.com/Rafiur/api-crud-assessment/internal/model"
	"github.com/Rafiur/api-crud-assessment/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	service usecase.TaskService
}

func NewTaskHandler(service usecase.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	res, err := h.service.Create(&task)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, res)
}
func (h *TaskHandler) GetByIDTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := h.service.GetById(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, task)
}

func (h *TaskHandler) ListTask(c *gin.Context) {
	tasks, err := h.service.List()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tasks)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	task.ID = id
	res, err := h.service.Update(&task)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	c.JSON(200, res)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Task deleted"})
}
