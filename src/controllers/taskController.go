package controllers

import (
	"database/sql"
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
	if err := json.NewDecoder(request.Body).Decode(&task); err != nil {
		errorJsonRespond(writer, http.StatusBadRequest, errJsonDecode)
		return
	}
	err := controller.taskService.CreateTask(&task)
	if err != nil {
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	respondJson(writer, http.StatusCreated, task)
}

func (controller *TaskController) GetTasks(writer http.ResponseWriter, request *http.Request) {
	result, err := controller.taskService.GetTasks()
	if err != nil {
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	respondJson(writer, http.StatusOK, result)
}

func (controller *TaskController) GetTaskByID(writer http.ResponseWriter, request *http.Request) {
	taskId, _ := strconv.Atoi(mux.Vars(request)["id"])
	result, err := controller.taskService.GetTaskByID(taskId)
	if err != nil {
		if err == sql.ErrNoRows {
			errorJsonRespond(writer, http.StatusNotFound, errNoTask)
			return
		}
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	respondJson(writer, http.StatusOK, result)
}

func (controller *TaskController) RemoveTaskByID(writer http.ResponseWriter, request *http.Request) {
	taskId, _ := strconv.Atoi(mux.Vars(request)["id"])
	err := controller.taskService.RemoveByID(taskId)
	if err != nil {
		if err == sql.ErrNoRows {
			errorJsonRespond(writer, http.StatusNotFound, errNoTask)
			return
		}
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	respondJson(writer, http.StatusOK, nil)
}

func (controller *TaskController) UpdateTask(writer http.ResponseWriter, request *http.Request) {
	taskId, _ := strconv.Atoi(mux.Vars(request)["id"])
	var task dto.Task
	if err := json.NewDecoder(request.Body).Decode(&task); err != nil {
		errorJsonRespond(writer, http.StatusBadRequest, errJsonDecode)
		return
	}
	task.Id = taskId
	err := controller.taskService.Update(&task)
	if err != nil {
		if err == sql.ErrNoRows {
			errorJsonRespond(writer, http.StatusNotFound, errNoTask)
			return
		}
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	respondJson(writer, http.StatusOK, task)
}
