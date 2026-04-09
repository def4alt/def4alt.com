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

func TestLoadPostsParsesOptionalImageField(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	postsDir := filepath.Join(root, "posts")
	if err := os.MkdirAll(postsDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	mustWrite(t, filepath.Join(postsDir, "image.md"), `---
title: Image Post
slug: image-post
date: 2024-04-01
description: post with images
tags: visual
image: /content/images/gallery-1.webp
image_alt: A colorful browser gallery scene
---
Body.
`)

	blog, err := content.Load(root)
	if err != nil {
		t.Fatalf("load: %v", err)
	}

	post, ok := blog.BySlug("image-post")
	if !ok {
		t.Fatalf("expected post by slug")
	}
	if post.Image != "/content/images/gallery-1.webp" {
		t.Fatalf("expected image field, got %q", post.Image)
	}
	if post.ImageAlt != "A colorful browser gallery scene" {
		t.Fatalf("expected image alt, got %q", post.ImageAlt)
	}
}

func TestSearchFiltersByQueryAndTag(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	postsDir := filepath.Join(root, "posts")
	if err := os.MkdirAll(postsDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	mustWrite(t, filepath.Join(postsDir, "alpha.md"), `---
title: Alpha Post
slug: alpha-post
date: 2024-02-01
description: Practical Alpha
tags: alpha, backend
---
This post covers server-side Alpha.
`)
	mustWrite(t, filepath.Join(postsDir, "beta.md"), `---
title: Beta Guide
slug: beta-guide
date: 2024-01-01
description: Beta patterns
tags: beta, frontend
---
This post covers live search.
`)

	blog, err := content.Load(root)
	if err != nil {
		t.Fatalf("load: %v", err)
	}

	filtered := blog.Search("server", "alpha")
	if got := len(filtered); got != 1 {
		t.Fatalf("expected 1 matching post, got %d", got)
	}
	if filtered[0].Slug != "alpha-post" {
		t.Fatalf("expected alpha-post, got %q", filtered[0].Slug)
	}

	filtered = blog.Search("", "beta")
	if got := len(filtered); got != 1 {
		t.Fatalf("expected 1 tag match, got %d", got)
	}
	if filtered[0].Slug != "beta-guide" {
		t.Fatalf("expected beta-guide, got %q", filtered[0].Slug)
	}
}

func mustWrite(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
