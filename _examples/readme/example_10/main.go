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
