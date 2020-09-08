package interfaces

type StoreProvider interface {
	TaskRepository() TaskRepositoryProvider
}
