package handler_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"def4alt/def4alt.com/internal/content"
	"def4alt/def4alt.com/internal/handler"
	"def4alt/def4alt.com/internal/render"
)

func TestSearchReturnsMatchingPostsAsFragment(t *testing.T) {
	t.Parallel()

	app := newTestApp(t)

	req := httptest.NewRequest(http.MethodGet, "/search?q=go", nil)
	req.Header.Set("HX-Request", "true")
	rr := httptest.NewRecorder()

	app.Routes().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	body := rr.Body.String()
	if !strings.Contains(body, "Go Tips") {
		t.Fatalf("expected go result, got %s", body)
	}
	if strings.Contains(body, "HTMX Guide") {
		t.Fatalf("unexpected non-matching post in %s", body)
	}
}

func TestTagRouteRendersFullPage(t *testing.T) {
	t.Parallel()

	app := newTestApp(t)

	req := httptest.NewRequest(http.MethodGet, "/tags/htmx", nil)
	rr := httptest.NewRecorder()

	app.Routes().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	body := rr.Body.String()
	if !strings.Contains(body, "Tag: htmx") {
		t.Fatalf("expected tag heading, got %s", body)
	}
	if !strings.Contains(body, "HTMX Guide") {
		t.Fatalf("expected tagged post, got %s", body)
	}
}

func TestRSSReturnsFeedWithPosts(t *testing.T) {
	t.Parallel()

	app := newTestApp(t)

	req := httptest.NewRequest(http.MethodGet, "/rss.xml", nil)
	rr := httptest.NewRecorder()

	app.Routes().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	body := rr.Body.String()
	if !strings.Contains(body, "<rss") {
		t.Fatalf("expected rss xml, got %s", body)
	}
	if !strings.Contains(body, "Go Tips") || !strings.Contains(body, "HTMX Guide") {
		t.Fatalf("expected posts in feed, got %s", body)
	}
}

type testApp struct {
	routes http.Handler
}

func (a testApp) Routes() http.Handler {
	return a.routes
}

func newTestApp(t *testing.T) testApp {
	t.Helper()

	root := t.TempDir()
	mustMkdir(t, filepath.Join(root, "content", "posts"))
	mustMkdir(t, filepath.Join(root, "templates", "layouts"))
	mustMkdir(t, filepath.Join(root, "templates", "pages"))
	mustMkdir(t, filepath.Join(root, "templates", "partials"))

	mustWrite(t, filepath.Join(root, "content", "posts", "go.md"), `---
title: Go Tips
slug: go-tips
date: 2024-02-01
description: Practical Go
tags: go, backend
---
This post covers server-side Go.
`)
	mustWrite(t, filepath.Join(root, "content", "posts", "htmx.md"), `---
title: HTMX Guide
slug: htmx-guide
date: 2024-01-01
description: HTMX patterns
tags: htmx, frontend
---
This post covers live search.
`)

	mustWrite(t, filepath.Join(root, "templates", "layouts", "base.html"), `{{define "base"}}<!doctype html><html><body>{{template "content" .}}</body></html>{{end}}`)
	mustWrite(t, filepath.Join(root, "templates", "pages", "index.html"), `{{define "content"}}<h1>Blog</h1>{{template "search_bar" .}}{{template "post_list" .}}{{end}}`)
	mustWrite(t, filepath.Join(root, "templates", "pages", "post.html"), `{{define "content"}}<article><h1>{{.Post.Title}}</h1>{{.Post.HTML}}</article>{{end}}`)
	mustWrite(t, filepath.Join(root, "templates", "pages", "tag.html"), `{{define "content"}}<h1>Tag: {{.Tag}}</h1>{{template "post_list" .}}{{end}}`)
	mustWrite(t, filepath.Join(root, "templates", "partials", "search_bar.html"), `{{define "search_bar"}}<form hx-get="/search" hx-target="#results" hx-push-url="true"><input name="q" value="{{.Query}}"><input type="hidden" name="tag" value="{{.Tag}}"></form>{{end}}`)
	mustWrite(t, filepath.Join(root, "templates", "partials", "post_list.html"), `{{define "post_list"}}<section id="results">{{if .Posts}}{{range .Posts}}{{template "post_list_item" .}}{{end}}{{else}}<p>No posts found.</p>{{end}}</section>{{end}}`)
	mustWrite(t, filepath.Join(root, "templates", "partials", "post_list_item.html"), `{{define "post_list_item"}}<article><a href="/posts/{{.Slug}}">{{.Title}}</a>{{template "post_tags" .}}</article>{{end}}`)
	mustWrite(t, filepath.Join(root, "templates", "partials", "post_tags.html"), `{{define "post_tags"}}{{if .Tags}}<ul>{{range .Tags}}<li><a href="/search?tag={{.}}">{{.}}</a></li>{{end}}</ul>{{end}}{{end}}`)

	blog, err := content.Load(filepath.Join(root, "content"))
	if err != nil {
		t.Fatalf("load content: %v", err)
	}

	r, err := render.New(root)
	if err != nil {
		t.Fatalf("render: %v", err)
	}

	h := handler.New(r, blog)
	return testApp{routes: h.Routes()}
}

func mustMkdir(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", path, err)
	}
}

func mustWrite(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
