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

const (
	contextKeyId = "userId"
)

type TaskController struct {
	taskService       interfaces.TaskServiceProvider
	validationService interfaces.ValidationServiceProvider
}

func NewTaskController(taskService interfaces.TaskServiceProvider, validationService interfaces.ValidationServiceProvider) *TaskController {
	return &TaskController{taskService, validationService}
}

func (controller *TaskController) CreateTask(writer http.ResponseWriter, request *http.Request) {
	var task dto.Task
	if err := json.NewDecoder(request.Body).Decode(&task); err != nil {
		errorJsonRespond(writer, http.StatusBadRequest, errJsonDecode)
		return
	}
	userId := request.Context().Value(contextKeyId).(string)
	task.UserId = userId
	if err := controller.validationService.ValidateTask(&task); err != nil {
		errorJsonRespond(writer, http.StatusBadRequest, err)
		return
	}
	if err := controller.taskService.CreateTask(&task); err != nil {
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	respondJson(writer, http.StatusCreated, task)
}

func (controller *TaskController) GetTasks(writer http.ResponseWriter, request *http.Request) {
	userId := request.Context().Value(contextKeyId).(string)
	result, err := controller.taskService.GetTasks(userId)
	if err != nil {
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	respondJson(writer, http.StatusOK, result)
}

func (controller *TaskController) GetTaskByID(writer http.ResponseWriter, request *http.Request) {
	taskId, _ := strconv.Atoi(mux.Vars(request)["id"])
	userId := request.Context().Value(contextKeyId).(string)
	result, err := controller.taskService.GetTaskByID(taskId)
	if err != nil {
		if err == sql.ErrNoRows {
			errorJsonRespond(writer, http.StatusNotFound, errNoTask)
			return
		}
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	if result.UserId != userId {
		errorJsonRespond(writer, http.StatusForbidden, errNoAccess)
		return
	}
	respondJson(writer, http.StatusOK, result)
}

func (controller *TaskController) RemoveTaskByID(writer http.ResponseWriter, request *http.Request) {
	taskId, _ := strconv.Atoi(mux.Vars(request)["id"])
	userId := request.Context().Value(contextKeyId).(string)
	if ok, err := controller.checkUsersAccess(userId, taskId); err != nil {
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
	if err := controller.taskService.RemoveByID(taskId); err != nil {
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
	userId := request.Context().Value(contextKeyId).(string)
	if ok, err := controller.checkUsersAccess(userId, taskId); err != nil {
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
	var task dto.Task
	if err := json.NewDecoder(request.Body).Decode(&task); err != nil {
		errorJsonRespond(writer, http.StatusBadRequest, errJsonDecode)
		return
	}
	if err := controller.validationService.ValidateTask(&task); err != nil {
		errorJsonRespond(writer, http.StatusBadRequest, err)
		return
	}
	task.Id = taskId
	if err := controller.taskService.Update(&task); err != nil {
		if err == sql.ErrNoRows {
			errorJsonRespond(writer, http.StatusNotFound, errNoTask)
			return
		}
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	respondJson(writer, http.StatusOK, task)
}

func (controller *TaskController) checkUsersAccess(userId string, taskId int) (bool, error) {
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
