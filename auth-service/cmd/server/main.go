package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
   router := mux.NewRouter()

	 // Add routes and middleware
   router.HandleFunc("/login", handleLogin).Methods("POST")

   log.Fatal(http.ListenAndServe(":8080", router))

   log.Fatal(http.ListenAndServe(":8080", router))
}

func handleLogin(w http.ResponseWriter, r *http.Request){
   
}
