package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"weakpass/checker"
	"weakpass/encode"
	"weakpass/models"
)

func main() {
	models.Init()
	encode.Enumeration()
	Api()

}
func Api() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/hash", checker.HashGate).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //localhost
	}

	fmt.Println(port)
	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8080/api
	if err != nil {
		fmt.Print(err)
	}
}
