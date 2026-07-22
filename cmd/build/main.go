package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	stdhtml "html"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v3"
)

// Content types
type Article struct {
	Title       string        `yaml:"title"`
	Date        string        `yaml:"date"`
	Description string        `yaml:"description"`
	Tags        []string      `yaml:"tags"`
	Draft       bool          `yaml:"draft"`
	Slug        string        `yaml:"-"`
	Content     template.HTML `yaml:"-"`
	ParsedDate  time.Time     `yaml:"-"`
	URL         string        `yaml:"-"`
	TagLinks    []Tag         `yaml:"-"`
	Headings    []Heading     `yaml:"-"`
	ReadMinutes int           `yaml:"-"`
}

// Heading is one entry in an article's table of contents.
type Heading struct {
	ID    string
	Title string
	Depth int // relative to the shallowest heading in the article, 0 or 1
}

type Tag struct {
	Name  string
	Slug  string
	URL   string
	Count int
}

type NewsItem struct {
	Date       string        `yaml:"date"`
	Content    template.HTML `yaml:"-"`
	ParsedDate time.Time     `yaml:"-"`
}

type CVSection struct {
	Title   string    `yaml:"title"`
	Entries []CVEntry `yaml:"entries"`
}

type CVEntry struct {
	Title       string `yaml:"title"`
	Institution string `yaml:"institution"`
	Location    string `yaml:"location"`
	Date        string `yaml:"date"`
	Description string `yaml:"description"`
}

type CV struct {
	Name     string      `yaml:"name"`
	Title    string      `yaml:"title"`
	Email    string      `yaml:"email"`
	Website  string      `yaml:"website"`
	Github   string      `yaml:"github"`
	LinkedIn string      `yaml:"linkedin"`
	Headshot string      `yaml:"headshot"`
	Bio      string      `yaml:"bio"`
	Sections []CVSection `yaml:"sections"`
}

type SiteConfig struct {
	Title       string `yaml:"title"`
	Author      string `yaml:"author"`
	Description string `yaml:"description"`
	BaseURL     string `yaml:"base_url"`
}

type Site struct {
	Config   SiteConfig
	Articles []Article
	News     []NewsItem
	CV       CV
	Tags     []Tag
}

// RSS types
type RSS struct {
	XMLName xml.Name   `xml:"rss"`
	Version string     `xml:"version,attr"`
	Channel RSSChannel `xml:"channel"`
}

