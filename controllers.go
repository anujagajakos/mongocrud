package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

    "github.com/dgrijalva/jwt-go"
	//"go.mongodb.org/mongo-driver/mongo/options"
    
   "strings"

   // "github.com/dgrijalva/jwt-go"
  //"github.com/gorilla/context"
    //"github.com/mitchellh/mapstructure"

)

var userCollection = db().Database("goTest").Collection("userinfo")

type user struct {
	Name string `json:"name"`
	City string `json:"city"`
	Age  int    `json:"age"`
}


type JwtToken struct {
    Token string `json:"token"`
}
type Exception struct {
    Message string `json:"message"`
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        authorizationHeader := req.Header.Get("authorization")
        if authorizationHeader != "" {
            bearerToken := strings.Split(authorizationHeader, " ")
            if len(bearerToken) == 2 {
                token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
                    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                        return nil, fmt.Errorf("There was an error")
                    }
                    return []byte("secret"), nil
                })
                if error != nil {
                    json.NewEncoder(w).Encode(Exception{Message: error.Error()})
                    return
                }
                if token.Valid {
                    //context.Set(req, "decoded", token.Claims)
                    next(w, req)
                } else {
                    json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
                }
            }
        } else {
            json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
        }
    })
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
	
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": person.Name,
	})
	
    tokenString, error := token.SignedString([]byte("secret"))
    if error != nil {
        fmt.Println(error)
    }
    json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}


func updateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//declare a new user struct
	var person user
	var updatedResult user
	//decode body of request as struct user
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// fmt.Fprintf(w, person.Name)
	result, err := userCollection.UpdateOne(
		context.TODO(),
		bson.M{"name": person.Name},
		bson.D{
			{"$set", bson.D{{"city", person.City}}},
			{"$set", bson.D{{"age", person.Age}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	//dispaly modified count
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	//send response of the updated user
	//find user
	err = userCollection.FindOne(context.TODO(), bson.M{"name": person.Name}).Decode(&updatedResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//send response
	json.NewEncoder(w).Encode(updatedResult)
}

