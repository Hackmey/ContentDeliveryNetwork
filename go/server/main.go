package main

import (
	"fmt"
	"log"
	"net/http"
)



const contentDir="./content"

func main(){
	http.HandleFunc("/",serveFile)
	fmt.Println(("CDN edge server is running on port 8000..."))
	log.Fatal(http.ListenAndServe(":8000",nil))
}

func serveFile(w http.ResponseWriter, r *http.Request){
	filePath := contentDir + r.URL.Path
	http.ServeFile(w,r,filePath)
}