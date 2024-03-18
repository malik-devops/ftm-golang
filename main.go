package main

import (
	"log"
	"net/http"
	"os"
	"test/controller"
	"test/database"
	"test/entity"

	"github.com/gorilla/handlers"

	godotenv "github.com/joho/godotenv"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	corsMiddleware := handlers.CORS(headersOk, methodsOk, originsOk)

	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", corsMiddleware(router)))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controller.Createproduct).Methods("POST")
	router.HandleFunc("/get", controller.GetAllproduct).Methods("GET")
	router.HandleFunc("/get/{id}", controller.GetproductByID).Methods("GET")
	router.HandleFunc("/update/{id}", controller.UpdateproductByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controller.DeletproductByID).Methods("DELETE")
}

func initDB() {
	error := godotenv.Load()
	if error != nil {
		log.Println("Error while loading .env")
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")

	config :=
		database.Config{
			ServerName: DB_HOST,
			Port:       DB_PORT,
			User:       DB_USER,
			Password:   DB_PASSWORD,
			DB:         DB_NAME,
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)

	if err != nil {
		panic(err.Error())

	}
	database.Migrate(&entity.Product{})

}
