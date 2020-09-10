package interfaces

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
)

type TaskServiceProvider interface {
	CreateTask(task *dto.Task) error
	GetTasks() (*[]dto.Task, error)
	GetTaskByID(taskId int) (*dto.Task, error)
}
