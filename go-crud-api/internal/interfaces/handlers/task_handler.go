package handlers

import (
	"go-crud-api/internal/entities"
	"go-crud-api/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskHandler struct {
	usecase usecases.TaskUseCase
}

func NewTaskHandler(usecase usecases.TaskUseCase) *TaskHandler {
	return &TaskHandler{usecase: usecase}
}

// @Summary 		Create a new task
// @Description Create a new task
// @Tags 				tasks
// @Accept 			json
// @Produce 		json
// @Param 			task 		body 			entities.Task 						true "Task data"
// @Success 		201 		{object} 	map[string]interface{}
// @Failure 		400 		{object} 	map[string]string
// @Failure 		500 		{object} 	map[string]string
// @Router 			/tasks 	[post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task entities.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.usecase.CreateTask(c.Request.Context(), &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id.Hex()})
}

// @Summary 		List all tasks
// @Description List all tasks
// @Tags 				tasks
// @Produce 		json
// @Success 		200 		{array} 	entities.Task
// @Failure 		500 		{object} 	map[string]string
// @Router 			/tasks 	[get]
func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.usecase.GetTasks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// @Summary 		Update a task
// @Description Update a task by ID
// @Tags 				tasks
// @Accept 			json
// @Produce 		json
// @Param 			id 		path 			string 						true "Task ID"
// @Param 			task 	body 			entities.Task 		true "Updated task data"
// @Success 		200 	{object} 	map[string]string
// @Failure 		400 	{object} 	map[string]string
// @Failure 		404 	{object} 	map[string]string
// @Failure 		500 	{object} 	map[string]string
// @Router 			/tasks/{id} 		[put]
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task entities.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.UpdateTask(c.Request.Context(), id, &task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

// @Summary 		Delete a task
// @Description Delete a task by ID
// @Tags 				tasks
// @Produce 		json
// @Param 			id 		path 			string 						true "Task ID"
// @Success 		200 	{object} 	map[string]string
// @Failure 		400 	{object} 	map[string]string
// @Failure 		404 	{object} 	map[string]string
// @Failure 		500 	{object} 	map[string]string
// @Router 			/tasks/{id} 		[delete]
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := h.usecase.DeleteTask(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}