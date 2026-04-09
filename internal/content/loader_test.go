package content_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"def4alt/def4alt.com/internal/content"
)

func TestLoadPostsParsesMarkdownAndSortsNewestFirst(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	postsDir := filepath.Join(root, "posts")
	if err := os.MkdirAll(postsDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	mustWrite(t, filepath.Join(postsDir, "old.md"), `---
title: Old Post
slug: old-post
date: 2024-01-01
description: old description
tags: go, htmx
---
# Old

Old body.
`)
	mustWrite(t, filepath.Join(postsDir, "new.md"), `---
title: New Post
slug: new-post
date: 2024-02-01
description: new description
tags: go
---
# New

New body.
`)
	mustWrite(t, filepath.Join(postsDir, "draft.md"), `---
title: Draft Post
slug: draft-post
date: 2024-03-01
description: draft description
draft: true
tags: go
---
# Draft

Hidden body.
`)

	blog, err := content.Load(root)
	if err != nil {
		t.Fatalf("load: %v", err)
	}

	posts := blog.Posts()
	if got := len(posts); got != 2 {
		t.Fatalf("expected 2 public posts, got %d", got)
	}
	if posts[0].Slug != "new-post" {
		t.Fatalf("expected newest post first, got %q", posts[0].Slug)
	}
	if posts[1].Slug != "old-post" {
		t.Fatalf("expected oldest post second, got %q", posts[1].Slug)
	}

	post, ok := blog.BySlug("new-post")
	if !ok {
		t.Fatalf("expected post by slug")
	}
	if post.Title != "New Post" {
		t.Fatalf("expected title %q, got %q", "New Post", post.Title)
	}
	if !strings.Contains(string(post.HTML), "<h1>New</h1>") {
		t.Fatalf("expected markdown to render to HTML, got %s", string(post.HTML))
	}
}

func TestSearchFiltersByQueryAndTag(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	postsDir := filepath.Join(root, "posts")
	if err := os.MkdirAll(postsDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	mustWrite(t, filepath.Join(postsDir, "go.md"), `---
title: Go Tips
slug: go-tips
date: 2024-02-01
description: Practical Go
tags: go, backend
---
This post covers server-side Go.
`)
	mustWrite(t, filepath.Join(postsDir, "htmx.md"), `---
title: HTMX Guide
slug: htmx-guide
date: 2024-01-01
description: HTMX patterns
tags: htmx, frontend
---
This post covers live search.
`)

	blog, err := content.Load(root)
	if err != nil {
		t.Fatalf("load: %v", err)
	}

	filtered := blog.Search("server", "go")
	if got := len(filtered); got != 1 {
		t.Fatalf("expected 1 matching post, got %d", got)
	}
	if filtered[0].Slug != "go-tips" {
		t.Fatalf("expected go-tips, got %q", filtered[0].Slug)
	}

	filtered = blog.Search("", "htmx")
	if got := len(filtered); got != 1 {
		t.Fatalf("expected 1 tag match, got %d", got)
	}
	if filtered[0].Slug != "htmx-guide" {
		t.Fatalf("expected htmx-guide, got %q", filtered[0].Slug)
	}
}

func mustWrite(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
