package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/content/{file}", serveContent).Methods("GET")

	fmt.Println("Origin Server running on port 9090")
	log.Fatal(http.ListenAndServe(":9090", r))
}

func serveContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file := vars["file"]

	// Simulate content based on file name
	content := fmt.Sprintf("Original content of file: %s from ORIGIN", file)
	w.Write([]byte(content))
}
