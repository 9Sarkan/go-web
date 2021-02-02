package main

import (
	jsonparse "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

// Args rpc
type Args struct {
	ID string
}

// Book rpc
type Book struct {
	ID     string `"json:string,omitempty"`
	Name   string `"json:name,omitempty"`
	Author string `"json:author,omitempty"`
}

// JSONServer rpc
type JSONServer struct{}

// GetBookDetail get a book detail with json RPC
func (b *JSONServer) GetBookDetail(req http.Request, args *Args, reply *Book) error {
	var books []Book
	raw, err := ioutil.ReadFile("books.json")
	if err != nil {
		log.Fatalln("read error: ", err)
	}
	err = jsonparse.Unmarshal(raw, &books)
	if err != nil {
		log.Fatalln("unmarshal error: ", err)
	}
	for _, book := range books {
		if book.ID == args.ID {
			*reply = book
			break
		}
	}
	return nil
}
func main() {
	server := new(JSONServer)
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterService(server, "JSONServer")
	r := mux.NewRouter()
	r.Handle("/rpc", rpcServer)
	loggingRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":1234", loggingRouter)
}
