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
		Addr:     "redis.railway.internal:6379",
		Password: "uElJcVUUNQNVZroVszeZsdPxuydgDDMv", // Ensure Redis is running locally
	})

	r := mux.NewRouter()
	r.HandleFunc("/", redir).Methods("GET")
	r.HandleFunc("/{file}", serveContent).Methods("GET")

	server := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("CDN Edge Server running on port 8080")
	log.Fatal(server.ListenAndServe())
}

func redir(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func serveContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file := vars["file"]

	// Check Redis Cache
	start := time.Now()
	content, err := redisClient.Get(ctx, file).Result()
	elapsed := time.Since(start).Microseconds()

	header := fmt.Sprintf("Fetched from US server cache in %d µs", elapsed)
	pageData := PageData{Header: header}

	if err == nil {
		// Parse the cached template string
		tmpl, err := template.New("cached").Parse(content)
		if err != nil {
			http.Error(w, "Template Parse Error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, pageData)
		return
	}

	// Cache Miss: Fetch from Origin
	// resp, err := http.Get("https://origin-production.up.railway.app/" + file)
	resp, err := http.Get("http://localhost:9090/" + file)
	if err != nil {
		http.Error(w, "Origin Server Down", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	content = string(body)

	header = fmt.Sprintf("Fetched from Origin Server in %d µs", elapsed)
	pageData = PageData{Header: header}

	// Parse the fetched HTML
	tmpl, err := template.New("origin").Parse(content)
	if err != nil {
		http.Error(w, "Origin Template Parse Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, pageData)

	// Store in Redis cache
	redisClient.Set(ctx, file, content, 10*time.Minute)
}
