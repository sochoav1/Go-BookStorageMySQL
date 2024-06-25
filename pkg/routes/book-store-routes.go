package routes

import (
	"github.com/gorilla/mux"
	"github.com/sochoav1/Go-BookStorageMySQL"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/books", controllers.CreateBooks).Methods("POST")
	router.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", controllers.DeleteBook).Methods("DELETE")

}
