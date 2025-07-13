package handlers

import (
	"encoding/json"
	"io"
	"meme-api/internal/meme-api/services"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusFound)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func MemeStreamHandler(w http.ResponseWriter, r *http.Request) {
	memeService := services.NewMemeService()
	meme, err := memeService.GetRandomMeme()
	if err != nil || meme.Url == "" {
		http.Error(w, "Failed to fetch meme", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("GET", meme.Url, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to fetch meme image", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Header().Set("Cache-Control", "no-cache,max-age=0,no-store,s-maxage=0,proxy-revalidate")
	w.Header().Set("Content-Disposition", "inline")
	w.WriteHeader(http.StatusOK)
	// Stream the image directly to the response
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Failed to stream image", http.StatusInternalServerError)
	}
}

func MemeBySubredditHandler(w http.ResponseWriter, r *http.Request) {
	subreddit := r.URL.Query().Get("subreddit")
	memeService := services.NewMemeService()
	meme, err := memeService.GetMemeBySubreddit(subreddit)
	if err != nil || meme.Url == "" {
		http.Error(w, "Failed to fetch meme", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("GET", meme.Url, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to fetch meme image", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Header().Set("Cache-Control", "no-cache,max-age=0,no-store,s-maxage=0,proxy-revalidate")
	w.Header().Set("Content-Disposition", "inline")
	w.WriteHeader(http.StatusOK)
	// Stream the image directly to the response
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Failed to stream image", http.StatusInternalServerError)
	}
}
