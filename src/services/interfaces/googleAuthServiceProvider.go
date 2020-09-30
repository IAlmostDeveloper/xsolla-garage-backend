package interfaces

import "github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"

type GoogleAuthServiceProvider interface {
	Authenticate()
	Register()
	ResolveUser(user *dto.User) error
}
