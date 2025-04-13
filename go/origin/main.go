package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type PageData struct {
	Header string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", redir).Methods("GET")
	r.HandleFunc("/home", serveMain).Methods("GET")
	r.HandleFunc("/architecture", serveArchitecture).Methods("GET")
	r.HandleFunc("/details", serveDetails).Methods("GET")

	fmt.Println("Origin Server running on port 9090")
	log.Fatal(http.ListenAndServe(":9090", r))
}

func redir(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func serveMain(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("pages/main.html"))
	pageData := PageData{
		Header: "",
	}
	tmpl.Execute(w, pageData)
}
func serveArchitecture(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("pages/arch.html"))
	pageData := PageData{
		Header: "",
	}
	tmpl.Execute(w, pageData)
}
func serveDetails(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("pages/desc.html"))
	pageData := PageData{
		Header: "",
	}
	tmpl.Execute(w, pageData)
}
