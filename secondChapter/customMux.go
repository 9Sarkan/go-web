package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

// CustomServeMux for route
type CustomServeMux struct{}

func (c *CustomServeMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		giveRandom(w, req)
		return
	}
	http.NotFound(w, req)
	return
}
func giveRandom(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "your random number is: %d\n", rand.Intn(100))
}
func main() {
	mux := &CustomServeMux{}
	http.ListenAndServe(":8080", mux)
}
