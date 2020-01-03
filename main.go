package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contact (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Para mayor informacion.. bla bla")
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
	r:=mux.NewRouter()
	r.NotFoundHandler=http.HandlerFunc(notFound)
	r.HandleFunc("/",home)
	r.HandleFunc("/contact",contact)
	r.HandleFunc("/faq",faq)
	http.ListenAndServe(":3000", r)
}