type RSSChannel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Items       []RSSItem `xml:"item"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

var frontmatterRegex = regexp.MustCompile(`(?s)^---\n(.+?)\n---\n(.*)$`)

func main() {
	serve := flag.Bool("serve", false, "Start a local development server")
	port := flag.String("port", "8080", "Port for development server")
	flag.Parse()

	if err := build(); err != nil {
		fmt.Fprintf(os.Stderr, "Build error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✓ Site built successfully → output/")

	if *serve {
		fmt.Printf("Starting server at http://localhost:%s\n", *port)
		startServer(*port)
	}
}

func build() error {
	site := &Site{}

	// Load config
	if err := loadConfig(site); err != nil {
		return fmt.Errorf("loading config: %w", err)
	}

	// Load content
	if err := loadArticles(site); err != nil {
		return fmt.Errorf("loading articles: %w", err)
	}
	if err := loadNews(site); err != nil {
		return fmt.Errorf("loading news: %w", err)
	}
	if err := loadCV(site); err != nil {
		return fmt.Errorf("loading CV: %w", err)
	}

	// Must run before rendering: templates receive copies of the articles
	buildTags(site)

	// Clean and create output directory
	os.RemoveAll("output")
	os.MkdirAll("output/articles", 0755)
	os.MkdirAll("output/tags", 0755)

	// Load templates
	tmpl, err := loadTemplates()
	if err != nil {
		return fmt.Errorf("loading templates: %w", err)
	}

	// Render pages
	if err := renderIndex(site, tmpl); err != nil {
		return fmt.Errorf("rendering index: %w", err)
	}
	if err := renderArticles(site, tmpl); err != nil {
		return fmt.Errorf("rendering articles: %w", err)
	}
	if err := renderArticleList(site, tmpl); err != nil {
		return fmt.Errorf("rendering article list: %w", err)
	}
	if err := renderTagPages(site, tmpl); err != nil {
		return fmt.Errorf("rendering tag pages: %w", err)
	}
	if err := renderTagIndex(site, tmpl); err != nil {
		return fmt.Errorf("rendering tag index: %w", err)
	}
	if err := renderCV(site, tmpl); err != nil {
		return fmt.Errorf("rendering CV: %w", err)
	}
	if err := renderRSS(site); err != nil {
		return fmt.Errorf("rendering RSS: %w", err)
	}

	// Copy static files
	if err := copyStatic(); err != nil {
		return fmt.Errorf("copying static files: %w", err)
	}

	return nil
}

func loadConfig(site *Site) error {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		// Default config
		site.Config = SiteConfig{
			Title:       "Jons Website",
			Author:      "Jonathan Dickinson",
			Description: "Personal website",
			BaseURL:     "jonathjd.github.io",
		}
		return nil
	}
	return yaml.Unmarshal(data, &site.Config)
}

func loadArticles(site *Site) error {
	entries, err := os.ReadDir("content/articles")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.Typographer),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		goldmark.WithRendererOptions(html.WithUnsafe()),
	)

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		data, err := os.ReadFile(filepath.Join("content/articles", entry.Name()))
		if err != nil {
			return err
		}

		article, err := parseArticle(data, md)
		if err != nil {
			return fmt.Errorf("parsing %s: %w", entry.Name(), err)
		}

		if article.Draft {
			continue
		}

		article.Slug = strings.TrimSuffix(entry.Name(), ".md")
		article.URL = "/articles/" + article.Slug + ".html"
		site.Articles = append(site.Articles, article)
	}

	// Sort by date, newest first
	sort.Slice(site.Articles, func(i, j int) bool {
		return site.Articles[i].ParsedDate.After(site.Articles[j].ParsedDate)
	})

	return nil
}

func parseArticle(data []byte, md goldmark.Markdown) (Article, error) {
	var article Article

	matches := frontmatterRegex.FindSubmatch(data)
	if matches == nil {
		return article, fmt.Errorf("invalid frontmatter")
	}

	if err := yaml.Unmarshal(matches[1], &article); err != nil {
		return article, err
	}

	var buf bytes.Buffer
	if err := md.Convert(matches[2], &buf); err != nil {
		return article, err
	}
	article.Content = template.HTML(buf.String())
	article.Headings = extractHeadings(buf.String())
	article.ReadMinutes = readingTime(buf.String())

	article.ParsedDate, _ = time.Parse("2006-01-02", article.Date)

	return article, nil
}

var (
	headingRegex = regexp.MustCompile(`(?is)<h([1-3]) id="([^"]*)"[^>]*>(.*?)</h[1-3]>`)
	tagRegex     = regexp.MustCompile(`<[^>]+>`)
	nonAlnum     = regexp.MustCompile(`[^a-z0-9]+`)
)

// extractHeadings pulls the table of contents out of rendered article HTML.
// Goldmark's WithAutoHeadingID gives every heading an id to link to. Depth is
// normalized against the shallowest heading present, since some articles start
// at h1 and others at h2.
func extractHeadings(content string) []Heading {
	matches := headingRegex.FindAllStringSubmatch(content, -1)
	if len(matches) == 0 {
		return nil
	}

	type raw struct {
		id, title string
		level     int
	}

	var found []raw
	minLevel := 3

	for _, m := range matches {
		level, _ := strconv.Atoi(m[1])
		title := strings.TrimSpace(stdhtml.UnescapeString(tagRegex.ReplaceAllString(m[3], "")))
		if title == "" || m[2] == "" {
			continue
		}
		if level < minLevel {
			minLevel = level
		}
		found = append(found, raw{id: m[2], title: title, level: level})
	}

	headings := make([]Heading, 0, len(found))
	for _, f := range found {
		depth := f.level - minLevel
		if depth > 1 {
			depth = 1
		}
		headings = append(headings, Heading{ID: f.id, Title: f.title, Depth: depth})
	}

	return headings
}

// readingTime estimates minutes to read at 200 words per minute, rounding up.
func readingTime(content string) int {
	words := len(strings.Fields(tagRegex.ReplaceAllString(content, " ")))
	minutes := (words + 199) / 200
	if minutes < 1 {
		return 1
	}
	return minutes
}

