package mysqlStorage

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagRepository_AddToTask(t *testing.T) {
	setup()
	task, err := makeTask()
	if err != nil{
		assert.Error(t, err)
	}
	if err := storage.TaskRepository().Create(task); err != nil{
		assert.Error(t, err)
	}
	tag := &dto.Tag{Id: 0, Name: "Test tag 1"}
	if err := storage.TagRepository().AddToTask(task.Id, tag); err != nil{
		assert.Error(t, err)
	}
	// Допилить тест, пока по факту не тестируется ничего
	tearDown()
}
func TestTagRepository_RemoveFromTask(t *testing.T) {
	setup()
	task, err := makeTask()
	if err != nil{
		assert.Error(t, err)
	}
	if err := storage.TaskRepository().Create(task); err != nil{
		assert.Error(t, err)
	}
	tag := &dto.Tag{Id: 0, Name: "Test tag 1"}
	if err := storage.TagRepository().AddToTask(task.Id, tag); err != nil{
		assert.Error(t, err)
	}
	if err := storage.TagRepository().RemoveFromTask(task.Id, tag.Id); err != nil {
		assert.Error(t, err)
	}
	// Допилить тест, пока по факту не тестируется ничего
	tearDown()
}

func TestTagRepository_GetByTaskId(t *testing.T) {
	setup()
	task, err := makeTask()
	if err != nil{
		assert.Error(t, err)
	}
	if err := storage.TaskRepository().Create(task); err != nil{
		assert.Error(t, err)
	}
	tag := &dto.Tag{Id: 0, Name: "Test tag 1"}
	if err := storage.TagRepository().AddToTask(task.Id, tag); err != nil{
		assert.Error(t, err)
	}
	result, err := storage.TagRepository().GetByTaskId(task.Id)
	if err != nil{
		assert.Error(t, err)
	}
	assert.Equal(t, []*dto.Tag{tag}, result)
	tearDown()
}