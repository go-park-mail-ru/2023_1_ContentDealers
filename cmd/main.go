package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello world!")
	}).Methods("GET")
	router.HandleFunc("/signup", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Sign Up!")
	}).Methods("GET")

	corsHandler := cors.New(cors.Options{
		// TODO: поменять настройки CORS, когда будет известен домен
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})

	router.Use(corsHandler.Handler)

	addr := ":8080"

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Println("start listening on", addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
