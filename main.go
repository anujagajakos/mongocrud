package main
import (
"github.com/gorilla/mux"
"log"
"net/http"
)

func main() {
s := mux.NewRouter()

//Routes
s.HandleFunc("/createProfile", createProfile).Methods("POST")
//s.HandleFunc("/getAllUsers", getAllUsers).Methods("GET")

//s.HandleFunc("/updateProfile", updateProfile).Methods("PUT")

log.Fatal(http.ListenAndServe(":8000", s)) 
}

