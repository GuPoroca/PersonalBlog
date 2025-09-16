package front

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Front() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Serve static files from ./static at /static/*
	FileServer(r, "/static", http.Dir("static"))
	FileServer(r, "/images", http.Dir("images"))

	// List posts (default = en, or ?lang=br, or /{lang}/blog)
	r.Get("/blog", BlogListHandler)
	r.Get("/{lang}/blog", BlogListHandler)

	// Single post
	r.Get("/{lang}/blog/{slug}", BlogPostHandler)

	log.Println("Server running on http://localhost:3000")
	http.ListenAndServe(":3000", r)
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	r.Get(path+"/*", func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}
