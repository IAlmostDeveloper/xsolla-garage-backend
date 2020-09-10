package interfaces

type StorageProvider interface {
	TaskRepository() TaskRepositoryProvider
}
