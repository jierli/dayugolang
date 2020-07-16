package main

import (
	"net/http"

	"todolist/handler"
)

func main()  {
	addr:=":8080"
	var task handler.Task
	http.HandleFunc("/",task.Index)

	http.ListenAndServe(addr,nil)
}