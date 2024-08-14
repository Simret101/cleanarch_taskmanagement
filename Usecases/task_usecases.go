package Usecases

import (
	"task/Domain"
)

type TaskUseCase struct {
	TaskRepo Domain.TaskRepository
}

func (uc *TaskUseCase) GetAllTasks() ([]Domain.Task, error) {
	return uc.TaskRepo.GetAllTasks()
}

func (uc *TaskUseCase) GetTaskByID(id int) (*Domain.Task, error) {
	return uc.TaskRepo.GetTaskByID(id)
}

func (uc *TaskUseCase) CreateTask(task *Domain.Task) error {
	return uc.TaskRepo.CreateTask(task)
}

func (uc *TaskUseCase) UpdateTask(id int, updatedTask *Domain.Task) error {
	return uc.TaskRepo.UpdateTask(id, updatedTask)
}

func (uc *TaskUseCase) DeleteTask(id int) error {
	return uc.TaskRepo.DeleteTask(id)
}
