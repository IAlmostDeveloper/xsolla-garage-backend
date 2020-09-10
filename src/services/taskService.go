package services

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/storage/interfaces"
)

type TaskService struct {
	storage interfaces.StorageProvider
}

func (s *TaskService) GetTaskByID(taskId int) (*dto.Task, error) {
	return s.storage.TaskRepository().GetByID(taskId)
}

func NewTaskService(storage interfaces.StorageProvider) *TaskService {
	return &TaskService{storage: storage}
}

func (s *TaskService) CreateTask(task *dto.Task) error {
	return s.storage.TaskRepository().Create(task)
}

func (s *TaskService) GetTasks() (*[]dto.Task, error) {
	return s.storage.TaskRepository().GetAll()
}

func (s *TaskService) RemoveByID(taskId int) error {
	return s.storage.TaskRepository().RemoveByID(taskId)
}
