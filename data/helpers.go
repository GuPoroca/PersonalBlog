package data

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"sort"
	"time"

	"github.com/pelletier/go-toml/v2"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type Meta struct {
	Title        string   `toml:"title"`
	Description  string   `toml:"description"`
	Created_at   string   `toml:"created_at"`
	Last_updated string   `toml:"last_updated"`
	Author       string   `toml:"author"`
	Tags         []string `toml:"tags"`
	Cover        string   `toml:"cover"`
}
type Post struct {
	Meta              Meta
	HTML              template.HTML // safe HTML to inject into templ
	parsedCreatedDate time.Time
	Slug              string
}

var md = goldmark.New(
	goldmark.WithParserOptions(parser.WithAutoHeadingID()),
	goldmark.WithRendererOptions(html.WithUnsafe()),
)

func ParseMarkdownFile(path string) (*Post, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	content := string(data)

	parts := strings.SplitN(content, "+++", 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid markdown file, missing front-matter: %s", path)
	}

	header := strings.TrimSpace(parts[1])
	body := strings.TrimSpace(parts[2])

	var meta Meta
	if err := toml.Unmarshal([]byte(header), &meta); err != nil {
		return nil, fmt.Errorf("failed to parse TOML in %s: %w", path, err)
	}

	var buf bytes.Buffer
	if err := md.Convert([]byte(body), &buf); err != nil {
		return nil, fmt.Errorf("failed to convert markdown: %w", err)
	}

	// Parse Created_at
	layout := "02 Jan 2006"
	var parsed time.Time
	if meta.Created_at != "" {
		if t, err := time.Parse(layout, meta.Created_at); err == nil {
			parsed = t
		}
	}

	// Generate slug from title if not set
	slug := Slugify(meta.Title)

	return &Post{
		Meta:              meta,
		HTML:              template.HTML(buf.String()),
		parsedCreatedDate: parsed,
		Slug:              slug,
	}, nil
}
func FilterAndSortPosts(posts []*Post, tag string) []*Post {
	layout := "02 Jan 2006"

	// Parse Created_at into time.Time
	for _, p := range posts {
		if p.Meta.Created_at != "" {
			t, err := time.Parse(layout, p.Meta.Created_at)
			if err == nil {
				p.parsedCreatedDate = t
			}
		}
	}

	// Filter by tag if provided
	var filtered []*Post
	for _, p := range posts {
		if tag == "" {
			filtered = append(filtered, p)
			continue
		}
		if slices.Contains(p.Meta.Tags, tag) {
			filtered = append(filtered, p)
		}
	}

	// Sort descending by date
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].parsedCreatedDate.After(filtered[j].parsedCreatedDate)
	})

	// Limit to 5
	if len(filtered) > 5 {
		filtered = filtered[:5]
	}
	return filtered
}

func Slugify(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "'", "")
	slug = strings.ReplaceAll(slug, "\"", "")
	slug = slug + ".md"
	return slug
}

func LoadPosts(lang string) ([]*Post, error) {
	dir := filepath.Join("content", lang)

	var posts []*Post
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".md") {
			return nil
		}
		post, err := ParseMarkdownFile(path)
		if err != nil {
			return err
		}
		posts = append(posts, post)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func LoadPost(lang, slug string) (*Post, error) {
	dir := filepath.Join("content", lang)
	files, err := os.ReadDir(dir)
	log.Print(dir)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".md") {
			post, err := ParseMarkdownFile(filepath.Join(dir, f.Name()))
			if err != nil {
				return nil, err
			}
			if post.Slug == slug {
				return post, nil
			}
		}
	}
	return nil, fmt.Errorf("post not found: %s/%s", lang, slug)
}
