# Aitch ("H") - for HTML (or SVG)
[![GoDoc](https://godoc.org/github.com/go-andiamo/aitch?status.svg)](https://pkg.go.dev/github.com/go-andiamo/aitch)
[![Latest Version](https://img.shields.io/github/v/tag/go-andiamo/aitch.svg?sort=semver&style=flat&label=version&color=blue)](https://github.com/go-andiamo/aitch/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-andiamo/aitch)](https://goreportcard.com/report/github.com/go-andiamo/aitch)

Aitch is a fluent, low/zero-GC pressure HTML and SVG rendering library for Go.  
Built for performance, correctness, and developers who care.

---

## Features

- Fluent, composable, declarative API for HTML and SVG
- Mixed content and attributes handled cleanly — order doesn’t matter
    - Repeated attributes are protected: either the last one wins (e.g., `id`), or values are merged (e.g., `class`, `style`)
- Attribute merging (e.g., multiple `Class()` or `Style()` calls)
- Virtually zero allocations during rendering<sup>†</sup>
- Conditional rendering (elements, attributes, values)
- Boolean attribute and void element aware
- Dynamic value support using runtime data context
- Imperative rendering for absolute dynamically built HTML/SVG
- Seamless integration with `html/template`
- 100% tested and production-ready

†: With totally static templates (no dynamic attributes, no conditional attributes, no imperative rendering) zero allocs (zero GC-pressure) is almost guaranteed!
Dynamic values or conditional logic may introduce some allocations — this is unavoidable in Go, but Aitch minimizes them wherever possible.

---

## Installation

```bash
go get github.com/go-andiamo/aitch
```

---

## Declarative Templates, Runtime Flexibility

Aitch separates **template declaration** from **template rendering**.

Templates are built once — as plain Go `var`s — and contain no runtime logic.
Rendering is done later, with context-aware data and **zero allocations**.

This means you can:

- Declare templates at package level
- Inject dynamic values or render conditionally *at runtime*
- Reuse templates safely and efficiently across renders
- Stream directly to any `io.Writer`

---

## Examples

<details>
    <summary><strong>1. Mixed content and attributes handled cleanly</strong></summary>

```go
package main

import (
    "github.com/go-andiamo/aitch/context"
    "github.com/go-andiamo/aitch/html"
    "os"
)

var template = html.P(
    "This is a para",
    html.Id(1), html.Class("foo"),
    html.Br(), " >>> and this is more",
    html.Id(2), html.Class("bar"))

func main() {
    ctx := context.New(os.Stdout, nil, nil)
    _ = template.Render(ctx)
}
```
correctly produces...
```html
<p id="2" class="foo bar">This is a para<br> &gt;&gt;&gt; and this is more</p>
```

</details><br>

<details>
    <summary><strong>2. Dynamic values</strong></summary>

Attribute and text can be injected dynamically at rendering...

```go
package main

import (
    "github.com/go-andiamo/aitch"
    "github.com/go-andiamo/aitch/context"
    "github.com/go-andiamo/aitch/html"
    "os"
)

var template = html.P(
    "Hello, ", aitch.DynamicValueKey("name"),  // resolves ctx.Data["name"] at render time
    html.Class("greeting"), html.Class(aitch.DynamicValueKey("greeting-class")))

func main() {
    data := map[string]any{
        "name":           "Aitch!",
        "greeting-class": "admin",
    }
    ctx := context.New(os.Stdout, data, nil)
    _ = template.Render(ctx)
}
```
produces...
```html
<p class="greeting admin">Hello, Aitch!</p>
```
</details><br>

<details>
    <summary><strong>3. Dynamic values - as functions</strong></summary>

```go
package main

import (
    "github.com/go-andiamo/aitch"
    "github.com/go-andiamo/aitch/context"
    "github.com/go-andiamo/aitch/html"
    "os"
    "strconv"
)

func idIncrement(ctx *context.Context) []byte {
    id := ctx.Data["id"].(int) + 1
    ctx.Data["id"] = id
    return []byte(strconv.Itoa(id))
}

var template = aitch.Collection(
    html.P("First", html.Id(idIncrement)),
    html.P("Second", html.Id(idIncrement)),
    html.P("Third", html.Id(idIncrement)))

func main() {
    data := map[string]any{
        "id": 0,
    }
    ctx := context.New(os.Stdout, data, nil)
    _ = template.Render(ctx)
}
```
produces...
```html
<p id="1">First</p><p id="2">Second</p><p id="3">Third</p>
```

</details><br>

<details>
    <summary><strong>4. Conditionals</strong></summary>

```go
package main

import (
    "github.com/go-andiamo/aitch"
    "github.com/go-andiamo/aitch/context"
    "github.com/go-andiamo/aitch/html"
    "os"
)

var template = html.P(
    "Hello, ", aitch.DynamicValueKey("username"),
    html.Class("base"),
    aitch.Conditional(func(ctx *context.Context) bool {
        return ctx.Data["isAdmin"].(bool)
    }, html.Class("admin")))

func main() {
    data := map[string]any{
        "username": "Aitch",
        "isAdmin":  true,
    }
    ctx := context.New(os.Stdout, data, nil)
    _ = template.Render(ctx)
}
```
produces...
```html
<p class="base admin">Hello, Aitch</p>
```
</details><br>

<details>
    <summary><strong>5. Imperative generation</strong></summary>

```go
package main

import (
    "github.com/go-andiamo/aitch"
    "github.com/go-andiamo/aitch/context"
    "github.com/go-andiamo/aitch/html"
    "os"
)

var template = html.Div(html.Select(
    aitch.Imperative(func(ctx aitch.ImperativeContext) error {
        items, _ := context.Get[[]string](ctx.Context(), "items")
        for _, item := range items {
            ctx.Start([]byte("option"), false)
            ctx.WriteRaw([]byte(item))
            ctx.End()
        }
        return nil
    })))

func main() {
    data := map[string]any{
        "items": []string{"First", "Second", "Third"},
    }
    ctx := context.New(os.Stdout, data, nil)
    _ = template.Render(ctx)
}
```
produces...
```html
<div><select><option>First</option><option>Second</option><option>Third</option></select></div>
```

</details><br>

---

## Not a replacement for `html/template`

Aitch complements `html/template` — it doesn’t try to replace it.

Use `html/template` for layout and full-page templates.  
Use Aitch for fragments, dynamic sections, or when you want correctness, composability, and **predictable performance**.

You can inject Aitch-rendered fragments using `template.HTML(...)` cleanly and safely.

Aitch includes a NewTemplate() wrapper for `html/template` that lets you bind Aitch nodes to `{{token .}}` markers — rendering them directly to the output stream.

This gives you clean integration without resorting to raw `template.HTML(...)` injection — while keeping layout logic in templates and dynamic fragments in code where they belong.

<details>
    <summary><strong>Example</strong></summary>

```go
package main

import (
    "github.com/go-andiamo/aitch"
    "github.com/go-andiamo/aitch/html"
    "os"
)

const baseHtml = `<!DOCTYPE html>
<html {{lang .}}>
<head>
{{headers .}}
</head>
<body>
{{body .}}
</body>
</html>`

var langAtt = html.Lang(aitch.DynamicValueKey("lang"))
var headers = aitch.Collection(
    html.Meta(html.Charset("utf-8")),
    html.TitleElement(aitch.DynamicValueKey("title")))
var body = aitch.Collection(
    html.P("Hello, ", aitch.DynamicValueKey("username")))
var template = aitch.MustNewTemplate("index", baseHtml, aitch.NodeMap{
    "lang":    langAtt,
    "headers": headers,
    "body":    body,
})

func main() {
    data := map[string]any{
        "lang":     "en",
        "title":    "Test rendered",
        "username": "Aitch",
    }
    _ = template.Execute(os.Stdout, data, nil)
}
```
produces...
```html
<!DOCTYPE html>
<html  lang="en">
<head>
<meta charset="utf-8"><title>Test rendered</title>
</head>
<body>
<p>Hello, Aitch</p>
</body>
</html>
```
</details>