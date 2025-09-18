+++
title = "Blog Writing Cheatsheet"
description = "Quick reference: Markdown, code, tables, embeds, and safe raw HTML usage for posts."
created_at = "17 Sep 2025"
last_updated = "17 Sep 2025"
author = "Poroca"
tags = []
cover = "/images/cover-sample.png"
+++

# Blog Writing Cheatsheet 📎

A compact reference of things you can do in posts—Markdown basics, code blocks, tables, images, and safe ways to include raw HTML without the chat frontend eating your tags.


## Headings & Text:

# H1 Title
## H2 Section
### H3 Subsection

**bold**, *italic*, ~~strike~~, `inline code`

Paragraph with a  
line break (two spaces at line end).


## Lists

- Item 1
- Item 2
  - Nested A
  - Nested B
- Item 3

1. First
2. Second
3. Third

- [ ] Task to do
- [x] Done task

## Links & Images

[Inline link](https://example.com "Optional title")

![Alt text](/images/cover-sample.png "Caption-ish tooltip")
I myself prefer the html way.

## Blockquotes & Callouts


> Simple quote or callout.
> Can span multiple lines.

> **Note:** You can style “callouts” by starting with bold labels.

## Code Block

```bash
# Shell snippet
make build && ./server
```

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello From Go CodeBlock")
}
```

## Tables

<table>
  <thead><tr><th>Lang</th><th>Level</th></tr></thead>
  <tbody>
    <tr><td>Go</td><td>Advanced</td></tr>
    <tr><td>Python</td><td>Intermediate</td></tr>
  </tbody>
</table>

## Horizontal Rule

---

## “Details / Summary” Toggles (HTML)

<details>
  <summary>Click to expand</summary>
  <p>Hidden content here.</p>
</details>

To use it in your post (will render as a toggle), place it as raw HTML (not in a code block):

<!-- use raw HTML directly in your .md file -->
<details>
  <summary>Notes on performance</summary>
  <ul>
    <li>Cache rendered pages</li>
    <li>Compress assets</li>
  </ul>
</details>

## Safe Raw HTML Patterns

When you need more control than Markdown:
Wrap HTML in fenced code to demonstrate it (so the chat doesn’t interpret).
Use HTML directly to render it on the blog.
Examples to demonstrate:

<figure>
  <img src="/images/cover-sample.png" alt="Figure image" />
  <figcaption>Caption under the image</figcaption>
</figure>



## Escaping HTML Quickly

Use &lt;div&gt;not rendered as HTML&lt;/div&gt; to keep tags literal.

## Metadata / Front Matter

Keep your front matter consistent across posts (title, dates, tags, cover). Example you can reuse:

+++
title = "Post Title"
description = "One-liner about the post."
created_at = "17 Sep 2025"
last_updated = "17 Sep 2025"
author = "Poroca"
tags = ["tag1", "tag2"]
cover = "/images/cover-sample.png"
+++
