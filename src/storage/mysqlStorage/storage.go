package mysqlStorage

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/storage/interfaces"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db             *sqlx.DB
	taskRepository interfaces.TaskRepositoryProvider
}

func New(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (storage *Storage) TaskRepository() interfaces.TaskRepositoryProvider {
	if storage.taskRepository != nil {
		return storage.taskRepository
	}

	storage.taskRepository = &TaskRepository{
		db: storage.db,
	}

	return storage.taskRepository
}