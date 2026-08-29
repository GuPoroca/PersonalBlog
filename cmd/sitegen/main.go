// sitegen renders the blog into plain HTML files for static hosting.
package main

import (
	"blog/components"
	"blog/data"
	"blog/front"
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"

	"github.com/a-h/templ"
)

func main() {
	out := flag.String("out", "dist", "directory for generated site")
	flag.Parse()
	if err := build(*out); err != nil {
		fmt.Fprintln(os.Stderr, "site build failed:", err)
		os.Exit(1)
	}
}

func build(out string) error {
	if err := os.RemoveAll(out); err != nil {
		return err
	}
	if err := os.MkdirAll(out, 0o755); err != nil {
		return err
	}
	for _, dir := range []string{"static", "images"} {
		if err := copyDir(dir, filepath.Join(out, dir)); err != nil {
			return err
		}
	}
	for _, lang := range []string{"en", "br"} {
		posts, err := data.LoadPosts(lang)
		if err != nil {
			return fmt.Errorf("load %s posts: %w", lang, err)
		}
		if err := render(filepath.Join(out, lang, "blog", "index.html"), components.Layout(components.PostList(posts, "", lang), lang)); err != nil {
			return err
		}
		for _, post := range posts {
			if err := render(filepath.Join(out, lang, "blog", post.Slug, "index.html"), components.Layout(components.Post(post), lang)); err != nil {
				return err
			}
		}
		for _, tag := range tags(posts) {
			content := components.PostList(posts, tag, lang)
			if len(data.FilterAndSortPosts(posts, tag)) == 0 {
				if lang == "br" {
					content = front.EmptyFallback_br
				} else {
					content = front.EmptyFallback_en
				}
			}
			if err := render(filepath.Join(out, lang, "blog", "tags", data.Slugify(tag), "index.html"), components.Layout(content, lang)); err != nil {
				return err
			}
		}
	}
	enPosts, err := data.LoadPosts("en")
	if err != nil {
		return err
	}
	if err := render(filepath.Join(out, "blog", "index.html"), components.Layout(components.PostList(enPosts, "", "en"), "en")); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(out, "_redirects"), []byte("/ /blog/ 302\n/blog /blog/ 308\n/en/blog /en/blog/ 308\n/br/blog /br/blog/ 308\n"), 0o644)
}

func render(path string, component templ.Component) error {
	var body bytes.Buffer
	if err := component.Render(context.Background(), &body); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, body.Bytes(), 0o644)
}

func tags(posts []*data.Post) []string {
	// These are the tag links in the navigation. Keep empty tag pages available
	// as well, so publishing a post under a new tag is never required first.
	set := map[string]struct{}{
		"apollo":     {},
		"athena":     {},
		"daedalus":   {},
		"hephaestus": {},
	}
	for _, post := range posts {
		for _, tag := range post.Meta.Tags {
			set[tag] = struct{}{}
		}
	}
	result := make([]string, 0, len(set))
	for tag := range set {
		result = append(result, tag)
	}
	sort.Strings(result)
	return result
}

func copyDir(src, dst string) error {
	return filepath.WalkDir(src, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if entry.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, contents, 0o644)
	})
}
