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
