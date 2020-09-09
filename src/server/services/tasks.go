package services

import (
	"fmt"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/store/mysqlStore"
	"github.com/jmoiron/sqlx"
)
var db, _ = sqlx.Connect("mysql", "XsollaGarage:XsollaGarage@(localhost:3306)/XsollaGarage")
var repoProvider = mysqlStore.New(db)

func CreateTask(task dto.Task){
	err := repoProvider.TaskRepository().Create(&task)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetTasks() *[]dto.Task{
	res, err := repoProvider.TaskRepository().GetAll()
	if err != nil{
		fmt.Println(err.Error())
	}
	return res
}

func GetTask(taskId int) *dto.Task{
	res, err := repoProvider.TaskRepository().GetByID(taskId)
	if err != nil{
		fmt.Println(err.Error())
	}
	return res
}
