package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
	dbutils "github.com/sarkan9/raveldbutils"
)

var db *sql.DB

// TrainResource is the model for holding rail information
type TrainResource struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

// StationResource holds information about locations
type StationResource struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

// ScheduleResource links both trains and stations
type ScheduleResource struct {
	ID          int
	TrainID     int
	StationID   int
	ArrivalTime time.Time
}

// Register container
func (t *TrainResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/train").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.saveTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.deleteTrain))
	ws.Route(ws.PUT("/{train-id}").To(t.updateTrain))
	container.Add(ws)
}
func (t *TrainResource) getTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	err := db.QueryRow("select ID, DRIVER_NAME, OPERATING_STATUS from train where ID=?", id).
		Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plan")
		response.WriteErrorString(http.StatusNotFound, "Train could not found!")
	} else {
		response.WriteEntity(t)
	}
}
func (t *TrainResource) saveTrain(request *restful.Request, response *restful.Response) {
	log.Println(request.Request.Body)
	decoder := json.NewDecoder(request.Request.Body)
	var trainResource TrainResource
	err := decoder.Decode(&trainResource)
	log.Println(trainResource.DriverName, trainResource.OperatingStatus)
	if err != nil {
		response.AddHeader("Content-Type", "text/plan")
		response.WriteHeader(http.StatusBadRequest)
		response.WriteErrorString(http.StatusBadRequest, "wrong data format")
		return
	}
	statement, _ := db.Prepare("insert into train (DRIVER_NAME, OPERATING_STATUS) values (?, ?)")
	result, errResult := statement.Exec(trainResource.DriverName, trainResource.OperatingStatus)
	if errResult == nil {
		newID, _ := result.LastInsertId()
		trainResource.ID = int(newID)
		response.WriteHeaderAndEntity(http.StatusCreated, trainResource)
	} else {
		response.AddHeader("Content-Type", "text/plan")
		response.WriteErrorString(http.StatusInternalServerError, errResult.Error())
	}
}
func (t *TrainResource) deleteTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	statement, _ := db.Prepare("delete from train where id=?")
	_, err := statement.Exec(id)
	if err != nil {
		response.AddHeader("Content-Type", "text/plan")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	} else {
		response.WriteHeader(http.StatusOK)
	}
}
func (t *TrainResource) updateTrain(request *restful.Request, response *restful.Response) {
	decoder := json.NewDecoder(request.Request.Body)
	var train TrainResource
	err := decoder.Decode(&train)
	if err != nil {
		response.AddHeader("Content-Type", "text/plan")
		response.WriteErrorString(http.StatusBadRequest, "invalid data format")
		return
	}
	statement, _ := db.Prepare("update train set DRIVER_NAME=?, OPERATING_STATUS=? where id=?")
	id := request.PathParameter("train-id")
	_, err = statement.Exec(train.DriverName, train.OperatingStatus, id)
	if err == nil {
		train.ID, _ = strconv.Atoi(id)
		response.WriteEntity(train)
	} else {
		response.AddHeader("Content-Type", "text/plan")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}
func main() {
	var err error
	db, err = sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println(err)
	}
	dbutils.Initialize(db)
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := TrainResource{}
	t.Register(wsContainer)
	server := http.Server{Addr: ":8000", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
