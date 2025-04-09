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
