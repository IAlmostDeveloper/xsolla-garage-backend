package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/services/interfaces"
	"net/http"
)

type TagController struct {
	tagService interfaces.TagServiceProvider
}

func NewTagController(tagService interfaces.TagServiceProvider) *TagController {
	return &TagController{tagService}
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

	if err := controller.tagService.RemoveFromTask(reqBody.TaskId, reqBody.TagId); err == sql.ErrNoRows {
		errorJsonRespond(writer, http.StatusNotFound, errNotFound)
		return
	} else if err != nil {
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	respondJson(writer, http.StatusOK, nil)
}
