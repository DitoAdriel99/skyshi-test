package router

import (
	"go-learn/controller/activity"
	"go-learn/controller/todo"
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/activity-groups", activity.Create).Methods("POST")
	router.HandleFunc("/activity-groups", activity.GetAll).Methods("GET")
	router.HandleFunc("/activity-groups/{id}", activity.Update).Methods("PUT")
	router.HandleFunc("/activity-groups/{id}", activity.GetOne).Methods("GET")
	router.HandleFunc("/activity-groups/{id}", activity.Delete).Methods("DELETE")

	router.HandleFunc("/todo-items", todo.Create).Methods("POST")
	router.HandleFunc("/todo-items", todo.GetAll).Methods("GET")
	router.HandleFunc("/todo-items/{id}", todo.Update).Methods("PUT")
	router.HandleFunc("/todo-items/{id}", todo.GetOne).Methods("GET")
	router.HandleFunc("/todo-items/{id}", todo.Delete).Methods("DELETE")
	return router
}
