package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var redisClient *redis.Client

func main() {
	// Initialize Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Ensure Redis is running locally
	})

	r := mux.NewRouter()
	r.HandleFunc("/content/{file}", serveContent).Methods("GET")
	r.HandleFunc("/purge/{file}", purgeCache).Methods("POST")

	server := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("CDN Edge Server running on port 8080")
	log.Fatal(server.ListenAndServe())
}

// Serve content (check cache, then origin)
func serveContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file := vars["file"]

	// Check Redis Cache
	content, err := redisClient.Get(ctx, file).Result()
	if err == nil {
		w.Write([]byte(content + " (from US cache)"))
		return
	}

	// Cache Miss: Fetch from Origin
	resp, err := http.Get("http://localhost:9090/content/" + file) // Using localhost for testing
	if err != nil {
		http.Error(w, "Origin Server Down", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	content = string(body)
	w.Write([]byte(content + " (from origin)"))

	// Store in cache
	redisClient.Set(ctx, file, content, 10*time.Minute)
}

// Purge cache
func purgeCache(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file := vars["file"]

	redisClient.Del(ctx, file)
	w.Write([]byte("Cache Purged for " + file))
}
