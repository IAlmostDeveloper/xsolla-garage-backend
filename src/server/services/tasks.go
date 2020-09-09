package services

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/store/interfaces"
)

var store interfaces.StoreProvider

func CreateTask(task *dto.Task) error{
	return store.TaskRepository().Create(task)
}

func GetTasks() (*[]dto.Task, error){
	return store.TaskRepository().GetAll()
}

func GetTask(taskId int) (*dto.Task, error){
	return store.TaskRepository().GetByID(taskId)
}
