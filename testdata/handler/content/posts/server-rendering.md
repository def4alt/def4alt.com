---
title: Server Rendering Matters
slug: server-rendering
date: 2026-04-08
description: Why server-rendered HTML still feels great.
tags: go, web
image: /content/images/gallery-1.webp
image_alt: A colorful browser gallery scene
---
# Server Rendering Matters

Fast pages, simple data flow, and fewer moving parts.

![A colorful browser gallery scene](/content/images/gallery-1.webp)

## Why it matters

- predictable data flow
- fewer client states
- easier caching

> Server rendering keeps the browser honest.

It also works well with **bold text**, *emphasis*, ~~strikethrough~~, `inline code`, and [links](https://developer.mozilla.org/en-US/docs/Web/Markdown).

```go
fmt.Println("hello from Go")
```

| Layer | Role |
| --- | --- |
| Content | Markdown source |
| Server | HTML rendering |

- [x] rendered
- [ ] deferred

