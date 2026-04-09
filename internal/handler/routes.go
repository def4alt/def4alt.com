package handler

import (
	"mime"
	"net/http"
	"path/filepath"
)

func init() {
	_ = mime.AddExtensionType(".webp", "image/webp")
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(h.root, "static")))))
	mux.Handle("GET /content/images/", http.StripPrefix("/content/images/", http.FileServer(http.Dir(filepath.Join(h.root, "content", "images")))))
	mux.HandleFunc("GET /healthz", h.HandleHealthz)
	mux.HandleFunc("GET /", h.HandleIndex)
	mux.HandleFunc("GET /search", h.HandleSearch)
	mux.HandleFunc("GET /tags/{tag}", h.HandleTag)
	mux.HandleFunc("GET /posts/{slug}", h.HandlePost)
	mux.HandleFunc("GET /rss.xml", h.HandleRSS)

	return mux
}
