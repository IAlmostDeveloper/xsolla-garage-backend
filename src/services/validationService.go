package services

import (
	"errors"
	"fmt"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"strings"
	"time"
)

type ValidationService struct {
}

func NewValidationService() *ValidationService {
	return &ValidationService{}
}

func (v *ValidationService) ValidateTask(task *dto.Task) error {
	return validation.ValidateStruct(task,
		validation.Field(&task.Title,
			validation.NotNil,
			validation.Required,
			validation.By(v.validateStringPtrRuneCount(1, 100))),
		validation.Field(&task.DateTarget,
			validation.Min(time.Now().AddDate(0, 0, -1))),
		validation.Field(&task.Tags,
			validation.Skip.When(len(task.Tags) == 0),
			validation.By(v.validateTagSlice)),
	)
}

func (v *ValidationService) ValidateTag(tag *dto.Tag) error {
	return validation.ValidateStruct(tag,
		validation.Field(&tag.Name,
			validation.Required,
			validation.By(v.validateStringRuneCount(1, 30))),
	)
}

func (v *ValidationService) ValidateFeedback(feedback dto.Feedback) error {
	return validation.ValidateStruct(feedback,
		validation.Field(&feedback.Content,
			validation.Required,
			validation.By(v.validateStringRuneCount(1,1000))))
}

func (v *ValidationService) validateStringPtrRuneCount(min int, max int) validation.RuleFunc {
	return func(value interface{}) error {
		stringPtr, ok := value.(*string)
		if !ok {
			return errors.New("not a string")
		}
		runeCount := len([]rune(*stringPtr))
		if runeCount > max || runeCount < min {
			return errors.New(fmt.Sprintf("the length must be between %d and %d.", min, max))
		}
		return nil
	}
}

func (v *ValidationService) validateStringRuneCount(min int, max int) validation.RuleFunc {
	return func(value interface{}) error {
		string, ok := value.(string)
		if !ok {
			return errors.New("not a string")
		}
		runeCount := len([]rune(string))
		if runeCount > max || runeCount < min {
			return errors.New(fmt.Sprintf("the length must be between %d and %d.", min, max))
		}
		return nil
	}
}

func (v *ValidationService) ValidateTags(tags []*dto.Tag) error {
	return validation.Validate(tags, validation.By(v.validateTagSlice))
}

func (v *ValidationService) validateTagSlice(value interface{}) error {
	tags, ok := value.([]*dto.Tag)
	if !ok {
		return errors.New("not a tag array")
	}
	for _, tag := range tags {
		if err := v.ValidateTag(tag); err != nil {
			return err
		}
	}
	for i := 0; i < len(tags)-1; i++ {
		for j := i + 1; j < len(tags); j++ {

			if strings.ToLower(strings.Trim(tags[i].Name, " ")) ==
				strings.ToLower(strings.Trim(tags[j].Name, " ")) {
				return errors.New("some tags are duplicated")
			}
		}
	}
	return nil
}
