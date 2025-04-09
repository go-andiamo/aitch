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
