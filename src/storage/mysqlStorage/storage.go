package mysqlStorage

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/storage/interfaces"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db                 *sqlx.DB
	taskRepository     interfaces.TaskRepositoryProvider
	tagRepository      interfaces.TagRepositoryProvider
	feedbackRepository interfaces.FeedbackRepositoryProvider
	userRepository     interfaces.UserRepositoryProvider
}

func (storage *Storage) TagRepository() interfaces.TagRepositoryProvider {
	if storage.tagRepository != nil {
		return storage.tagRepository
	}

	storage.tagRepository = &TagRepository{
		db: storage.db,
	}

	return storage.tagRepository
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

func (storage *Storage) FeedbackRepository() interfaces.FeedbackRepositoryProvider {
	if storage.feedbackRepository != nil {
		return storage.feedbackRepository
	}

	storage.feedbackRepository = &FeedbackRepository{
		db: storage.db,
	}

	return storage.feedbackRepository
}

func (storage *Storage) UserRepository() interfaces.UserRepositoryProvider {
	if storage.userRepository != nil {
		return storage.userRepository
	}

	storage.userRepository = &UserRepository{
		db: storage.db,
	}

	return storage.userRepository
}
