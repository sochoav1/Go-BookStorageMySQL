package main

import (
	"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sochoav1/Go-BookStorageMySQL/pkg/routes"
	"log"
	"net/http"
)

func main(){
	r := routes.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil)
}