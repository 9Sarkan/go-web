package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func bookHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Fprintf(w, "requested book name: %s\n", vars["book"])
	fmt.Fprintf(w, "requested page: %s", vars["page"])
}

func main() {
	newMux := mux.NewRouter()
	newMux.HandleFunc("/book/{book}/{page:[0-9]+}", bookHandler).Name("BookPage")
	serv := &http.Server{
		Handler:      newMux,
		Addr:         "localhost:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	url, _ := newMux.Get("BookPage").URL("book", "hello-world", "page", "10")
	fmt.Printf("Book Page Url: %s\n", url.String())
	log.Fatal(serv.ListenAndServe())
}
