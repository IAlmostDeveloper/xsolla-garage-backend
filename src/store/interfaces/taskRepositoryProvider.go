package interfaces

import "github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"

type TaskRepositoryProvider interface {
	Create(*dto.Task) error
	GetByID(int) (*dto.Task, error)
	GetAll() (*[]dto.Task, error)
}
