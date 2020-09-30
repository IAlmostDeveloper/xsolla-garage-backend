package services

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/storage/interfaces"
)

type GoogleAuthService struct {
	storage interfaces.StorageProvider
}

func NewGoogleAuthService(storage interfaces.StorageProvider) *GoogleAuthService {
	return &GoogleAuthService{storage: storage}
}

func (s *GoogleAuthService) Authenticate() {

}

func (s *GoogleAuthService) Register(user *dto.User) error {
	return s.storage.UserRepository().Create(user)

}

// creates user in database if not exist
func (s *GoogleAuthService) ResolveUser(user *dto.User) error {
	storageUser, err := s.storage.UserRepository().GetById(user.Id)
	if err != nil {
		return err
	}
	if storageUser == nil {
		if err := s.Register(user); err != nil {
			return err
		}
	}
	return nil
}
