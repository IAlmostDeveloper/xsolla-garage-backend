package interfaces

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"time"
)

type GoogleAuthServiceProvider interface {
	Authenticate(accessToken string) (string, error)
	LoginUser(user *dto.User) (string, error)
	ResolveUser(user *dto.User) error
	GetAccessTokenTTL() time.Duration
}
