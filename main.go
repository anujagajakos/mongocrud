package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	//"encoding/json"
   // "strings"

   // "github.com/dgrijalva/jwt-go"
  // "github.com/gorilla/context"
    //"github.com/mitchellh/mapstructure"
)



func main() {

	fmt.Println("Starting the application ....")
	s := mux.NewRouter()

	//Routes
	s.HandleFunc("/createProfile", createProfile).Methods("POST")
	s.HandleFunc("/getAllUsers", ValidateMiddleware(getAllUsers)).Methods("GET")

	s.HandleFunc("/updateProfile", updateProfile).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", s))
}
