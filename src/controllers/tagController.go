package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/services/interfaces"
	"net/http"
)

type TagController struct {
	tagService        interfaces.TagServiceProvider
	taskService       interfaces.TaskServiceProvider
	validationService interfaces.ValidationServiceProvider
}

func NewTagController(
	tagService interfaces.TagServiceProvider,
	validationService interfaces.ValidationServiceProvider,
	taskService interfaces.TaskServiceProvider) *TagController {
	return &TagController{
		tagService:        tagService,
		taskService:       taskService,
		validationService: validationService,
	}
}

func (controller *TagController) AddToTask(writer http.ResponseWriter, request *http.Request) {
	type requestBody struct {
		TaskId int        `json:"task_id"`
		Tags   []*dto.Tag `json:"Tags"`
	}
	reqBody := &requestBody{}
	if err := json.NewDecoder(request.Body).Decode(reqBody); err != nil {
		errorJsonRespond(writer, http.StatusBadRequest, errJsonDecode)
		return
	}

	userId := request.Context().Value(contextKeyId).(string)
	if ok, err := controller.checkUsersAccess(userId, reqBody.TaskId); err != nil {
		if err == sql.ErrNoRows {
			errorJsonRespond(writer, http.StatusNotFound, errNoTask)
			return
		}
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	} else if !ok {
		errorJsonRespond(writer, http.StatusForbidden, errNoAccess)
		return
	}

	if err := controller.validationService.ValidateTags(reqBody.Tags); err != nil {
		errorJsonRespond(writer, http.StatusBadRequest, err)
		return
	}
	if err := controller.tagService.AddToTask(reqBody.TaskId, reqBody.Tags); err != nil {
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	respondJson(writer, http.StatusCreated, reqBody)
}

func (controller *TagController) RemoveFromTask(writer http.ResponseWriter, request *http.Request) {
	type requestBody struct {
		TaskId int `json:"task_id"`
		TagId  int `json:"tag_id"`
	}

	reqBody := &requestBody{}
	if err := json.NewDecoder(request.Body).Decode(reqBody); err != nil {
		errorJsonRespond(writer, http.StatusBadRequest, errJsonDecode)
		return
	}

	userId := request.Context().Value(contextKeyId).(string)
	if ok, err := controller.checkUsersAccess(userId, reqBody.TaskId); err != nil {
		if err == sql.ErrNoRows {
			errorJsonRespond(writer, http.StatusNotFound, errNoTask)
			return
		}
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	} else if !ok {
		errorJsonRespond(writer, http.StatusForbidden, errNoAccess)
		return
	}

	if err := controller.tagService.RemoveFromTask(reqBody.TaskId, reqBody.TagId); err == sql.ErrNoRows {
		errorJsonRespond(writer, http.StatusNotFound, errNotFound)
		return
	} else if err != nil {
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	respondJson(writer, http.StatusOK, nil)
}

func (controller *TagController) checkUsersAccess(userId string, taskId int) (bool, error) {
	task, err := controller.taskService.GetTaskByID(taskId)
	if err != nil {
		return false, err
	}
	if task.UserId == userId {
		return true, nil
	} else {
		return false, nil
	}
}
