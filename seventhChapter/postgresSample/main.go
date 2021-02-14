package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"

	"github.com/sarkan9/base62"
	models "github.com/sarkan9/shorterModels"

	"github.com/gorilla/mux"
)

// DBClient ...
type DBClient struct {
	db *sql.DB
}

// Record of our db
type Record struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

// GetOrginalLink ...
func (driver *DBClient) GetOrginalLink(w http.ResponseWriter, req *http.Request) {
	var url string
	vars := mux.Vars(req)
	// get Id from base62 string
	id := base62.ToBase10(vars["encoded_string"])
	err := driver.db.QueryRow("Select url from url_shorter where id=$1", id).Scan(&url)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		responseMap := map[string]interface{}{"url": url}
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

// GenerateShortURL ...
func (driver *DBClient) GenerateShortURL(w http.ResponseWriter, req *http.Request) {
	var id int
	var record Record
	postBody, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(postBody, &record)
	err := driver.db.QueryRow("insert into url_shorter(url) values ($1) returning id", record.URL).Scan(&id)
	base62ID := base62.ToBase62(id)
	log.Println("id: ", id)
	responseMap := map[string]interface{}{"encoded_string": base62ID}
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func main() {
	db, err := models.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	driver := &DBClient{db: db}
	router := mux.NewRouter()
	router.HandleFunc("/v1/short/{encoded_string:[0-9a-zA-Z]*}", driver.GetOrginalLink).Methods("GET")
	router.HandleFunc("/v1/short", driver.GenerateShortURL).Methods("POST")
	loggingRouter := handlers.LoggingHandler(os.Stdout, router)
	srv := &http.Server{
		Handler:      loggingRouter,
		Addr:         ":8000",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
