package server

import (
	"fmt"
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
	getRouter.HandleFunc("/google-auth", s.authController.GoogleLogin)
	getRouter.HandleFunc("/google-callback", s.authController.GoogleCallback)

	postRouter.HandleFunc("/task", s.taskController.CreateTask)
	postRouter.HandleFunc("/tag", s.tagController.AddToTask)

	deleteRouter.HandleFunc("/task/{id:[0-9]+}", s.taskController.RemoveTaskByID)
	deleteRouter.HandleFunc("/tag", s.tagController.RemoveFromTask)

	putRouter.HandleFunc("/task/{id:[0-9]+}", s.taskController.UpdateTask)
}

func HelloWorld(writer http.ResponseWriter, request *http.Request) {
	var html = `<html><body><a href="/google-auth">Hello world!</a></body></html>`
	fmt.Fprint(writer,  html)
}
