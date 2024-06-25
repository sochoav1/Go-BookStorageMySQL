package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	DB *gorm.DB
)

func Connect() {
	// Carga las variables de entorno desde el archivo .env
	err := godotenv.Load("C:\\Users\\Santiago\\Desktop\\Proyectos\\Go\\Go-BookStorageMySQL\\.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Obtiene las variables de entorno
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	// Verificar que todas las variables de entorno están cargadas
	if user == "" || password == "" || name == "" || host == "" || port == "" {
		log.Fatalf("Missing one or more environment variables: DB_USER=%s, DB_PASSWORD=%s, DB_NAME=%s, DB_HOST=%s, DB_PORT=%s", user, password, name, host, port)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, name)
	database, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = database
	log.Println("Database connection established")
}
