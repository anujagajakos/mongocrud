package main
import (
"github.com/gorilla/mux"
"log"
"net/http"

)

func main() {
s := mux.NewRouter()

//Routesgo 
s.HandleFunc("/createProfile", createProfile).Methods("POST")

s.HandleFunc("/getAllUsers", getAllUsers).Methods("GET")

//s.HandleFunc("/updateProfile", updateProfile).Methods("PUT")

log.Fatal(http.ListenAndServe(":8000", s)) 
}

//run the program as -- go run main.go db.go controllers.go