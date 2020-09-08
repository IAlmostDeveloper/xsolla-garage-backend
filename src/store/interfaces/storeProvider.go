package interfaces

type StoreProvider interface {
	Task() TaskRepositoryProvider
}
