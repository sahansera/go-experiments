package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/api/v1/kernal-version", kernelVer)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getCommandOutput(command string, arguments ...string) string {
	out, _ := exec.Command(command, arguments...).Output()
	return string(out)
}

func kernelVer(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response := getCommandOutput("uname", "-a")
	io.WriteString(w, response)
}
