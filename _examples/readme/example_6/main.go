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
