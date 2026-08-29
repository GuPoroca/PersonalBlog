# GuPoroca blog

This repository contains both a local Go server and a static-site build for
Cloudflare Pages. Posts live in `content/en` and `content/br` as Markdown files
with TOML front matter.

## Publish through Cloudflare Pages

1. Commit and push this repository to GitHub.
2. In Cloudflare, open **Workers & Pages** → **Create application** → **Pages**
   → **Import an existing Git repository** and select the repository.
3. Use these build settings:

   | Setting | Value |
   | --- | --- |
   | Production branch | `main` |
   | Build command | `make site` |
   | Build output directory | `dist` |
   | Root directory | leave empty |

4. Set `GO_VERSION` to `1.24.7` for both the Production and Preview
   environments under **Settings** → **Environment variables**.
5. Deploy. Cloudflare will give the project a `*.pages.dev` URL.
6. Under **Custom domains**, add your apex domain (for example,
   `example.com`) and optionally `www.example.com`. Since the zone is already
   managed by Cloudflare, Pages creates the required DNS records and TLS
   certificate for you.

Every push to `main` now rebuilds and publishes production automatically. A
pull request receives an isolated preview deployment, so you can review a post
before merging it. To rebuild the exact current `main` commit without changing
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
