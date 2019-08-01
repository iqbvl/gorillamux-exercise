package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hey, this is homepage")
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hey, this is product")
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hey, this is article")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/products", productHandler)
	r.HandleFunc("/articles", articleHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
