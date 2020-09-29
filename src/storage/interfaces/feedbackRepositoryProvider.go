package interfaces

import "github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"

type FeedbackRepositoryProvider interface {
	AddFeedback(feedback *dto.Feedback) error
	GetAllFeedback() ([]*dto.Feedback, error)
}
