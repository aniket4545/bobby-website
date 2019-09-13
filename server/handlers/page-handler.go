package handlers

import (
	"html/template"
	"net/http"
)

//InitialRequest this is the first empty handler
func InitialRequest(wr http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("client/sample.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(wr, nil)
}
