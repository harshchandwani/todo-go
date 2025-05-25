package routes

import (
	"todo-app/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/todos", handlers.GetTodos).Methods("GET")
	router.HandleFunc("/api/todos", handlers.CreateTodo).Methods("POST")
	router.HandleFunc("/api/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/api/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

	return router
}
