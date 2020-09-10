package server

import (
	"github.com/gorilla/handlers"
	"net/http"
)

func (s *server) ConfigureRouter() {

	getRouter := s.router.Methods(http.MethodGet).Subrouter()
	postRouter := s.router.Methods(http.MethodPost).Subrouter()

	getRouter.HandleFunc("/", HelloWorld)
	getRouter.HandleFunc("/task/{id:[0-9]+}", s.taskController.GetTaskByID)
	getRouter.HandleFunc("/task", s.taskController.GetTasks)

	postRouter.HandleFunc("/task", s.taskController.CreateTask)

	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"POST", "GET"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	s.router.Use(handlers.CORS(headers, methods, origins))
}

func HelloWorld(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello world"))
}
