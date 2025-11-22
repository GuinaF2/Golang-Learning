package routes

import (
	handler "Api-Aula1/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/books/search", handler.HandleSearch).Methods(http.MethodGet)
	print("Bateu Aqui")
	r.HandleFunc("/books/save", handler.HandleSave).Methods(http.MethodPost)
}
