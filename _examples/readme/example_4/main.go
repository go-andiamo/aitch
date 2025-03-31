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
