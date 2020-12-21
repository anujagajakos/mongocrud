package main

import (


func main() {

	//Routes
	s.HandleFunc("/createProfile", createProfile).Methods("POST")
	s.HandleFunc("/getAllUsers", getAllUsers).Methods("GET")

	s.HandleFunc("/updateProfile", updateProfile).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", s))
}


