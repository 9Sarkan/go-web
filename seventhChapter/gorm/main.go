package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	"github.com/sarkan9/models"
)

type dbClient struct {
	db *gorm.DB
}

// UserResponse struct
type UserResponse struct {
	User models.User `json:"user"`
	Data interface{} `json:"data"`
}

// GetUserByName function
func (driver *dbClient) GetUserByName(w http.ResponseWriter, req *http.Request) {
	var users []models.User
	name := req.FormValue("firstName")
	var query = "select * from \"user\" where data->>'first_name'=$1"
	driver.db.Raw(query, name).Scan(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resposneJSON, _ := json.Marshal(users)
	w.Write(resposneJSON)
}

// GetUser by id
func (driver *dbClient) GetUser(w http.ResponseWriter, req *http.Request) {
	var user models.User
	vars := mux.Vars(req)
	driver.db.First(&user, vars["id"])
	var userData interface{}
	json.Unmarshal([]byte(user.Data), &userData)
	var response = UserResponse{User: user, Data: userData}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseJSON, _ := json.Marshal(response)
	w.Write(responseJSON)
}

// PostUser adds URL to DB and gives back shortened string
func (driver *dbClient) PostUser(w http.ResponseWriter, req *http.Request) {
	var user = models.User{}
	postBody, _ := ioutil.ReadAll(req.Body)
	user.Data = string(postBody)
	driver.db.Save(&user)
	responseMap := map[string]interface{}{"id": user.ID}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal(responseMap)
	w.Write(response)
}

func main() {
	db, err := models.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	driver := &dbClient{db: db}
	router := mux.NewRouter()
	router.HandleFunc("/user", driver.GetUserByName).Methods("GET")
	router.HandleFunc("/user/{id:[a-zA-Z0-9]*}", driver.GetUser).Methods("GET")
	router.HandleFunc("/user", driver.PostUser).Methods("POST")
	handler := handlers.LoggingHandler(os.Stdout, router)
	srv := http.Server{
		Handler: handler,
		Addr:    ":8000",
	}
	log.Fatal(srv.ListenAndServe())
}
