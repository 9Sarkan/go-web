package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func getCommandOutput(command string, arguments ...string) string {
	cmd := exec.Command(command, arguments...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
	return out.String()
}
func getGoVersion(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "%s", getCommandOutput("go", "version"))
}
func cat(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	fmt.Fprint(w, getCommandOutput("/bin/cat", params.ByName("name")))
}
func main() {
	newRouter := httprouter.New()
	newRouter.GET("/api/v1/go-version", getGoVersion)
	newRouter.GET("/api/v1/cat/:name", cat)
	log.Fatal(http.ListenAndServe(":8080", newRouter))
}
