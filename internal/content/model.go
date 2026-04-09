package content

import (
	"html/template"
	"sort"
	"strings"
	"time"
)

type Post struct {
	Title       string
	Slug        string
	Date        time.Time
	Description string
	Tags        []string
	Draft       bool
	Body        string
	HTML        template.HTML
	SearchText  string
}

type Blog struct {
	posts  []Post
	bySlug map[string]Post
	tags   []string
}

func (b *Blog) Posts() []Post {
	out := make([]Post, len(b.posts))
	copy(out, b.posts)
	return out
}

func (b *Blog) BySlug(slug string) (Post, bool) {
	post, ok := b.bySlug[strings.ToLower(slug)]
	return post, ok
}

func (b *Blog) Tags() []string {
	out := make([]string, len(b.tags))
	copy(out, b.tags)
	return out
}

func (b *Blog) Search(query, tag string) []Post {
	query = strings.ToLower(strings.TrimSpace(query))
	tag = strings.ToLower(strings.TrimSpace(tag))

	matches := make([]Post, 0, len(b.posts))
	for _, post := range b.posts {
		if query != "" && !strings.Contains(post.SearchText, query) {
			continue
		}
		if tag != "" && !hasTag(post.Tags, tag) {
			continue
		}
		matches = append(matches, post)
	}
	return matches
}

func hasTag(tags []string, want string) bool {
	for _, tag := range tags {
		if strings.EqualFold(tag, want) {
			return true
		}
	}
	return false
}

func uniqueTags(posts []Post) []string {
	seen := make(map[string]struct{})
	for _, post := range posts {
		for _, tag := range post.Tags {
			if tag == "" {
				continue
			}
			seen[strings.ToLower(tag)] = struct{}{}
		}
	}

	tags := make([]string, 0, len(seen))
	for tag := range seen {
		tags = append(tags, tag)
	}
	sort.Strings(tags)
	return tags
}
