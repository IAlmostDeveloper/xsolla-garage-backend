package interfaces

import "github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"

type TagRepositoryProvider interface {
	AddToTask(taskID int, tag *dto.Tag) error
	RemoveFromTask(taskId int, tagId int) error
	GetByTaskId(taskId int) ([]*dto.Tag, error)
}
