package interfaces

import "github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"

type FeedbackServiceProvider interface {
	AddFeedback(feedback *dto.Feedback) error
	GetAllFeedback() ([]*dto.Feedback, error)
}
