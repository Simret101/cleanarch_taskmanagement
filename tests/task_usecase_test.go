package tests

import (
	"task/Domain"
	"task/Repositories"
	"task/Usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskUseCase(t *testing.T) {
	taskRepo := Repositories.NewTaskRepository()
	taskUseCase := Usecases.TaskUseCase{TaskRepo: taskRepo}

	task := &Domain.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     "2024-08-09",
		Status:      "pending",
	}
	err := taskUseCase.CreateTask(task)
	assert.NoError(t, err)
	assert.Equal(t, 1, task.ID)

	tasks, err := taskUseCase.GetAllTasks()
	assert.NoError(t, err)
	assert.Len(t, tasks, 1)

	retrievedTask, err := taskUseCase.GetTaskByID(1)
	assert.NoError(t, err)
	assert.Equal(t, task.Title, retrievedTask.Title)

	task.Title = "Updated Task"
	err = taskUseCase.UpdateTask(1, task)
	assert.NoError(t, err)
	updatedTask, err := taskUseCase.GetTaskByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Task", updatedTask.Title)

	err = taskUseCase.DeleteTask(1)
	assert.NoError(t, err)
	deletedTask, err := taskUseCase.GetTaskByID(1)
	assert.Nil(t, deletedTask)
	assert.Error(t, err)
}
