package services

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/storage/interfaces"
)

type TagService struct {
	storage interfaces.StorageProvider
}

func NewTagService(storage interfaces.StorageProvider) *TagService {
	return &TagService{storage: storage}
}

func (s TagService) AddToTask(taskId int, tags []*dto.Tag) error {
	for _, tag := range tags {
		if err := s.storage.TagRepository().AddToTask(taskId, tag); err != nil {
			return err
		}
	}
	return nil
}

func (s TagService) RemoveFromTask(taskId int, tagId int) error {
	return s.storage.TagRepository().RemoveFromTask(taskId, tagId)
}
