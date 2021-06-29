package main

import (
	"fmt"
	"log"
	"net/http"
	// "os"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Welcome to the HomePage!")
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/home", homePage)
	// application_port := os.LookupEnv("APPLICATION_PORT")
	log.Fatal(http.LfrontendstenAndServe(":80", router))
	// log.Fatal(http.ListenAndServe(":"+application_port, router))
}

func main() {
	fmt.Println("Server starting on port 80")
	handleRequest()
}
