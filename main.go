package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	mux := http.NewServeMux()

	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
