# Aitch ("H") - for HTML (or SVG)
[![GoDoc](https://godoc.org/github.com/go-andiamo/aitch?status.svg)](https://pkg.go.dev/github.com/go-andiamo/aitch)
[![Latest Version](https://img.shields.io/github/v/tag/go-andiamo/aitch.svg?sort=semver&style=flat&label=version&color=blue)](https://github.com/go-andiamo/aitch/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-andiamo/aitch)](https://goreportcard.com/report/github.com/go-andiamo/aitch)

Aitch is a fluent, low/zero-GC pressure HTML and SVG rendering library for Go.

Built for developers who care about performance, scalability and clean code, Aitch gives you the power to generate dynamic HTML/SVG with surgical precision — while playing perfectly with `html/template`.

---

## Features

- Fluent, composable, declarative API for HTML and SVG
- Mixed content and attributes handled cleanly — order doesn’t matter
    - Attributes don't have to be specified before element contents  
    - Repeated attributes are protected: either the last one wins (e.g., `id`), or values are merged (e.g., `class`, `style`)
- Attribute merging (e.g., multiple `Class()` or `Style()` declarations)
- Virtually zero allocations during rendering<sup>†</sup>
- Conditional rendering (elements, attributes, values)
- Boolean attribute and void element aware
- Dynamic value support using render context data
- Smart value rendering (specify element content & attribute values as almost any type<sup>‡</sup>)
- Imperative rendering for absolute dynamically built HTML/SVG
- Seamless integration with `html/template`
- Full godocs (including HTML & SVG links to MDN docs)
- Inline style generation via fluent `css` module
- 100% tested and production-ready

†: With totally static templates (no dynamic attributes, no conditional attributes, no imperative rendering) zero allocs (zero GC-pressure) is almost guaranteed!
Dynamic values or conditional logic may introduce some allocations — this is unavoidable in Go, but Aitch minimizes them wherever possible.<br>
‡: Types rendered: `bool`,`[]byte`,`string`,`int`,`int8`,`int16`,`int32`,`int64`,`uint`,`unit8`,`uint16`,`uint32`,`uint64`,`float32`,`float64` - or any type that implements `fmt.Stringer`

---

## Installation

```bash
go get github.com/go-andiamo/aitch
```

---

## Declarative Templates, Runtime Flexibility

Aitch separates **template declaration** from **template rendering**.

Templates can be declared once — as plain Go `var`s (everything, as far as possible, is prepped for fast output writing)
Rendering is done later, with context-aware data and **zero or low allocations**.

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

[try on go-playground](https://go.dev/play/p/zQ4t_pvTewD)

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

[try on go-playground](https://go.dev/play/p/S0n5mMeYtzJ)

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

[try on go-playground](https://go.dev/play/p/MyxxnMZmCKt)

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

[try on go-playground](https://go.dev/play/p/xwRVLSRQBK2)

Or a simplified version using `aitch.When` rather than `aitch.Conditional`...
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
    aitch.When("isAdmin", html.Class("admin")),
)

func main() {
    data := map[string]any{
        "username": "Aitch",
        "isAdmin":  true,
    }
    ctx := context.New(os.Stdout, data, nil)
    _ = template.Render(ctx)
}
```

[try on go-playground](https://go.dev/play/p/Ue2WFQ5cE44)

</details><br>

<details>
    <summary><strong>5. Imperative generation</strong></summary>

Because sometimes you may want to take complete control of writing markup...

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

[try on go-playground](https://go.dev/play/p/2QcIv3bsQez)

</details><br>

<details>
    <summary><strong>6. Iterate (key)</strong></summary>

```go
package main

import (
    "github.com/go-andiamo/aitch"
    "github.com/go-andiamo/aitch/context"
    "github.com/go-andiamo/aitch/html"
    "os"
)

var template = html.Div(
    html.Select(
        aitch.IterateKey("items",
            html.Option(html.Value(aitch.DynamicValueKey("value")), aitch.DynamicValueKey("label"))),
    ))

func main() {
    data := map[string]any{
        "items": []map[string]any{
            {
                "label": "First",
                "value": 1,
            },
            {
                "label": "Second",
                "value": 2,
            },
            {
                "label": "Third",
                "value": 3,
            },
        },
    }
    ctx := context.New(os.Stdout, data, nil)
    _ = template.Render(ctx)
}
```
produces...
```html
<div><select><option value="1">First</option><option value="2">Second</option><option value="3">Third</option></select></div>
```

[try on go-playground](https://go.dev/play/p/7OKpYrtkuZB)

</details><br>

<details>
    <summary><strong>7. Iterate (func)</strong></summary>

```go
package main

import (
    "github.com/go-andiamo/aitch"
    "github.com/go-andiamo/aitch/context"
    "github.com/go-andiamo/aitch/html"
    "os"
)

var template = html.Div(
    html.Select(
        aitch.Iterate(func(ctx *context.Context) []any {
            return []any{
                map[string]any{
                    "label": "First",
                    "value": 1,
                },
                map[string]any{
                    "label": "Second",
                    "value": 2,
                },
                map[string]any{
                    "label": "Third",
                    "value": 3,
                },
            }
        }, html.Option(html.Value(aitch.DynamicValueKey("value")), aitch.DynamicValueKey("label"))),
    ))

func main() {
    ctx := context.New(os.Stdout, nil, nil)
    _ = template.Render(ctx)
}

```
produces...
```html
<div><select><option value="1">First</option><option value="2">Second</option><option value="3">Third</option></select></div>
```

[try on go-playground](https://go.dev/play/p/cqjrUfkevr-)

</details><br>

<details>
    <summary><strong>8. Iterate (yield)</strong></summary>

```go
package main

import (
    "github.com/go-andiamo/aitch"
    "github.com/go-andiamo/aitch/context"
    "github.com/go-andiamo/aitch/html"
    "os"
)

var template = html.Div(
    html.Select(
        aitch.IterateYield(func(ctx *context.Context, yield func(any)) {
            labels, _ := context.Get[[]string](ctx, "options")
            for i, label := range labels {
                yield(map[string]any{
                    "label":    label,
                    "value":    i + 1,
                    "selected": i == 1,
                })
            }
        }, html.Option(html.Value(aitch.DynamicValueKey("value")), aitch.DynamicValueKey("label"), aitch.When("selected", html.Selected()))),
    ))

func main() {
    data := map[string]any{
        "options": []string{"First", "Second", "Third"},
    }
    ctx := context.New(os.Stdout, data, nil)
    _ = template.Render(ctx)
}
```
produces...
```html
<div><select><option value="1">First</option><option value="2" selected>Second</option><option value="3">Third</option></select></div>
```

[try on go-playground](https://go.dev/play/p/VbNhTy_Hz10)

</details><br>

<details>
    <summary><strong>9. Smart value rendering</strong></summary>

```go
package main

import (
    "github.com/go-andiamo/aitch/context"
    "github.com/go-andiamo/aitch/html"
    "os"
)

var template = html.P(
    "Hello ", []byte("there "), true, " ", 1, " ", 1.2, " ", uint16(3), func() any { return " something" },
    html.Class(true), html.Class(1), html.Class(2.2))

func main() {
    ctx := context.New(os.Stdout, nil, nil)
    _ = template.Render(ctx)
}
```
produces...
```html
<p class="true 1 2.2">Hello there true 1 1.2 3 something</p>
```

[try on go-playground](https://go.dev/play/p/LNM82Gka_3f)

</details><br>

<details>
    <summary><strong>10. Inline CSS (using <code>css</code> module)</strong></summary>

Although it's not encouraged, sometimes inline CSS is handy — for dynamic fragments, micro-layout tweaks, or when you're not working with external stylesheets...

```go
package main

import (
    "github.com/go-andiamo/aitch"
    "github.com/go-andiamo/aitch/context"
    "github.com/go-andiamo/aitch/css"
    "github.com/go-andiamo/aitch/html"
    "os"
)

var template = html.P(
    css.Width(css.Px(aitch.DynamicValueKey("width")), css.Important),
    html.Style("scrollbar-gutter: stable"),
    aitch.When("isAdmin", css.BackgroundColor("red")),
    "My paragraph with inline-style",
)

func main() {
    data := map[string]any{
        "width":   10,
        "isAdmin": true,
    }
    ctx := context.New(os.Stdout, data, nil)
    _ = template.Render(ctx)
}
```
produces...
```html
<p style="width:10px !important; scrollbar-gutter: stable; background-color:red">My paragraph with inline-style</p>
```

</details><br>

---

## Not a replacement for `html/template`

Aitch complements `html/template` — it doesn’t try to replace it.

Use `html/template` for layout and full-page templates.  
Use Aitch for fragments, dynamic sections, or when you want correctness, composability, and **predictable performance**.

Aitch includes a NewTemplate() wrapper for `html/template` that lets you bind Aitch nodes to `{{token .}}` markers — rendering them directly to the template output stream.

This gives you clean integration without resorting to raw `template.HTML(...)` injection — while keeping layout logic in templates and real dynamic sections in code where they belong.

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

[try on go-playground](https://go.dev/play/p/V2ro2ZnBQcC)

</details><br>
