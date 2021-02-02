package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
)

func pingTime(req *restful.Request, response *restful.Response) {
	io.WriteString(response, fmt.Sprintf("%s", time.Now()))
}
func main() {
	restservice := new(restful.WebService)
	restservice.Route(restservice.GET("/ping").To(pingTime))
	restful.Add(restservice)
	http.ListenAndServe(":8000", nil)
}
