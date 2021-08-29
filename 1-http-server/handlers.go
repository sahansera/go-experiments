package main

import (
	"fmt"
	"log"
	"net/http"
)

type FooHandler struct{}

func (h *FooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Foo!")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bar!")
}

func main() {
	fooHandler := &FooHandler{}
	http.Handle("/foo", fooHandler)
	http.HandleFunc("/bar", barHandler)

	log.Println("Starting a server at port 8080")
	http.ListenAndServe(":8080", nil)
}
