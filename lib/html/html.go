package html

import (
	"fmt"
	"net/http"
	"html/template"
)

type Page struct {
	Title string
}

func execute(tmp string, statement interface{}, w http.ResponseWriter, r *http.Request) {
	fmap := template.FuncMap{}

	fmt.Println()

	h := template.Must(template.New("header.html").Funcs(fmap).ParseFiles("templates/header.html"))
	err := h.Execute(w, statement)
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New(tmp).Funcs(fmap).ParseFiles("templates/" + tmp))
	err = t.Execute(w, statement)
	if err != nil {
		panic(err)
	}

	f := template.Must(template.New("footer.html").Funcs(fmap).ParseFiles("templates/footer.html"))
	err = f.Execute(w, statement)
	if err != nil {
		panic(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	statement := Page{
		Title: "Заголовок",
	}
	execute("index.html", statement, w, r)
}

func Case(w http.ResponseWriter, r *http.Request) {
	statement := Page{
		Title: "Заголовок",
	}
	execute("case.html", statement, w, r)
}

