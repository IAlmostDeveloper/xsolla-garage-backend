package controllers

import (
	"encoding/json"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/server/services"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateTask(writer http.ResponseWriter, request *http.Request) {
	var task dto.Task
	json.NewDecoder(request.Body).Decode(&task)
	err := services.CreateTask(&task)
	if err != nil{
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	//TODO Make response
}

func GetTasks(writer http.ResponseWriter, request *http.Request) {
	result, err := services.GetTasks()
	if err != nil{
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	js, _ := json.Marshal(result)
	writer.Write(js)
}

func GetTask(writer http.ResponseWriter, request *http.Request) {
	taskId, _ := strconv.Atoi(mux.Vars(request)["id"])
	result, err := services.GetTask(taskId)
	if err != nil{
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	js, _ := json.Marshal(result)
	writer.Write(js)
}
