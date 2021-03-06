package interfaces

import "github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"

type TaskRepositoryProvider interface {
	Create(*dto.Task) error
	GetByID(int) (*dto.Task, error)
	GetByUserID(userId string) ([]*dto.Task, error)
	GetAll() ([]*dto.Task, error)
	RemoveByID(int) error
	Update(*dto.Task) error
}
