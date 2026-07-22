# Academic Site Generator

A minimal static site generator for academic personal websites. Built in Go, renders Markdown to clean HTML.

## Structure

```
├── config.yaml           # Site configuration
├── content/
│   ├── articles/         # Markdown articles with YAML frontmatter
│   ├── news.yaml         # News feed entries
│   └── cv.yaml           # CV data
├── templates/            # HTML templates
├── static/               # CSS, images, etc.
└── output/               # Generated site (git-ignored)
```

## Usage

```bash
# Install dependencies
make install

# Build the site
make build

# Build and serve locally at localhost:8080
make serve

# Clean output
make clean
```

## Writing Articles

Create a markdown file in `content/articles/` with YAML frontmatter:

```markdown
---
title: "Your Article Title"
date: "2024-01-15"
description: "A brief description for listings and RSS."
tags: ["research", "ml"]
draft: false
---

Your content here. Supports **Markdown** formatting.
```

Set `draft: true` to exclude from build.

## News Feed

Edit `content/news.yaml`:

```yaml
- date: "2024-01-15"
  content: "Paper accepted to **ICML 2024**!"

- date: "2024-01-01"
  content: "Happy new year."
```

Supports inline Markdown.

## CV

Edit `content/cv.yaml` with your information:

```yaml
name: "Your Name"
title: "Your Title"
email: "you@example.com"
github: "yourusername"
linkedin: "yourusername"
headshot: "/headshot.jpg"  # Place image in static/headshot.jpg
# ... see cv.yaml for full structure
```

**Headshot:** Place your photo at `static/headshot.jpg` (or any name, just update the path in cv.yaml). Recommended: square image, at least 200x200px.

## Customization

- Edit `static/style.css` for styling
- Modify templates in `templates/`
- Add new template types by extending `cmd/build/main.go`

## Deployment

### GitHub Pages (CI/CD)

This repo includes a GitHub Actions workflow that automatically builds and deploys on push to `main`.

**Setup:**

1. Push this repo to `jonathjd.github.io` (or your `username.github.io`)
2. Go to repo Settings → Pages
3. Under "Build and deployment", select **GitHub Actions** as the source
4. Push to `main` - the site will deploy automatically

**Manual deployment:**

```bash
make build
# Copy output/ contents to your gh-pages branch or hosting
```

### Netlify/Vercel

Set build command to `make build` and publish directory to `output`.

### Manual

Copy `output/` contents to your web server.

## Features

- Markdown with GitHub Flavored Markdown extensions
- Automatic RSS feed generation (`/feed.xml`)
- Dark mode support (follows system preference)
- Responsive design
- Zero JavaScript required
- ~400 lines of Go, ~300 lines of CSS

## Dependencies

- Go 1.21+
- [goldmark](https://github.com/yuin/goldmark) - Markdown parser
- [yaml.v3](https://gopkg.in/yaml.v3) - YAML parser
