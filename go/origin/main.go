package main

import (
	"fmt"
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

	http.ServeFile(w, r, "pages/main.html")
}
func serveArchitecture(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pages/arch.html")
}
func serveDetails(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pages/desc.html")

}