func slugify(s string) string {
	return strings.Trim(nonAlnum.ReplaceAllString(strings.ToLower(s), "-"), "-")
}

// buildTags collects every tag in use into site.Tags and gives each article the
// linkable version of its own tags. Articles are already sorted newest first.
func buildTags(site *Site) {
	bySlug := map[string]*Tag{}

	for _, article := range site.Articles {
		for _, name := range article.Tags {
			slug := slugify(name)
			if slug == "" {
				continue
			}

			tag, ok := bySlug[slug]
			if !ok {
				tag = &Tag{Name: name, Slug: slug, URL: "/tags/" + slug + ".html"}
				bySlug[slug] = tag
			}
			tag.Count++
		}
	}

	// Second pass, so each article's tags carry the final counts
	for i := range site.Articles {
		for _, name := range site.Articles[i].Tags {
			if tag, ok := bySlug[slugify(name)]; ok {
				site.Articles[i].TagLinks = append(site.Articles[i].TagLinks, *tag)
			}
		}
	}

	for _, tag := range bySlug {
		site.Tags = append(site.Tags, *tag)
	}
	sort.Slice(site.Tags, func(i, j int) bool {
		return site.Tags[i].Name < site.Tags[j].Name
	})
}

func articlesWithTag(site *Site, slug string) []Article {
	var out []Article
	for _, article := range site.Articles {
		for _, tag := range article.TagLinks {
			if tag.Slug == slug {
				out = append(out, article)
				break
			}
		}
	}
	return out
}

func loadNews(site *Site) error {
	data, err := os.ReadFile("content/news.yaml")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	var items []struct {
		Date    string `yaml:"date"`
		Content string `yaml:"content"`
	}

	if err := yaml.Unmarshal(data, &items); err != nil {
		return err
	}

	md := goldmark.New(goldmark.WithExtensions(extension.GFM))

	for _, item := range items {
		var buf bytes.Buffer
		md.Convert([]byte(item.Content), &buf)

		parsed, _ := time.Parse("2006-01-02", item.Date)
		site.News = append(site.News, NewsItem{
			Date:       item.Date,
			Content:    template.HTML(buf.String()),
			ParsedDate: parsed,
		})
	}

	// Sort by date, newest first
	sort.Slice(site.News, func(i, j int) bool {
		return site.News[i].ParsedDate.After(site.News[j].ParsedDate)
	})

	return nil
}

func loadCV(site *Site) error {
	data, err := os.ReadFile("content/cv.yaml")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return yaml.Unmarshal(data, &site.CV)
}

type Templates struct {
	Index    *template.Template
	Article  *template.Template
	Articles *template.Template
	CV       *template.Template
	Tag      *template.Template
	Tags     *template.Template
}

func loadTemplates() (*Templates, error) {
	funcMap := template.FuncMap{
		"formatDate": func(t time.Time) string {
			return t.Format("January 2, 2006")
		},
		"shortDate": func(t time.Time) string {
			return t.Format("Jan 2006")
		},
		"year": func() int {
			return time.Now().Year()
		},
	}

	parse := func(files ...string) (*template.Template, error) {
		return template.New("").Funcs(funcMap).ParseFiles(files...)
	}

	index, err := parse("templates/base.html", "templates/index.html")
	if err != nil {
		return nil, fmt.Errorf("index: %w", err)
	}

	article, err := parse("templates/base.html", "templates/article.html")
	if err != nil {
		return nil, fmt.Errorf("article: %w", err)
	}

	articles, err := parse("templates/base.html", "templates/articles.html")
	if err != nil {
		return nil, fmt.Errorf("articles: %w", err)
	}

	cv, err := parse("templates/base.html", "templates/cv.html")
	if err != nil {
		return nil, fmt.Errorf("cv: %w", err)
	}

	tag, err := parse("templates/base.html", "templates/tag.html")
	if err != nil {
		return nil, fmt.Errorf("tag: %w", err)
	}

	tags, err := parse("templates/base.html", "templates/tags.html")
	if err != nil {
		return nil, fmt.Errorf("tags: %w", err)
	}

	return &Templates{
		Index:    index,
		Article:  article,
		Articles: articles,
		CV:       cv,
		Tag:      tag,
		Tags:     tags,
	}, nil
}

