package routes

import (
	"meme-api/internal/meme-api/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/meme", handlers.MemeStreamHandler)
	http.HandleFunc("/meme/", func(w http.ResponseWriter, r *http.Request) {
		prefix := "/meme/"
		subreddit := r.URL.Path[len(prefix):]
		if subreddit == "" || r.URL.Path == "/meme/" {
			http.NotFound(w, r)
			return
		}
		// Optionally validate subreddit here
		r.URL.RawQuery = "subreddit=" + subreddit
		handlers.MemeBySubredditHandler(w, r)
	})
}
