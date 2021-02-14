package main

import (
	"log"

	"github.com/levigross/grequests"
)

func main() {
	response, err := grequests.Get("http://httpbin.org/get", nil)
	if err != nil {
		log.Fatal(err)
	}
	var responseMap map[string]interface{}
	response.JSON(&responseMap)
	log.Println(responseMap)
}
