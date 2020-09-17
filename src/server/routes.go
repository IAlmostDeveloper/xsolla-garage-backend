package server

import (
	"github.com/gorilla/handlers"
	"net/http"
)

func (s *server) ConfigureRouter() {
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	s.router.Use(handlers.CORS(headers, methods, origins))

	getRouter := s.router.Methods(http.MethodGet, http.MethodOptions).Subrouter()
	postRouter := s.router.Methods(http.MethodPost, http.MethodOptions).Subrouter()
	deleteRouter := s.router.Methods(http.MethodDelete, http.MethodOptions).Subrouter()
	putRouter := s.router.Methods(http.MethodPut, http.MethodOptions).Subrouter()

	getRouter.HandleFunc("/", HelloWorld)
	getRouter.HandleFunc("/task/{id:[0-9]+}", s.taskController.GetTaskByID)
	getRouter.HandleFunc("/task", s.taskController.GetTasks)

	postRouter.HandleFunc("/task", s.taskController.CreateTask)

	deleteRouter.HandleFunc("/task/{id:[0-9]+}", s.taskController.RemoveTaskByID)

	putRouter.HandleFunc("/task/{id:[0-9]+}", s.taskController.UpdateTask)
}

func HelloWorld(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello world"))
}
