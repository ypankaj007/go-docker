package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	fmt.Println("Application started")
	router.HandleFunc("/", welcome)
	log.Fatal(http.ListenAndServe(":8080", router))

}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Golang & Docker demo application")
}
