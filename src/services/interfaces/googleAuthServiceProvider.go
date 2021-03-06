package interfaces

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"time"
)

type GoogleAuthServiceProvider interface {
	GetUser(id string) (*dto.User, error)
	Authorize(accessToken string) (string, error)
	Authenticate(user *dto.User) (string, error)
	ResolveUser(user *dto.User) error
	GetAccessTokenTTL() time.Duration
}
