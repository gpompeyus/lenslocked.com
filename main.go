package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	homeTemplate *template.Template
	contactTemplate *template.Template
)


func home (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err:= homeTemplate.Execute(w,nil);err != nil {
		panic(err)
	}
}

func contact (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err:= contactTemplate.Execute(w,nil);err != nil {
		panic(err)
	}
}

func faq (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Coloca aqui tus dudas y te responderemos lo antes posible")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "No pudimos encontrar la pagina buscada")
}

func main() {
	var err error
	homeTemplate, err= template.ParseFiles("views/home.gohtml")
	if err != nil{
		panic(err)
	}

	contactTemplate, err= template.ParseFiles("views/contact.gohtml")
	if err != nil{
		panic(err)
	}

	router:=mux.NewRouter()
	router.NotFoundHandler=http.HandlerFunc(notFound)
	router.HandleFunc("/",home)
	router.HandleFunc("/contact",contact)
	router.HandleFunc("/faq",faq)
	http.ListenAndServe(":3000", router)
}