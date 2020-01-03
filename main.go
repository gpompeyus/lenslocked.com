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
		w.Header().Set("Content-Type","text/html")
		fmt.Fprint(w, "Para mayor informacion.. bla bla")

}

func main() {
	r:=mux.NewRouter()
	r.HandleFunc("/",home)
	r.HandleFunc("/contact",contact)
	http.ListenAndServe(":3000", r)
}
