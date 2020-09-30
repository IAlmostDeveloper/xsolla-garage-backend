package interfaces

import "github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"

type GoogleAuthServiceProvider interface {
	Authenticate()
	Register(user *dto.User) error
	ResolveUser(user *dto.User) error
}
