package main

import (
	"net/http"
	"github.com/sdaf47/stt/lib/html"
)

const (
	TemplateIndex = "/"
	TemplateStatic = "/static/"
	TemplateCase = "/case/"
)

func main() {

	fs := http.FileServer(http.Dir("."))
	http.Handle(TemplateStatic, fs)

	http.HandleFunc(TemplateIndex, html.Index)
	http.HandleFunc(TemplateCase, html.Case)

	err := http.ListenAndServe(":9988", nil)
	if err != nil {
		panic(err)
	}
}
