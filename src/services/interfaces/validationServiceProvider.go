package interfaces

import "github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"

type ValidationServiceProvider interface {
	ValidateTask(task *dto.Task) error
}