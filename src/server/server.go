package server

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/controllers"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/services"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/store/interfaces"
	"github.com/gorilla/mux"
	"net/http"
)

type server struct {
	router         *mux.Router
	store          interfaces.StoreProvider
	taskController *controllers.TaskController
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(store interfaces.StoreProvider) *server {
	server := &server{
		router:         mux.NewRouter(),
		store:          store,
		taskController: controllers.NewTaskController(services.NewTaskService(store)),
	}

	server.ConfigureRouter()
	return server
}
