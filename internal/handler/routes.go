package handler

import (
	"net/http"
)

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("GET /", h.HandleIndex)
	mux.HandleFunc("GET /search", h.HandleSearch)
	mux.HandleFunc("GET /tags/{tag}", h.HandleTag)
	mux.HandleFunc("GET /posts/{slug}", h.HandlePost)
	mux.HandleFunc("GET /rss.xml", h.HandleRSS)

	return mux
}
