package front

import (
	"blog/components"
	"blog/data"
	"context"
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
)

// GET /blog?tag=golang&lang=en

// in front/handlers.go (or in components if you prefer)
var ErrorFallback = templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
	_, err := io.WriteString(w, `<div class="text-center text-red-400">something went wrong :(</div>`)
	return err
})

var EmptyFallback_en = templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
	_, err := io.WriteString(w, `<div class="text-center text-gray-400">no posts found</div>`)
	return err
})

var EmptyFallback_br = templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
	_, err := io.WriteString(w, `<div class="text-center text-gray-400">nenhum post encontrado</div>`)
	return err
})

func BlogListHandler(w http.ResponseWriter, r *http.Request) {
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		lang = chi.URLParam(r, "lang")
	}
	if lang == "" {
		lang = "en"
	}

	tag := r.URL.Query().Get("tag")

	posts, err := data.LoadPosts(lang)

	var component templ.Component
	if len(posts) == 0 {
		switch lang {
		case "en":
			component = EmptyFallback_en
		case "br":
			component = EmptyFallback_br
		}
	} else if err != nil {
		log.Printf("error loading posts: %v", err)
		component = ErrorFallback
	} else {
		component = components.PostList(posts, tag, lang)
	}

	if r.Header.Get("HX-Request") == "true" {
		templ.Handler(component).ServeHTTP(w, r)
		return
	}
	templ.Handler(components.Layout(component)).ServeHTTP(w, r)
}

// GET /blog/{lang}/{slug}

func BlogPostHandler(w http.ResponseWriter, r *http.Request) {
	lang := chi.URLParam(r, "lang")
	if lang == "" {
		lang = "en"
	}
	slug := chi.URLParam(r, "slug")
	log.Print(slug)

	post, err := data.LoadPost(lang, slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	component := components.Post(post)
	templ.Handler(components.Layout(component)).ServeHTTP(w, r)
}
