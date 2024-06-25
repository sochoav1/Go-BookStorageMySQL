package main

import (
	"log"
	"net/http"

	"github.com/sochoav1/Go-BookStorageMySQL/pkg/routes"
)

func main() {
	r := routes.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
