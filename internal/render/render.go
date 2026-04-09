package render

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Renderer struct {
	pages map[string]*template.Template
}

func New(root string) (*Renderer, error) {
	base := filepath.Join(root, "templates", "layouts", "base.html")

	partialFiles, err := glob(filepath.Join(root, "templates", "partials", "*.html"))
	if err != nil {
		return nil, err
	}

	pageFiles, err := glob(filepath.Join(root, "templates", "pages", "*.html"))
	if err != nil {
		return nil, err
	}

	funcs := template.FuncMap{}
	pages := make(map[string]*template.Template, len(pageFiles))

	for _, pageFile := range pageFiles {
		name := strings.TrimSuffix(filepath.Base(pageFile), filepath.Ext(pageFile))
		files := append([]string{base}, partialFiles...)
		files = append(files, pageFile)

		tmpl, err := parseTemplateSet(funcs, files...)
		if err != nil {
			return nil, err
		}

		pages[name] = tmpl
	}

	return &Renderer{pages: pages}, nil
}

func (r *Renderer) Page(w http.ResponseWriter, name string, data any) error {
	tmpl, ok := r.pages[name]
	if !ok {
		http.Error(w, "page not found", http.StatusNotFound)
		return nil
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tmpl.ExecuteTemplate(w, "base", data)
}

func glob(pattern string) ([]string, error) {
	paths, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	sort.Strings(paths)
	return paths, nil
}

func parseTemplateSet(funcs template.FuncMap, files ...string) (*template.Template, error) {
	files = normalizeFiles(files)
	if len(files) == 0 {
		return template.New("").Funcs(funcs), nil
	}

	return template.New("").Funcs(funcs).ParseFiles(files...)
}

func normalizeFiles(files []string) []string {
	cleaned := make([]string, 0, len(files))
	seen := make(map[string]struct{}, len(files))

	for _, file := range files {
		if file == "" {
			continue
		}
		if _, err := os.Stat(file); err != nil {
			continue
		}
		if _, ok := seen[file]; ok {
			continue
		}

		seen[file] = struct{}{}
		cleaned = append(cleaned, file)
	}

	return cleaned
}
