package services

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/store/interfaces"
)

type TaskService struct {
	store interfaces.StoreProvider
}

func (s *TaskService) GetTaskByID(taskId int) (*dto.Task, error) {
	return s.store.TaskRepository().GetByID(taskId)
}

func NewTaskService(store interfaces.StoreProvider) *TaskService {
	return &TaskService{store: store}
}

func (s *TaskService) CreateTask(task *dto.Task) error {
	return s.store.TaskRepository().Create(task)
}

func (s *TaskService) GetTasks() (*[]dto.Task, error) {
	return s.store.TaskRepository().GetAll()
}
