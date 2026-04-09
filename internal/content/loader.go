package content

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
)

type frontMatter struct {
	Title       string
	Slug        string
	Date        time.Time
	Description string
	Tags        []string
	Draft       bool
}

func Load(root string) (*Blog, error) {
	paths, err := filepath.Glob(filepath.Join(root, "posts", "*.md"))
	if err != nil {
		return nil, err
	}

	posts := make([]Post, 0, len(paths))
	for _, path := range paths {
		post, err := parsePost(path)
		if err != nil {
			return nil, err
		}
		if post.Draft {
			continue
		}
		posts = append(posts, post)
	}

	sort.Slice(posts, func(i, j int) bool {
		if posts[i].Date.Equal(posts[j].Date) {
			return posts[i].Slug < posts[j].Slug
		}
		return posts[i].Date.After(posts[j].Date)
	})

	index := make(map[string]Post, len(posts))
	for _, post := range posts {
		index[strings.ToLower(post.Slug)] = post
	}

	return &Blog{
		posts:  posts,
		bySlug: index,
		tags:   uniqueTags(posts),
	}, nil
}

func parsePost(path string) (Post, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Post{}, err
	}

	fm, body, err := splitFrontMatter(string(data))
	if err != nil {
		return Post{}, fmt.Errorf("%s: %w", path, err)
	}

	meta, err := parseFrontMatter(fm)
	if err != nil {
		return Post{}, fmt.Errorf("%s: %w", path, err)
	}

	if meta.Slug == "" {
		meta.Slug = strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	}

	var html bytes.Buffer
	if err := goldmark.Convert([]byte(body), &html); err != nil {
		return Post{}, fmt.Errorf("%s: %w", path, err)
	}

	searchText := strings.ToLower(strings.Join([]string{
		meta.Title,
		meta.Description,
		strings.Join(meta.Tags, " "),
		body,
	}, " "))

	return Post{
		Title:       meta.Title,
		Slug:        meta.Slug,
		Date:        meta.Date,
		Description: meta.Description,
		Tags:        meta.Tags,
		Draft:       meta.Draft,
		Body:        body,
		HTML:        template.HTML(html.String()),
		SearchText:  searchText,
	}, nil
}

func splitFrontMatter(input string) (string, string, error) {
	trimmed := strings.TrimLeft(input, "\ufeff\n\r\t ")
	if !strings.HasPrefix(trimmed, "---\n") && !strings.HasPrefix(trimmed, "---\r\n") {
		return "", "", fmt.Errorf("missing front matter")
	}

	lines := strings.Split(trimmed, "\n")
	if len(lines) < 3 {
		return "", "", fmt.Errorf("invalid front matter")
	}

	if strings.TrimSpace(lines[0]) != "---" {
		return "", "", fmt.Errorf("missing front matter delimiter")
	}

	end := -1
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "---" {
			end = i
			break
		}
	}
	if end == -1 {
		return "", "", fmt.Errorf("unterminated front matter")
	}

	frontMatter := strings.Join(lines[1:end], "\n")
	body := strings.Join(lines[end+1:], "\n")
	body = strings.TrimLeft(body, "\r\n")
	return frontMatter, body, nil
}

func parseFrontMatter(input string) (frontMatter, error) {
	meta := frontMatter{}
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, ":")
		if !ok {
			return meta, fmt.Errorf("invalid front matter line %q", line)
		}
		key = strings.TrimSpace(strings.ToLower(key))
		value = strings.TrimSpace(value)
		value = strings.Trim(value, "\"")

		switch key {
		case "title":
			meta.Title = value
		case "slug":
			meta.Slug = value
		case "date":
			parsed, err := parseDate(value)
			if err != nil {
				return meta, err
			}
			meta.Date = parsed
		case "description":
			meta.Description = value
		case "tags":
			meta.Tags = parseTags(value)
		case "draft":
			meta.Draft = strings.EqualFold(value, "true")
		}
	}

	if meta.Title == "" {
		return meta, fmt.Errorf("missing title")
	}
	if meta.Date.IsZero() {
		return meta, fmt.Errorf("missing date")
	}
	return meta, nil
}

func parseTags(input string) []string {
	parts := strings.Split(input, ",")
	tags := make([]string, 0, len(parts))
	for _, part := range parts {
		tag := strings.ToLower(strings.TrimSpace(part))
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	return tags
}

func parseDate(value string) (time.Time, error) {
	layouts := []string{
		time.RFC3339,
		"2006-01-02",
		"2006-01-02 15:04",
	}
	for _, layout := range layouts {
		if parsed, err := time.Parse(layout, value); err == nil {
			return parsed, nil
		}
	}
	return time.Time{}, fmt.Errorf("invalid date %q", value)
}
