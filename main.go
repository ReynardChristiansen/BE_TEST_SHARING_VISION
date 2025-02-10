package main

import (
	"BE/config"
	"BE/handlers"
	corsHandlers "github.com/gorilla/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	router := mux.NewRouter()

	router.HandleFunc("/article", handlers.CreateArticle).Methods("POST")
	router.HandleFunc("/article/{limit}/{offset}", handlers.GetArticles).Methods("GET")
	router.HandleFunc("/article/{id}", handlers.GetArticleByID).Methods("GET")
	router.HandleFunc("/article/{id}", handlers.UpdateArticle).Methods("POST", "PUT", "PATCH")
	router.HandleFunc("/article/{id}", handlers.DeleteArticle).Methods("POST", "DELETE")

	corsHandler := corsHandlers.CORS(
		corsHandlers.AllowedOrigins([]string{"*"}),
		corsHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
		corsHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	log.Println("Server berjalan di port 8080...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", corsHandler(router)))
}
