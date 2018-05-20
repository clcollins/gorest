package main

import (
  //"encoding/json"
	"fmt"
	"os"
	"log"
	"net/http"
	//"html/template"
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

	response := fmt.Sprintf(
		"I'll gladly pay you tuesday for a %s today!\n", meal)

  w.Write([]byte(response))

	for key, value := range query_params {
		fmt.Printf("Key: %s\tValue: %v\n", key, value)
	}
}

func GetCheeseburger(w http.ResponseWriter, r *http.Request) {
  // Unsure why this returns string(?) of random numbers...
  // eg:
  //
  // [123 34 73 68 34 58 49 44 34 78 97 109 101 34 58 34 99 104 101 101 115 101 115 34 44 34 84 111 112 112 105 110 103 84 121 112 101 115 34 58 91 34 97 109 101 114 105 99 97 110 34 44 34 99 104 101 100 100 97 114 34 44 34 115 119 105 115 115 34 44 34 98 114 105 101 34 44 34 112 114 111 118 111 108 111 110 101 34 44 34 109 111 122 122 97 114 101 108 108 97 34 93 125]
  //
  //  type ToppingGroup struct {
  //  	ID     int
  //  	Name   string
  //  	ToppingTypes []string
  //  }
  //
  //  cheeses := ToppingGroup{
  //  	ID: 1,
  //  	Name: "cheeses",
  //  	ToppingTypes: []string{
  //  		"american",
  //  		"cheddar",
  //  		"swiss",
  //  		"brie",
  //  		"provolone",
  //  		"mozzarella"},
  //  	}
  //
  //  out, err := json.Marshal(cheeses)
  //	if err != nil {
  //   	fmt.Println("error:", err)
  //  }
  //
  //	fmt.Println(out)
	response := "CheeseburgerCheeseburger\nWe got the cheeses!"
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
