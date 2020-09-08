package mysqlStore

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/store/interfaces"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db             *sqlx.DB
	taskRepository interfaces.TaskRepositoryProvider
}

func New(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

func (store *Store) Task() interfaces.TaskRepositoryProvider {
	if store.taskRepository != nil {
		return store.taskRepository
	}

	store.taskRepository = &TaskRepository{
		db: store.db,
	}

	return store.taskRepository
}
