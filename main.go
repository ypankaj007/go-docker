package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	fmt.Println("Application started")
	router.HandleFunc("/", welcome)
	log.Fatal(http.ListenAndServe(":8080", router))

}

func welcome(w http.ResponseWriter, r *http.Request) {
	str := fmt.Sprintf("%s%d%s\n", "Welcome to Golang & Docker demo application. \n Application is running at ", os.Getegid(), " process ID.")
	fmt.Fprintln(w, str)
}
