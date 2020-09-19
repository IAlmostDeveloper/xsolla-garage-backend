package interfaces

import "github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"

type TagServiceProvider interface {
	AddToTask(taskId int, tags []*dto.Tag) error
	RemoveFromTask(taskId int, tagId int) error
}
