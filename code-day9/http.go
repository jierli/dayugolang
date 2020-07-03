package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Dayu1 struct {
}

func (h *Dayu1) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "dayu1")
}

type Dayu2 struct {
}

func (h *Dayu2) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Printf("%T,,,,,,%#v\n", response, response)
	fmt.Fprintln(response, "dayu2")
}

type Dayu3 struct {
}

func (h *Dayu3) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("dayu3"))
}
func main() {
	http.Handle("/dayu/", &Dayu1{})
	http.Handle("/dayu2/", &Dayu2{})
	http.Handle("/dayu3/", &Dayu3{})
	http.ListenAndServe(":8080", nil)
}
