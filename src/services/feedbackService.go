package services

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/storage/interfaces"
)

type FeedbackService struct {
	storage interfaces.StorageProvider
}

func NewFeedbackService(storage interfaces.StorageProvider) *FeedbackService{
	return &FeedbackService{storage: storage}
}

func (f *FeedbackService) AddFeedback(feedback *dto.Feedback) error {
	return f.storage.FeedbackRepository().AddFeedback(feedback)
}

func (f *FeedbackService) GetAllFeedback() ([]*dto.Feedback, error){
	return f.storage.FeedbackRepository().GetAllFeedback()
}
