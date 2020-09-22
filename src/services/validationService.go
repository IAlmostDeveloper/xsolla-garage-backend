package services

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"time"
)

type ValidationService struct {
}

func NewValidationService() *ValidationService {
	return &ValidationService{}
}

func (v *ValidationService) ValidateTask(task *dto.Task) error {
	if err := validateTitle(task.Title); err != nil {
		return err
	}
	if err := validateDate(task.DateTarget); err != nil {
		return err
	}
	return nil
}

func validateTitle(text *string) error {
	if *text == "" || text == nil {
		return errValidateTitle
	}
	return nil
}

func validateDate(date *dto.TimeJson) error {
	dateRaw, err := date.Value()
	if err != nil {
		return err
	}
	if dateRaw.(time.Time).Before(time.Now()) {
		return errValidateDate
	}
	return nil
}
