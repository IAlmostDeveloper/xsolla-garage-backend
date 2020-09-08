package controllers

import (
	"encoding/json"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/services"
	"net/http"
)

func CreateTask(writer http.ResponseWriter, request *http.Request) {
	var task dto.Task
	json.NewDecoder(request.Body).Decode(task)
	services.CreateTask(task)
	//TODO Make response
}

func GetTasks(writer http.ResponseWriter, request *http.Request) {
	result := services.GetTasks()
	js, _ := json.Marshal(result)
	writer.Write(js)
}

func GetTask(writer http.ResponseWriter, request *http.Request) {
	taskId := services.GetIdFromPath(request)
	result := services.GetTask(taskId)
	js, _ := json.Marshal(result)
	writer.Write(js)
}
