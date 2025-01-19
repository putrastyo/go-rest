package handler

import (
	"net/http"
	"praktik-todo/internal/entity"
	"praktik-todo/internal/usecase"
	"praktik-todo/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type TaskHandler struct {
	taskUsecase usecase.TaskUsecase
	validator   *validator.Validate
	log         *logrus.Logger
}

func NewTaskHandler(taskUsecase usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{
		taskUsecase: taskUsecase,
		validator:   validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task entity.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validation
	if err := h.validator.Struct(task); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	if err := h.taskUsecase.CreateTask(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.log.Infof("Task created: %v", task)
	c.JSON(http.StatusCreated, gin.H{"task": task})
}

func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.taskUsecase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": tasks})
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing id"})
		return
	}
	task, err := h.taskUsecase.GetTaskByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	var task entity.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		utils.HandleValidationError(c, err)
		return
	}
	if err := h.validator.Struct(task); err != nil {
		utils.HandleValidationError(c, err)
		return
	}
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing id"})
		return
	}
	task.ID = uint(id)
	if err := h.taskUsecase.UpdateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing id"})
		return
	}
	if err := h.taskUsecase.DeleteTask(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}
