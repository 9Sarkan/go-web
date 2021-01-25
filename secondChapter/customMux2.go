package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func randomInt(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "your random int: %d\n", rand.Intn(100))
}
func randomFloat(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "your random float: %v\n", rand.Float32())
}
func main() {
	customMux := http.NewServeMux()
	customMux.HandleFunc("/random-int", randomInt)
	customMux.HandleFunc("/random-float", randomFloat)
	http.ListenAndServe(":8080", customMux)
}
