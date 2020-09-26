package services

import (
	"errors"
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
			validation.Length(1, 100)),
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
			validation.Length(1, 30)),
	)
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
