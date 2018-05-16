package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateCheeseburger(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Made a CheesebugerCheeseburger!")
}

func GetCheeseburger(w http.ResponseWriter, r *http.Request) {
	response := "CheeseburgerCheeseburger"
	fmt.Println("Got a", response)
	json.NewEncoder(w).Encode(response)
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("gorest!")
}

func main() {
	port := "8080"

	router := mux.NewRouter()
	router.HandleFunc("/", GetIndex).Methods("GET")
	router.HandleFunc("/ichc", GetCheeseburger).Methods("GET")
	router.HandleFunc("/ichc", CreateCheeseburger).Methods("POST")

  fmt.Println("Listening on " + "0.0.0.0" + ":" + port)
	log.Fatal(http.ListenAndServe(":8080", router))
}
