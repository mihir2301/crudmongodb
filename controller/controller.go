package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mihirproject/model"
	connection "mihirproject/mongodb"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func createOneMovie(movie model.Netflix) {

	insert, err := connection.Collect.InsertOne(context.Background(), movie)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("inserted one movie ", insert.InsertedID)
}

func updateOneMovie(movieid string) {
	id, _ := primitive.ObjectIDFromHex(movieid)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := connection.Collect.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("updated id is", result.ModifiedCount)

}
func deleteOneMovie(movieid string) {
	id, _ := primitive.ObjectIDFromHex(movieid)
	filter := bson.M{"_id": id}
	result, err := connection.Collect.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("deleted one movie wih id", result)
}

func deleteallmovies() {
	result, err := connection.Collect.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Printf("error ocuured %v", err)
	}
	fmt.Println("Deleted movie count is", result.DeletedCount)
}

func getallmovies() []bson.M {
	movie, err := connection.Collect.Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("error is %v", err)
	}
	var movies []bson.M

	for movie.Next(context.Background()) {
		var movied bson.M
		err := movie.Decode(&movied)
		if err != nil {
			log.Println("error found")
		}
		movies = append(movies, movied)
	}
	return movies
}

func Getallmovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	movies := getallmovies()
	json.NewEncoder(w).Encode(movies)
}

func CreateOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("allow-control-allow-methods", "POST")
	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	createOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkedAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("allow-control-allow-methods", "PUT")
	params := mux.Vars(r)
	id := params["id"]
	updateOneMovie(id)
	json.NewEncoder(w).Encode(id)
}
func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	deleteallmovies()
	json.NewEncoder(w).Encode("deleted all movies")
}
