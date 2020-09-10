package controllers

import (
	"encoding/json"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/services/interfaces"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TaskController struct {
	taskService interfaces.TaskServiceProvider
}

func NewTaskController(taskService interfaces.TaskServiceProvider) *TaskController {
	return &TaskController{taskService}
}

func (controller *TaskController) CreateTask(writer http.ResponseWriter, request *http.Request) {
	var task dto.Task
	json.NewDecoder(request.Body).Decode(&task)
	err := controller.taskService.CreateTask(&task)
	if err != nil {
		errorRespond(writer, request, http.StatusInternalServerError, err)
		return
	}
	respond(writer, request, http.StatusCreated, task)
}

func (controller *TaskController) GetTasks(writer http.ResponseWriter, request *http.Request) {
	result, err := controller.taskService.GetTasks()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	respond(writer, request, http.StatusOK, result)
}

func (controller *TaskController) GetTaskByID(writer http.ResponseWriter, request *http.Request) {
	taskId, _ := strconv.Atoi(mux.Vars(request)["id"])
	result, err := controller.taskService.GetTaskByID(taskId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	respond(writer, request, http.StatusOK, result)

}
