package Domain

type Task struct {
	ID          int
	Title       string
	Description string
	DueDate     string
	Status      string
}
type TaskRepository interface {
	GetAllTasks() ([]Task, error)

	GetTaskByID(id int) (*Task, error)

	CreateTask(task *Task) error

	UpdateTask(id int, updatedTask *Task) error

	DeleteTask(id int) error
}
