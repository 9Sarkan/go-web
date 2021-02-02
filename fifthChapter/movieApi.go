package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	mgo "gopkg.in/mgo.v2"
)

// Movie model
type Movie struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name      string        `bson:"name"`
	Year      string        `bson:"year"`
	Directors []string      `bson:"directors"`
	Writers   []string      `bson:"writers"`
	BoxOffice BoxOfficeType `bson:"boxOffice"`
}

// BoxOfficeType model
type BoxOfficeType struct {
	Budget uint64 `bson:"budget"`
	Gross  uint64 `bson:"gross"`
}

// DB model
type DB struct {
	session    *mgo.Session
	collection *mgo.Collection
}

func (db *DB) getMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movie := Movie{}
	err := db.collection.Find(bson.M{"_id": bson.ObjectIdHex(vars["id"])}).One(&movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response, _ := json.Marshal(movie)
		w.Write(response)
	}
}
func (db *DB) saveMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &movie)
	movie.ID = bson.NewObjectId()
	err := db.collection.Insert(movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(movie)
		w.Write(response)
	}
}
func (db *DB) updateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var movie Movie
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &movie)
	err := db.collection.Update(bson.M{"_id": bson.ObjectIdHex(vars["id"])}, bson.M{"$set": &movie})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(movie)
		w.Write(response)
	}
}
func main() {
	info := &mgo.DialInfo{
		Addrs:    []string{"127.0.0.1:27017"},
		Username: "mongo",
		Password: "mongo",
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		log.Fatal(err)
	}
	collection := session.DB("MovieStar").C("Movie")
	db := DB{
		session:    session,
		collection: collection,
	}
	r := mux.NewRouter()
	r.HandleFunc("/movie/{id:[a-zA-Z0-9]*}", db.getMovie).Methods("GET")
	r.HandleFunc("/movie/{id:[a-zA-Z0-9]*}", db.updateMovie).Methods("PUT")
	r.HandleFunc("/movie", db.saveMovie).Methods("POST")
	loggingHandler := handlers.LoggingHandler(os.Stdout, r)
	server := &http.Server{
		Addr:    ":8000",
		Handler: loggingHandler,
	}
	log.Fatal(server.ListenAndServe())
}
