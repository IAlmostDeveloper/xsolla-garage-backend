package interfaces

type StorageProvider interface {
	TaskRepository() TaskRepositoryProvider
	TagRepository() TagRepositoryProvider
	FeedbackRepository() FeedbackRepositoryProvider
	UserRepository() UserRepositoryProvider
}
