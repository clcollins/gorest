package main

import (
  //"encoding/json"
	"fmt"
	"os"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ajays20078/go-http-logger"
)

func CreateCheeseburger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	query_params := r.URL.Query()

	var meal string
	if _, cheese := query_params["cheese"]; cheese {
		meal = "cheeseburger"
	} else {
		meal = "hamburger"
	}

	fmt.Println(meal)

	response := fmt.Sprintf("I'll gladly pay you tuesday for a %s today!\n", meal)


  w.Write([]byte(response))

	for key, value := range query_params {
		fmt.Printf("Key: %s\tValue: %v\n", key, value)
	}
}

func GetCheeseburger(w http.ResponseWriter, r *http.Request) {
	response := "CheeseburgerCheeseburger\n"
	w.Write([]byte(response))
	// I'm not yet sure how to work with json encoding
	//json.NewEncoder(w).Encode(response)
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	response := "There's not really supposed to be an index"
	w.Write([]byte(response))
}

func main() {
	port := "8080"
	iface := "0.0.0.0"

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", GetIndex).Methods("GET")
	router.HandleFunc("/ichc", GetCheeseburger).Methods("GET")
	router.HandleFunc("/ichc", CreateCheeseburger).Methods("POST")

	// Am I *really* listening on this interface?
  fmt.Println("Listening on " + iface + ":" + port)
	log.Fatal(http.ListenAndServe(":" + port, httpLogger.WriteLog(
		router, os.Stdout)))
}
