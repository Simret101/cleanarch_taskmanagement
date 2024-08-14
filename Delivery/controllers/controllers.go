package controllers

import (
	"net/http"
	"strconv"
	"task/Domain"
	"task/Usecases"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUseCase Usecases.TaskUseCase
}

func (c *TaskController) GetAllTasks(ctx *gin.Context) {

	tasks, err := c.TaskUseCase.GetAllTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskController) CreateTask(ctx *gin.Context) {

	var task Domain.Task
	if err := ctx.BindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := c.TaskUseCase.CreateTask(&task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, task)
}

func (c *TaskController) GetTaskByID(ctx *gin.Context) {

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	task, err := c.TaskUseCase.GetTaskByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if task == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task Domain.Task
	if err := ctx.BindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	task.ID = id

	if err := c.TaskUseCase.UpdateTask(id, &task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	if err := c.TaskUseCase.DeleteTask(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

type UserController struct {
	UserUseCase Usecases.UserUseCase
}

func (u *UserController) Register(ctx *gin.Context) {

	var user Domain.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := u.UserUseCase.Register(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (u *UserController) Login(ctx *gin.Context) {

	var credentials Domain.Credentials
	if err := ctx.BindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	token, err := u.UserUseCase.Login(&credentials)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
