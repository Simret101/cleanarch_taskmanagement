package Repositories

import (
	"errors"

	"task/Domain"
)

type taskRepository struct {
	tasks  []Domain.Task
	lastID int
}

func NewTaskRepository() Domain.TaskRepository {
	return &taskRepository{tasks: []Domain.Task{}, lastID: 0}
}

func (r *taskRepository) GetAllTasks() ([]Domain.Task, error) {
	return r.tasks, nil
}

func (r *taskRepository) GetTaskByID(id int) (*Domain.Task, error) {
	for _, task := range r.tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}

func (r *taskRepository) CreateTask(task *Domain.Task) error {
	r.lastID++
	task.ID = r.lastID
	r.tasks = append(r.tasks, *task)
	return nil
}

func (r *taskRepository) UpdateTask(id int, updatedTask *Domain.Task) error {
	for i, task := range r.tasks {
		if task.ID == id {
			r.tasks[i] = *updatedTask
			r.tasks[i].ID = id
			return nil
		}
	}
	return errors.New("task not found")
}

func (r *taskRepository) DeleteTask(id int) error {
	for i, task := range r.tasks {
		if task.ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
