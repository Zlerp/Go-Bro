package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

type WebData struct {
	Title string
	Name  string
}

func home(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "You home, mah nigga!!")
}

func about(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "we are Team Supreme, stay clean ;)")
}

func contact(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Contact us right hurrrrrr !!!")
}

func html(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("layout.gohtml", "tpl.gohtml" )
	if err != nil {
		log.Fatalln(err)
	}
	wd := WebData{
		Title: "Home",
		Name:  "Zach",
	}

	tpl.Execute(res, &wd)
}

func main() {

	http.HandleFunc("/home", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/html", html)

	http.ListenAndServe(":8080", nil)
}