func renderIndex(site *Site, tmpl *Templates) error {
	f, err := os.Create("output/index.html")
	if err != nil {
		return err
	}
	defer f.Close()

	// Show only recent items on index
	data := struct {
		Site           *Site
		RecentArticles []Article
		RecentNews     []NewsItem
	}{
		Site:           site,
		RecentArticles: take(site.Articles, 5),
		RecentNews:     takeNews(site.News, 5),
	}

	return tmpl.Index.ExecuteTemplate(f, "base", data)
}

func renderArticles(site *Site, tmpl *Templates) error {
	for _, article := range site.Articles {
		f, err := os.Create(filepath.Join("output/articles", article.Slug+".html"))
		if err != nil {
			return err
		}

		data := struct {
			Site    *Site
			Article Article
		}{site, article}

		if err := tmpl.Article.ExecuteTemplate(f, "base", data); err != nil {
			f.Close()
			return err
		}
		f.Close()
	}
	return nil
}

func renderArticleList(site *Site, tmpl *Templates) error {
	f, err := os.Create("output/articles/index.html")
	if err != nil {
		return err
	}
	defer f.Close()

	data := struct {
		Site     *Site
		Articles []Article
	}{site, site.Articles}

	return tmpl.Articles.ExecuteTemplate(f, "base", data)
}

func renderTagPages(site *Site, tmpl *Templates) error {
	for _, tag := range site.Tags {
		f, err := os.Create(filepath.Join("output/tags", tag.Slug+".html"))
		if err != nil {
			return err
		}

		data := struct {
			Site     *Site
			Tag      Tag
			Articles []Article
		}{site, tag, articlesWithTag(site, tag.Slug)}

		if err := tmpl.Tag.ExecuteTemplate(f, "base", data); err != nil {
			f.Close()
			return err
		}
		f.Close()
	}
	return nil
}

func renderTagIndex(site *Site, tmpl *Templates) error {
	f, err := os.Create("output/tags/index.html")
	if err != nil {
		return err
	}
	defer f.Close()

	data := struct {
		Site *Site
		Tags []Tag
	}{site, site.Tags}

	return tmpl.Tags.ExecuteTemplate(f, "base", data)
}

func renderCV(site *Site, tmpl *Templates) error {
	f, err := os.Create("output/cv.html")
	if err != nil {
		return err
	}
	defer f.Close()

	data := struct {
		Site *Site
		CV   CV
	}{site, site.CV}

	return tmpl.CV.ExecuteTemplate(f, "base", data)
}

func renderRSS(site *Site) error {
	var items []RSSItem
	for _, article := range take(site.Articles, 20) {
		items = append(items, RSSItem{
			Title:       article.Title,
			Link:        site.Config.BaseURL + article.URL,
			Description: article.Description,
			PubDate:     article.ParsedDate.Format(time.RFC1123Z),
		})
	}

	rss := RSS{
		Version: "2.0",
		Channel: RSSChannel{
			Title:       site.Config.Title,
			Link:        site.Config.BaseURL,
			Description: site.Config.Description,
			Items:       items,
		},
	}

	f, err := os.Create("output/feed.xml")
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(xml.Header)
	f.WriteString(`<?xml-stylesheet type="text/xsl" href="/feed.xsl"?>` + "\n")
	enc := xml.NewEncoder(f)
	enc.Indent("", "  ")
	return enc.Encode(rss)
}

func copyStatic() error {
	return filepath.Walk("static", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		relPath, _ := filepath.Rel("static", path)
		destPath := filepath.Join("output", relPath)

		os.MkdirAll(filepath.Dir(destPath), 0755)

		src, err := os.Open(path)
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		return err
	})
}

func take(articles []Article, n int) []Article {
	if len(articles) < n {
		return articles
	}
	return articles[:n]
}

func takeNews(news []NewsItem, n int) []NewsItem {
	if len(news) < n {
		return news
	}
	return news[:n]
}

func startServer(port string) {
	http.Handle("/", http.FileServer(http.Dir("output")))
	http.ListenAndServe(":"+port, nil)
}
