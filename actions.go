// Se definen todos los metodos del proyecto

package main

import (
	"encoding/json" //permite realizar respuestas con json
	"fmt"           //permite manipular objetos
	"log"
	"net/http" //libreria para implementar el servidor

	"github.com/gorilla/mux" //importamos la libreria gorilla mux para el routing de webs en GO
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}

func responseMovie(rw http.ResponseWriter, status int, results Movie) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(results)
}

func responseMovies(rw http.ResponseWriter, status int, results []Movie) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(results)
}

var collection = getSession().DB("Golang_MongoDB").C("movies")

func Index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Home desde el server GOLANG")
}

func MovieList(rw http.ResponseWriter, r *http.Request) {
	var results []Movie
	err := collection.Find(nil).Sort("-_id").All(&results)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados: ", results)
	}

	responseMovies(rw, 200, results)
}
func MovieShow(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		rw.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	results := Movie{}
	err := collection.FindId(oid).One(&results)

	if err != nil {
		rw.WriteHeader(404)
		return
	}

	responseMovie(rw, 200, results)
}

func MovieAdd(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	err = collection.Insert(movie_data)

	if err != nil {
		rw.WriteHeader(500)
		return
	}

	responseMovie(rw, 200, movie_data)
}

func MovieUpdate(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		rw.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(movie_id)
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
		rw.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	document := bson.M{"_id": oid}
	change := bson.M{"$set": movie_data}

	err = collection.Update(document, change)

	if err != nil {
		rw.WriteHeader(404)
		return
	}

	responseMovie(rw, 200, movie_data)
}

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func MovieRemove(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		rw.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	err := collection.RemoveId(oid)
	if err != nil {
		rw.WriteHeader(404)
		return
	}

	results := Message{"success", "El documento de la pelicula con ID" + movie_id + " ha sido eliminado correctamente"}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(results)

}
