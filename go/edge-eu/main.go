package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var redisClient *redis.Client

type PageData struct {
	Header string
}

func main() {
	// Initialize Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis-copy.railway.internal:6379",
		Password: "rLCBzTqNcJJJtTIRnzLPGnEsElMbvosI",
	})

	r := mux.NewRouter()
	r.HandleFunc("/{file}", serveContent).Methods("GET")

	server := &http.Server{
		Handler:      r,
		Addr:         ":8090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("CDN Edge Server running on port 8090")
	log.Fatal(server.ListenAndServe())
}

// Serve content (check cache, then origin)
func serveContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file := vars["file"]

	// Check Redis Cache
	start := time.Now()
	content, err := redisClient.Get(ctx, file).Result()
	if err == nil {
		elapsed := time.Since(start).Microseconds()
		tmpl := template.Must(template.ParseFiles("content"))
		header := fmt.Sprintf("Fetched from EU server cache in %d ms", elapsed)
		pageData := PageData{
			Header: header,
		}
		tmpl.Execute(w, pageData)
		return
	}

	// Cache Miss: Fetch from Origin
	resp, err := http.Get("https://origin-production.up.railway.app/" + file) // Using localhost for testing
	if err != nil {
		http.Error(w, "Origin Server Down", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	content = string(body)
	elapsed := time.Since(start).Microseconds()
	header := fmt.Sprintf("Fetched from Origin Server in %d ms", elapsed)
	template.Must(template.ParseFiles(content)).Execute(w, header)

	// Store in cache
	redisClient.Set(ctx, file, content, 10*time.Minute)
}
