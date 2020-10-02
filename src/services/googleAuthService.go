package services

import (
	"context"
	"errors"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/storage/interfaces"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

type GoogleAuthService struct {
	storage        interfaces.StorageProvider
	accessSecret   string
	AccessTokenTTL time.Duration
	redis          *redis.Client
}

func NewGoogleAuthService(storage interfaces.StorageProvider, redis *redis.Client) *GoogleAuthService {
	return &GoogleAuthService{
		storage:        storage,
		accessSecret:   os.Getenv("ACCESS_SECRET"),
		AccessTokenTTL: time.Hour * 24,
		redis:          redis,
	}
}

func (s *GoogleAuthService) GetAccessTokenTTL() time.Duration {
	return s.AccessTokenTTL
}

func (s *GoogleAuthService) Authenticate(accessToken string) (string, error) {
	claims := jwt.MapClaims{}
	if _, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.accessSecret), nil
	}); err != nil {
		return "", err
	}

	if result := s.redis.Get(context.Background(), accessToken); result.Val() == claims["user_id"].(string) {
		return "", errors.New("stored user id and token user id do not match")
	}
	return claims["user_id"].(string), nil
}

func (s *GoogleAuthService) register(user *dto.User) error {
	return s.storage.UserRepository().Create(user)
}

// creates user in database if not exist
func (s *GoogleAuthService) ResolveUser(user *dto.User) error {
	storageUser, err := s.storage.UserRepository().GetById(user.Id)
	if err != nil {
		return err
	}
	if storageUser == nil {
		if err := s.register(user); err != nil {
			return err
		}
	}
	return nil
}

func (s *GoogleAuthService) LoginUser(user *dto.User) (string, error) {
	accessToken, err := s.createAccessToken(user)
	if err != nil {
		return "", err
	}
	if err := s.saveToken(user, accessToken); err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *GoogleAuthService) createAccessToken(user *dto.User) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = user.Id
	atClaims["exp"] = time.Now().Add(s.AccessTokenTTL).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err := at.SignedString([]byte(s.accessSecret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *GoogleAuthService) saveToken(user *dto.User, accessToken string) error {
	s.redis.Set(context.Background(), accessToken, user.Id, s.AccessTokenTTL)
	res := s.redis.Save(context.Background())
	return res.Err()
}
