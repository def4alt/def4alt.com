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

func TestHealthzReturnsOk(t *testing.T) {
	t.Parallel()

	app := newTestApp(t)

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rr := httptest.NewRecorder()

	app.Routes().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	if body := rr.Body.String(); body != "ok" {
		t.Fatalf("expected body ok, got %q", body)
	}
}

func TestContentImagesAreServed(t *testing.T) {
	t.Parallel()

	app := newTestApp(t)

	req := httptest.NewRequest(http.MethodGet, "/content/images/gallery-1.webp", nil)
	rr := httptest.NewRecorder()

	app.Routes().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); !strings.Contains(ct, "image/webp") {
		t.Fatalf("expected webp content type, got %q", ct)
	}
	body := rr.Body.Bytes()
	if len(body) < 12 || string(body[:4]) != "RIFF" || string(body[8:12]) != "WEBP" {
		t.Fatalf("expected webp bytes, got %x", body[:12])
	}
}

func TestSearchFiltersResults(t *testing.T) {
	t.Parallel()

	app := newTestApp(t)

	req := httptest.NewRequest(http.MethodGet, "/search?q=missing", nil)
	req.Header.Set("HX-Request", "true")
	rr := httptest.NewRecorder()

	app.Routes().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	body := rr.Body.String()
	if !strings.Contains(body, "No posts found.") {
		t.Fatalf("expected empty search results, got %s", body)
	}
}

func TestIndexRendersPosts(t *testing.T) {
	t.Parallel()

	app := newTestApp(t)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	app.Routes().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	body := rr.Body.String()
	if !strings.Contains(body, "server-rendering") {
		t.Fatalf("expected posts on index, got %s", body)
	}
}

func TestSearchShowsActiveTagAndClearsItWhenClicked(t *testing.T) {
	t.Parallel()

	app := newTestApp(t)

	req := httptest.NewRequest(http.MethodGet, "/search?q=server&tag=go", nil)
	rr := httptest.NewRecorder()

	app.Routes().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	body := rr.Body.String()
	if !strings.Contains(body, `href="/search?q=server"`) {
		t.Fatalf("expected active tag to clear tag filter, got %s", body)
	}
	if !strings.Contains(body, "server-rendering") {
		t.Fatalf("expected matching post, got %s", body)
	}
}

func TestPostRendersMarkdownFeatures(t *testing.T) {
	t.Parallel()

	app := newTestApp(t)

	req := httptest.NewRequest(http.MethodGet, "/posts/server-rendering", nil)
	rr := httptest.NewRecorder()

	app.Routes().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	body := rr.Body.String()
	if !strings.Contains(body, "gallery-1.webp") {
		t.Fatalf("expected inline image, got %s", body)
	}
	if !strings.Contains(body, "<blockquote>") || !strings.Contains(body, "<table>") || !strings.Contains(body, "<pre><code") || !strings.Contains(body, "<del>") || !strings.Contains(body, "type=\"checkbox\"") {
		t.Fatalf("expected markdown features, got %s", body)
	}
	if !strings.Contains(body, "Apr 8, 2026") {
		t.Fatalf("expected date in header, got %s", body)
	}
}

func TestPostMissingReturns404(t *testing.T) {
	t.Parallel()

	app := newTestApp(t)

	req := httptest.NewRequest(http.MethodGet, "/posts/missing", nil)
	rr := httptest.NewRecorder()

	app.Routes().ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", rr.Code)
	}
}

func TestTagRouteRendersFullPage(t *testing.T) {
	t.Parallel()

	app := newTestApp(t)

	req := httptest.NewRequest(http.MethodGet, "/tags/go", nil)
	rr := httptest.NewRecorder()

	app.Routes().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	body := rr.Body.String()
	if !strings.Contains(body, "server-rendering") {
		t.Fatalf("expected tagged post, got %s", body)
	}
	if strings.Contains(body, "No posts found.") {
		t.Fatalf("unexpected empty state in %s", body)
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
	if !strings.Contains(body, "/posts/server-rendering") {
		t.Fatalf("expected post in feed, got %s", body)
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
	copyFixtureTree(t, filepath.Join("..", "..", "testdata", "handler"), root)
	copyFixtureTree(t, filepath.Join("..", "..", "content", "images"), filepath.Join(root, "content", "images"))

	blog, err := content.Load(filepath.Join(root, "content"))
	if err != nil {
		t.Fatalf("load content: %v", err)
	}

	r, err := render.New(root)
	if err != nil {
		t.Fatalf("render: %v", err)
	}

	h := handler.New(root, r, blog)
	return testApp{routes: h.Routes()}
}

func copyFixtureTree(t *testing.T, src, dst string) {
	t.Helper()

	entries, err := os.ReadDir(src)
	if err != nil {
		t.Fatalf("read fixtures: %v", err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		if entry.IsDir() {
			copyFixtureTree(t, srcPath, dstPath)
			continue
		}

		data, err := os.ReadFile(srcPath)
		if err != nil {
			t.Fatalf("read fixture %s: %v", srcPath, err)
		}
		if err := os.MkdirAll(filepath.Dir(dstPath), 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", filepath.Dir(dstPath), err)
		}
		if err := os.WriteFile(dstPath, data, 0o644); err != nil {
			t.Fatalf("write fixture %s: %v", dstPath, err)
		}
	}
}
