package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/services"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HelloWorld(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello world"))
}

func CreateTask(writer http.ResponseWriter, request *http.Request) {
	var task dto.Task
	json.NewDecoder(request.Body).Decode(task)
	services.CreateTask(task)
	//TODO Make response
}

func Handle(){
	router := mux.NewRouter()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	postRouter := router.Methods(http.MethodPost).Subrouter()

	getRouter.HandleFunc("/", HelloWorld)

	postRouter.HandleFunc("/task", CreateTask)


	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"POST", "GET"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	port := ":8081" // ":" + os.Getenv("PORT") // for env var $PORT
	fmt.Println("Port " + port)

	log.Fatal(http.ListenAndServe(port, handlers.CORS(headers, methods, origins)(router)))
}
