package main

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/setup"
	"log"
	"net/http"
)

const addr = ":8080"

func main() {

	router := setup.Routes()

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Println("start listening on", addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
