package controllers

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HelloWorld(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello world"))
}

func Handle(){
	router := mux.NewRouter()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	_ = router.Methods(http.MethodPost).Subrouter()

	getRouter.HandleFunc("/", HelloWorld)


	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"POST", "GET"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	port := ":8081" // ":" + os.Getenv("PORT") // for env var $PORT
	fmt.Println("Port " + port)

	log.Fatal(http.ListenAndServe(port, handlers.CORS(headers, methods, origins)(router)))
}
