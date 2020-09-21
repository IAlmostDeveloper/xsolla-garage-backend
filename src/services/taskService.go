package services

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/storage/interfaces"
)

type TaskService struct {
	storage interfaces.StorageProvider
}

func NewTaskService(storage interfaces.StorageProvider) *TaskService {
	return &TaskService{storage: storage}
}

func (s *TaskService) GetTaskByID(taskId int) (*dto.Task, error) {
	if task, err := s.storage.TaskRepository().GetByID(taskId); err != nil {
		return nil, err
	} else{
		if task.Tags, err = s.storage.TagRepository().GetByTaskId(taskId) ; err != nil {
			return nil, err
		}
		return task, nil
	}
}

func (s *TaskService) CreateTask(task *dto.Task) error {
	if err := s.storage.TaskRepository().Create(task); err != nil {
		return err
	}
	for _, tag := range task.Tags {
		if err := s.storage.TagRepository().AddToTask(task.Id, tag); err != nil {
			return err
		}
	}
	return nil
}

func (s *TaskService) GetTasks() ([]*dto.Task, error) {
	if tasks, err := s.storage.TaskRepository().GetAll() ; err != nil {
		return nil, err
	} else {
		for _, task := range tasks {
			task.Tags, err = s.storage.TagRepository().GetByTaskId(task.Id)
			if err != nil {
				return nil, err
			}
		}
		return tasks, nil
	}
}

func (s *TaskService) RemoveByID(taskId int) error {
	return s.storage.TaskRepository().RemoveByID(taskId)
}

func (s *TaskService) Update(task *dto.Task) error {
	return s.storage.TaskRepository().Update(task)
}
