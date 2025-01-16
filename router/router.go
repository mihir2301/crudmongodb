package user

import (
	"mihirproject/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/create", controller.CreateOneMovie).Methods("POST")
	r.HandleFunc("/update/{id}", controller.MarkedAsWatched).Methods("PUT")
	r.HandleFunc("/deleteone/{id}", controller.DeleteOneMovie).Methods("DELETE")
	r.HandleFunc("/deletemany", controller.DeleteAllMovies).Methods("DELETE")
	r.HandleFunc("/getallmovies", controller.Getallmovies).Methods("GET")
	return r
}
