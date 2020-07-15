package handler

import (
	"html/template"
	"net/http"
)

type QueryHandler struct {
}

func (h *QueryHandler) ServeHTTP(response http.ResponseWriter, request *http.Request)  {
	request.ParseForm()
	tpl := template.Must(template.ParseFiles("views/index.html"))
	tpl.ExecuteTemplate(response, "index.html", []int{1, 2, 3})
}
