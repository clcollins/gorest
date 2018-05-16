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
}

func GetCheeseburger(w http.ResponseWriter, r *http.Request) {
	response := "CheeseburgerCheeseburger"
	w.Write([]byte(response))
	// I'm not yet sure how to work with json encoding
	//json.NewEncoder(w).Encode(response)
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
}

func main() {
	port := "8080"
	iface := "0.0.0.0"

	router := mux.NewRouter()
	router.HandleFunc("/", GetIndex).Methods("GET")
	router.HandleFunc("/ichc", GetCheeseburger).Methods("GET")
	router.HandleFunc("/ichc", CreateCheeseburger).Methods("POST")

	// Am I *really* listening on this interface?
  fmt.Println("Listening on " + iface + ":" + port)
	log.Fatal(http.ListenAndServe(":" + port, httpLogger.WriteLog(
		router, os.Stdout)))
}
