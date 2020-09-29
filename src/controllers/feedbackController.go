package controllers

import (
	"encoding/json"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/services/interfaces"
	"net/http"
	"time"
)

type FeedbackController struct {
	feedbackService interfaces.FeedbackServiceProvider
	validationService interfaces.ValidationServiceProvider
}

func NewFeedbackController (feedbackService interfaces.FeedbackServiceProvider, validationService interfaces.ValidationServiceProvider) *FeedbackController{
	return &FeedbackController{feedbackService: feedbackService, validationService: validationService}
}

func (controller *FeedbackController) AddFeedback(writer http.ResponseWriter, request *http.Request){
	var feedback dto.Feedback
	if err := json.NewDecoder(request.Body).Decode(&feedback); err != nil{
		errorJsonRespond(writer, http.StatusBadRequest, err)
		return
	}
	if err := controller.validationService.ValidateFeedback(feedback) ; err != nil{
		errorJsonRespond(writer, http.StatusBadRequest, err)
		return
	}
	var timeNow dto.TimeJson
	if err := timeNow.Scan(time.Now().Format(dto.DateFormat)) ; err != nil{
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	feedback.DateCreate = timeNow
	if err := controller.feedbackService.AddFeedback(&feedback); err != nil {
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	respondJson(writer, http.StatusCreated, feedback)
}


func (controller *FeedbackController) GetAllFeedback(writer http.ResponseWriter, request *http.Request){
	result, err := controller.feedbackService.GetAllFeedback()
	if err != nil {
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	respondJson(writer, http.StatusOK, result)
}


