package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//go:embed templates/* static/*
var assets embed.FS

type Post struct {
	ID    int
	Title string
	Body  string
}

var posts = []Post{
	{ID: 1, Title: "First post", Body: "This is the first post."},
	{ID: 2, Title: "Second post", Body: "This is the second post."},
}

var templates = template.Must(template.ParseFS(assets, "templates/*.html"))

func main() {
	mux := http.NewServeMux()

	staticFS, err := fs.Sub(assets, "static")
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))
	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("GET /posts/", postHandler)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if err := templates.ExecuteTemplate(w, "index.html", map[string]any{"Posts": posts}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/posts/"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	for _, post := range posts {
		if post.ID == id {
			if err := templates.ExecuteTemplate(w, "post.html", post); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	http.NotFound(w, r)
}
