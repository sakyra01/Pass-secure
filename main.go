package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"weakpass/checker"
	"weakpass/models"
)

func main() {
	models.Init() // connection DB with auto migration
	Api()         // API functionality
}

func Api() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/hash", checker.HashGate).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "227" //localhost 227 port
	}

	fmt.Println(port)
	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:227/api
	if err != nil {
		fmt.Print(err)
	}
}
