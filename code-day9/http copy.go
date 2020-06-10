package main

import (
	"io"
	"net/http"
)

// 处理器
// Handler接口
// ServeHTTP(http.ResponseWriter, *http.Request)

type TimeHandler struct {
}

func (h *TimeHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "dayu")
}

func main() {
	http.Handle("/time/", &TimeHandler{})

	http.ListenAndServe(":9998", nil)
}
