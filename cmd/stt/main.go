package main

import (
	"net/http"
	"github.com/sdaf47/stt/lib/html"
	"github.com/sdaf47/stt/lib/app"
	"github.com/sdaf47/stt/lib/db"
	"time"
)

const (
	TemplateIndex  = "/"
	TemplateStatic = "/static/"
	TemplateCase   = "/case/"
)

var App = &app.Config{}

func main() {

	App.Load()

	Sess, err := db.NewClient(App.Mongo.HostPort)
	if err != nil {
		panic(err)
	}
	defer Sess.Close()



	fs := http.FileServer(http.Dir("."))
	http.Handle(TemplateStatic, fs)

	http.HandleFunc(TemplateIndex, html.Index)
	http.HandleFunc(TemplateCase, html.Case)

	err = http.ListenAndServe(":9988", nil)
	if err != nil {
		panic(err)
	}
}
