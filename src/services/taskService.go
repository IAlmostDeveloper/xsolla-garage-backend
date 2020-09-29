package services

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/storage/interfaces"
	"strings"
)

type TaskService struct {
	storage interfaces.StorageProvider
}

func NewTaskService(storage interfaces.StorageProvider) *TaskService {
	return &TaskService{storage: storage}
}

func (s *TaskService) GetTaskByID(taskId int) (*dto.Task, error) {
	task, err := s.storage.TaskRepository().GetByID(taskId)
	if err != nil {
		return nil, err
	}
	task.Tags, err = s.storage.TagRepository().GetByTaskId(taskId)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) CreateTask(task *dto.Task) error {
	if err := s.storage.TaskRepository().Create(task); err != nil {
		return err
	}
	for _, tag := range task.Tags {
		tag.Name = strings.ToLower(strings.Trim(tag.Name, " "))
	}
	for _, tag := range task.Tags {
		if err := s.storage.TagRepository().AddToTask(task.Id, tag); err != nil {
			return err
		}
	}
	return nil
}

func (s *TaskService) GetTasks() ([]*dto.Task, error) {
	tasks, err := s.storage.TaskRepository().GetAll()
	if err != nil {
		return nil, err
	}
	for _, task := range tasks {
		task.Tags, err = s.storage.TagRepository().GetByTaskId(task.Id)
		if err != nil {
			return nil, err
		}
	}
	return tasks, nil
}

func (s *TaskService) RemoveByID(taskId int) error {
	return s.storage.TaskRepository().RemoveByID(taskId)
}

func (s *TaskService) Update(task *dto.Task) error {
	return s.storage.TaskRepository().Update(task)
}
