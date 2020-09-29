package interfaces

import "github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"

type ValidationServiceProvider interface {
	ValidateTask(task *dto.Task) error
	ValidateTags(tags []*dto.Tag) error
	ValidateFeedback(feedback *dto.Feedback) error
}
