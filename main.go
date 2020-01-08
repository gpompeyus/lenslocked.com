package main

import (
	"fmt"
	"lenslocked.com/views"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	homeView *views.View
	contactView *views.View
)


func home (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err:= homeView.Template.Execute(w,nil)
	if err != nil {
		panic(err)
	}
}

func contact (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err:= contactView.Template.Execute(w,nil)
	if err != nil {
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
	homeView=views.NewView("views/home.gohtml")
	contactView=views.NewView("views/contact.gohtml")

	router:=mux.NewRouter()
	router.NotFoundHandler=http.HandlerFunc(notFound)
	router.HandleFunc("/",home)
	router.HandleFunc("/contact",contact)
	router.HandleFunc("/faq",faq)
	http.ListenAndServe(":3000", router)
}