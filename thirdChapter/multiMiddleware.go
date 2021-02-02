package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"

	"github.com/justinas/alice"
)

type city struct {
	Name string
	Area uint64
}

func addCity(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		var tempCity city
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&tempCity)
		defer req.Body.Close()
		if err != nil {
			log.Println(err)
		}
		log.Printf("City Name: %s\tCity Area: %d", tempCity.Name, tempCity.Area)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed\n"))
	}
}
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type\n"))
			return
		}
		handler.ServeHTTP(w, req)
	})
}
func addTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		handler.ServeHTTP(w, req)
		cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)
	})
}
func main() {
	// http.Handle("/city", filterContentType(addTimeCookie(cityHandler)))
	// http.Handle("/city", chain)
	cityHandler := http.HandlerFunc(addCity)
	chain := alice.New(filterContentType, addTimeCookie).Then(cityHandler)
	r := mux.NewRouter()
	r.Handle("/city", chain)
	loggingRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":8000", loggingRouter)
}
