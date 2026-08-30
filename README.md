# GuPoroca blog

This repository contains both a local Go server and a static-site build for
Cloudflare Pages. Posts live in `content/en` and `content/br` as Markdown files
with TOML front matter.

Every push to `master` now rebuilds and publishes production automatically. A
pull request receives an isolated preview deployment, so you can review a post
before merging it. To rebuild the exact current `master` commit without changing
code, choose **Deployments** → **Retry deployment** in the Pages dashboard.

## Local build

Run:

```sh
make site
```

The generated deployable site is in `dist/` (ignored by Git). It includes the
assets, English and Portuguese indexes, posts, and static tag pages. Run the
same command after editing Markdown, templates, images, or styles.

For the dynamic development server, first regenerate templates with `make
templ`, then run your usual Go entrypoint.
