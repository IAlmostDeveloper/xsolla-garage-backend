package controllers

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func HelloWorld(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello world"))
}

func Handle() error {
	router := mux.NewRouter()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	postRouter := router.Methods(http.MethodPost).Subrouter()

	getRouter.HandleFunc("/", HelloWorld)

	getRouter.HandleFunc("/task/{id:[0-9]+}", GetTask)
	getRouter.HandleFunc("/task", GetTasks)

	postRouter.HandleFunc("/task", CreateTask)


	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"POST", "GET"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	port := ":8081" // ":" + os.Getenv("PORT") // for env var $PORT
	fmt.Println("Port " + port)

	return http.ListenAndServe(port, handlers.CORS(headers, methods, origins)(router))
}
