package interfaces

import "github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"

type UserRepositoryProvider interface {
	Create(user *dto.User) error
	GetById(id string) (*dto.User, error)
	RemoveById(id string) error
}
