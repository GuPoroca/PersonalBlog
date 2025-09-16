.PHONY: all build css templ watch dev

# Run templ generate once
templ:
	templ generate

# Watch mode: Tailwind + Templ side by side
watch:
	templ generate --watch

# For development (alias to watch)
dev: watch
