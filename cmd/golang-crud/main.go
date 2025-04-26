package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/suryanshvermaa/golang-crud/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//database setup
	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello welcome to crud apis"))
	})
	//setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	fmt.Println("Server started")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start the server")
	}
}
