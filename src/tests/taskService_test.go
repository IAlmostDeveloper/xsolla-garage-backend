package tests

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/services"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/storage/mysqlStorage"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockedDB struct {
	*sqlx.DB
	mock.Mock
}

func TestCreateTask(t *testing.T){
	db := new(MockedDB)
	storage := mysqlStorage.New(db.DB)
	taskService := services.NewTaskService(storage)
}

func TestGetTaskById(t *testing.T){

}

func TestGetTasks(t *testing.T){

}

func TestUpdateTask(t *testing.T){

}

func TestDeleteTask(t *testing.T){

}
