package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	
	/*"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"*/
)

var userCollection = db().Database("goTest").Collection("userinfo")

type user struct{
	Name string `json:"name"`
	City string `json:"city"`
	Age int `json:"age"`
	}

func createProfile(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json") 
var person user
err := json.NewDecoder(r.Body).Decode(&person) 
if err != nil {
fmt.Print(err)
}
insertResult, err := userCollection.InsertOne(context.TODO(),person)
if err != nil {
log.Fatal(err)
}
json.NewEncoder(w).Encode(insertResult.InsertedID) 
}

