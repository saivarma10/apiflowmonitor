package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	tm.Start()
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the Homepage!")
	})
	router.HandleFunc("/create", addTask).Methods("POST")
	router.HandleFunc("/update", UpdateTask).Methods("POST")
	router.HandleFunc("/delete", DeleteTask).Methods("POST")
	router.HandleFunc("/get", GetTask).Methods("POST")
	return router
}
