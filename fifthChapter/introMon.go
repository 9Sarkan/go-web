package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

// Movie model
type Movie struct {
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

func (m *Movie) toString() string {
	return fmt.Sprintf("Name: %s\nYear:%s\nDirectors: %v\nWriters: %v\nBoxOffice: \n\tBudget: %d\n\tGross: %d\n",
		m.Name, m.Year, m.Directors, m.Writers, m.BoxOffice.Budget, m.BoxOffice.Gross)
}

// mongodb const
const (
	host     = "127.0.0.1:27017"
	username = "mongo"
	password = "mongo"
)

func main() {
	info := &mgo.DialInfo{
		Addrs:    []string{host},
		Username: username,
		Password: password,
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	// create a movie
	darkNight := Movie{
		Name:      "Dark Night",
		Year:      "2012",
		Directors: []string{"Christopher Nolan"},
		Writers:   []string{"Jonathan Nolan", "Christopher Nolan"},
		BoxOffice: BoxOfficeType{
			Budget: 200000,
			Gross:  19232,
		},
	}
	movieCollection := session.DB("MovieStar").C("Movie")
	err = movieCollection.Insert(darkNight)
	if err != nil {
		log.Fatal(err)
	}

	// query the movie
	result := Movie{}
	err = movieCollection.Find(bson.M{"boxOffice.gross": bson.M{"$gt": 15000}}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.toString())
}
