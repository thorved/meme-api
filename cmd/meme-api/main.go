package main

import (
	"meme-api/internal/meme-api/routes"
	"net/http"
)

func main() {
	routes.RegisterRoutes()
	http.ListenAndServe(":3000", nil)
}
