package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-gonic/gin"
	dbutils "github.com/sarkan9/raveldbutils"
)

var db *sql.DB

// StationResource struct
type StationResource struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

func getStation(c *gin.Context) {
	var station StationResource
	id := c.Param("station_id")
	err := db.QueryRow("select ID, NAME, CAST(OPENING_TIME as CHAR), CAST(CLOSING_TIME as CHAR) from station where ID=?", id).
		Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"result": station,
		})
	}

}
func getAllStation(c *gin.Context) {
	var stations []StationResource
	raw, err := db.Query("select ID, NAME, CAST(OPENING_TIME as CHAR), CAST(CLOSING_TIME as CHAR) from station")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	var tempStation StationResource
	for raw.Next() {
		raw.Scan(&tempStation.ID, &tempStation.Name, &tempStation.OpeningTime, &tempStation.ClosingTime)
		stations = append(stations, tempStation)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": stations,
	})

}
func saveStation(c *gin.Context) {
	var station StationResource
	if err := c.BindJSON(&station); err == nil {
		statement, _ := db.Prepare("insert into station (NAME, OPENING_TIME, CLOSING_TIME) values (?, ?, ?)")
		result, err := statement.Exec(station.Name, station.OpeningTime, station.ClosingTime)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			id, _ := result.LastInsertId()
			station.ID = int(id)
			c.JSON(201, gin.H{
				"result": station,
			})
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}
func deleteStation(c *gin.Context) {
	id := c.Param("station_id")
	statement, err := db.Prepare("delete from station where ID=?")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		_, execErr := statement.Exec(id)
		if execErr == nil {
			c.String(204, "deleted successfully")
		}
		c.String(500, execErr.Error())
	}
}
func main() {
	var err error
	db, err = sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Fatal(err)
	}
	dbutils.Initialize(db)
	r := gin.Default()
	// add routes
	r.GET("station", getAllStation)
	r.POST("station", saveStation)
	r.GET("station/:station_id", getStation)
	r.DELETE("station/:station_id", deleteStation)
	r.Run(":8000")
}
