package handlers

import (
	"github.com/gorilla/mux"
)

type Server struct {
}
	
func ConfigureRouter(handler GuitarHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/guitar", handler.CreateGuitar).Methods("POST")
	r.HandleFunc("/guitar", handler.ReadAll).Methods("GET")
	r.HandleFunc("/guitar/{Id}", handler.GetByGuitarId).Methods("GET")
	r.HandleFunc("/guitar/{Id}", handler.DeleteGuitarId).Methods("DELETE")
	r.HandleFunc("/guitar/{Id}", handler.UpdateGuitarInfo).Methods("PUT")

	return r
}