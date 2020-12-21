package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"



var userCollection = db().Database("goTest").Collection("userinfo")

type user struct {
	Name string `json:"name"`
	City string `json:"city"`
	Age  int    `json:"age"`
}


func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M                                   //slice for multiple documents
	cur, err := userCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
	if err != nil {

		fmt.Println(err)

	}
	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem) // appending document pointed by Next()
	}
	cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
	json.NewEncoder(w).Encode(results)
}

func createProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person user
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		fmt.Print(err)
	}
	insertResult, err := userCollection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(insertResult.InsertedID)
}

