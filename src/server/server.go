package server

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/controllers"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/services"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/storage/interfaces"
	"github.com/gorilla/mux"
	"net/http"
)

type server struct {
	router         *mux.Router
	storage        interfaces.StorageProvider
	taskController *controllers.TaskController
	tagController  *controllers.TagController
	authController *controllers.AuthController
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(storage interfaces.StorageProvider) *server {
	server := &server{
		router:  mux.NewRouter(),
		storage: storage,
		taskController: controllers.NewTaskController(
			services.NewTaskService(storage),
			services.NewValidationService()),
		tagController: controllers.NewTagController(
			services.NewTagService(storage),
			services.NewValidationService()),
			authController: controllers.NewAuthController(services.NewGoogleAuthService()),
	}

	server.ConfigureRouter()
	return server
}
