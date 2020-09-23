package mysqlStorage

import (
	"database/sql"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

var testingDbConnection = "root:root@(localhost:3306)/XsollaGarage"

var storage *Storage

func TestMain(m *testing.M){
	code := m.Run()
	os.Exit(code)
}

func setup(){
	db, err := sqlx.Open("mysql", testingDbConnection)
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := migrate(testingDbConnection, "up"); err != nil{
		log.Fatal(err.Error())
	}
	storage = New(db)
}

func tearDown(){
	//if err := migrate(testingDbConnection, "down"); err != nil{
	//	log.Fatal(err.Error())
	//}
	if _, err := storage.db.Exec("DELETE FROM tasks"); err != nil{
		log.Fatal(err.Error())
	}
	// Пока что хз как чистить бд через миграцию, пока хардкод
}

func migrate(dbConnection string, command string) error {
	dir := "../../../migrations"
	db, err := sql.Open("mysql", dbConnection)
	if err != nil {
		return err
	}
	defer db.Close()
	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	if err := goose.Run(command, db, dir); err != nil {
		return err
	}
	return nil
}
func makeTask() (*dto.Task, error) {
	title := "Title"
	content := "Main content"
	dateCreate := new(dto.TimeJson)
	err := dateCreate.UnmarshalJSON([]byte(time.Now().Format(dto.DateFormat)))
	if err != nil{
		return nil, err
	}
	dateTarget := new(dto.TimeJson)
	if err = dateTarget.UnmarshalJSON([]byte(time.Now().Add(5).Format(dto.DateFormat))); err != nil{
		return nil, err
	}
	return &dto.Task{
		Title: &title,
		TextContent: &content,
		DateCreate: dateCreate,
		DateTarget: dateTarget,
	}, nil
}

func TestTaskRepository_Create_NoErrors(t *testing.T) {
	setup()
	task, err := makeTask()
	if err != nil{
		assert.Error(t, err)
	}
	assert.Nil(t, storage.TaskRepository().Create(task))
	tearDown()
}

func TestTaskRepository_CreateAndGetAll(t *testing.T) {
	setup()
	length := 5
	for i:=0;i<length;i++{
		task, err := makeTask()
		if err != nil{
			assert.Error(t, err)
		}
		if err := storage.TaskRepository().Create(task); err != nil{
			assert.Error(t, err)
		}
	}
	result, err := storage.TaskRepository().GetAll()
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, len(result), length)
	tearDown()
}

func TestTaskRepository_GetByID(t *testing.T) {
	setup()
	task, err := makeTask()
	if err != nil{
		assert.Error(t, err)
	}
	if err := storage.TaskRepository().Create(task); err != nil{
		assert.Error(t, err)
	}
	result, err := storage.TaskRepository().GetByID(task.Id)
	assert.Equal(t, task, result)
	tearDown()
}

func TestTaskRepository_Update(t *testing.T) {
	setup()
	task, err := makeTask()
	if err != nil{
		assert.Error(t, err)
	}
	if err := storage.TaskRepository().Create(task); err != nil{
		assert.Error(t, err)
	}
	updatedTask := *task
	updatedContent := "Updated task content"
	updatedTask.TextContent = &updatedContent
	if err := storage.TaskRepository().Update(&updatedTask); err != nil{
		assert.Error(t, err)
	}
	result, err := storage.TaskRepository().GetByID(task.Id)
	if err != nil{
		assert.Error(t, err)
	}
	assert.Equal(t, updatedContent, *result.TextContent)
	tearDown()
}

func TestTaskRepository_RemoveByID(t *testing.T) {
	setup()
	task, err := makeTask()
	if err != nil{
		assert.Error(t, err)
	}
	if err := storage.TaskRepository().Create(task); err != nil{
		assert.Error(t, err)
	}
	if err := storage.TaskRepository().RemoveByID(task.Id); err != nil {
		assert.Error(t, err)
	}
	result, err := storage.TaskRepository().GetAll()
	if err != nil{
		assert.Error(t, err)
	}
	assert.Equal(t, 0,len(result))
	tearDown()
}
