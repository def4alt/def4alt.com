package handler

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"def4alt/def4alt.com/internal/content"
	"def4alt/def4alt.com/internal/render"
)

type Handler struct {
	render *render.Renderer
	blog   *content.Blog
}

type ViewData struct {
	Title       string
	Description string
	Query       string
	Tag         string
	Tags        []string
	Posts       []content.Post
	Post        content.Post
}

func New(r *render.Renderer, blog *content.Blog) *Handler {
	return &Handler{render: r, blog: blog}
}

func (h *Handler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	if err := h.render.Page(w, "index", ViewData{
		Description: "A small Go + HTMX blog",
		Tags:        h.blog.Tags(),
		Posts:       h.blog.Posts(),
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) HandleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	tag := r.URL.Query().Get("tag")
	posts := h.blog.Search(query, tag)

	data := ViewData{
		Title:       "Search",
		Description: "Search results",
		Query:       query,
		Tag:         tag,
		Tags:        h.blog.Tags(),
		Posts:       posts,
	}

	if err := h.render.Page(w, "index", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) HandleTag(w http.ResponseWriter, r *http.Request) {
	tag := strings.ToLower(strings.TrimSpace(r.PathValue("tag")))
	posts := h.blog.Search("", tag)

	data := ViewData{
		Title:       fmt.Sprintf("Tag: %s", tag),
		Description: fmt.Sprintf("Posts tagged %s", tag),
		Tag:         tag,
		Tags:        h.blog.Tags(),
		Posts:       posts,
	}

	if err := h.render.Page(w, "tag", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) HandlePost(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimSpace(r.PathValue("slug"))
	post, ok := h.blog.BySlug(slug)
	if !ok {
		http.NotFound(w, r)
		return
	}

	if err := h.render.Page(w, "post", ViewData{
		Title:       post.Title,
		Description: post.Description,
		Tags:        h.blog.Tags(),
		Post:        post,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) HandleRSS(w http.ResponseWriter, r *http.Request) {
	baseURL := "http://" + r.Host
	if r.Host == "" {
		baseURL = "http://localhost:8080"
	}

	items := make([]rssItem, 0, len(h.blog.Posts()))
	for _, post := range h.blog.Posts() {
		items = append(items, rssItem{
			Title:       post.Title,
			Link:        baseURL + "/posts/" + url.PathEscape(post.Slug),
			Description: post.Description,
			PubDate:     post.Date.Format(time.RFC1123Z),
		})
	}

	feed := rssFeed{
		Version: "2.0",
		Channel: rssChannel{
			Title:       "Blog",
			Link:        baseURL,
			Description: "Blog posts",
			Items:       items,
		},
	}

	w.Header().Set("Content-Type", "application/rss+xml; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = xml.NewEncoder(w).Encode(feed)
}

type rssFeed struct {
	XMLName xml.Name   `xml:"rss"`
	Version string     `xml:"version,attr"`
	Channel rssChannel `xml:"channel"`
}

type rssChannel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Items       []rssItem `xml:"item"`
}

type rssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}
