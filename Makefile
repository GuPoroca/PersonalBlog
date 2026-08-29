.PHONY: all build css templ watch dev site

# Run templ generate once
templ:
	go run github.com/a-h/templ/cmd/templ@v0.3.943 generate

# Watch mode: Tailwind + Templ side by side
watch:
	templ generate --watch

# For development (alias to watch)
dev: watch

# Generate a static site suitable for Cloudflare Pages.
site: templ
	go run ./cmd/sitegen
