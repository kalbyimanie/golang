package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Welcome to the HomePage!")
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/home", homePage)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	fmt.Println("Server starting on port :8000")
	handleRequest()
}
